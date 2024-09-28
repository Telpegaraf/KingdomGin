package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"kingdom/config"
	"kingdom/database"
	"kingdom/router"
	"log"
	"os"
)

// Swagger
//
// @title Kingdom Api
// @version 0.1.0
// @description Api for Pet Project
// @schemes http https
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	conf := config.Get()

	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=" + os.Getenv("DB_SSLMODE")
	username := os.Getenv("ADMIN_USERNAME")
	password := os.Getenv("ADMIN_PASSWORD")
	email := os.Getenv("ADMIN_EMAIL")

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	db, err := database.New(dsn, username, password, &email, conf.PassStrength, true)
	if err != nil {
		panic(err)
	}

	engine, _ := router.Create(db, conf)

	err = engine.Run(":8080")
	if err != nil {
		panic(err)
	}

	//if err := runner.Run(engine, conf); err != nil {
	//	fmt.Println("Server error: ", err)
	//	os.Exit(1)
	//}
}
