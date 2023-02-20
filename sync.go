package ankiconnect

import "github.com/privatesquare/bkst-go-utils/utils/errors"

const (
	ActionSync = "sync"
)

type (
	// SyncManager describes the interface that can be used to perform sync operations on Anki.
	SyncManager interface {
		Trigger() *errors.RestErr
	}

	// syncManager implements SyncManager
	syncManager struct {
		Client *Client
	}
)

// Trigger syncs local Anki data to Anki web.
// The method returns an error if:
//   - the api request to ankiconnect fails.
//   - the api returns a http error.
func (sm *syncManager) Trigger() *errors.RestErr {
	_, restErr := post[string, ParamsDefault](sm.Client, ActionSync, nil)
	if restErr != nil {
		return restErr
	}
	return nil
}
