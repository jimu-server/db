package db

import (
	"database/sql"
	"fmt"
	"github.com/jimu-os/gobatis"
	"github.com/jimu-server/config"
	"github.com/jimu-server/logger"
	_ "github.com/mattn/go-sqlite3"
)

// LocalGoBatis 实例
var LocalGoBatis *gobatis.GoBatis
var LocalDB *sql.DB

func init() {
	var err error
	conStr := config.Evn.App.FileDb
	LocalDB, err = sql.Open("sqlite3", conStr)
	if err != nil {
		logger.Logger.Error(fmt.Sprint("Database connection failure,", err.Error()))
		panic(err)
	}
	if err = DB.Ping(); err != nil {
		logger.Logger.Error(fmt.Sprint("Ping the database connection failed. Procedure,", err.Error()))
		panic(err)
	}
	LocalGoBatis = gobatis.New(LocalDB)
	LocalGoBatis.DbType(gobatis.Sqlite)
	LocalGoBatis.Logs(logger.Logger)
}
