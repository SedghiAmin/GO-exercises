// cmd/main.go.go
package main

import (
	"github.com/SedghiAmin/GO-exercises/user-service-go/internal/config"
	"github.com/SedghiAmin/GO-exercises/user-service-go/internal/database"
	"github.com/SedghiAmin/GO-exercises/user-service-go/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	database.ConnectDB(cfg)

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
