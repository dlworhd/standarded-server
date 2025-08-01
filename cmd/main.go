package main

import (
	"github.com/dlworhd/standarded/handler"
	"github.com/dlworhd/standarded/model"
	"github.com/dlworhd/standarded/util"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL 드라이버
)

func main() {
	godotenv.Load(".env.local")

	postgres := model.PostgreSQL{}
	db, err := postgres.Connect()
	util.ErrorHandler(err)
	defer db.Close()

	router := gin.Default()

	profileGroup := router.Group("/profiles")
	{
		profile := handler.ProfileHandler{Repository: db}
		profileGroup.GET("/:id", profile.ReadHandler)
		profileGroup.GET("", profile.ReadAllHandler)
	}

	router.Run(":8080")
}
