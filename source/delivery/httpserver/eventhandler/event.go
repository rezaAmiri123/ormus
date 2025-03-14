package eventhandler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezaAmiri123/ormus/pkg/errmsg"
	"github.com/rezaAmiri123/ormus/pkg/httpmsg"
	"github.com/rezaAmiri123/ormus/pkg/httputil"
	"github.com/rezaAmiri123/ormus/source/params"
)

func (h Handler) NewEvent(ctx echo.Context) error {
	var req params.TrackEventRequest
	if err := ctx.Bind(&req); err != nil {
		return httputil.NewError(ctx, http.StatusBadRequest, errmsg.ErrBadRequest)
	}
	resp, err := h.eventSvc.CreateNewEvent(context.Background(), req)
	if err != nil {
		msg, code := httpmsg.Error(err)

		return ctx.JSON(code, echo.Map{
			"message": msg,
			"errors":  err,
		})
	}

	return ctx.JSON(http.StatusCreated, resp)
}
