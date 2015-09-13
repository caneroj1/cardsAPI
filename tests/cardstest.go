package tests

import "github.com/revel/revel/testing"

var contentType = "application/json; charset=utf-8"

// CardsTest encapsulates testing functionaltiy for the cards controller
type CardsTest struct {
	testing.TestSuite
}

// Before utility function for any set up that occurs before a test.
func (t *CardsTest) Before() {
	println("Set up")
}

// TestThatIndexRouteWorks checks to see if the response is json
// and that there is a response body
func (t *CardsTest) TestThatIndexRouteWorks() {
	t.Get("/cards")
	t.AssertOk()
	t.AssertContentType(contentType)
	cards := getCardsFromResponse(t.ResponseBody)
	t.Assert(len(cards) > 0)
	t.AssertStatus(200)
}

// TestThatClassicRouteWorks checks to see if the response is json
// and that there is a response body
func (t *CardsTest) TestThatClassicRouteWorks() {
	t.Get("/cards/classic")
	t.AssertOk()
	t.AssertContentType(contentType)
	cards := getCardsFromResponse(t.ResponseBody)
	t.Assert(len(cards) > 0)

	for _, card := range cards {
		t.Assert(card.Classic)
	}

	t.AssertStatus(200)
}

// TestThatCreatedRouteWorks checks to see if the response is json,
// that there is a response body, and that all the cards meet the criteria
// of a created card
func (t *CardsTest) TestThatCreatedRouteWorks() {
	t.Get("/cards/created")
	t.AssertOk()
	t.AssertContentType(contentType)
	cards := getCardsFromResponse(t.ResponseBody)
	if len(cards) > 0 {
		for _, card := range cards {
			t.Assert(card.Approved)
			t.Assert(!card.Classic)
		}
	}
}

// TestThatNewRouteWorks checks to see if the response is json,
// that there is a response body, and that all the cards meet the criteria
// of a new card
func (t *CardsTest) TestThatNewRouteWorks() {
	t.Get("/cards/new")
	t.AssertOk()
	t.AssertContentType(contentType)
	cards := getCardsFromResponse(t.ResponseBody)
	if len(cards) > 0 {
		for _, card := range cards {
			t.Assert(!card.Approved)
			t.Assert(!card.Classic)
		}
	}
}

// TestThatShowRouteWorks checks to see if the response from the show
// page is a json card
func (t *CardsTest) TestThatShowRouteWorks() {
	t.Get("/cards/1")
	t.AssertOk()
	t.AssertContentType(contentType)
	card := getCardFromResponse(t.ResponseBody)
	t.Assert(card.CardBody != "")
	t.AssertStatus(200)
}

// After utility function for any tear down that occurs after a test.
func (t *CardsTest) After() {
	println("Tear down")
}
