package ankiconnect

import "github.com/privatesquare/bkst-go-utils/utils/errors"

const (
	ActionSync = "sync"
)

type (
	SyncManager interface {
		Trigger() *errors.RestErr
	}

	syncManager struct{
		Client *Client
	}
)

func (sm *syncManager) Trigger() *errors.RestErr {
	_, restErr := post[string, ParamsDefault](sm.Client, ActionSync, nil)
	if restErr != nil {
		return restErr
	}
	return nil
}
