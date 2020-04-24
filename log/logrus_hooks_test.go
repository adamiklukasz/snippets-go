package log

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
	"math/rand"
	"testing"
	"time"
)

var logF = logrus.WithFields(logrus.Fields{"test": "adamik"})

/*****************************************************************************/

type MyHook struct{}

func (*MyHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (*MyHook) Fire(e *logrus.Entry) error {
	fmt.Printf("!!!!!! FIELDS=%#v\n", e.Data)
	fmt.Printf("@@@@@@ MSG=%#v\n", e.Message)
	return nil
}

/*****************************************************************************/

func TestLoggerHooks(t *testing.T) {
	InitLogger()
	ProduceLogs()

	select {}
}

func InitLogger() {
	formatter := logrus.TextFormatter{
		TimestampFormat: time.RFC3339Nano,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "timestamp",
			logrus.FieldKeyFile: "location",
			logrus.FieldKeyFunc: "block",
		},
	}
	logrus.SetFormatter(&formatter)
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.TraceLevel)
	logrus.AddHook(&MyHook{})

	_, _ = logrus.ParseLevel("warn")

	lf := &logrustash.LogstashFormatter{}
	logrus.SetFormatter(lf)
}

func ProduceLogs() {
	for i := 0; i < 10; i++ {
		logF.Infof(randomdata.Paragraph())
		time.Sleep(time.Duration(rand.Int63n(100)) * time.Millisecond)
		logF.Warnf(randomdata.Paragraph())
	}
	logF.Log(logrus.TraceLevel, "Some message")
}
