package models

import (
	"database/sql"
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
	Approved   bool
	CreatedBy  int64
	CreatedOn  time.Time
	ModifiedOn time.Time
}

// ValidateCard validates whether the card's values are appropriate
// when creating a new card.
func ValidateCard(validator *revel.Validation, params url.Values) {
	validator.Clear()

	id, err := strconv.ParseInt(params.Get("CreatorID"), 0, 0)
	if err != nil {
		validator.Error("The ID of the creator should be a number").Key("creatorID")
	}
	creatorID := int(id)

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
	validator.Min(creatorID, 0).Message("The creator ID must be greater than 0.")
}

// SaveCard saves a card to the database and sets its id
func SaveCard(card Card, creator int) Card {
	sql := "insert into %s (cardbody, cardtype, cardblanks, createdby) VALUES ($1, $2, $3, $4) returning id, createdon"
	var id int64

	err := database.QueryRow(sql, card.CardBody, card.CardType, card.CardBlanks, creator).Scan(&id, &card.CreatedOn)
	if err != nil {
		log.Fatal(err)
		return card
	}

	card.ID = id
	card.CreatedBy = int64(creator)
	return card
}

// getCardsWithWhereQuery returns all of the cards that match a certain WHERE query.
func getCardsWithWhereQuery(query string, params ...interface{}) []Card {
	var cards []Card
	var err error

	sqlQ := "select cardbody, cardtype, cardblanks, classic, id, createdon, rating, raters, approved, createdby from %s " + query
	var rows *sql.Rows
	if len(params) == 0 {
		rows = database.GetByQuery(sqlQ)
	} else {
		rows = database.GetByQuery(sqlQ, params...)
	}
	defer rows.Close()

	var card Card
	for rows.Next() {
		err = rows.Scan(&card.CardBody, &card.CardType,
			&card.CardBlanks, &card.Classic, &card.ID,
			&card.CreatedOn, &card.Rating, &card.Raters,
			&card.Approved, &card.CreatedBy)
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

// GetAllCreatedCards returns all of the cards in the db that were created by users and were approved.
// e.g. classic = false
func GetAllCreatedCards() []Card {
	query := "where classic = false and approved = true"
	cards := getCardsWithWhereQuery(query)
	return cards
}

// GetAllClassicCards returns all of the classic cards in the db
func GetAllClassicCards() []Card {
	query := "where classic = true"
	cards := getCardsWithWhereQuery(query)
	return cards
}

// GetAllNewCards returns all of the cards that were created by users and were not approved
func GetAllNewCards() []Card {
	query := "where classic = false and approved = false"
	cards := getCardsWithWhereQuery(query)
	return cards
}

// GetAllCards returns all of the cards in the db
func GetAllCards() []Card {
	query := "where 1 = 1"
	cards := getCardsWithWhereQuery(query)
	return cards
}

// GetCardByID gets the cards from the db by id
func GetCardByID(id int64) Card {
	query := "where id = $1"
	cards := getCardsWithWhereQuery(query, id)
	return cards[0]
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
