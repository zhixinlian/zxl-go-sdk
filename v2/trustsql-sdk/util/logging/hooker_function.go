package logging

import (
	"path/filepath"
	"runtime"
	"strings"

	logrus "github.com/sirupsen/logrus"
)

type functionHooker struct {
	innerLogger *logrus.Logger
	file        string
}

func (h *functionHooker) Fire(entry *logrus.Entry) error {
	pc := make([]uintptr, 10)
	runtime.Callers(6, pc)
	for i := 0; i < 10; i++ {
		if pc[i] == 0 {
			break
		}
		f := runtime.FuncForPC(pc[i])
		file, line := f.FileLine(pc[i])
		if strings.Contains(file, "sirupsen") {
			continue
		}
		fname := f.Name()
		if strings.Contains(fname, "/") {
			index := strings.LastIndex(fname, "/")
			entry.Data["func"] = fname[index+1:]
			// entry.Data["package"] = fname[0:index]
		} else {
			entry.Data["func"] = fname
		}
		entry.Data["line"] = line
		entry.Data["file"] = filepath.Base(file)
		break
	}
	return nil
}

func (h *functionHooker) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}

// LoadFunctionHooker loads a function hooker to the logger
func LoadFunctionHooker(logger *logrus.Logger) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
	}
	inst := &functionHooker{
		innerLogger: logger,
		file:        file,
	}
	logger.Hooks.Add(inst)
}
