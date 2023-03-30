package conf

import (
	"fmt"

	"github.com/jinzhu/configor"
)

type Config struct {
	LocalMysql ConfigMysql `yaml:"local_mysql"`
}

// ConfigMysql sets the MySQL
type ConfigMysql struct {
	Username       string `yaml:"username"`
	Password       string `yaml:"password"`
	Host           string `yaml:"host"`
	Port           int    `yaml:"port"`
	DBName         string `yaml:"db_name"`
	MaxIdle        int    `yaml:"max_idle"`
	MaxConn        int    `yaml:"max_conn"`
	LogType        string `yaml:"log_type"`
	NotPrintSql    bool   `yaml:"not_print_sql"`
	NotCreateTable bool   `yaml:"not_create_table"`
	AutoMerge      bool   `yaml:"auto_merge"`
	Charset        string `yaml:"charset"`
}

func LoadConf(dest interface{}, path string) error {
	if err := configor.Load(dest, path); err != nil {
		panic(fmt.Sprintf("failed to load local config file: %v", err))
	}

	return nil
}
