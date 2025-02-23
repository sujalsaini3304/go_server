package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sujalsaini3304/go_server/routes"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	g := gin.Default()
	routes.RegisterRoutes(g)
	g.ServeHTTP(w, r)
}
