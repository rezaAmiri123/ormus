package eventhandler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezaAmiri123/ormus/logger"
	"github.com/rezaAmiri123/ormus/pkg/errmsg"
	"github.com/rezaAmiri123/ormus/pkg/httpmsg"
	"github.com/rezaAmiri123/ormus/pkg/httputil"
	"github.com/rezaAmiri123/ormus/source/params"
)

func (h Handler) NewEvent(ctx echo.Context) error {
	var req []params.TrackEventRequest
	if err := ctx.Bind(&req); err != nil {
		logger.L().Error(err.Error())
		return httputil.NewError(ctx, http.StatusBadRequest, errmsg.ErrBadRequest)
	}
	resp, err := h.eventSvc.CreateNewEvent(context.Background(), req, ctx.Get("invalid_write_keys").([]string))
	if err != nil {
		msg, code := httpmsg.Error(err)
		logger.L().Error(err.Error())
		return ctx.JSON(code, echo.Map{
			"message": msg,
			"errors":  err,
		})
	}

	return ctx.JSON(http.StatusCreated, resp)
}
