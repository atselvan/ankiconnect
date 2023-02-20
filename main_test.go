package ankiconnect

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/privatesquare/bkst-go-utils/utils/fileutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testDataPath = "data/test/"
	jsonExt      = ".json"
)

var (
	client           = NewClient()
	errorResponse    = Result[string]{}
	genericErrorJson = []byte(`{
    "result": null,
    "error": "some error message"
}`)
	genericSuccessJson = []byte(`{
    "result": null,
    "error": null
}`)
)

func TestMain(m *testing.M) {

	httpmock.ActivateNonDefault(client.httpClient.GetClient())
	defer httpmock.DeactivateAndReset()

	os.Exit(m.Run())
}

func registerErrorResponse(t *testing.T) {
	json.Unmarshal(genericErrorJson, &errorResponse)
	responder, err := httpmock.NewJsonResponder(http.StatusOK, errorResponse)
	assert.NoError(t, err)

	httpmock.RegisterResponder(http.MethodPost, ankiConnectUrl, responder)
}
func loadTestData(t *testing.T, path string, out interface{}) {
	err := fileutils.ReadJsonFile(path, &out)
	assert.NoError(t, err)
}

func loadTestPayload(t *testing.T, ankiConnectAction string) []byte {
	bytes, err := fileutils.ReadFile(
		testDataPath + ankiConnectAction + "Payload.json")
	assert.NoError(t, err)
	return bytes
}

func loadTestResult(t *testing.T, ankiConnectAction string) []byte {
	bytes, err := fileutils.ReadFile(
		testDataPath + ankiConnectAction + "Result.json")
	assert.NoError(t, err)
	return bytes
}

func registerVerifiedPayload(t *testing.T, payloadJson []byte, responseJson []byte) {
	httpmock.RegisterResponder(http.MethodPost, ankiConnectUrl,
		func(req *http.Request) (*http.Response, error) {

			bodyBytes, err := io.ReadAll(req.Body)
			assert.NoError(t, err)

			require.JSONEq(t, string(payloadJson), string(bodyBytes))

			// We cannot use NewJsonResponse since that serializes an interface
			// Instead we just craft our own response with the right headers
			resp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBuffer(responseJson)),
				Header:     http.Header{"Content-Type": []string{"application/json"}},
			}

			return resp, nil
		},
	)
}
