package ankiconnect

import "github.com/privatesquare/bkst-go-utils/utils/errors"

const (
	ActionFindCards = "findCards"
	ActionCardsInfo = "cardsInfo"
)

type (
	// Notes manager describes the interface that can be used to perform operation on the notes in a deck.
	CardsManager interface {
		Search(query string) (*[]int64, *errors.RestErr)
		Get(query string) (*[]ResultCardsInfo, *errors.RestErr)
	}

	// notesManager implements NotesManager.
	cardsManager struct {
		Client *Client
	}

	ParamsFindCards struct {
		Query string `json:"query,omitempty"`
	}

	ResultCardsInfo struct {
		Answer     string               `json:"answer,omitempty"`
		Question   string               `json:"question,omitempty"`
		DeckName   string               `json:"deckName,omitempty"`
		ModelName  string               `json:"modelName,omitempty"`
		FieldOrder int64                `json:"fieldOrder,omitempty"`
		Fields     map[string]FieldData `json:"fields,omitempty"`
		Css        string               `json:"css,omitempty"`
		CardId     int64                `json:"cardId,omitempty"`
		Interval   int64                `json:"interval,omitempty"`
		Note       int64                `json:"note,omitempty"`
		Ord        int64                `json:"ord,omitempty"`
		Type       int64                `json:"type,omitempty"`
		Queue      int64                `json:"queue,omitempty"`
		Due        int64                `json:"due,omitempty"`
		Reps       int64                `json:"reps,omitempty"`
		Lapses     int64                `json:"lapses,omitempty"`
		Left       int64                `json:"left,omitempty"`
		Mod        int64                `json:"mod,omitempty"`
	}

	// ParamsCardsInfo represents the ankiconnect API params for getting card info.
	ParamsCardsInfo struct {
		Cards *[]int64 `json:"cards,omitempty"`
	}
)

func (cm *cardsManager) Search(query string) (*[]int64, *errors.RestErr) {
	findParams := ParamsFindCards{
		Query: query,
	}
	return post[[]int64](cm.Client, ActionFindCards, &findParams)
}

func (cm *cardsManager) Get(query string) (*[]ResultCardsInfo, *errors.RestErr) {
	cardIds, restErr := cm.Search(query)
	if restErr != nil {
		return nil, restErr
	}
	infoParams := ParamsCardsInfo{
		Cards: cardIds,
	}
	return post[[]ResultCardsInfo](cm.Client, ActionCardsInfo, &infoParams)
}
