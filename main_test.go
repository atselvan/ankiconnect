package ankiconnect

import (
	"os"
	"testing"
  "encoding/json"
	"net/http"

	"github.com/jarcoal/httpmock"
	"github.com/privatesquare/bkst-go-utils/utils/fileutils"
	"github.com/stretchr/testify/assert"
)

const (
	testDataPath          = "data/test/"
	errorTestDataFileName = "error.json"
	jsonExt               = ".json"
  genericErrorJson      = `{
    "result": null,
    "error": "some error message"
}`
)

var (
	client = NewClient()
  errorResponse = Result[string]{}
)

func TestMain(m *testing.M) {
  json.Unmarshal([]byte(genericErrorJson), &errorResponse)

	httpmock.ActivateNonDefault(client.httpClient.GetClient())
	defer httpmock.DeactivateAndReset()

	os.Exit(m.Run())
}

func registerErrorResponse(t *testing.T) {
		responder, err := httpmock.NewJsonResponder(http.StatusOK, errorResponse)
		assert.NoError(t, err)

		httpmock.RegisterResponder(http.MethodPost, ankiConnectUrl, responder)
}
func loadTestData(t *testing.T, path string, out interface{}) {
	err := fileutils.ReadJsonFile(path, &out)
	assert.NoError(t, err)
}
