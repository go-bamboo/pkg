package gormx

import (
	"github.com/go-bamboo/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type DB = gorm.DB

func MustNew(c *Conf) *DB {
	db, err := New(c)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func New(c *Conf) (*DB, error) {
	var dialector gorm.Dialector
	if DBType(c.Driver) == DBType_mysql {
		dialector = mysql.Open(c.Source)
	} else if DBType(c.Driver) == DBType_postgres {
		dialector = postgres.Open(c.Source)
	} else if DBType(c.Driver) == DBType_sqlite {
		dialector = sqlite.Open(c.Source)
	} else if DBType(c.Driver) == DBType_sqlserver {
		dialector = sqlserver.Open(c.Source)
	}
	core := log.GetCore()
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: NewLogger(c.Logger, core),
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
	if err = db.Use(NewGormTracer()); err != nil {
		return nil, err
	}
	if err = db.Use(NewGormError()); err != nil {
		return nil, err
	}
	return db, nil
}
