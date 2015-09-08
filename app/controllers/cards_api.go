package controllers

import (
	"github.com/caneroj1/cardsAPI/app/models"
	"github.com/revel/revel"
)

// CardsAPI is the controller for all card-related api operations
type CardsAPI struct {
	*revel.Controller
}

// Index returns all of the cards
func (c CardsAPI) Index() revel.Result {
	return c.RenderJson(models.GetAllCards())
}

// Classic returns all of the classic cards
func (c CardsAPI) Classic() revel.Result {
	return c.RenderJson(models.GetAllClassicCards())
}

// Created returns all of the cards that were created by users
func (c CardsAPI) Created() revel.Result {
	return c.RenderJson(models.GetAllCreatedCards())
}

// Show returns the json for a specific card
func (c CardsAPI) Show(id int64) revel.Result {
	return c.RenderJson(models.GetCardByID(id))
}

// Create allows a card to be POSTed and created
func (c CardsAPI) Create(CardBody string, CardType, CardBlanks int) revel.Result {
	models.ValidateCard(c.Validation, c.Params.Form)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.Response.Status = 400
		return c.RenderJson(c.Validation.Errors)
	}

	var card models.Card
	c.Params.Bind(&card, "card")
	return c.RenderJson(models.SaveCard(card))
}
