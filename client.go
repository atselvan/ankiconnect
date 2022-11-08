package ankiconnect

import (
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"github.com/privatesquare/bkst-go-utils/utils/httputils"
	"github.com/privatesquare/bkst-go-utils/utils/logger"
)

const (
	ankiConnectUrl     = "http://localhost:8765"
	ankiConnectVersion = 6

	ankiConnectPingErrMsg = "AnkiConnect api is not accessible. Check if anki is running and the ankiconnect add-on is installed correctly"
)

type (
	// Client represents the anki connect api client.
	Client struct {
		Url        string
		Version    int
		httpClient *resty.Client

		// supported interfaces
		Decks DecksManager
		Notes NotesManager
		Sync  SyncManager
		Cards CardsManager
	}

	// RequestPayload represents the request payload for anki connect api.
	// [P any] represents a generic type that can accept any type for the Params field.
	RequestPayload[P any] struct {
		Action  string `json:"action,omitempty"`
		Version int    `json:"version,omitempty"`
		Params  *P     `json:"params,omitempty"`
	}

	// ParamsDefault represent the default parameters for a action to be executed by the anki connect api.
	ParamsDefault struct{}

	// Result is the result returned by anki connect api.
	// The Result filed will return different typed values depending on the action hence it is defined as a generic.
	// [T any] represents a generic type that can accept any type for the Result field.
	Result[T any] struct {
		Result T      `json:"result,omitempty"`
		Error  string `json:"error,omitempty"`
	}
)

// NewClient returns a instance of Client with the default settings.
func NewClient() *Client {
	c := &Client{
		Url:        ankiConnectUrl,
		Version:    ankiConnectVersion,
		httpClient: resty.New(),
	}

	c.Decks = &decksManager{Client: c}
	c.Notes = &notesManager{Client: c}
	c.Sync = &syncManager{Client: c}
	c.Cards = &cardsManager{Client: c}

	return c
}

// SetHTTPClient can be used set a custom httpClient.
func (c *Client) SetHTTPClient(httpClient *resty.Client) *Client {
	c.httpClient = httpClient
	return c
}

// SetURL can be used to set a custom url for the ankiconnect api.
func (c *Client) SetURL(url string) *Client {
	c.Url = url
	return c
}

// SetVersion can be used to set a custom version for the ankiconnect api.
func (c *Client) SetVersion(version int) *Client {
	c.Version = version
	return c
}

// SetDecksManager can be used to set a custom DecksManager interface.
// This function is added for testing the DecksManager interface.
func (c *Client) SetDecksManager(dm DecksManager) *Client {
	c.Decks = dm
	return c
}

// SetNotesManager can be used to set a custom NotesManager interface.
// This function is added for testing the NotesManager interface.
func (c *Client) SetNotesManager(nm NotesManager) *Client {
	c.Notes = nm
	return c
}

// SetSyncManager can be used to set a custom SyncManager interface.
// This function is added for testing the SyncManager interface.
func (c *Client) SetSyncManager(sm SyncManager) *Client {
	c.Sync = sm
	return c
}

// request formats and returns a base http request that can be extended later.
// as part of this the baseUrl and the default headers are set in the http client.
func (c *Client) request() *resty.Request {
	c.httpClient.SetBaseURL(c.Url)
	c.httpClient.SetDisableWarn(true)

	return c.httpClient.R().SetHeader(httputils.ContentTypeHeaderKey, httputils.ApplicationJsonMIMEType).
		SetHeader(httputils.AcceptHeaderKey, httputils.ApplicationJsonMIMEType)
}

// Ping checks if the anki connect api is online and healthy.
// If there is no response from the anki connect api a error will be returned.
func (c *Client) Ping() *errors.RestErr {
	resp, err := c.request().Get("")
	logger.RestyDebugLogs(resp)
	if err != nil {
		return &errors.RestErr{
			Message:    ankiConnectPingErrMsg,
			StatusCode: http.StatusServiceUnavailable,
			Error:      http.StatusText(http.StatusServiceUnavailable),
		}
	}
	logger.Info("ping: " + string(resp.Body()))
	return nil
}

// post makes a POST request to the anki connect API with a payload of the action to be executed.
// The Params field in the RequestPayload struct and the Result filed of the Result struct can be different
// based on the action that needs to be executed. Hence post is defined with some generic types.
// R any - represents the type of the result that will be returned by the API.
// P any - represents the type of the params that will be sent along with the action to be executed to the API.
func post[R any, P any](c *Client, action string, params *P) (*R, *errors.RestErr) {
	payload := RequestPayload[P]{
		Action:  action,
		Version: c.Version,
		Params:  params,
	}
	result := new(Result[R])
	resp, err := c.request().SetBody(payload).SetResult(result).Post("")
	logger.RestyDebugLogs(resp)
	if err != nil {
		return nil, &errors.RestErr{
			Message:    http.StatusText(http.StatusInternalServerError),
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
		}
	}
	if result.Error != "" {
		return nil, &errors.RestErr{
			Message:    result.Error,
			StatusCode: http.StatusBadRequest,
			Error:      result.Error,
		}
	}
	return &result.Result, nil
}
