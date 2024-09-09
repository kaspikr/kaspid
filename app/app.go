package app

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/kaspikr/kaspid/infrastructure/config"
	"github.com/kaspikr/kaspid/infrastructure/db/database"
	"github.com/kaspikr/kaspid/infrastructure/db/database/ldb"
	"github.com/kaspikr/kaspid/infrastructure/logger"
	"github.com/kaspikr/kaspid/infrastructure/os/execenv"
	"github.com/kaspikr/kaspid/infrastructure/os/limits"
	"github.com/kaspikr/kaspid/infrastructure/os/signal"
	"github.com/kaspikr/kaspid/infrastructure/os/winservice"
	"github.com/kaspikr/kaspid/util/panics"
	"github.com/kaspikr/kaspid/util/profiling"
	"github.com/kaspikr/kaspid/version"
)

const (
	leveldbCacheSizeMiB = 256
	defaultDataDirname  = "datadir2"
)

var desiredLimits = &limits.DesiredLimits{
	FileLimitWant: 2048,
	FileLimitMin:  1024,
}

var serviceDescription = &winservice.ServiceDescription{
	Name:        "kaspidsvc",
	DisplayName: "Kaspid Service",
	Description: "Downloads and stays synchronized with the Kaspi blockDAG and " +
		"provides DAG services to applications.",
}

type kaspidApp struct {
	cfg *config.Config
}

// StartApp starts the kaspid app, and blocks until it finishes running
func StartApp() error {
	execenv.Initialize(desiredLimits)

	// Load configuration and parse command line. This function also
	// initializes logging and configures it accordingly.
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	defer logger.BackendLog.Close()
	defer panics.HandlePanic(log, "MAIN", nil)

	app := &kaspidApp{cfg: cfg}

	// Call serviceMain on Windows to handle running as a service. When
	// the return isService flag is true, exit now since we ran as a
	// service. Otherwise, just fall through to normal operation.
	if runtime.GOOS == "windows" {
		isService, err := winservice.WinServiceMain(app.main, serviceDescription, cfg)
		if err != nil {
			return err
		}
		if isService {
			return nil
		}
	}

	return app.main(nil)
}

func (app *kaspidApp) main(startedChan chan<- struct{}) error {
	// Get a channel that will be closed when a shutdown signal has been
	// triggered either from an OS signal such as SIGINT (Ctrl+C) or from
	// another subsystem such as the RPC server.
	interrupt := signal.InterruptListener()
	defer log.Info("Shutdown complete")

	// Show version at startup.
	log.Infof("Version %s", version.Version())

	// Enable http profiling server if requested.
	if app.cfg.Profile != "" {
		profiling.Start(app.cfg.Profile, log)
	}
	profiling.TrackHeap(app.cfg.AppDir, log)

	// Return now if an interrupt signal was triggered.
	if signal.InterruptRequested(interrupt) {
		return nil
	}

	if app.cfg.ResetDatabase {
		err := removeDatabase(app.cfg)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	// Open the database
	databaseContext, err := openDB(app.cfg)
	if err != nil {
		log.Errorf("Loading database failed: %+v", err)
		return err
	}

	defer func() {
		log.Infof("Gracefully shutting down the database...")
		err := databaseContext.Close()
		if err != nil {
			log.Errorf("Failed to close the database: %s", err)
		}
	}()

	// Return now if an interrupt signal was triggered.
	if signal.InterruptRequested(interrupt) {
		return nil
	}

	// Create componentManager and start it.
	componentManager, err := NewComponentManager(app.cfg, databaseContext, interrupt)
	if err != nil {
		log.Errorf("Unable to start kaspid: %+v", err)
		return err
	}

	defer func() {
		log.Infof("Gracefully shutting down kaspid...")

		shutdownDone := make(chan struct{})
		go func() {
			componentManager.Stop()
			shutdownDone <- struct{}{}
		}()

		const shutdownTimeout = 2 * time.Minute

		select {
		case <-shutdownDone:
		case <-time.After(shutdownTimeout):
			log.Criticalf("Graceful shutdown timed out %s. Terminating...", shutdownTimeout)
		}
		log.Infof("Kaspid shutdown complete")
	}()

	componentManager.Start()

	if startedChan != nil {
		startedChan <- struct{}{}
	}

	// Wait until the interrupt signal is received from an OS signal or
	// shutdown is requested through one of the subsystems such as the RPC
	// server.
	<-interrupt
	return nil
}

// dbPath returns the path to the block database given a database type.
func databasePath(cfg *config.Config) string {
	return filepath.Join(cfg.AppDir, defaultDataDirname)
}

func removeDatabase(cfg *config.Config) error {
	dbPath := databasePath(cfg)
	return os.RemoveAll(dbPath)
}

func openDB(cfg *config.Config) (database.Database, error) {
	dbPath := databasePath(cfg)

	err := checkDatabaseVersion(dbPath)
	if err != nil {
		return nil, err
	}

	log.Infof("Loading database from '%s'", dbPath)
	db, err := ldb.NewLevelDB(dbPath, leveldbCacheSizeMiB)
	if err != nil {
		return nil, err
	}

	return db, nil
}
