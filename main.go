package main

import (
	"go-ecommerce-api/database"
	"go-ecommerce-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	r := gin.Default()
	routes.ProductRoutes(r)
	r.Run()
}
