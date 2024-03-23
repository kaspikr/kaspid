package rpc

import (
	"github.com/kaspikr/kaspid/infrastructure/logger"
	"github.com/kaspikr/kaspid/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
