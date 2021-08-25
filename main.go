package main

import (
	"context"
	"log"
	"os"

	"github.com/aivative/login-dev/config"
	"github.com/aivative/login-dev/controller"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)


func init() {
	configBytes, err := os.ReadFile("config.json")
	if err != nil {
		logrus.Fatalln("CAN'T READ CONFIG FILE")
	}

	config.MongoConf = config.ParseMongoConfig(string(configBytes))
	config.SVCConf = config.ParseServiceConfig(string(configBytes))
	config.APIKeyConf = config.ParseAPIKeyConfig(string(configBytes))
}

func main() {
	ctrl, err := controller.New(context.Background())
	if err != nil {
		logrus.Fatalln(err)
		return
	}

	r := gin.Default()
	r.POST("/login", ctrl.Login)
	r.GET("/revoketoken/:uid", ctrl.RevokeToken)
	r.POST("/refreshtoken", ctrl.RefreshToken)

	r.GET("/check", ctrl.Check)

	if err := r.Run(config.SVCConf["user-service"].Host + ":" + config.SVCConf["user-service"].Port); err != nil {
		log.Fatalln("Error running services")
		return
	}
}
