package ankiconnect

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestNotesManager_Add(t *testing.T) {
	createNoteStruct := Note{
		DeckName:  "test",
		ModelName: "Basic-a39a1",
		Fields: Fields{
			"Front": "front content",
			"Back":  "back content",
		},
		Options: &Options{
			AllowDuplicate: false,
			DuplicateScope: "deck",
			DuplicateScopeOptions: &DuplicateScopeOptions{
				DeckName:       "test",
				CheckChildren:  false,
				CheckAllModels: false,
			},
		},
		Tags: []string{
			"yomichan",
		},
		Audio: []Audio{
			{
				URL:      "https://assets.languagepod101.com/dictionary/japanese/audiomp3.php?kanji=猫&kana=ねこ",
				Filename: "yomichan_ねこ_猫.mp3",
				SkipHash: "7e2c2f954ef6051373ba916f000168dc",
				Fields: []string{
					"Front",
				},
			},
		},
		Video: []Video{
			{
				URL:      "https://cdn.videvo.net/videvo_files/video/free/2015-06/small_watermarked/Contador_Glam_preview.mp4",
				Filename: "countdown.mp4",
				SkipHash: "4117e8aab0d37534d9c8eac362388bbe",
				Fields: []string{
					"Back",
				},
			},
		},
		Picture: []Picture{
			{
				URL:      "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c7/A_black_cat_named_Tilly.jpg/220px-A_black_cat_named_Tilly.jpg",
				Filename: "black_cat.jpg",
				SkipHash: "8d6e4646dfae812bf39651b59d7429ce",
				Fields: []string{
					"Back",
				},
			},
		},
	}

	addNoteResult := []byte(`{
    "result": 1659294247478,
    "error": null
  }`)

	t.Run("success", func(t *testing.T) {
		defer httpmock.Reset()

		registerVerifiedPayload(t,
			loadTestPayload(t, ActionAddNote),
			addNoteResult)

		note := createNoteStruct
		restErr := client.Notes.Add(note)
		assert.Nil(t, restErr)
	})

	t.Run("error", func(t *testing.T) {
		defer httpmock.Reset()

		registerErrorResponse(t)

		note := createNoteStruct
		restErr := client.Notes.Add(note)
		assert.NotNil(t, restErr)
		assert.Equal(t, http.StatusBadRequest, restErr.StatusCode)
		assert.Equal(t, "some error message", restErr.Message)
	})
}

func TestNotesManager_Get(t *testing.T) {
	notesInfoPayload := []byte(`{
    "action": "notesInfo",
    "version": 6,
    "params": {
        "notes": [1502298033753]
    }
  }`)
	t.Run("success", func(t *testing.T) {
		defer httpmock.Reset()

		registerMultipleVerifiedPayloads(t,
			[][2][]byte{
				// Get will do two api calls, first findNotes to get the card id's
				{
					loadTestPayload(t, ActionFindNotes),
					loadTestResult(t, ActionFindNotes),
				},
				// Then notesInfo to transform those into actual anki cards
				{
					notesInfoPayload,
					loadTestResult(t, ActionNotesInfo),
				},
			})

		payload := "deck:current"
		notes, restErr := client.Notes.Get(payload)
		assert.Nil(t, restErr)
		assert.Equal(t, len(*notes), 1)

	})

	t.Run("errorFailSearch", func(t *testing.T) {
		defer httpmock.Reset()

		registerErrorResponse(t)

		_, restErr := client.Notes.Get("deck:current")
		assert.NotNil(t, restErr)
		assert.Equal(t, http.StatusBadRequest, restErr.StatusCode)
		assert.Equal(t, "some error message", restErr.Message)
	})
}

func TestNotesManager_Update(t *testing.T) {
	updateNoteStruct := UpdateNote{
		Id: 1514547547030,
		Fields: Fields{
			"Front": "new front content",
			"Back":  "new back content",
		},
		Audio: []Audio{
			{
				URL:      "https://assets.languagepod101.com/dictionary/japanese/audiomp3.php?kanji=猫&kana=ねこ",
				Filename: "yomichan_ねこ_猫.mp3",
				SkipHash: "7e2c2f954ef6051373ba916f000168dc",
				Fields: []string{
					"Front",
				},
			},
		},
	}

	t.Run("success", func(t *testing.T) {
		defer httpmock.Reset()

		registerVerifiedPayload(t,
			loadTestPayload(t, ActionUpdateNoteFields),
			genericSuccessJson)

		restErr := client.Notes.Update(updateNoteStruct)
		assert.Nil(t, restErr)

	})

}
