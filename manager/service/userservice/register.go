package userservice

import (
	"github.com/rezaAmiri123/ormus/manager/entity"
	"github.com/rezaAmiri123/ormus/param"
	"github.com/rezaAmiri123/ormus/pkg/password"
	"github.com/rezaAmiri123/ormus/pkg/richerror"
)

func (s Service) Register(req param.RegisterRequest) (*param.RegisterResponse, error) {
	hashedPassword, err := password.HashPassword(req.Password)
	if err != nil {
		return nil, richerror.New("register.hash").WithWrappedError(err)
	}

	user := entity.User{
		DeletedAt: nil,
		Email:     req.Email,
		Password:  hashedPassword,
		IsActive:  false,
	}

	createdUser, err := s.repo.Register(user)
	if err != nil {
		return nil, richerror.New("register.repo").WithWrappedError(err)
	}

	// TODO: we have to trigger an event of registration in this phase of function
	// return create new user
	return &param.RegisterResponse{
		ID:    createdUser.ID,
		Email: createdUser.Email,
	}, nil
}
