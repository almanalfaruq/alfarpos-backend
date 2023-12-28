package logger

import (
	"os"

	"github.com/almanalfaruq/alfarpos-backend/util"
	"github.com/kataras/golog"
)

var Log *golog.Logger

func New(cfg *util.Config) (func(), error) {
	Log = golog.New()
	golog.SetLevel("info")

	if cfg.Env == "dev" || cfg.Debug {
		golog.SetLevel("debug")
		Log.SetLevel("debug")
	}

	if cfg.Env == "local" {
		return func() {}, nil
	}

	// use  debug.log and info.log files for the example.
	debugFile, err := os.OpenFile(cfg.Log.PathDebug, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return func() {}, nil
	}

	infoFile, err := os.OpenFile(cfg.Log.PathInfo, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return func() {}, nil
	}

	Log.SetLevelOutput("info", infoFile)
	golog.SetLevelOutput("info", infoFile)
	Log.SetLevelOutput("debug", debugFile)
	golog.SetLevelOutput("debug", debugFile)
	return func() {
		infoFile.Close()
		debugFile.Close()
	}, nil
}
