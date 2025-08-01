package main

import (
	"time"

	"github.com/dlworhd/standarded/handler"
	"github.com/dlworhd/standarded/model"
	"github.com/dlworhd/standarded/util"
	"github.com/gin-contrib/cors"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                   // 허용할 Origin 리스트
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 허용할 HTTP 메서드
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 허용할 헤더
		AllowCredentials: true,                                                // 쿠키 전달 등 크리덴셜 허용 여부
		MaxAge:           12 * time.Hour,                                      // preflight 요청 캐시 시간
	}))

	profileGroup := router.Group("/profiles")
	{
		profile := handler.ProfileHandler{Repository: db}
		profileGroup.GET("/:id", profile.ReadHandler)
		profileGroup.GET("", profile.ReadAllHandler)
	}

	router.Run(":8080")
}
