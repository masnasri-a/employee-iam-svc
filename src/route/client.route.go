package route

import (
	"emp-iam/src/service"

	"github.com/gin-gonic/gin"
)

func ClientRoute(r *gin.Engine) {
	api := r.Group("/api/v1/client")
	{
		api.POST("/create", service.CreateClient)
		api.POST("/login", service.GetClient)
	}
}
