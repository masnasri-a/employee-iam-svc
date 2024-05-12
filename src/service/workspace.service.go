package service

import (
	"emp-iam/src/config"
	"emp-iam/src/model"
	"emp-iam/src/util"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description create client
// @Tags workspace
// @Accept json
// @Produce json
// @Param user body model.WorkspaceInput true "User details"
// @Success 200 {string} object
// @Router /workspace/create [post]
func CreateWorkspace(c *gin.Context) {
	var workspaceInput model.WorkspaceInput
	if err := c.ShouldBindJSON(&workspaceInput); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	client, err := config.ConnectMongo()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer client.Disconnect(c.Request.Context())

	var inputDB model.WorkspaceInputDB
	inputDB.ID = util.GenerateID()
	inputDB.Name = workspaceInput.Name
	inputDB.WorkspaceEmail = workspaceInput.WorkspaceEmail
	inputDB.Address = workspaceInput.Address
	inputDB.Description = workspaceInput.Description
	inputDB.CreatedBy = workspaceInput.CreatedBy
	inputDB.CreatedAt = util.GetTimeNow()
	inputDB.UpdatedAt = util.GetTimeNow()
	collection := client.Database("emp-iam").Collection("workspace")

	// check if email already exist
	isExist, err := util.CheckFieldDB(client, "workspace_email", inputDB.WorkspaceEmail, "workspace")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if isExist {
		c.JSON(400, gin.H{"error": "workspace email already exist"})
		return
	}

	_, err = collection.InsertOne(c.Request.Context(), inputDB)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    workspaceInput,
	})
}
