package db

import (
	"github.com/jimu-server/config"
	logs "github.com/jimu-server/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)
import "gorm.io/driver/postgres"

var Gorm *gorm.DB

func init() {
	var err error
	conStr := config.Evn.App.Database
	Gorm, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  conStr,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), getConfig())
	if err != nil {
		logs.Logger.Panic(err.Error())
	}
	sqlDB, err := Gorm.DB()
	if err != nil {
		logs.Logger.Panic(err.Error())
		return
	}

	if err := sqlDB.Ping(); err != nil {
		logs.Logger.Panic(err.Error())
		return
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func getConfig() *gorm.Config {
	newLogger := logger.New(
		log.New(logs.MultiWriteSyncer, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
	return &gorm.Config{
		Logger:                                   newLogger,
		DisableAutomaticPing:                     true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   true, // skip the snake_casing of names
		},
	}
}
