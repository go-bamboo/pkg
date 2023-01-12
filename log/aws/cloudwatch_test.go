package aws_test

import (
	"testing"
	"time"

	"github.com/gofor-little/env"
)

func TestCloudWatchLogger(t *testing.T) {
	//is := is.New(t)

	if err := env.Load(".env"); err != nil {
		t.Log(".env file not found, ignore this if running in CI/CD Pipeline")
	}

	//var err error
	// AWS_PROFILE

	//c, err = NewCloudWatchCore(context.Background(), env.Get("AWS_PROFILE", ""), env.Get("AWS_REGION", ""), "CloudWatchLoggerTest", log.Fields{
	//	"tag": "cloudWatchLoggerTest",
	//})
	//is.NoErr(err)
	//
	//err = log.Info(log.Fields{
	//	"string": "test info string",
	//	"bool":   true,
	//	"int":    64,
	//	"float":  3.14159,
	//})
	//is.NoErr(err)
	//
	//err = log.Error(log.Fields{
	//	"string": "test error string",
	//	"bool":   true,
	//	"int":    64,
	//	"float":  3.14159,
	//})
	//is.NoErr(err)
	//
	//err = log.Debug(log.Fields{
	//	"string": "test debug string",
	//	"bool":   true,
	//	"int":    64,
	//	"float":  3.14159,
	//})
	//is.NoErr(err)

	time.Sleep(time.Second)
}
