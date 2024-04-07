package db

import (
	"database/sql"
	"fmt"
	"github.com/jimu-os/gobatis"
	"github.com/jimu-server/config/config"
	"github.com/jimu-server/logger/logger"
	_ "github.com/lib/pq"
)

// GoBatis 实例
var GoBatis *gobatis.GoBatis
var DB *sql.DB

func init() {
	var err error
	conStr := config.Evn.App.Database
	DB, err = sql.Open("postgres", conStr)
	if err != nil {
		logger.Logger.Error(fmt.Sprint("数据库连接失败,", err.Error()))
		panic(err)
	}
	if err = DB.Ping(); err != nil {
		logger.Logger.Error(fmt.Sprint("Ping 数据库连接失败,", err.Error()))
		panic(err)
	}
	GoBatis = gobatis.New(DB)
	// PostgreSQL 就必须指定 Type 属性 兼容模版参数解析
	GoBatis.Type = gobatis.PostgreSQL
	GoBatis.Logs(logger.Logger)
}
