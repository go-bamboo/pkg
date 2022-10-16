package gormx

import (
	"github.com/emberfarkas/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
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
	dialector := mysql.Open(c.Source)
	gormlogConfig := gormlog.Config{Colorful: true, LogLevel: gormlog.LogLevel(c.LogLevel)}
	core := log.GetCore()
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: NewLogger(gormlogConfig, core),
	})
	if err != nil {
		return nil, err
	}
	if err = db.Use(NewGormTracingHook()); err != nil {
		return nil, err
	}
	//if err := db.Use(tracing.NewPlugin()); err != nil {
	//	return nil, err
	//}
	return db, nil
}
