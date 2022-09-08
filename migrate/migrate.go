package migrate

import (
	"database/sql"

	"github.com/emberfarkas/pkg/ecode"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Config struct {
	SourceURL string `protobuf:"bytes,3,opt,name=sourceURL,proto3" json:"SourceURL,omitempty"`
	Driver    string `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Source    string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Version   uint
}

func MigrateUp(c *Config) error {
	db, err := sql.Open(c.Driver, c.Source)
	if err != nil {
		err = ecode.WrapStdError(err)
		return err
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		err = ecode.WrapStdError(err)
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(c.SourceURL, c.Driver, driver)
	if err != nil {
		err = ecode.WrapStdError(err)
		return err
	}
	if err = m.Up(); err != nil {
		if err.Error() == "no change" {
			goto handleDone
		}
		err = ecode.WrapStdError(err)
		return err
	}
handleDone:
	err1, err2 := m.Close()
	if err1 != nil {
		err = ecode.WrapStdError(err1)
		return err
	}
	if err2 != nil {
		err = ecode.WrapStdError(err2)
		return err
	}
	return nil
}

func Migrate(c *Config) error {
	db, err := sql.Open(c.Driver, c.Source)
	if err != nil {
		err = ecode.WrapStdError(err)
		return err
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		err = ecode.WrapStdError(err)
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(c.SourceURL, c.Driver, driver)
	if err != nil {
		err = ecode.WrapStdError(err)
		return err
	}
	if err = m.Migrate(c.Version); err != nil {
		if err.Error() == "no change" {
			goto handleDone
		}
		err = ecode.WrapStdError(err)
		return err
	}
handleDone:
	err1, err2 := m.Close()
	if err1 != nil {
		err = ecode.WrapStdError(err1)
		return err
	}
	if err2 != nil {
		err = ecode.WrapStdError(err2)
		return err
	}
	return nil
}
