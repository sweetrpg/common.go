package logging

import (
	"log"
	"time"

	"github.com/sweetrpg/common/constants"
	"github.com/sweetrpg/common/util"
	"github.com/zerodha/logf"
)

var Logger logf.Logger

// Initialize logging setup.
func Init() {
	level := logf.InfoLevel
	switch util.GetEnv(constants.LOG_LEVEL, constants.INFO) {
	case constants.DEBUG:
		log.Print("Setting log level to DEBUG")
		level = logf.DebugLevel
	case constants.WARN:
		log.Print("Setting log level to WARN")
		level = logf.WarnLevel
	case constants.ERROR:
		log.Print("Setting log level to ERROR")
		level = logf.ErrorLevel
	}
	log.Printf("Log level: %v", level)

	Logger = logf.New(logf.Opts{
		EnableColor:          true,
		Level:                level,
		CallerSkipFrameCount: 3,
		EnableCaller:         true,
		TimestampFormat:      time.RFC3339,
		DefaultFields:        []any{},
	})
}
