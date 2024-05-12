package model

// Workspace model Input Body
type WorkspaceInput struct {
	Name           string `json:"name" binding:"required"`
	WorkspaceEmail string `json:"workspace_email" binding:"email,required"`
	Address        string `json:"address" binding:"required"`
	Description    string `json:"description" binding:"required"`
	CreatedBy      string `json:"created_by" binding:"required"`
}

type WorkspaceInputDB struct {
	ID             string `bson:"_id"`
	Name           string `bson:"name"`
	WorkspaceEmail string `bson:"workspace_email"`
	Address        string `bson:"address"`
	Description    string `bson:"description"`
	CreatedBy      string `bson:"created_by"`
	CreatedAt      string `bson:"created_at"`
	UpdatedAt      string `bson:"updated_at"`
}

type CreateWorkspaceDataResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	WorkspaceEmail string `json:"workspace_email"`
	Address        string `json:"address"`
	Description    string `json:"description"`
}
