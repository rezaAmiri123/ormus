package sourceparam

import "github.com/rezaAmiri123/ormus/manager/entity"

type RotateWriteKeyRequest struct {
	UserID   string `json:"-"`
	SourceID string `json:"-" param:"SourceID"`
}

type RotateWriteKeyResponse struct {
	Source entity.Source `json:"source"`
}
