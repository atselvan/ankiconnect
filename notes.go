package ankiconnect

import "github.com/privatesquare/bkst-go-utils/utils/errors"

const (
	ActionFindNotes        = "findNotes"
	ActionNotesInfo        = "notesInfo"
	ActionAddNote          = "addNote"
	ActionAddNotes         = "addNotes"
	ActionDeleteNotes      = "deleteNotes"
	ActionUpdateNoteFields = "updateNoteFields"
)

type (
	// Notes manager describes the interface that can be used to perform operation on the notes in a deck.
	NotesManager interface {
		Add(note Note) *errors.RestErr
		Get(query string) (*[]ResultNotesInfo, *errors.RestErr)
		Update(note UpdateNote) *errors.RestErr
	}

	// notesManager implements NotesManager.
	notesManager struct {
		Client *Client
	}

	// ParamsCreateNote represents the ankiconnect API params for creating a note.
	ParamsCreateNote struct {
		Note *Note `json:"note,omitempty"`
	}

	// ParamsCreateNote represents the ankiconnect API params for updating a note.
	ParamsUpdateNote struct {
		Note *UpdateNote `json:"note,omitempty"`
	}

	// ParamsGetNotes represents the ankiconnect API params for querying notes.
	ParamsFindNotes struct {
		Query string `json:"query,omitempty"`
	}

	// ResultFindNotes represents the return value of querying notes
	ResultNotesInfo struct {
		NoteId    int64                `json:"noteId,omitempty"`
		ModelName string               `json:"modelName,omitempty"`
		Fields    map[string]FieldData `json:"fields,omitempty"`
		Tags      []string             `json:"tags,omitempty"`
	}

	// ParamsNotesInfo represents the ankiconnect API params for getting note info.
	ParamsNotesInfo struct {
		Notes *[]int64 `json:"notes,omitempty"`
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

	UpdateNote struct {
		Id      int64     `json:"id,omitempty"`
		Fields  Fields    `json:"fields,omitempty"`
		Audio   []Audio   `json:"audio,omitempty"`
		Video   []Video   `json:"video,omitempty"`
		Picture []Picture `json:"picture,omitempty"`
	}

	// Fields represents the main fields for a Anki Note
	Fields map[string]string

	// FieldData represents the format of a field returned by ankiconnect
	FieldData struct {
		Value string `json:"value,omitempty"`
		Order int64  `json:"order,omitempty"`
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
		Data     string   `json:"data,omitempty"`
		Path     string   `json:"path,omitempty"`
		Filename string   `json:"filename,omitempty"`
		SkipHash string   `json:"skipHash,omitempty"`
		Fields   []string `json:"fields,omitempty"`
	}

	// Video can be used to add a video file to a Anki Note.
	Video struct {
		URL      string   `json:"url,omitempty"`
		Data     string   `json:"data,omitempty"`
		Path     string   `json:"path,omitempty"`
		Filename string   `json:"filename,omitempty"`
		SkipHash string   `json:"skipHash,omitempty"`
		Fields   []string `json:"fields,omitempty"`
	}

	// Picture can be used to add a picture to a Anki Note.
	Picture struct {
		URL      string   `json:"url,omitempty"`
		Data     string   `json:"data,omitempty"`
		Path     string   `json:"path,omitempty"`
		Filename string   `json:"filename,omitempty"`
		SkipHash string   `json:"skipHash,omitempty"`
		Fields   []string `json:"fields,omitempty"`
	}
)

// Add adds a new note in Anki.
// The method returns an error if:
//   - the api request to ankiconnect fails.
//   - the api returns a http error.
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

func (nm *notesManager) Get(query string) (*[]ResultNotesInfo, *errors.RestErr) {
	findParams := ParamsFindNotes{
		Query: query,
	}
	noteIds, restErr := post[[]int64](nm.Client, ActionFindNotes, &findParams)
	if restErr != nil {
		return nil, restErr
	}

	infoParams := ParamsNotesInfo{
		Notes: noteIds,
	}

	notes, restErr := post[[]ResultNotesInfo](nm.Client, ActionNotesInfo, &infoParams)
	if restErr != nil {
		return nil, restErr
	}
	return notes, nil
}

func (nm *notesManager) Update(note UpdateNote) *errors.RestErr {
	params := ParamsUpdateNote{
		Note: &note,
	}
	// The return of this should always be 'null' int64 may not be the best
	// type here
	_, restErr := post[int64](nm.Client, ActionUpdateNoteFields, &params)
	if restErr != nil {
		return restErr
	}

	return nil
}
