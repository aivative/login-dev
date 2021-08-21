package main

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/aivative/login-dev/config"
	"github.com/aivative/login-dev/controller"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func initializeStruct(t reflect.Type, v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		ft := t.Field(i)
		switch ft.Type.Kind() {
		case reflect.Map:
			f.Set(reflect.MakeMap(ft.Type))
		case reflect.Slice:
			f.Set(reflect.MakeSlice(ft.Type, 0, 0))
		case reflect.Chan:
			f.Set(reflect.MakeChan(ft.Type, 0))
		case reflect.Struct:
			initializeStruct(ft.Type, f)
		case reflect.Ptr:
			fv := reflect.New(ft.Type.Elem())
			initializeStruct(ft.Type.Elem(), fv.Elem())
			f.Set(fv)
		default:
		}
	}
}

func AppendStruct(values ...interface{}) {

	for _, value := range values {
		vs := reflect.ValueOf(value)
		ts := vs.Type()
		for i := 0; i < vs.NumField(); i++ {
			fmt.Printf("%v --> %v\n", ts.Field(i).Name, vs.Field(i).Interface())
		}
	}
}

func init() {}

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
