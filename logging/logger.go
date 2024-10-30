package logging

import (
	"time"

	"github.com/sweetrpg/common/constants"
	"github.com/sweetrpg/common/util"
	"github.com/zerodha/logf"
)

var Logger logf.Logger

func Init() {
	level := logf.InfoLevel
	switch util.GetEnv(constants.LOG_LEVEL, constants.INFO) {
	case constants.DEBUG:
		level = logf.DebugLevel
	case constants.WARN:
		level = logf.WarnLevel
	case constants.ERROR:
		level = logf.ErrorLevel
	}

	Logger = logf.New(logf.Opts{
		EnableColor:          true,
		Level:                level,
		CallerSkipFrameCount: 3,
		EnableCaller:         true,
		TimestampFormat:      time.RFC3339,
		DefaultFields:        []any{},
	})
}
