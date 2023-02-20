package ankiconnect

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestNotesManager_Add(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		defer httpmock.Reset()

		result := new(Result[int64])
		loadTestData(t, testDataPath+ActionAddNote+"Result"+jsonExt, result)
		responder, err := httpmock.NewJsonResponder(http.StatusOK, result)
		assert.NoError(t, err)

		httpmock.RegisterResponder(http.MethodPost, ankiConnectUrl, responder)

		note := new(Note)
		// This doesn't actually do anything
		loadTestData(t, testDataPath+ActionAddNote+"Payload"+jsonExt, note)
		restErr := client.Notes.Add(*note)
		assert.Nil(t, restErr)
	})

	t.Run("error", func(t *testing.T) {
		defer httpmock.Reset()

		result := new(Result[string])
		loadTestData(t, testDataPath+errorTestDataFileName, result)
		responder, err := httpmock.NewJsonResponder(http.StatusOK, result)
		assert.NoError(t, err)

		httpmock.RegisterResponder(http.MethodPost, ankiConnectUrl, responder)

		note := new(Note)
		loadTestData(t, testDataPath+ActionAddNote+"Payload"+jsonExt, result)
		restErr := client.Notes.Add(*note)
		assert.NotNil(t, restErr)
		assert.Equal(t, http.StatusBadRequest, restErr.StatusCode)
		assert.Equal(t, "some error message", restErr.Message)
	})
}

func TestNotesManager_Get(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		defer httpmock.Reset()

		// Get will do two api calls, first findNotes to get the note id's
		findResult := new(Result[[]int64])
		loadTestData(t, testDataPath+ActionFindNotes+"Result"+jsonExt, findResult)
		findResponse, err := httpmock.NewJsonResponse(http.StatusOK, findResult)
		assert.NoError(t, err)

		// Then notesInfo to transform those into actual anki notes
		infoResult := new(Result[[]ResultNotesInfo])
		loadTestData(t, testDataPath+ActionNotesInfo+"Result"+jsonExt, infoResult)
		assert.Equal(t, infoResult.Result[0].ModelName, "Basic")
		infoResponse, err := httpmock.NewJsonResponse(http.StatusOK, infoResult)
		assert.NoError(t, err)

		responder := httpmock.ResponderFromMultipleResponses(
			[]*http.Response{
				findResponse,
				infoResponse,
			},
		)

		httpmock.RegisterResponder(http.MethodPost, ankiConnectUrl, responder)

		// If possible test to make sure the payload is properly transformed into
		// findNotesPayload.json (which seems to be attempted in above tests but is not
		// actually working)
		payload := "deck:current"
		notes, restErr := client.Notes.Get(payload)
		assert.Nil(t, restErr)
		assert.Equal(t, len(*notes), 1)

	})

	t.Run("errorFailSearch", func(t *testing.T) {
		defer httpmock.Reset()

		result := new(Result[string])
		loadTestData(t, testDataPath+errorTestDataFileName, result)
		responder, err := httpmock.NewJsonResponder(http.StatusOK, result)
		assert.NoError(t, err)

		httpmock.RegisterResponder(http.MethodPost, ankiConnectUrl, responder)

		_, restErr := client.Notes.Get("deck:current")
		assert.NotNil(t, restErr)
		assert.Equal(t, http.StatusBadRequest, restErr.StatusCode)
		assert.Equal(t, "some error message", restErr.Message)
	})
}

func TestNotesManager_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		defer httpmock.Reset()

		result := new(Result[int64])
		loadTestData(t, testDataPath+ActionUpdateNoteFields+"Result"+jsonExt, result)
		responder, err := httpmock.NewJsonResponder(http.StatusOK, result)
		assert.NoError(t, err)

		httpmock.RegisterResponder(http.MethodPost, ankiConnectUrl, responder)

		noteRequest := UpdateNote{
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
		restErr := client.Notes.Update(noteRequest)
		assert.Nil(t, restErr)

	})

}
