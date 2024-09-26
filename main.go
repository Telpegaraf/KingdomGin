package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"kingdom/config"
	"kingdom/database"
	"kingdom/router"
	"os"
)

var db *gorm.DB
var err error

func main() {
	err = godotenv.Load()
	conf := config.Get()

	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=" + os.Getenv("DB_SSLMODE")
	fmt.Println(dsn)
	db, err := database.New(dsn, conf.DefaultUser.Name, conf.DefaultUser.Pass, conf.PassStrength, true)
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

//func getUsers(c *gin.Context) {
//	var users []model.User
//	result := db.Find(&users)
//	if result.Error != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
//		return
//	}
//	c.IndentedJSON(http.StatusOK, users)
//}

//func createUser(c *gin.Context) {
//	var user model.User
//	c.BindJSON(&user)
//	db.Create(&user)
//	c.IndentedJSON(http.StatusCreated, user)
//}

//func getUserByID(c *gin.Context) {
//	var user model.User
//	result := db.First(&user, c.Param("id"))
//	if result.Error != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
//		return
//	}
//	c.IndentedJSON(http.StatusOK, user)
//}
