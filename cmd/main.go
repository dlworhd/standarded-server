package main

import (
	"github.com/dlworhd/standarded/model"
	"github.com/dlworhd/standarded/util"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL 드라이버
)

func main() {
	godotenv.Load(".env.local")

	postgres := model.PostgreSQL{}
	db, err := postgres.Connect()
	util.ErrorHandler(err)

	defer db.Close()
}
