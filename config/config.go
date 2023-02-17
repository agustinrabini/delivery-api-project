package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	MysqlUser     string `yaml:"delivery_api_mysql_username"`
	MysqlPassword string `yaml:"delivery_api_mysql_password"`
	MysqlHost     string `yaml:"delivery_api_mysql_host"`
	MysqlPort     string `yaml:"delivery_api_mysql_port"`
	DataBase      string `yaml:"delivery_api_mysql_database"`
}

var ConfMap Configuration

func LoadConfigs() Configuration {

	ConfMap.MysqlUser = os.Getenv("MYSQL_USER")
	ConfMap.MysqlPassword = os.Getenv("MYSQL_PASS")
	ConfMap.MysqlHost = os.Getenv("MYSQL_HOST")
	ConfMap.MysqlPort = os.Getenv("MYSQL_PORT")
	ConfMap.MysqlUser = os.Getenv("MYSQL_DATABASE")

	//sets default config values
	if ConfMap.MysqlUser == "" {

		yfile, err := ioutil.ReadFile("./config/config.yaml")
		if err != nil {
			log.Fatal(err)
		}

		err = yaml.Unmarshal(yfile, &ConfMap)
		if err != nil {
			log.Fatal("error loading default config values: ", err)
		}

		fmt.Println("Loading default configs: %v", ConfMap)

	}

	return ConfMap
}
