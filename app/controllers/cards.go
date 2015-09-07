package controllers

import (
	"encoding/json"
	"github.com/caneroj1/cardsAPI/app/models"
	"github.com/revel/revel"
)

// Cards is the controller that handles requests for non-api card operations
type Cards struct {
	*revel.Controller
}

// Index renders the index page for cards
func (c Cards) Index() revel.Result {
	if cards, err := json.Marshal(models.GetAllCards()); err == nil {
		jsonCards := string(cards[:])
		return c.Render(jsonCards)
	}

	return c.Render()
}
