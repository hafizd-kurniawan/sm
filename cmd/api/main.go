package main

import (
	"boilerplate/config"
	"boilerplate/internal/server"
	"boilerplate/pkg/infra/db"
	"boilerplate/pkg/infra/logger"
)

// @title Smart Device Management API
// @version 1.0
// @description This is a simple API for managing IoT devices and user access control for a technical test.
// @description All protected endpoints require a Bearer token for authorization.

// @contact.name Hafizd Kurniawan
// @contact.url https://github.com/hafizd-kurniawan
// @contact.email hafizdkurniawan@gmail.com

// @host localhost:3000
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and the JWT token.
func main() {

	//* ====================== Config ======================

	conf := config.InitConfig("local")

	//* ====================== Logger ======================

	//* Loggrus
	appLogger := logger.NewLogrusLogger(&conf.Logger.Logrus)

	//* Grafana Loki
	// if conf.Grafana.IsActive {
	// 	if conf.App.Env != "local" {
	// 		err := logger.InitLoki(conf, appLogger)
	// 		if err != nil {
	// 			appLogger.Errorf("Grafana Loki err: %s", err.Error())
	// 		}
	// 	}
	// }

	//* ====================== Connection DB ======================

	//var dbList db.MongoInstance

	var dbList db.DatabaseList
	dbList.DatabaseApp = db.NewSQLConnection(&conf.Connection.DatabaseApp, appLogger)
	//? Wab Fondasi Mongo DB

	//* ====================== Running Server ======================

	server.Run(conf, &dbList, appLogger)
}
