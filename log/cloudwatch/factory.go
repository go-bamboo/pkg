package cloudwatch

import (
	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/log/core"
	"go.uber.org/zap/zapcore"
)

func init() {
	log.Register("aliyun", Create)
}

func Create(c *log.Conf) (core.Logger, error) {
	cc, err := NewCloudWatchCore(
		Level(zapcore.Level(c.CloudWatch.Level)),
		WithRegion(c.CloudWatch.Region),
		WithAccessKey(c.CloudWatch.Key),
		WithAccessSecret(c.CloudWatch.Secret),
		WithLogGroupName(c.CloudWatch.LogGroupName),
	)
	if err != nil {
		return nil, err
	}
	return cc, nil
}
