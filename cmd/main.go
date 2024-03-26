package main

import (
	"main/pkg/api"

	"main/pkg/common/database/mongodb"
	"main/pkg/common/database/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
)


func main(){

	db, err := sqlc.OpenPostgresConnection()
	
	if err != nil {
		logrus.Fatalf("error in connection to psql: %s", err)
		return
	}
	psgreClient := sqlc.New(db)
	logrus.Info("The Postgres is connected to the server!")
	client, err := mongodb.Init("mongodb://localhost:27017")
	validate := validator.New()

	r := gin.Default()

	api.RegisterRouter(r,validate, psgreClient, client)

	err = r.Run()
	
	if err != nil {
		logrus.Fatal("Can't start the server on the port: 8080")
		return
	}

	logrus.Info("The server is up!")
}