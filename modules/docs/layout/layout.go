package layout

import (
	"github.com/templwind/templwind"
	"github.com/templwind/templwind/components/apprail"
	"github.com/templwind/templwind/components/link"
	"github.com/templwind/templwind/components/sidenav"
)

var RailMenu = apprail.WithMenuItems(
	link.WithProps(
		link.WithHXGet("/docs"),
		link.WithTitle("Docs"),
		link.WithIcon(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="currentColor" d="M16 15H9v-2h7m3-2H9V9h10m0-2H9V5h10m2-4H7c-1.11 0-2 .89-2 2v14a2 2 0 0 0 2 2h14c1.11 0 2-.89 2-2V3a2 2 0 0 0-2-2M3 5v16h16v2H3a2 2 0 0 1-2-2V5z"/></svg>`),
		link.WithHXTarget("#content"),
		link.WithXOnTrigger("load-docs-menu"),
	),
	link.WithProps(
		link.WithHXGet("/docs/components"),
		link.WithTitle("Components"),
		link.WithIcon(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="currentColor" d="M5 2a1 1 0 0 0-1-1a1 1 0 0 0-1 1v4H1v6h6V6H5zm4 14c0 1.3.84 2.4 2 2.82V23h2v-4.18c1.16-.41 2-1.51 2-2.82v-2H9zm-8 0c0 1.3.84 2.4 2 2.82V23h2v-4.18C6.16 18.4 7 17.3 7 16v-2H1zM21 6V2a1 1 0 0 0-1-1a1 1 0 0 0-1 1v4h-2v6h6V6zm-8-4a1 1 0 0 0-1-1a1 1 0 0 0-1 1v4H9v6h6V6h-2zm4 14c0 1.3.84 2.4 2 2.82V23h2v-4.18c1.16-.41 2-1.51 2-2.82v-2h-6z"/></svg>`),
		link.WithHXTarget("#content"),
		link.WithXOnTrigger("load-component-menu"),
	),
	link.WithProps(
		link.WithHXGet("/docs/utilties"),
		link.WithTitle("Utilities"),
		link.WithIcon(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="currentColor" d="m13.78 15.3l6 6l2.11-2.16l-6-6zm3.72-5.2c-.39 0-.81-.05-1.14-.19L4.97 21.25l-2.11-2.11l7.41-7.4L8.5 9.96l-.72.7l-1.45-1.41v2.86l-.7.7l-3.52-3.56l.7-.7h2.81l-1.4-1.41l3.56-3.56a2.976 2.976 0 0 1 4.22 0L9.89 5.74l1.41 1.4l-.71.71l1.79 1.78l1.82-1.88c-.14-.33-.2-.75-.2-1.12a3.49 3.49 0 0 1 3.5-3.52c.59 0 1.11.14 1.58.42L16.41 6.2l1.5 1.5l2.67-2.67c.28.47.42.97.42 1.6c0 1.92-1.55 3.47-3.5 3.47"/></svg>`),
		link.WithHXTarget("#content"),
		link.WithXOnTrigger("load-utilties-menu"),
	),
	link.WithProps(
		link.WithHXGet("/docs/demos"),
		link.WithTitle("Demos"),
		link.WithIcon(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="currentColor" d="M7.5 9.5C6.1 9.5 5 10.6 5 12s1.1 2.5 2.5 2.5S10 13.4 10 12S8.9 9.5 7.5 9.5m9 0c-1.4 0-2.5 1.1-2.5 2.5s1.1 2.5 2.5 2.5S19 13.4 19 12s-1.1-2.5-2.5-2.5M12 4c4.4 0 8 3.6 8 8s-3.6 8-8 8s-8-3.6-8-8s3.6-8 8-8m0-2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2m0 3c-1.4 0-2.5 1.1-2.5 2.5S10.6 10 12 10s2.5-1.1 2.5-2.5S13.4 5 12 5m1.5 2.5c0 .2-.1.4-.2.6l-.6-.6l.6-.6c.1.2.2.4.2.6m-.9-1.3l-.6.6l-.6-.6c.2-.1.4-.2.6-.2s.4.1.6.2m-1.9 1.9c-.1-.2-.2-.4-.2-.6s.1-.4.2-.6l.6.6zm.7.7l.6-.6l.6.6c-.2.1-.4.2-.6.2s-.4-.1-.6-.2M12 14c-1.4 0-2.5 1.1-2.5 2.5S10.6 19 12 19s2.5-1.1 2.5-2.5S13.4 14 12 14m1.5 2.5c0 .2-.1.4-.2.6l-.6-.6l.6-.6c.1.2.2.4.2.6m-.9-1.3l-.6.6l-.6-.6c.2-.1.4-.2.6-.2s.4.1.6.2m-1.9 1.9c-.1-.2-.2-.4-.2-.6s.1-.4.2-.6l.6.6zm.7.7l.6-.6l.6.6c-.2.1-.4.2-.6.2s-.4-.1-.6-.2"/></svg>`),
		link.WithHXTarget("#content"),
		link.WithXOnTrigger("load-demos-menu"),
	),
)

var SideMenu map[string]templwind.OptFunc[sidenav.Props] = map[string]templwind.OptFunc[sidenav.Props]{
	"docs": sidenav.WithSubmenu(
		link.WithProps(
			link.WithTitle("Docs"),
			link.WithSubmenu(
				link.WithProps(
					link.WithHXGet("/docs/getting-started"),
					link.WithTitle("Getting Started"),
					link.WithHXTarget("#content"),
				),
			),
		),
	),
	"components": sidenav.WithSubmenu(
		link.WithProps(
			link.WithTitle("Components"),
			link.WithSubmenu(
				link.WithProps(link.WithHXGet("/docs/components/accordian"), link.WithTitle("Accordian"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/appbar"), link.WithTitle("Appbar"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/apprail"), link.WithTitle("Apprail"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/appshell"), link.WithTitle("Appshell"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/autocomplete"), link.WithTitle("Autocomplete"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/avatar"), link.WithTitle("Avatar"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/card"), link.WithTitle("Card"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/drawer"), link.WithTitle("Drawer"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/dropdown"), link.WithTitle("Dropdown"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/filebutton"), link.WithTitle("Filebutton"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/filedropzone"), link.WithTitle("Filedropzone"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/indicator"), link.WithTitle("Indicator"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/inputchip"), link.WithTitle("Inputchip"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/link"), link.WithTitle("Link"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/listbox"), link.WithTitle("Listbox"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/navbar"), link.WithTitle("Navbar"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/paginator"), link.WithTitle("Paginator"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/progressbar"), link.WithTitle("Progressbar"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/progressradial"), link.WithTitle("Progressradial"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/radio"), link.WithTitle("Radio"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/rangeslider"), link.WithTitle("Rangeslider"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/ratings"), link.WithTitle("Ratings"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/sidenav"), link.WithTitle("Sidenav"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/slidetoggle"), link.WithTitle("Slidetoggle"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/stepper"), link.WithTitle("Stepper"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/tab"), link.WithTitle("Tab"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/table"), link.WithTitle("Table"), link.WithHXTarget("#content")),
				link.WithProps(link.WithHXGet("/docs/components/treeview"), link.WithTitle("Treeview"), link.WithHXTarget("#content")),
			),
		),
	),
	"utilities": sidenav.WithSubmenu(),
	"demo":      sidenav.WithSubmenu(),
}
