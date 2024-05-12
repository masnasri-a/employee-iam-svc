package route

import (
	"emp-iam/src/service"

	"github.com/gin-gonic/gin"
)

func WorkspaceRoute(r *gin.Engine) {
	api := r.Group("/api/v1/workspace")
	{
		api.POST("/create", service.CreateWorkspace)
	}
}
