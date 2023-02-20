package ankiconnect

import (
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/privatesquare/bkst-go-utils/utils/fileutils"
	"github.com/stretchr/testify/assert"
)

const (
	testDataPath          = "data/test/"
	errorTestDataFileName = "error.json"
	jsonExt               = ".json"
)

func TestMain(m *testing.M) {
	httpmock.ActivateNonDefault(client.httpClient.GetClient())
	defer httpmock.DeactivateAndReset()

	os.Exit(m.Run())
}

func loadTestData(t *testing.T, path string, out interface{}) {
	err := fileutils.ReadJsonFile(path, &out)
	assert.NoError(t, err)
}
