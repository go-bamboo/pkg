package cloudwatch

import (
	"github.com/go-bamboo/pkg/log/core"
	"go.uber.org/zap/zapcore"
)

func init() {
	core.Register("CloudWatch", Create)
}

func Create(c *core.Conf) (core.Logger, error) {
	cc, err := NewCloudWatchCore(
		Level(zapcore.Level(c.Level)),
		WithRegion(c.Region),
		WithAccessKey(c.Key),
		WithAccessSecret(c.Secret),
		WithLogGroupName(c.LogGroupName),
	)
	if err != nil {
		return nil, err
	}
	return cc, nil
}
