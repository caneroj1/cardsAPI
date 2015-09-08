package binders

import (
	"github.com/caneroj1/cardsAPI/app/models"
	"github.com/revel/revel"
	"reflect"
)

func bindCard(params *revel.Params, name string, typ reflect.Type) reflect.Value {
	var card models.Card
	if name == "card" && typ == reflect.TypeOf(card) {
		params.Bind(&card.CardBody, "CardBody")
		params.Bind(&card.CardBlanks, "CardBlanks")
		params.Bind(&card.CardType, "CardType")

		return reflect.ValueOf(card)
	}
	return reflect.ValueOf(nil)
}

func unbindCard(output map[string]string, name string, val interface{}) {
	if name == "card" {
		card := val.(models.Card)
		output["CardBody"] = card.CardBody
		output["CardType"] = string(card.CardType)
		output["CardBlanks"] = string(card.CardBlanks)
	}
}

// CardBinder is used to bind and unbind query strings and structs related to the Card struct
var CardBinder = revel.Binder{
	Bind:   bindCard,
	Unbind: unbindCard,
}
