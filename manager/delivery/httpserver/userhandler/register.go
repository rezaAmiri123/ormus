package userhandler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezaAmiri123/ormus/manager/validator"
	"github.com/rezaAmiri123/ormus/param"
	"github.com/rezaAmiri123/ormus/pkg/echomsg"
	"github.com/rezaAmiri123/ormus/pkg/errmsg"
	"github.com/rezaAmiri123/ormus/pkg/httpmsg"
)

func (h Handler) RegisterUser(ctx echo.Context) error {
	var Req param.RegisterRequest
	if err := ctx.Bind(&Req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echomsg.DefaultMessage(errmsg.ErrBadRequest))
	}

	// TODO: should we return service error? or should we only return bad request error?
	resp, err := h.userSvc.Register(Req)

	var vErr *validator.Error
	if errors.As(err, &vErr) {
		msg, code := httpmsg.Error(vErr.Err)

		return ctx.JSON(code, echo.Map{
			"message": msg,
			"errors":  vErr.Fields,
		})
	}
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echomsg.DefaultMessage(errmsg.ErrBadRequest))
	}

	return ctx.JSON(http.StatusCreated, resp)
}
