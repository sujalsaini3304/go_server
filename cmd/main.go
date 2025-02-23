package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading env file data")
	}
	server_port := os.Getenv("PORT")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Author":  "Sujal Kumar Saini",
			"Message": "Go Server Started",
			"Status":  "O.K.",
		})
	})
	r.Run(server_port)
	
}
