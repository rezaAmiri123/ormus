package eventvalidator

import (
	"context"

	"github.com/rezaAmiri123/ormus/logger"
	"github.com/rezaAmiri123/ormus/pkg/errmsg"
	"github.com/rezaAmiri123/ormus/pkg/richerror"
)

func (v Validator) ValidateWriteKey(ctx context.Context, writeKey string) (bool, error) {
	isValid, err := v.repo.IsWriteKeyValid(ctx, writeKey, v.config.WriteKeyRedisExpiration)
	if err != nil {
		logger.L().Error(err.Error())

		return false, richerror.New("source.service").WithMessage(errmsg.ErrSomeThingWentWrong).WhitKind(richerror.KindUnexpected).WithWrappedError(err)
	}
	if !isValid {
		return false, nil
	}

	return true, nil
}
