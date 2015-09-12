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

// New renders the page where users can submit new cards
func (c Cards) New() revel.Result {
	return c.Render()
}

// Show renders the page where users can see a specific card
func (c Cards) Show(id int64) revel.Result {
	if card, err := json.Marshal(models.GetCardByID(id)); err == nil {
		jsonCard := string(card[:])
		return c.Render(jsonCard)
	}

	return c.Render()
}
