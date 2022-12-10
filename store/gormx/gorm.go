package gormx

import (
	"github.com/go-bamboo/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
)

// DBType database type
type DBType string

const (
	// dbMySQL Gorm Drivers mysql || postgres || sqlite || sqlserver
	dbMySQL     DBType = "mysql"
	dbPostgres  DBType = "postgres"
	dbSQLite    DBType = "sqlite"
	dbSQLServer DBType = "sqlserver"
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
	if DBType(c.Driver) == dbMySQL {
		dialector = mysql.Open(c.Source)
	} else if DBType(c.Driver) == dbPostgres {
		dialector = postgres.Open(c.Source)
	} else if DBType(c.Driver) == dbSQLite {
		dialector = sqlite.Open(c.Source)
	} else if DBType(c.Driver) == dbSQLServer {
		dialector = sqlserver.Open(c.Source)
	}
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
