package docs

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/templwind/docs/internal/svc"
	"github.com/templwind/docs/modules/docs/index"
	"github.com/templwind/docs/modules/docs/layout"
	"github.com/templwind/templwind"
	"github.com/templwind/templwind/components/sidenav"
)

func Routes(svcCtx *svc.ServiceContext, g *echo.Group) {
	// index
	index.Routes(svcCtx, g)

	// load the menus
	g.GET("/menu/:section", loadDocsMenu(svcCtx))
}

func loadDocsMenu(svcCtx *svc.ServiceContext) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		section := ctx.Param("section")
		fmt.Println("loadDocsMenu", section)
		if menuItems, ok := layout.SideMenu[section]; ok {
			return templwind.Render(ctx, http.StatusOK, sidenav.New(menuItems))
		}
		return ctx.NoContent(http.StatusNotFound)
	}
}
