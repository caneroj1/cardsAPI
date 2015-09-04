package controllers

import (
	"github.com/caneroj1/cardsAPI/app/models"
	"github.com/revel/revel"
)

// Cards is the controller for all card-related operations
type Cards struct {
	*revel.Controller
}

// Index returns all of the cards
func (c Cards) Index() revel.Result {
	return c.RenderJson(models.GetAllCards())
}

// Show returns the json for a specific card
func (c Cards) Show(id int64) revel.Result {
	return c.RenderJson(models.GetCardByID(id))
}

// Create allows a card to be POSTed and created
func (c Cards) Create(cardBody string, cardType, cardBlanks int) revel.Result {
	models.ValidateCard(c.Validation, cardBody, cardType, cardBlanks)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		return c.RenderJson(c.Validation.Errors)
	}

	var card models.Card
	c.Params.Bind(&card, "card")
	return c.RenderJson(models.SaveCard(card))
}
