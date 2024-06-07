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
		logger.Logger.Error(fmt.Sprint("数据库连接失败,", err.Error()))
		panic(err)
	}
	if err = DB.Ping(); err != nil {
		logger.Logger.Error(fmt.Sprint("Ping 数据库连接失败,", err.Error()))
		panic(err)
	}
	LocalGoBatis = gobatis.New(LocalDB)
	// PostgreSQL 就必须指定 Type 属性 兼容模版参数解析
	LocalGoBatis.DbType(gobatis.Sqlite)
	LocalGoBatis.Logs(logger.Logger)
}
