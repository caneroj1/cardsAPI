package models

import (
	"fmt"
	"github.com/caneroj1/cardsAPI/app/database"
	"github.com/caneroj1/hush"
	"github.com/revel/revel"
	"log"
)

// Card struct that maps cards in the db to Go struct
type Card struct {
	CardBody   string
	CardBlanks int
	CardType   int
	Classic    bool
	ID         int64
}

// ValidateCard validates whether the card's values are appropriate
// when creating a new card.
func ValidateCard(validator *revel.Validation, cardBody string, cardType, cardBlanks int) {
	validator.Required(cardBody)

	if cardType == 0 {
		validator.Max(cardBlanks, 0).Message("You cannot have blank spaces in a card unless it is a black card (type = 1)")
	}
}

// SaveCard saves a card to the database and sets its id
func SaveCard(card Card) Card {
	secrets := hush.Hushfile()
	table, ok := secrets.GetString("table")
	if !ok {
		panic("No table name exists")
	}

	sql := fmt.Sprintf("insert into %s (cardbody, cardtype, cardblanks) VALUES ($1, $2, $3) returning id", table)

	var id int64
	err := database.Database.QueryRow(sql, card.CardBody, card.CardType, card.CardBlanks).Scan(&id)
	if err != nil {
		log.Fatal(err)
		return card
	}

	card.ID = id
	return card
}

// GetAllCards returns all of the cards in the db
func GetAllCards() []Card {
	var cards []Card
	secrets := hush.Hushfile()
	table, ok := secrets.GetString("table")
	if !ok {
		panic("No table name exists")
	}

	sql := fmt.Sprintf("select * from %s", table)
	rows, err := database.Database.Query(sql)
	if err != nil {
		log.Fatal(err)
		return cards
	}
	defer rows.Close()

	var card Card
	for rows.Next() {
		err := rows.Scan(&card.CardBody, &card.CardType,
			&card.CardBlanks, &card.Classic, &card.ID)
		if err != nil {
			log.Fatal(err)
		}

		cards = append(cards, card)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return cards
}

// GetCardByID gets the cards from the db by id
func GetCardByID(id int64) Card {
	var card Card
	secrets := hush.Hushfile()
	table, ok := secrets.GetString("table")
	if !ok {
		panic("No table name exists")
	}
	sql := fmt.Sprintf("select * from %s where id = $1", table)

	rows, err := database.Database.Query(sql, id)
	if err != nil {
		log.Fatal(err)
		return card
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&card.CardBody, &card.CardType,
			&card.CardBlanks, &card.Classic, &card.ID)

		if err != nil {
			log.Fatal(err)
			return card
		}
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return card
}
