//go:build all || race
// +build all race

package peer_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/honeycombio/refinery/config"
	"github.com/honeycombio/refinery/internal/peer"
	"github.com/stretchr/testify/require"
)

func TestNewPeersMemberList(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	cbCh := make(chan struct{}, 10)
	p0, err := peer.NewPeers(ctx, &config.MockConfig{
		GetListenAddrVal:       "127.0.0.1:8180",
		PeerManagementType:     "member-list",
		GetPeerListenAddrVal:   "127.0.0.1:8080",
		MemberListListenAddr:   "127.0.0.1:8519",
		MemberListKnownMembers: []string{"127.0.0.1"},
	})
	require.NoError(t, err)
	require.NotNil(t, p0)
	defer p0.Close(context.Background())

	// This is a race condition, Since `peer.NewPeers()` launches async go routines
	// for both `redis` and `member-list`. It is possible that a callback could occur
	// before we call `RegisterUpdatedPeersCallback()` to register the call back,
	// therefore, missing the initial callback.
	p0.RegisterUpdatedPeersCallback(func() {
		t.Logf("CB for p0 called")
		cbCh <- struct{}{}
	})

	p1, err := peer.NewPeers(ctx, &config.MockConfig{
		GetListenAddrVal:       "127.0.0.2:8181",
		PeerManagementType:     "member-list",
		GetPeerListenAddrVal:   "127.0.0.2:8081",
		MemberListListenAddr:   "127.0.0.2:8519",
		MemberListKnownMembers: []string{"127.0.0.1"},
	})
	require.NoError(t, err)
	require.NotNil(t, p1)
	defer p1.Close(context.Background())

	p1.RegisterUpdatedPeersCallback(func() {
		t.Logf("CB for p1 called")
	})

	p2, err := peer.NewPeers(ctx, &config.MockConfig{
		GetListenAddrVal:       "127.0.0.3:8182",
		PeerManagementType:     "member-list",
		GetPeerListenAddrVal:   "127.0.0.3:8082",
		MemberListListenAddr:   "127.0.0.3:8519",
		MemberListKnownMembers: []string{"127.0.0.1"},
	})
	require.NoError(t, err)
	require.NotNil(t, p2)

	p2.RegisterUpdatedPeersCallback(func() {
		t.Logf("CB for p2 called")
	})
	defer p2.Close(context.Background())

	// Wait for at least one callback
	<-cbCh

	for {
		peers, _ := p0.GetPeers()
		if len(peers) != 3 {
			select {
			case <-ctx.Done():
				require.NoError(t, fmt.Errorf("expected 3 peers got '%d' instead", len(peers)))
			case <-time.After(300 * time.Millisecond):
				continue
			}
		}
		break
	}

	peers, _ := p0.GetPeers()
	t.Logf("Peers: %s", peers)
}
