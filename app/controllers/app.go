package controllers

import (
	"github.com/revel/revel"
)

// App is the controller for the index page
type App struct {
	*revel.Controller
}

// Index returns the index page.
func (c App) Index() revel.Result {
	return c.Render()
}
