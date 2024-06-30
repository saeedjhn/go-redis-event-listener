package bootstrap

import (
	"github.com/saeedjhn/go-redis-pubsub-message-broker/internal/infrastructure/persistance/db/mysql"
	"log"
)

func NewMysqlConnection(config mysql.Config) mysql.DB {
	return mysql.New(config)
}

func CloseMysqlConnection(mysqlDB mysql.DB) {
	err := mysqlDB.Conn().Close()

	if err != nil {
		log.Fatalf("don`t close mysql connection: %s", err.Error())
	}
}
