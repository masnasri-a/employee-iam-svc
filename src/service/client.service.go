package service

import (
	"emp-iam/src/config"
	"emp-iam/src/model"
	"emp-iam/src/util"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// PingExample godoc
// @Summary Login
// @Schemes
// @Description Login Endpoint
// @Tags client
// @Accept json
// @Produce json
// @Param user body model.LoginInput true "Login Form"
// @Success 200 {string} Login
// @Router /client/login [post]
func GetClient(c *gin.Context) {
	var clientInput model.LoginInput
	err := c.ShouldBindJSON(&clientInput)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New(validator.WithRequiredStructEnabled())

	err = validate.Struct(clientInput)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	client, err := config.ConnectMongo()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer client.Disconnect(c.Request.Context())
	collection := client.Database("emp-iam").Collection("client-iam")
	var result model.LoginDataResponseDB
	err = collection.FindOne(c.Request.Context(), map[string]string{"email": clientInput.Email, "password": util.HashString(clientInput.Password)}).Decode(&result)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid email or password"})
		return
	}
	token, err := util.GenerateAccessToken(result, "client")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message":      "success",
		"access_token": token,
		"data":         result,
	})
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description create client
// @Tags client
// @Accept json
// @Produce json
// @Param user body model.ClientInput true "User details"
// @Success 200 {string} object
// @Router /client/create [post]
func CreateClient(c *gin.Context) {
	var clientInput model.ClientInput
	err := c.ShouldBindJSON(&clientInput)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New(validator.WithRequiredStructEnabled())

	err = validate.Struct(clientInput)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("clientInput", clientInput)
	client, err := config.ConnectMongo()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var inputDB model.ClientInputDB
	inputDB.ID = util.GenerateID()
	inputDB.Email = clientInput.Email
	inputDB.FullName = clientInput.FullName
	inputDB.Password = util.HashString(clientInput.Password)
	isExist, err := util.CheckFieldDB(client, "email", inputDB.Email, "client-iam")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if isExist {
		c.JSON(400, gin.H{"error": "email already exist"})
		return
	}
	defer client.Disconnect(c.Request.Context())
	collection := client.Database("emp-iam").Collection("client-iam")
	_, err = collection.InsertOne(c.Request.Context(), inputDB)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	result := model.CreateClientDataResponse{
		ID:       inputDB.ID,
		Email:    inputDB.Email,
		FullName: inputDB.FullName,
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    result,
	})
}
