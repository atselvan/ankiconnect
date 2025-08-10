package ankiconnect

import (
	"github.com/privatesquare/bkst-go-utils/utils/errors"
)

const (
	ActionRetrieveMedia = "retrieveMediaFile"
	ActionStoreMedia    = "storeMediaFile"
	ActionGetMediaNames = "getMediaFileNames"
	ActionDeleteMedia   = "deleteMediaFile"
)

type (
	// Media describes the interface that can be used to perform operations stored media.
	MediaManager interface {
		// Returns the contents of the file encoded in base64
		RetrieveMediaFile(filename string) (*string, *errors.RestErr)
		StoreMediaFile(filename string, encodedMediaContent string) (*string, *errors.RestErr)
		GetMediaFileNames(pattern string) (*[]string, *errors.RestErr)
		DeleteMediaFile(filename string) (*string, *errors.RestErr)
	}

	ParamsRetrieveMediaFile struct {
		Filename string `json:"filename,omitempty"`
	}

	ParamsStoreMediaFile struct {
		Filename string `json:"filename,omitempty"`
		Data     string `json:"data,omitempty"`
	}

	ParamsGetMediaFileNames struct {
		Pattern string `json:"pattern,omitempty"`
	}

	ParamsDeleteMediaFile struct {
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

// StoreMediaFile store media file to Anki storage.
// Method is expecting content of media file encoded in base64.
// The result is a name already stored media file.
// The method returns an error if:
//   - the api request to ankiconnect fails.
//   - the api returns a http error.
func (mm *mediaManager) StoreMediaFile(filename string, encodedMediaContent string) (*string, *errors.RestErr) {
	params := ParamsStoreMediaFile{
		Filename: filename,
		Data:     encodedMediaContent,
	}

	savedFileName, restErr := post[string](mm.Client, ActionStoreMedia, &params)
	return savedFileName, restErr
}

// GetMediaFileNames get array of media file names which match by pattern from Anki storage
// The result is array of media file names
// The method returns an error if:
//   - the api request to ankiconnect fails.
//   - the api returns a http error.
func (mm *mediaManager) GetMediaFileNames(pattern string) (*[]string, *errors.RestErr) {
	params := ParamsGetMediaFileNames{
		Pattern: pattern,
	}
	foundFileNames, restErr := post[[]string](mm.Client, ActionGetMediaNames, &params)
	if restErr != nil {
		return nil, restErr
	}
	return foundFileNames, nil
}

// DeleteMediaFile delete media file from Anki storage
// The result is deleted media file name
// The method returns an error if:
//   - the api request to ankiconnect fails.
func (mm *mediaManager) DeleteMediaFile(filename string) (*string, *errors.RestErr) {
	params := ParamsDeleteMediaFile{
		Filename: filename,
	}
	deletedFilename, restErr := post[string](mm.Client, ActionDeleteMedia, &params)
	if restErr != nil {
		return nil, restErr
	}
	return deletedFilename, nil
}
