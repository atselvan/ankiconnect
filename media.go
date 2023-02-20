package ankiconnect

import (
	"github.com/privatesquare/bkst-go-utils/utils/errors"
)

const (
	ActionRetrieveMedia = "retrieveMediaFile"
	// TODO
	// storeMediaFile
	// getMediaFileNames
	// deleteMediaFile
)

type (
	// Media describes the interface that can be used to perform operations stored media.
	MediaManager interface {
		// Returns the contents of the file encoded in base64
		RetrieveMediaFile(filename string) (*string, *errors.RestErr)
	}

	ParamsRetrieveMediaFile struct {
		Filename string `json:"filename,omitempty"`
	}

	// mediaManager implements MediaManager.
	mediaManager struct {
		Client *Client
	}
)

// RetrieveMediaFile retrieve the contents of the named file from Anki.
// The result is a string with the base64-encoded contents.
// The method returns an error if:
//   - the api request to ankiconnect fails.
//   - the api returns a http error.
func (mm *mediaManager) RetrieveMediaFile(filename string) (*string, *errors.RestErr) {
	params := ParamsRetrieveMediaFile{
		Filename: filename,
	}
	result, restErr := post[string](mm.Client, ActionRetrieveMedia, &params)
	if restErr != nil {
		return nil, restErr
	}
	return result, nil
}
