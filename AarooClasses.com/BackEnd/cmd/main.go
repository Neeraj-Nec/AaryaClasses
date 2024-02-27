package main

import (
	"arooClasses/pkg/userM"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Main struct {
	usrM *userM.UserManagement
}

func main() {
	var m Main
	viper.SetConfigName("config")
	viper.AddConfigPath("../../BackEnd")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}
	userRouter := gin.Default()

	userRouter.POST("/aarya/login", m.usrM.LoginHandler)
	userRouter.POST("aarya/signup", m.usrM.Register)
	userRouter.Run(":8090")
}
