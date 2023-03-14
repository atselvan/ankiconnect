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

// For each api call implemented, you test it as follows
//
//  1. Get the expected request and response from api description found at
//     https://github.com/FooSoft/anki-connect
//
//  2. If the json blob is tiny, insert it in the test as a multiline string
//     if its pretty large, it can be saved under
//     data/test/{apiName}{Result/Payload}.json
//
//  3. You can then call some of the below helper functions to properly mock
//     the http requests / responses
//
// 4) Do some assertions on the result of the api call
//
// The idea behind the this is that we want our tests to construct a go struct
// in the same way as the end user would using our defined structs. We then ensure
// that that go struct get correctly transformed into the json format expected by
// anki connect
func TestMain(m *testing.M) {

	httpmock.ActivateNonDefault(client.httpClient.GetClient())
	defer httpmock.DeactivateAndReset()

	os.Exit(m.Run())
}

// If you want your test to fail, call this before calling the api call
func registerErrorResponse(t *testing.T) {
	json.Unmarshal(genericErrorJson, &errorResponse)
	responder, err := httpmock.NewJsonResponder(http.StatusOK, errorResponse)
	assert.NoError(t, err)

	httpmock.RegisterResponder(http.MethodPost, ankiConnectUrl, responder)
}

// If you need to load a http request from a file, use this
func loadTestPayload(t *testing.T, ankiConnectAction string) []byte {
	bytes, err := fileutils.ReadFile(
		testDataPath + ankiConnectAction + "Payload.json")
	assert.NoError(t, err)
	return bytes
}

// If you need to load a http response from a file, use this
func loadTestResult(t *testing.T, ankiConnectAction string) []byte {
	bytes, err := fileutils.ReadFile(
		testDataPath + ankiConnectAction + "Result.json")
	assert.NoError(t, err)
	return bytes
}

// The loaded json strings then can be passed to this function, which will
// for each pair, verify that the client produces the correct json encoding
// and send the desired response
func registerMultipleVerifiedPayloads(t *testing.T, pairs [][2][]byte) {
	currentIndex := 0
	httpmock.RegisterResponder(http.MethodPost, ankiConnectUrl,
		func(req *http.Request) (*http.Response, error) {

			// If this assert fails the test is misconfigured
			assert.Less(t, currentIndex, len(pairs), "Responder called too many times")

			currentPair := pairs[currentIndex]
			payloadJson := currentPair[0]
			responseJson := currentPair[1]
			currentIndex += 1

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

// If you only are doing one Payload/Response, you can use this function
func registerVerifiedPayload(t *testing.T, payloadJson []byte, responseJson []byte) {
	registerMultipleVerifiedPayloads(t, [][2][]byte{{payloadJson, responseJson}})
}
