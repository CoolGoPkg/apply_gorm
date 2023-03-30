package model

import (
	"CoolGoPkg/apply_gorm/v2/conf"
	"CoolGoPkg/apply_gorm/v2/model"
	"CoolGoPkg/apply_gorm/v2/utils"
	"fmt"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	cnf := conf.Config{}
	err := conf.LoadConf(&cnf, "../conf/config.yaml")
	if err != nil {
		panic(fmt.Sprintf("load config err : %v", err))
	}
	DB = utils.CreateDB(cnf.LocalMysql)
	err = DB.AutoMigrate(
		&model.Stock{},
		&model.Fiance{},
	)
	if err != nil {
		panic(fmt.Sprintf("DB.AutoMigrate err : %v", err))
	}
}
