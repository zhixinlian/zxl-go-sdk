package logging

import (
	"os"
	"path/filepath"
	"time"

	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// NewFileRotateHooker enable log file output
func NewFileRotateHooker(path string, age uint32) logrus.Hook {
	if len(path) == 0 {
		panic("Failed to parse logger folder:" + path + ".")
	}
	if !filepath.IsAbs(path) {
		path, _ = filepath.Abs(path)
	}
	if err := os.MkdirAll(path, 0700); err != nil {
		panic("Failed to create logger folder:" + path + ". err:" + err.Error())
	}
	filePath := path + "/neb-%Y%m%d-%H.log"
	linkPath := path + "/neb.log"
	writer, err := rotatelogs.New(
		filePath,
		rotatelogs.WithLinkName(linkPath),
		rotatelogs.WithRotationTime(time.Duration(3600)*time.Second),
	)
	// set log file max age
	if age > 0 {
		rotatelogs.WithMaxAge(time.Duration(age) * time.Second).Configure(writer)
	}

	if err != nil {
		panic("Failed to create rotate logs. err:" + err.Error())
	}

	hook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
	}, nil)
	return hook
}
