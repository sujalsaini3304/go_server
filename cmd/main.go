package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sujalsaini3304/go_server/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading env file data")
	}
	server_port := os.Getenv("PORT")
	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(server_port)
}

