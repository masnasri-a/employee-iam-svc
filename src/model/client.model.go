package model

// Create a struct that will be used to create a new client
type ClientInput struct {
	Email    string `json:"email" binding:"required,email"`
	FullName string `json:"full_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ClientInputDB struct {
	ID       string `bson:"_id"`
	Email    string `bson:"email"`
	FullName string `bson:"full_name"`
	Password string `bson:"password"`
}

type CreateClientDataResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type GetClientDataResponse struct {
	Message    string                   `json:"message"`
	StatusCode int                      `json:"status_code"`
	Result     CreateClientDataResponse `json:"result"`
}

// Login struct to be used to login a client
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginDataResponseDB struct {
	ID       string `bson:"_id"`
	Email    string `bson:"email"`
	FullName string `bson:"full_name"`
}
