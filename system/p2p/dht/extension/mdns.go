package extension

import (
	"context"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	discovery "github.com/libp2p/go-libp2p/p2p/discovery/mdns"
)

// MDNS mdns
type MDNS struct {
	Service discovery.Service
	notifee *discoveryNotifee
}

// 用于局域网内节点发现
type discoveryNotifee struct {
	PeerChan chan peer.AddrInfo
}

// HandlePeerFound is a interface to be called when new peer is found
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	n.PeerChan <- pi
}

// NewMDNS Initialize the MDNS service
func NewMDNS(ctx context.Context, peerhost host.Host, serviceTag string) (*MDNS, error) {
	//register with service so that we get notified about peer discovery
	notifee := &discoveryNotifee{}
	notifee.PeerChan = make(chan peer.AddrInfo, 1)
	ser := discovery.NewMdnsService(peerhost, serviceTag, notifee)
	err := ser.Start()
	if err != nil {
		log.Error("NewMDNS", "start mdns service err", err)
		return nil, err
	}
	mnds := new(MDNS)
	mnds.Service = ser
	mnds.notifee = notifee
	return mnds, nil
}

// PeerChan returns a peer channel
func (m *MDNS) PeerChan() chan peer.AddrInfo {
	return m.notifee.PeerChan
}
