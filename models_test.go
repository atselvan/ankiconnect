package ankiconnect

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestModelsManager_Create(t *testing.T) {
	newModel := Model{
		ModelName:     "newModelName",
		InOrderFields: []string{"Field1", "Field2", "Field3"},
		Css:           "Optional CSS with default to builtin css",
		IsCloze:       false,
		CardTemplates: []CardTemplate{
			{
				Name:  "My Card 1",
				Front: "Front html {{Field1}}",
				Back:  "Back html  {{Field2}}",
			},
		},
	}

	t.Run("success", func(t *testing.T) {
		defer httpmock.Reset()

		registerVerifiedPayload(t,
			loadTestPayload(t, ActionCreateModel),
			loadTestResult(t, ActionCreateModel))

		restErr := client.Models.Create(newModel)
		assert.Nil(t, restErr)
	})

	t.Run("error", func(t *testing.T) {
		defer httpmock.Reset()

		registerErrorResponse(t)

		restErr := client.Models.Create(newModel)
		assert.NotNil(t, restErr)
		assert.Equal(t, http.StatusBadRequest, restErr.StatusCode)
		assert.Equal(t, "some error message", restErr.Message)
	})
}

func TestModelsManager_GetAll(t *testing.T) {
	modelNamesPayload := []byte(`{
    "action": "modelNames",
    "version": 6
  }`)
	modelNamesResult := []byte(`{
    "result": ["Basic", "Basic (and reversed card)"],
    "error": null
  }`)

	t.Run("success", func(t *testing.T) {
		defer httpmock.Reset()

		registerVerifiedPayload(t,
			modelNamesPayload,
			modelNamesResult)

		names, restErr := client.Models.GetAll()
		assert.Nil(t, restErr)
		assert.Len(t, *names, 2)
	})

	t.Run("error", func(t *testing.T) {
		defer httpmock.Reset()

		registerErrorResponse(t)

		names, restErr := client.Models.GetAll()
		assert.Nil(t, names)
		assert.NotNil(t, restErr)
		assert.Equal(t, http.StatusBadRequest, restErr.StatusCode)
		assert.Equal(t, "some error message", restErr.Message)
	})

}

func TestModelsManager_GetFields(t *testing.T) {
	modelFieldsPayload := []byte(`{
    "action": "modelFieldNames",
    "version": 6,
    "params": {
        "modelName": "Basic"
    }
  }`)
	modelFieldsResult := []byte(`{
    "result": ["Front", "Back"],
    "error": null
  }`)

	t.Run("success", func(t *testing.T) {
		defer httpmock.Reset()

		registerVerifiedPayload(t,
			modelFieldsPayload,
			modelFieldsResult)

		fields, restErr := client.Models.GetFields("Basic")
		assert.Nil(t, restErr)
		assert.Len(t, *fields, 2)
	})

	t.Run("error", func(t *testing.T) {
		defer httpmock.Reset()

		registerErrorResponse(t)

		fields, restErr := client.Models.GetFields("Basic")
		assert.Nil(t, fields)
		assert.NotNil(t, restErr)
		assert.Equal(t, http.StatusBadRequest, restErr.StatusCode)
		assert.Equal(t, "some error message", restErr.Message)
	})

}
