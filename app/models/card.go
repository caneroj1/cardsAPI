package models

import (
	"github.com/caneroj1/cardsAPI/app/database"
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
	validator.Range(cardType, 0, 1).Message("The card type can only be 0 for a white card, or 1 for a black card.")
	validator.Min(cardBlanks, 0).Message("Card blank cannot be negative.")
	if cardType == 0 {
		validator.Max(cardBlanks, 0).Message("You cannot have blank spaces in a card unless it is a black card (type = 1)")
	}
}

// SaveCard saves a card to the database and sets its id
func SaveCard(card Card) Card {
	sql := "insert into %s (cardbody, cardtype, cardblanks) VALUES ($1, $2, $3) returning id"
	var id int64

	err := database.QueryRow(sql, card.CardBody, card.CardType, card.CardBlanks).Scan(&id)
	if err != nil {
		log.Fatal(err)
		return card
	}

	card.ID = id
	return card
}

// GetAllClassicCards returns all of the classic cards in the db
func GetAllClassicCards() []Card {
	var cards []Card
	var err error

	sql := "select * from %s where classic = true"
	rows := database.GetByQuery(sql)
	defer rows.Close()

	var card Card
	for rows.Next() {
		err = rows.Scan(&card.CardBody, &card.CardType,
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

// GetAllCreatedCards returns all of the cards in the db that were created by users.
// e.g. classic = false
func GetAllCreatedCards() []Card {
	var cards []Card
	var err error

	sql := "select * from %s where classic = false"
	rows := database.GetByQuery(sql)
	defer rows.Close()

	var card Card
	for rows.Next() {
		err = rows.Scan(&card.CardBody, &card.CardType,
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

// GetAllCards returns all of the cards in the db
func GetAllCards() []Card {
	var cards []Card
	var err error

	sql := "select * from %s"
	rows := database.GetByQuery(sql)
	defer rows.Close()

	var card Card
	for rows.Next() {
		err = rows.Scan(&card.CardBody, &card.CardType,
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
	var err error

	sql := ("select * from %s where id = $1")
	rows := database.GetByQuery(sql, id)
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&card.CardBody, &card.CardType,
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
