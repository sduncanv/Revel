package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	datos := map[string]interface{}{
		"titulo": "Página de inicio",
		"mensaje": "Hola desde Revel!",
		"companyName": "RevelApplication",
	}

	return c.Render(datos)
}
