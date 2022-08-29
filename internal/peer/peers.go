package peer

import (
	"context"
	"errors"
	"github.com/honeycombio/refinery/config"
)

// Peers holds the collection of peers for the cluster
type Peers interface {
	GetPeers() ([]string, error)

	RegisterUpdatedPeersCallback(callback func())

	Close(ctx context.Context) error
}

func NewPeers(ctx context.Context, c config.Config) (Peers, error) {
	t, err := c.GetPeerManagementType()

	if err != nil {
		return nil, err
	}

	switch t {
	case "file":
		return newFilePeers(c), nil
	case "redis":
		return newRedisPeers(ctx, c)
	case "member-list":
		return newMemberList(ctx, c)
	default:
		return nil, errors.New("invalid config option 'PeerManagement.Type'")
	}
}
