package ankiconnect

import (
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

const (
	ankiConnectTestUrl     = "http://ankiconnect.com:8765"
	ankiConnectTestVersion = 2
)

var (
	client = NewClient()
)

func TestClient_NewClient(t *testing.T) {
	c := NewClient()
	assert.NotNil(t, c)
	assert.NotNil(t, c.Decks)
	assert.NotNil(t, c.Notes)
}

func TestSetHTTPClient(t *testing.T) {
	httpClient := resty.New()
	c := NewClient().SetHTTPClient(httpClient)
	assert.Exactly(t, httpClient, c.httpClient)
}

func TestClient_SetURL(t *testing.T) {
	c := NewClient().SetURL(ankiConnectTestUrl)
	assert.Equal(t, ankiConnectTestUrl, c.Url)
}

func TestClient_SetVersion(t *testing.T) {
	c := NewClient().SetVersion(ankiConnectTestVersion)
	assert.Equal(t, ankiConnectTestVersion, c.Version)
}

func TestClient_SetDecksManager(t *testing.T) {
	dm := &decksManager{}
	c := NewClient().SetDecksManager(dm)
	assert.Exactly(t, dm, c.Decks)
}

func TestClient_SetNotesManager(t *testing.T) {
	nm := &notesManager{}
	c := NewClient().SetNotesManager(nm)
	assert.Exactly(t, nm, c.Notes)
}

func TestClient_SetSyncManager(t *testing.T) {
	sm := &syncManager{}
	c := NewClient().SetSyncManager(sm)
	assert.Exactly(t, sm, c.Sync)
}

func TestClient_Ping(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		c := NewClient()
		httpmock.ActivateNonDefault(c.httpClient.GetClient())
		defer httpmock.DeactivateAndReset()

		responder, err := httpmock.NewJsonResponder(http.StatusOK, "")
		assert.NoError(t, err)
		httpmock.RegisterResponder(http.MethodGet, ankiConnectUrl, responder)

		restErr := c.Ping()
		assert.Nil(t, restErr)
	})

	t.Run("failure", func(t *testing.T) {
		c := NewClient()
		httpmock.ActivateNonDefault(c.httpClient.GetClient())
		defer httpmock.DeactivateAndReset()

		restErr := c.Ping()
		assert.NotNil(t, restErr)
		assert.Equal(t, http.StatusServiceUnavailable, restErr.StatusCode)
		assert.Equal(t, ankiConnectPingErrMsg, restErr.Message)
	})
}
