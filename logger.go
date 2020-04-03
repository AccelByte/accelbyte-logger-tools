package loggergo

import (
	"fmt"
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

func InitLogger(
	serviceName string,
	realm string,
) *logrus.Entry {
	logrus.SetFormatter(new(logrus.TextFormatter))

	data := logrus.Fields{
		"service": serviceName,
		"realm":   realm,
	}

	logrus.AddHook(&ContextHook{Data: data})

	return logrus.NewEntry(logrus.StandardLogger())
}

type ContextHook struct {
	Data logrus.Fields
}

func (hook ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook ContextHook) Fire(entry *logrus.Entry) error {
	pc := make([]uintptr, 3)
	cnt := runtime.Callers(6, pc)

	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i] - 1)
		name := fu.Name()

		if !strings.Contains(name, "github.com/sirupsen/logrus") &&
			!strings.Contains(name, "adapter.go") &&
			!strings.Contains(name, "logging.(*logTraceLogger)") &&
			!strings.Contains(name, "logrus.(*adapter)") &&
			!strings.Contains(name, "logrus.(*Entry)") &&
			!strings.Contains(name, "CloseWithLog") {
			file, line := fu.FileLine(pc[i] - 1)

			data := make(map[string]interface{})
			for k, v := range entry.Data {
				data[k] = v
			}

			data["file"] = fmt.Sprintf("%v:%v", path.Base(file), line)
			data["func"] = path.Base(name)

			for k, v := range hook.Data {
				data[k] = v
			}

			entry.Data = data

			break
		}
	}

	return nil
}
