package config

import (
	"fmt"
	"os"
)

type Configuration struct {
	MysqlUser     string `mapstructure:"jopit_api_mysql_username"`
	MysqlPassword string `mapstructure:"jopit_api_mysql_password"`
	MysqlHost     string `mapstructure:"jopit_api_mysql_host"`
	MysqlPort     string `mapstructure:"jopit_api_mysql_port"`
	DataBase      string `mapstructure:"jopit_api_mysql_database"`
}

var ConfMap Configuration

func LoadConfigs() {

	fmt.Println("Loading configuration...")

	ConfMap.MysqlUser = os.Getenv("MYSQL_USER")
	ConfMap.MysqlPassword = os.Getenv("MYSQL_PASS")
	ConfMap.MysqlHost = os.Getenv("MYSQL_HOST")
	ConfMap.MysqlPort = os.Getenv("MYSQL_PORT")
	ConfMap.MysqlUser = os.Getenv("MYSQL_DATABASE")
}
