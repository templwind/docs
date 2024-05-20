package index

import (
	"github.com/labstack/echo/v4"
	"github.com/templwind/docs/internal/svc"
	"github.com/templwind/docs/modules/frontend/index/views"
	"github.com/templwind/templwind"
)

type Controller struct {
	svcCtx *svc.ServiceContext
}

func NewController(svcCtx *svc.ServiceContext) *Controller {
	return &Controller{
		svcCtx: svcCtx,
	}
}

func (c *Controller) Index(ctx echo.Context) error {
	return templwind.Render(ctx, 200, views.Index(ctx.Request()))
}
