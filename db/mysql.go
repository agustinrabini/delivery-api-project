package db

import (
	"database/sql"
	config "delivery-api-project/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitSQL() *sql.DB {

	conString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.ConfMap.MysqlUser, config.ConfMap.MysqlPassword, config.ConfMap.MysqlHost, config.ConfMap.MysqlPort, config.ConfMap.DataBase)
	fmt.Println("db conecttions stirng HEREEEE", conString)
	db, err := sql.Open("mysql", conString)
	if err != nil {
		panic(err.Error())
	}

	return db
}
