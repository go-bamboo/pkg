package gormx

import (
	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/store/gormx/conf"
	"github.com/go-bamboo/pkg/store/gormx/logger"
	"github.com/go-bamboo/pkg/store/gormx/plugins"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type DB = gorm.DB

func MustNew(c *conf.Conf) *DB {
	db, err := New(c)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func New(c *conf.Conf) (*DB, error) {
	var dialector gorm.Dialector
	if conf.DBType(c.Driver) == conf.DBType_mysql {
		dialector = mysql.Open(c.Source)
	} else if conf.DBType(c.Driver) == conf.DBType_postgres {
		dialector = postgres.Open(c.Source)
	} else if conf.DBType(c.Driver) == conf.DBType_sqlite {
		dialector = sqlite.Open(c.Source)
	} else if conf.DBType(c.Driver) == conf.DBType_sqlserver {
		dialector = sqlserver.Open(c.Source)
	}
	core := log.GetCore()
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.NewLogger(c.Logger, core),
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if c.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(int(c.MaxOpenConns))
	}
	if c.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(int(c.MaxIdleConns))
	}
	if c.ConnMaxLifetime.AsDuration() > 0 {
		sqlDB.SetConnMaxLifetime(c.ConnMaxLifetime.AsDuration())
	}
	if err = db.Use(plugins.NewGormTracer()); err != nil {
		return nil, err
	}
	if err = db.Use(plugins.NewGormError()); err != nil {
		return nil, err
	}
	return db, nil
}
