package ready

import (
	"github.com/kaspikr/kaspid/infrastructure/logger"
	"github.com/kaspikr/kaspid/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.GoroutineWrapperFunc(log)
