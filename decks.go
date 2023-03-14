package ankiconnect

import (
	"github.com/privatesquare/bkst-go-utils/utils/errors"
)

const (
	ActionDeckNames    = "deckNames"
	ActionCreateDeck   = "createDeck"
	ActionGetDeckStats = "getDeckStats"
	ActionDeleteDecks  = "deleteDecks"
)

type (
	// DecksManager describes the interface that can be used to perform operations on anki decks.
	DecksManager interface {
		GetAll() (*[]string, *errors.RestErr)
		Create(name string) *errors.RestErr
		Delete(name string) *errors.RestErr
	}

	// ParamsCreateDeck represents the ankiconnect API params required for creating a new deck.
	ParamsCreateDeck struct {
		Deck string `json:"deck,omitempty"`
	}

	// ParamsDeleteDeck represents the ankiconnect API params required for deleting one or more decks
	ParamsDeleteDecks struct {
		Decks    *[]string `json:"decks,omitempty"`
		CardsToo bool      `json:"cardsToo,omitempty"`
	}

	// decksManager implements DecksManager.
	decksManager struct {
		Client *Client
	}
)

// GetAll retrieves all the decks from Anki.
// The result is a slice of string with the names of the decks.
// The method returns an error if:
//   - the api request to ankiconnect fails.
//   - the api returns a http error.
func (dm *decksManager) GetAll() (*[]string, *errors.RestErr) {
	result, restErr := post[[]string, ParamsDefault](dm.Client, ActionDeckNames, nil)
	if restErr != nil {
		return nil, restErr
	}
	return result, nil
}

// Create creates a new deck in Anki.
// The method returns an error if:
//   - the api request to ankiconnect fails.
//   - the api returns a http error.
func (dm *decksManager) Create(name string) *errors.RestErr {
	params := ParamsCreateDeck{
		Deck: name,
	}
	_, restErr := post[int64](dm.Client, ActionCreateDeck, &params)
	if restErr != nil {
		return restErr
	}
	return nil
}

// Delete deletes a deck from Anki
// The method returns an error if:
//   - the api request to ankiconnect fails.
//   - the api returns a http error.
func (dm *decksManager) Delete(name string) *errors.RestErr {
	params := ParamsDeleteDecks{
		Decks:    &[]string{name},
		CardsToo: true,
	}
	_, restErr := post[string](dm.Client, ActionDeleteDecks, &params)
	if restErr != nil {
		return restErr
	}
	return nil
}
