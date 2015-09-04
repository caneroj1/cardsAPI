package controllers

import (
	// "github.com/caneroj1/cardsAPI/app/models"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}
