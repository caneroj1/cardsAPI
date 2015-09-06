package tests

import "github.com/revel/revel/testing"

// CardsTest encapsulates testing functionaltiy for the cards controller
type CardsTest struct {
	testing.TestSuite
}

// Before utility function for any set up that occurs before a test.
func (t *CardsTest) Before() {
	println("Set up")
}

// TestThatIndexPageWorks checks to see if the response is json
// and that there is a response body
func (t *CardsTest) TestThatIndexPageWorks() {
	t.Get("/cards")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
	cards := getCardsFromResponse(t.ResponseBody)
	t.Assert(len(cards) > 0)
	t.AssertStatus(200)
}

// TestThatClassicPageWorks checks to see if the response is json
// and that there is a response body
func (t *CardsTest) TestThatClassicPageWorks() {
	t.Get("/cards/classic")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
	cards := getCardsFromResponse(t.ResponseBody)
	t.Assert(len(cards) > 0)

	for _, card := range cards {
		t.Assert(card.Classic)
	}

	t.AssertStatus(200)
}

// TestThatShowPageWorks checks to see if the response from the show
// page is a json card
func (t *CardsTest) TestThatShowPageWorks() {
	t.Get("/cards/1")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
	card := getCardFromResponse(t.ResponseBody)
	t.Assert(card.CardBody != "")
	t.AssertStatus(200)
}

// After utility function for any tear down that occurs after a test.
func (t *CardsTest) After() {
	println("Tear down")
}
