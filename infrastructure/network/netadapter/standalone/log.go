package standalone

import (
	"github.com/kaspikr/kaspid/infrastructure/logger"
	"github.com/kaspikr/kaspid/util/panics"
)

var log = logger.RegisterSubSystem("NTAR")
var spawn = panics.GoroutineWrapperFunc(log)
