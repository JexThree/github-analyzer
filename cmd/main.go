package main

import (
	"github-analyzer/internal/database"
	"github-analyzer/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	database.ConnectMongo()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
