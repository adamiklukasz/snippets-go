package log

import (
	"math/rand"
	"testing"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/sirupsen/logrus"
)

/*****************************************************************************/

type MyHook struct{}

func (*MyHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (*MyHook) Fire(e *logrus.Entry) error {
	e.Data["password"] = "<OBFUSCATED>"
	return nil
}

/*****************************************************************************/

func TestLoggerHooks(t *testing.T) {
	InitLogger()
	ProduceLogs()

	select {}
}

func InitLogger() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.AddHook(&MyHook{})
}

func ProduceLogs() {
	logF := logrus.WithField("password", "asdasd")

	for i := 0; i < 10; i++ {
		logF.Infof(randomdata.Paragraph())
		time.Sleep(time.Duration(rand.Int63n(100)) * time.Millisecond)
	}
}
