package models

import (
	"github.com/caneroj1/cardsAPI/app/database"
	"github.com/revel/revel"
	"log"
	"net/url"
	"strconv"
	"time"
)

// Card struct that maps cards in the db to Go struct
type Card struct {
	CardBody   string
	CardBlanks int
	CardType   int
	Classic    bool
	ID         int64
	Raters     int
	Rating     float32
	CreatedOn  time.Time
	ModifiedOn time.Time
}

// ValidateCard validates whether the card's values are appropriate
// when creating a new card.
func ValidateCard(validator *revel.Validation, params url.Values) {
	validator.Clear()

	cardBody := params.Get("CardBody")
	t, err := strconv.ParseInt(params.Get("CardType"), 0, 0)
	if err != nil {
		validator.Error("Card Type should be a number").Key("cardType")
	}
	cardType := int(t)

	b, err := strconv.ParseInt(params.Get("CardBlanks"), 0, 0)
	if err != nil {
		validator.Error("Card Blanks should be a number").Key("cardBlanks")
	}
	cardBlanks := int(b)

	validator.Required(cardBody)
	validator.Range(int(cardType), 0, 1).Message("The card type can only be 0 for a white card, or 1 for a black card.")
	validator.Range(int(cardBlanks), 0, 3).Message("Card blanks must be in the range of 0 - 3.")
	if cardType == 0 {
		validator.Max(int(cardBlanks), 0).Message("You cannot have blank spaces in a card unless it is a black card.)")
	}
}

// SaveCard saves a card to the database and sets its id
func SaveCard(card Card) Card {
	sql := "insert into %s (cardbody, cardtype, cardblanks) VALUES ($1, $2, $3) returning id, createdon"
	var id int64

	err := database.QueryRow(sql, card.CardBody, card.CardType, card.CardBlanks).Scan(&id, &card.CreatedOn)
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

	sql := "select cardbody, cardtype, cardblanks, classic, id, createdon, rating, raters from %s where classic = true"
	rows := database.GetByQuery(sql)
	defer rows.Close()

	var card Card
	for rows.Next() {
		err = rows.Scan(&card.CardBody, &card.CardType,
			&card.CardBlanks, &card.Classic, &card.ID,
			&card.CreatedOn, &card.Rating, &card.Raters)
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

	sql := "select cardbody, cardtype, cardblanks, classic, id, createdon, rating, raters from %s where classic = false"
	rows := database.GetByQuery(sql)
	defer rows.Close()

	var card Card
	for rows.Next() {
		err = rows.Scan(&card.CardBody, &card.CardType,
			&card.CardBlanks, &card.Classic, &card.ID,
			&card.CreatedOn, &card.Rating, &card.Raters)
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

	sql := "select cardbody, cardtype, cardblanks, classic, id, createdon, rating, raters from %s"
	rows := database.GetByQuery(sql)
	defer rows.Close()

	var card Card
	for rows.Next() {
		err = rows.Scan(&card.CardBody, &card.CardType,
			&card.CardBlanks, &card.Classic, &card.ID,
			&card.CreatedOn, &card.Rating, &card.Raters)
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

	sql := ("select cardbody, cardtype, cardblanks, classic, id, createdon, rating, raters from %s where id = $1")
	rows := database.GetByQuery(sql, id)
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&card.CardBody, &card.CardType,
			&card.CardBlanks, &card.Classic, &card.ID,
			&card.CreatedOn, &card.Rating, &card.Raters)

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

// RateCard changes the card's rating and returns the new rating.
// The modifiedon timestamo and the modifiedby field will be updated appropriately.
// TODO: once a user system is in place, add this rating to the user's list of rated cards.
func RateCard(newRating, ID, raterID int) float32 {
	sql := `update %s set
	rating = ((rating * raters + $1) / (raters + 1)),
	raters = raters + 1,
	modifiedon = current_timestamp,
	modifiedby = $2
	where
	id = $3
	returning rating`
	var cardRating float32
	if err := database.QueryRow(sql, newRating, raterID, ID).Scan(&cardRating); err != nil {
		log.Fatal(err)
		return -1
	}

	return cardRating
}
