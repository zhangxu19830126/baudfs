package raftstore

import (
	"github.com/tiglabs/raft/proto"
)

// Constants for network port definition.
const (
	DefaultHeartbeatPort = 5901
	DefaultReplicatePort = 5902
	DefaultRetainLogs    = 20000
)

// Config defined necessary configuration properties for raft store.
type Config struct {
	NodeID        uint64 // Identity of raft server instance.
	WalPath       string // Path of WAL(Write after Log)
	IpAddr        string // IP address of node
	HeartbeatPort int
	ReplicatePort int
	RetainLogs    uint64 // // RetainLogs controls how many logs we leave after truncate. The default value is 20000.
}

type PeerAddress struct {
	proto.Peer
	Address       string
	HeartbeatPort int
	ReplicatePort int
}

// PartitionConfig defined necessary configuration properties for raft store partition.
type PartitionConfig struct {
	ID      uint64
	Applied uint64
	Leader  uint64
	Term    uint64
	Peers   []PeerAddress
	SM      PartitionFsm
}
