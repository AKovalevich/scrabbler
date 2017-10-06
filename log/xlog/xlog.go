package xlog

import (
	"os"

	"github.com/AKovalevich/scrabbler/log"
	"github.com/rs/xlog"
)

var (
	Do	log.Logger
)

func init() {
	Do = createNew(&log.Config{
		Level: log.LevelDebug,
		Time:  true,
		UTC:   true,
	})
}

// newXLog creates "github.com/rs/xlog" logger
func createNew(config *log.Config) log.Logger {
	var out xlog.Output
	switch config.Err {
	// We should find more matches between types of output
	case nil, os.Stderr:
		out = xlog.NewConsoleOutput()
	default:
		out = xlog.NewConsoleOutput()
	}
	return xlog.New(xlog.Config{
		Level:  xlog.Level(config.Level),
		Fields: config.Fields,
		Output: out,
	})
}