package logging

import (
	"os"
	"testing"

	"github.com/sweetrpg/common.go/constants"
)

func TestInitConfiguresLoggerForEachLevel(t *testing.T) {
	levels := []string{constants.DEBUG, constants.WARN, constants.ERROR, constants.INFO, ""}

	for _, level := range levels {
		t.Run(level, func(t *testing.T) {
			if level == "" {
				os.Unsetenv(constants.LOG_LEVEL)
			} else {
				os.Setenv(constants.LOG_LEVEL, level)
			}
			defer os.Unsetenv(constants.LOG_LEVEL)

			Init()
			Logger.Info("logger initialized", "level", level)
		})
	}
}
