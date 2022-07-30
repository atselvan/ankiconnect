package ankiconnect

import (
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"github.com/privatesquare/bkst-go-utils/utils/httputils"
	"github.com/privatesquare/bkst-go-utils/utils/logger"
)

// TODO: Check if ankiconnect is online
// if not launch anki?

const (
	ankiConnectUrl = "http://localhost:8765"

	ankiConnectVersion  = 6
	ankiVersionCheckFmt = "AnkiConnect v.%d"

	ActionFindNotes   = "findNotes"
	ActionNotesInfo   = "notesInfo"
	ActionAddNote     = "addNote"
	ActionAddNotes    = "addNotes"
	ActionDeleteNotes = "deleteNotes"

	AnkiConnectPingErrMsg = "AnkiConnect api is not accessible. Check if anki is running and the ankiconnect add-on is installed correctly"
)

type (
	// Client represents the anki connect api client.
	Client struct {
		Url        string
		Version    int
		httpClient *resty.Client

		Deck DecksManager
	}

	// ClientOption are additional settings that can be passed to the Client.
	ClientOption func(*Client)

	// RequestPayload represents the request payload for anki connect api.
	RequestPayload[P any] struct {
		Action  string  `json:"action,omitempty"`
		Version int     `json:"version,omitempty"`
		Params  *P `json:"params,omitempty"`
	}

	// Params represent the parameters for a action to be executed by the anki connect api.
	ParamsDefault struct {}

	// TODO: Result type can be diff for diff requests. Use generics?
	// Result is the result returned by anki connect api.
	Result[T any] struct {
		Result T `json:"result,omitempty"`
		Error  string   `json:"error,omitempty"`
	}
)

// NewClient returns a instance of Client with the default settings.
// These settings can be optionally modified using ClientOptions that can be passed to NewClient.
func NewClient(opts ...ClientOption) *Client {
	c := &Client{
		Url:        ankiConnectUrl,
		Version:    ankiConnectVersion,
		httpClient: resty.New(),
	}

	c.Deck = &decksManager{Client: c}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// WithHTTPClient can be used as a ClientOption to pass a custom httpClient.
func WithHTTPClient(httpClient *resty.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// TODO: Add a ClientOption to update URL and Version

// request formats and returns a base http request that can be extended later.
// as part of this the baseUrl and the default headers are set in the http client.
func (c *Client) request() *resty.Request {
	c.httpClient.SetBaseURL(c.Url)
	c.httpClient.SetDisableWarn(true)

	return c.httpClient.R().SetHeader(httputils.ContentTypeHeaderKey, httputils.ApplicationJsonMIMEType).
		SetHeader(httputils.AcceptHeaderKey, httputils.ApplicationJsonMIMEType)
}

// pingAnkiConnect checks ff the anki connect api is online and healthy.
// If there is no response from the anki connect api a error is returned.
func (c *Client) ping() *errors.RestErr {
	resp, err := c.request().Get("")
	logger.RestyDebugLogs(resp)
	if err != nil {
		return &errors.RestErr{
			Message:    AnkiConnectPingErrMsg,
			StatusCode: http.StatusServiceUnavailable,
			Error:      http.StatusText(http.StatusServiceUnavailable),
		}
	}
	logger.Info("ping: " + string(resp.Body()))
	return nil
}

func post[R any, P any](c *Client, action string, params *P) (*R, *errors.RestErr) {
	payload := RequestPayload[P]{
		Action: action,
		Version: c.Version,
		Params: params,
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
