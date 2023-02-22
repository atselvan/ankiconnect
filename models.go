package ankiconnect

import (
	"encoding/json"
	"github.com/privatesquare/bkst-go-utils/utils/errors"
)

const (
	ActionModelNames      = "modelNames"
	ActionModelFieldNames = "modelFieldNames"
	ActionCreateModel     = "createModel"
)

type (
	// Models Manager is used for creating various card and note types within Anki
	ModelsManager interface {
		Create(model Model) *errors.RestErr
		GetAll() (*[]string, *errors.RestErr)
		GetFields(model string) (*[]string, *errors.RestErr)
	}

	// ParamsCreateModel is used for creating a new Note type to add a new card to an
	// existing Note type will require the implementation of updateModelTemplates
	Model struct {
		ModelName     string   `json:"modelName,omitempty"`
		InOrderFields []string `json:"inOrderFields,omitempty"`
		Css           string   `json:"css,omitempty"`
		// Will default to false
		IsCloze       bool           `json:"isCloze"`
		CardTemplates []CardTemplate `json:"cardTemplates,omitempty"`
	}

	// CardTemplate contains the actual fields that will determine the
	// front and back of the anki card
	CardTemplate struct {
		Name  string `json:"Name,omitempty"`
		Front string `json:"Front,omitempty"`
		Back  string `json:"Back,omitempty"`
	}

	// modelsManager implements ModelsManager.
	modelsManager struct {
		Client *Client
	}

	// ParamsModelNames represents the ankiconnect API params required for
	// querying the Model Names avaliable
	ParamsModelNames struct {
		ModelName string `json:"modelName"`
	}

	// ResultCreateModel represents the ankiconnect API result from
	// creating a new model (Note type)
	//
	// The example given has some empty arrays. Since we dont know the data
	// types, and its not currently important we leave them as []interface{}
	// We do not technically need this, and could replace it with a [interface{}]
	// in Create, but it may be used later
	ResultCreateModel struct {
		Sortf     int64         `json:"sortf"`
		Did       int64         `json:"did"`
		LatexPre  string        `json:"latexPre"`
		LatexPost string        `json:"latexPost"`
		Mod       int64         `json:"mod"`
		Usn       int64         `json:"usn"`
		Vers      []interface{} `json:"vers"`
		Type      int64         `json:"type"`
		Css       string        `json:"css"`
		Name      string        `json:"name"`
		Flds      []struct {
			Name   string        `json:"name"`
			Ord    int64         `json:"ord"`
			Sticky bool          `json:"sticky"`
			Rtl    bool          `json:"rtl"`
			Font   string        `json:"font"`
			Size   int64         `json:"size"`
			Media  []interface{} `json:"media"`
		} `json:"flds"`
		Tmpls []struct {
			Name  string      `json:"name"`
			Ord   int64       `json:"ord"`
			Qfmt  string      `json:"qfmt"`
			Afmt  string      `json:"afmt"`
			Did   interface{} `json:"did"`
			Bqfmt string      `json:"bqfmt"`
			Bafmt string      `json:"bafmt"`
		} `json:"tmpls"`
		Tags []interface{} `json:"tags"`
		// The api description describes the id as being a "number" (eg "23443")
		// Where as in practice anki seems to return an actual int (eg 23443).
		// json.Number handles both instances
		Id  json.Number     `json:"id"`
		Req [][]interface{} `json:"req"`
	}
)

// The method returns an error if:
//   - the api request to ankiconnect fails.
//   - the api returns a http error.
func (mm *modelsManager) Create(model Model) *errors.RestErr {
	// This thing returns a large complicated struct back, ignore it for now
	_, restErr := post[ResultCreateModel](mm.Client, ActionCreateModel, &model)
	if restErr != nil {
		return restErr
	}
	return nil
}

func (mm *modelsManager) GetAll() (*[]string, *errors.RestErr) {
	modelNames, restErr := post[[]string, ParamsDefault](mm.Client, ActionModelNames, nil)
	if restErr != nil {
		return nil, restErr
	}
	return modelNames, nil
}
func (mm *modelsManager) GetFields(model string) (*[]string, *errors.RestErr) {
	modelName := ParamsModelNames{
		ModelName: model,
	}
	modelFields, restErr := post[[]string](mm.Client, ActionModelFieldNames, &modelName)
	if restErr != nil {
		return nil, restErr
	}
	return modelFields, nil

}
