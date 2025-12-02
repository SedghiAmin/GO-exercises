// internal/routes/routes.go
package routes

import (
	"github.com/SedghiAmin/GO-exercises/user-service-go/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.POST("/register", handlers.Register)
		// بعداً: login, profile, ...
	}
}
