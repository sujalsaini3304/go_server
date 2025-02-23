package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sujalsaini3304/go_server/routes"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.ServeHTTP(w, r)
}
