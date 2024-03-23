package main

import (
	"github.com/kaspikr/kaspid/infrastructure/logger"
	"github.com/kaspikr/kaspid/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("ORPH")
	spawn      = panics.GoroutineWrapperFunc(log)
)
