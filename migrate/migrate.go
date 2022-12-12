package migrate

import (
	"github.com/go-bamboo/pkg/log"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database"
	_ "github.com/golang-migrate/migrate/v4/source"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	SourceURL   string `protobuf:"bytes,3,opt,name=sourceURL,proto3" json:"SourceURL,omitempty"`
	databaseURL string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Version     uint
}

func Up(c *Config) error {
	m, err := migrate.New(c.SourceURL, c.databaseURL)
	defer func() {
		if err != nil {
			if _, err := m.Close(); err != nil {
				log.Error(err)
			}
		}
	}()
	if err != nil {
		return err
	}
	m.Log = NewLogger(zapcore.DebugLevel, log.GetCore())
	if err = m.Up(); err != nil {
		if err.Error() == "no change" {
		}
		return err
	}
	return nil
}

func Goto(c *Config) error {
	m, err := migrate.New(c.SourceURL, c.databaseURL)
	defer func() {
		if err != nil {
			if _, err := m.Close(); err != nil {
				log.Error(err)
			}
		}
	}()
	if err != nil {
		return err
	}
	m.Log = NewLogger(zapcore.DebugLevel, log.GetCore())
	if err = m.Migrate(c.Version); err != nil {
		if err.Error() == "no change" {
		}
		return err
	}
	return nil
}
