package client

import (
	"github.com/tinyci/ci-agents/gen/client/uisvc/client/operations"
)

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// New creates a new *Client. Passing a cert will enable client/server
// certificate authentication; otherwise pass nil for no auth.
func New(baseURL, token string) (*operations.Client, error) {
	return operations.New(baseURL, token)
}
