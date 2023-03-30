package model

import (
	"CoolGoPkg/apply_gorm/v1/conf"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/url"
)

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

func DB() *gorm.DB {
	return db
}

var (
	db *gorm.DB
)

func InitDB() {
	cnf := conf.Config{}
	err := conf.LoadConf(&cnf, "../conf/config.yaml")
	if err != nil {
		panic(fmt.Sprintf("load config err : %v", err))
	}

	InitModel(cnf.LocalMysql)
	db.AutoMigrate(
		&Fund{},
		&Index{},
		&Tag{},
	)

}

func InitModel(config conf.ConfigMysql) {
	db = NewMysqlDB(config)
}

func NewMysqlDB(config conf.ConfigMysql) *gorm.DB {
	var (
		username = config.Username
		password = config.Password
		host     = config.Host
		port     = config.Port
		dbName   = config.DBName
		maxIdle  = config.MaxIdle
		maxOpen  = config.MaxConn
	)

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s&allowNativePasswords=true",
		username,
		password,
		host,
		port,
		dbName,
		url.QueryEscape("Asia/Shanghai"),
	)

	fmt.Println("Try to connect to MYSQL host: ", host, ", port: ", port)
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		panic(fmt.Sprintf("failed to connect MYSQL %s:%d/%s: %s", host, port, dbName, err.Error()))
	}
	fmt.Println("Connected to MYSQL: ", host, ", port: ", port)

	if !config.NotPrintSql {
		db.LogMode(true)
	}

	db.DB().SetMaxIdleConns(maxIdle)
	db.DB().SetMaxOpenConns(maxOpen)

	return db
}
