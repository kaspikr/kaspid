package common

import (
	"fmt"
	"github.com/kaspikr/kaspid/domain/dagconfig"
	"os"
	"sync/atomic"
	"syscall"
	"testing"
)

// RunKaspidForTesting runs kaspid for testing purposes
func RunKaspidForTesting(t *testing.T, testName string, rpcAddress string) func() {
	appDir, err := TempDir(testName)
	if err != nil {
		t.Fatalf("TempDir: %s", err)
	}

	kaspidRunCommand, err := StartCmd("KASPID",
		"kaspid",
		NetworkCliArgumentFromNetParams(&dagconfig.DevnetParams),
		"--appdir", appDir,
		"--rpclisten", rpcAddress,
		"--loglevel", "debug",
	)
	if err != nil {
		t.Fatalf("StartCmd: %s", err)
	}
	t.Logf("Kaspid started with --appdir=%s", appDir)

	isShutdown := uint64(0)
	go func() {
		err := kaspidRunCommand.Wait()
		if err != nil {
			if atomic.LoadUint64(&isShutdown) == 0 {
				panic(fmt.Sprintf("Kaspid closed unexpectedly: %s. See logs at: %s", err, appDir))
			}
		}
	}()

	return func() {
		err := kaspidRunCommand.Process.Signal(syscall.SIGTERM)
		if err != nil {
			t.Fatalf("Signal: %s", err)
		}
		err = os.RemoveAll(appDir)
		if err != nil {
			t.Fatalf("RemoveAll: %s", err)
		}
		atomic.StoreUint64(&isShutdown, 1)
		t.Logf("Kaspid stopped")
	}
}
