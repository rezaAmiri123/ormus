package sourceparam

import "github.com/rezaAmiri123/ormus/manager/entity"

type ShowRequest struct {
	UserID   string `json:"-"`
	SourceID string `json:"-" param:"SourceID"`
}

type ShowResponse struct {
	Source entity.Source `json:"source"`
}
