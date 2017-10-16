package logger

import (
	"os"
	"testing"

	"github.com/albert-widi/go_common/errors"

	klog "github.com/go-kit/kit/log"

	"github.com/Sirupsen/logrus"
)

var log *Logger
var llog klog.Logger

func init() {
	log = fake()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	f, _ := os.Open("/dev/null")
	logrus.StandardLogger().Out = f
	llog = klog.NewJSONLogger(f)
}

func BenchmarkSimpleLogger(b *testing.B) {
	for n := 0; n < b.N; n++ {
		log.Error("Something is wrong")
	}
}

func BenchmarkLoggerWithFields(b *testing.B) {
	for n := 0; n < b.N; n++ {
		log.WithFields(Fields{"field1": "value1"}).Info("This is a info with fields")
	}
}

func BenchmarkErrors(b *testing.B) {
	err := errors.New("This is new errors", errors.Fields{"err1": "err_value_1"})
	for n := 0; n < b.N; n++ {
		log.Errors(err)
	}
}

func BenchmarkErrorsWithFields(b *testing.B) {
	err := errors.New("This is new errors", errors.Fields{"err1": "err_value_1"})
	for n := 0; n < b.N; n++ {
		log.WithFields(Fields{"field1": "value1"}).Errors(err)
	}
}

func BenchmarkGokitLog(b *testing.B) {
	for n := 0; n < b.N; n++ {
		llog.Log("Something is wrong")
	}
}

// func BenchmarkSimpleLogrus(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		logrus.Error("Something is wrong")
// 	}
// }

// func BenchmarkLogrusWithFields(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		logrus.WithFields(logrus.Fields{"field1": "value1"}).Info("This is a info with fields")
// 	}
// }
