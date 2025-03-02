package userservice

import (
	"github.com/rezaAmiri123/ormus/pkg/errmsg"
	"github.com/rezaAmiri123/ormus/pkg/richerror"
)

func (s Service) IsUserIDValid(email string) (bool, error) {
	const op = "userservice.IsUserIDValid"

	user, rErr := s.repo.GetUserByEmail(email)
	if rErr != nil {
		return false, richerror.New(op).WithWrappedError(rErr).WhitMessage(errmsg.ErrSomeThingWentWrong)
	}

	if user == nil {
		return false, nil
	}

	return true, nil
}
