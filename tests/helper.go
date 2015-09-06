package tests

import (
	"encoding/json"
	"fmt"
	"github.com/caneroj1/cardsAPI/app/models"
)

func unmarshal(data []byte, val interface{}) {
	err := json.Unmarshal(data, val)
	if err != nil {
		fmt.Println(err)
	}
}

func getCardsFromResponse(bytes []byte) []models.Card {
	var cards []models.Card
	unmarshal(bytes, &cards)

	return cards
}

func getCardFromResponse(bytes []byte) models.Card {
	var card models.Card
	unmarshal(bytes, &card)

	return card
}
