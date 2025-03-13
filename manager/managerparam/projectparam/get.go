package projectparam

import "github.com/rezaAmiri123/ormus/manager/entity"

type GetRequest struct {
	UserID    string `json:"-"`
	ProjectID string `json:"-" param:"projectID"`
}

type GetResponse struct {
	Project entity.Project `json:"project"`
}
