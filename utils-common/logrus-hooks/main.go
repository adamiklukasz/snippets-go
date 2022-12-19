package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/sirupsen/logrus"
)

var logF = logrus.WithFields(logrus.Fields{"test": "adamik"})

const (
	paragraphProduced = 100
	exitDelay         = 5 * time.Second
)

/*****************************************************************************/

type MyHook struct{}

func (*MyHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (*MyHook) Fire(e *logrus.Entry) error {
	fmt.Printf("[Hook] FIELDS=%#v\n", e.Data)
	fmt.Printf("[Hook] MSG=%#v\n", e.Message)
	return nil
}

/*****************************************************************************/

func main() {
	ProduceLogs()

	time.Sleep(exitDelay)
}

func init() {
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
}

func ProduceLogs() {
	for i := 0; i < paragraphProduced; i++ {
		logF.Infof(randomdata.Paragraph())
		time.Sleep(time.Duration(rand.Int63n(100)) * time.Millisecond)
		logF.Warnf(randomdata.Paragraph())
	}
	logF.Log(logrus.TraceLevel, "Logs production completed")
}
