package grpc

import (
	"github.com/blink-io/x/session"
)

type sm = session.Manager

type Manager struct {
	*sm
}

func NewManager() *Manager {
	m := &Manager{
		sm: session.NewManager(),
	}
	return m
}
