package frontend

import (
	"github.com/labstack/echo/v4"
	"github.com/templwind/docs/internal/svc"
	"github.com/templwind/docs/modules/frontend/index"
)

func Routes(svcCtx *svc.ServiceContext, g *echo.Group) {
	// index
	index.Routes(svcCtx, g)
}
