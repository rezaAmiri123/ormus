package source

import (
	"net/http"

	"github.com/rezaAmiri123/ormus/cli/api/types"
)

func (c Client) Disable(sourceID string) types.Request {
	return types.Request{
		Path:                  "sources/%s/disable",
		Method:                http.MethodGet,
		AuthorizationRequired: true,
		Header:                nil,
		URLParams: []any{
			sourceID,
		},
	}
}
