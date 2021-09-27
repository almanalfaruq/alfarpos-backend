package logger

import (
	"os"

	"github.com/almanalfaruq/alfarpos-backend/util"
	"github.com/kataras/golog"
)

var Log *golog.Logger

func New(cfg *util.Config) (error, func()) {
	Log = golog.New()

	if cfg.Env == "dev" || cfg.Debug {
		Log.SetLevel("debug")
	}

	if cfg.Env == "local" {
		return nil, func() {}
	}

	// use  debug.log and info.log files for the example.
	debugFile, err := os.OpenFile(cfg.Log.PathDebug, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err, func() {}
	}

	infoFile, err := os.OpenFile(cfg.Log.PathInfo, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err, func() {}
	}

	Log.SetLevelOutput("info", infoFile)
	Log.SetLevelOutput("debug", debugFile)
	return nil, func() {
		infoFile.Close()
		debugFile.Close()
	}
}
