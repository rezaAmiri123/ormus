package userhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
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
	result := h.userValidator.ValidateRegisterRequest(Req)
	if result != nil {
		msg, code := httpmsg.Error(result.Err)

		return ctx.JSON(code, echo.Map{
			"message": msg,
			"errors":  result.Fields,
		})
	}
	// TODO: should we return service error? or should we only return bad request error?
	resp, err := h.userSvc.Register(Req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echomsg.DefaultMessage(errmsg.ErrBadRequest))
	}

	return ctx.JSON(http.StatusCreated, resp)
}
