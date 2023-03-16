package main

import (
	"fmt"
	"log"
	"net/http"
	"user_api/controller"
	"user_api/database"
	"user_api/model"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	initDatabase()
	serveApplication()
}

func loadDatabase() {
	fmt.Println("Connecting to database")
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
}

func initDatabase() {
	fmt.Println("Initializing admin user")
	admin := model.User{
		Username: "admin",
		Password: "admin",
		Admin:    true,
	}

	err := database.Database.FirstOrCreate(&admin).Error
	if err != nil {
		log.Fatal("Failed to create admin user")
	} else {
		fmt.Println("Successfully created admin user")
	}
}

func loadEnv() {
	fmt.Println("Loading environment variables")
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		fmt.Println("Successfully loaded environment variables")
	}
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Pong"})
}

func serveApplication() {
	router := gin.Default()

	v1 := router.Group("/api/v1")

	{
		v1.GET("ping", Pong)
		v1.GET("users", controller.FindAll)
		v1.POST("users", controller.Create)
		v1.PUT("users", controller.Update)
	}

	router.Run(":3000")
	fmt.Println("Server running on port 3000")
}
