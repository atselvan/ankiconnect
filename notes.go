package ankiconnect

import "github.com/privatesquare/bkst-go-utils/utils/errors"

const (
	ActionFindNotes   = "findNotes"
	ActionNotesInfo   = "notesInfo"
	ActionAddNote     = "addNote"
	ActionAddNotes    = "addNotes"
	ActionDeleteNotes = "deleteNotes"
)

type (
	// Notes manager describes the interface that can be used to perform operation on the notes in a deck.
	NotesManager interface {
		Add(note Note) *errors.RestErr
	}

	// notesManager implements NotesManager.
	notesManager struct {
		Client *Client
	}

	// ParamsCreateNote represents the ankiconnect API params for creating a note.
	ParamsCreateNote struct {
		Note *Note `json:"note,omitempty"`
	}

	// Note represents a Anki Note.
	Note struct {
		DeckName  string    `json:"deckName,omitempty"`
		ModelName string    `json:"modelName,omitempty"`
		Fields    Fields    `json:"fields,omitempty"`
		Options   *Options  `json:"options,omitempty"`
		Tags      []string  `json:"tags,omitempty"`
		Audio     []Audio   `json:"audio,omitempty"`
		Video     []Video   `json:"video,omitempty"`
		Picture   []Picture `json:"picture,omitempty"`
	}

	// Fields represents the main fields for a Anki Note
	Fields struct {
		Front string `json:"Front,omitempty"`
		Back  string `json:"Back,omitempty"`
	}

	// Options represents note options.
	Options struct {
		AllowDuplicate        bool                   `json:"allowDuplicate,omitempty"`
		DuplicateScope        string                 `json:"duplicateScope,omitempty"`
		DuplicateScopeOptions *DuplicateScopeOptions `json:"duplicateScopeOptions,omitempty"`
	}

	// DuplicateScopeOptions represents the options that control the duplication of a Anki Note.
	DuplicateScopeOptions struct {
		DeckName       string `json:"deckName,omitempty"`
		CheckChildren  bool   `json:"checkChildren,omitempty"`
		CheckAllModels bool   `json:"checkAllModels,omitempty"`
	}

	// Audio can be used to add a audio file to a Anki Note.
	Audio struct {
		URL      string   `json:"url,omitempty"`
		Filename string   `json:"filename,omitempty"`
		SkipHash string   `json:"skipHash,omitempty"`
		Fields   []string `json:"fields,omitempty"`
	}

	// Video can be used to add a video file to a Anki Note.
	Video struct {
		URL      string   `json:"url,omitempty"`
		Filename string   `json:"filename,omitempty"`
		SkipHash string   `json:"skipHash,omitempty"`
		Fields   []string `json:"fields,omitempty"`
	}

	// Picture can be used to add a picture to a Anki Note.
	Picture struct {
		URL      string   `json:"url,omitempty"`
		Filename string   `json:"filename,omitempty"`
		SkipHash string   `json:"skipHash,omitempty"`
		Fields   []string `json:"fields,omitempty"`
	}
)

// Add adds a new note in Anki.
// The method returns an error if:
//	- the api request to ankiconnect fails.
//	- the api returns a http error.
func (nm *notesManager) Add(note Note) *errors.RestErr {
	params := ParamsCreateNote{
		Note: &note,
	}
	_, restErr := post[int64](nm.Client, ActionAddNote, &params)
	if restErr != nil {
		return restErr
	}
	return nil
}
