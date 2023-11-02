package shared

import (
	"crypto/tls"
	"time"
)

const (
	NoExpiration   = time.Duration(0)
	DefaultLockTTL = 60 * time.Second
)

// Config the Redis configuration.
type Config struct {
	TLS      *tls.Config
	Username string
	Password string
	DB       int
	Sentinel *Sentinel
}

// Sentinel holds the Redis Sentinel configuration.
type Sentinel struct {
	MasterName string
	Username   string
	Password   string

	// ClusterClient indicates whether to use the NewFailoverClusterClient to build the client.
	ClusterClient bool

	// Allows routing read-only commands to the closest master or replica node.
	// This option only works with NewFailoverClusterClient.
	RouteByLatency bool

	// Allows routing read-only commands to the random master or replica node.
	// This option only works with NewFailoverClusterClient.
	RouteRandomly bool

	// Route all commands to replica read-only nodes.
	ReplicaOnly bool

	// Use replicas disconnected with master when cannot get connected replicas
	// Now, this option only works in RandomReplicaAddr function.
	UseDisconnectedReplicas bool
}
