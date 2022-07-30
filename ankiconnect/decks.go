package ankiconnect

import (
	"github.com/privatesquare/bkst-go-utils/utils/errors"
)

const (
	ActionDeckNames    = "deckNames"
	ActionCreateDeck   = "createDeck"
	ActionGetDeckStats = "getDeckStats"
	ActionDeleteDecks = "deleteDecks"
)

type (
	DecksManager interface {
		GetAll() (*[]string, *errors.RestErr)
		Create(name string) *errors.RestErr
		Delete(name string) *errors.RestErr
	}

	ParamsCreateDeck struct {
		Deck string `json:"deck,omitempty"`
	}

	ParamsDeleteDeck struct {
		Decks *[]string `json:"decks,omitempty"`
		CardsToo bool `json:"cardsToo,omitempty"`
	}

	decksManager struct {
		Client *Client
	}
)

func (dm *decksManager) GetAll() (*[]string, *errors.RestErr) {
	result, restErr := post[[]string, ParamsDefault](dm.Client, ActionDeckNames, nil)
	if restErr != nil {
		return nil, restErr
	}
	return result, nil
}

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

func (dm *decksManager) Delete(name string) *errors.RestErr {
	params := ParamsDeleteDeck{
		Decks:    &[]string{name},
		CardsToo: true,
	}
	_, restErr := post[string](dm.Client, ActionDeleteDecks, &params)
	if restErr != nil {
		return restErr
	}
	return nil
}
