package db

import (
	"boilerplate/config"
	"boilerplate/pkg/utils"
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/sirupsen/logrus"
)

type DatabaseList struct {
	DatabaseApp *sql.DB
}

func NewSQLConnection(conf *config.DatabaseAccount, log *logrus.Logger) *sql.DB {
	dbName := utils.GetDBNameFromDriverSource(conf.DriverSource)

	db, err := sql.Open(conf.DriverName, conf.DriverSource)
	if err != nil {
		log.Fatal("Failed to connect database " + dbName + ", err: " + err.Error())
	}
	log.Info("Connection Opened to Database " + dbName)
	return db
}
