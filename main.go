package main

import (
	"context"
	_ "embed"
	"log"

	"github.com/aivative/login-dev/config"
	"github.com/aivative/login-dev/controller"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//go:embed config.json
var configStr string

func init() {
	config.MongoConf = config.ParseMongoConfig(configStr)
	config.SVCConf = config.ParseServiceConfig(configStr)
	config.APIKeyConf = config.ParseAPIKeyConfig(configStr)
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
