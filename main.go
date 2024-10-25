package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"kingdom/config"
	consumer2 "kingdom/consumer"
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
//
//	@securityDefinitions.apiKey  JWT
//	@in                          header
//	@name                        Authorization
//	@description                 JWT security accessToken. Please add it in the format "Bearer {AccessToken}" to authorize your requests.
func main() {
	err := godotenv.Load()
	log.Println(err)

	consumer, err := consumer2.New(os.Getenv("RMQ_URL"))
	if err != nil {
		log.Fatal(err)
	}

	//userQueue := consumer.RMQConsumer{
	//	"UserQueue",
	//	os.Getenv("RMQ_URL"),
	//	consumer.Handler,
	//}
	//
	//forever := make(chan bool)
	//
	//go func() {
	//	go userQueue.Consume()
	//
	//	<-forever
	//}()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	conf := config.Get()

	dsn := "host=" + os.Getenv("POSTGRES_HOST") +
		" user=" + os.Getenv("POSTGRES_USER") +
		" password=" + os.Getenv("POSTGRES_PASSWORD") +
		" dbname=" + os.Getenv("POSTGRES_DB") +
		" port=" + os.Getenv("POSTGRES_PORT") +
		" sslmode=" + os.Getenv("POSTGRES_SSLMODE")
	username := os.Getenv("ADMIN_USERNAME")
	password := os.Getenv("ADMIN_PASSWORD")
	email := os.Getenv("ADMIN_EMAIL")

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	db, err := database.New(dsn, username, password, email, conf.PassStrength, true)
	if err != nil {
		panic(err)
	}

	engine, _ := router.Create(db, conf, consumer)

	err = engine.Run(":8080")
	if err != nil {
		panic(err)
	}

	//if err := runner.Run(engine, conf); err != nil {
	//	fmt.Println("Server error: ", err)
	//	os.Exit(1)
	//}
}
