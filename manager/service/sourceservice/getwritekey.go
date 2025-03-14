package sourceservice

import (
	"github.com/rezaAmiri123/ormus/contract/go/source"
	"github.com/rezaAmiri123/ormus/pkg/richerror"
)

func (s Service) IsWriteKeyValid(req *source.ValidateWriteKeyReq) (*source.ValidateWriteKeyResp, error) {
	op := "sourceService.IsWriteKeyValid"
	resp, err := s.repo.IsWriteKeyValid(req.WriteKey)
	if err != nil {
		return nil, richerror.New(op).WithKind(richerror.KindUnexpected).WithWrappedError(err)
	}

	return resp, nil
}
