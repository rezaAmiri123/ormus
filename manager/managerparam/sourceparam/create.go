package sourceparam

import "github.com/rezaAmiri123/ormus/manager/entity"

type CreateRequest struct {
	UserID      string `json:"-"`
	ProjectID   string `json:"project_id"`
	Name        string `json:"name"  example:"test name"`
	Description string `json:"description"  example:"test description"`
}

type CreateResponse struct {
	Source entity.Source `json:"source"`
}
