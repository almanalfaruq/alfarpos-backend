package logger

import (
	"os"

	"github.com/almanalfaruq/alfarpos-backend/util"
	"github.com/kataras/golog"
)

var Log *golog.Logger

func New(cfg *util.Config) error {
	if cfg.Env == "dev" {
		Log.SetLevel("debug")
	}
	// use  debug.log and info.log files for the example.
	debugFile, err := os.OpenFile(cfg.Log.PathDebug, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer debugFile.Close()

	infoFile, err := os.OpenFile(cfg.Log.PathInfo, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer infoFile.Close()

	Log = golog.New()

	Log.SetLevelOutput("info", infoFile)
	Log.SetLevelOutput("debug", debugFile)
	return nil
}
