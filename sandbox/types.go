// Package sandbox defines the sandbox aggregate and persistence port.
package sandbox

import (
	"context"
	"time"
)

// NetworkRule is a CIDR or domain/protocol/port egress rule. Domain and CIDR
// are mutually exclusive.
type NetworkRule struct {
	Protocol string
	CIDR     string
	Port     uint16
	Domain   string
	RuleID   string
}

// Meta is non-lifecycle sandbox metadata.
type Meta struct {
	Agent         string
	Detached      bool
	PolicyProfile string
	LastUsedAt    time.Time
}

// NetworkPolicy is the create-time allow/deny policy snapshot.
type NetworkPolicy struct {
	Allow []NetworkRule
	Deny  []NetworkRule
}

// Port maps a host loopback address/port to a guest port.
type Port struct {
	Host     uint16
	Guest    uint16
	Protocol string
	HostIP   string
}

// Workspace is a host path mounted into the guest at the same absolute path.
type Workspace struct {
	Path     string
	ReadOnly bool
}

// Sandbox is the durable sandbox record.
type Sandbox struct {
	ID         string
	Name       string
	State      State
	Generation int64
	Image      string
	Workspace  string
	Workspaces []Workspace
	Ports      []Port
	Policy     *NetworkPolicy
	Meta       *Meta
	CPUs       uint32
	MemoryMiB  uint32
	StoppedAt  time.Time
	Error      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// ListOptions controls pagination.
type ListOptions struct {
	PageSize  int
	PageToken string
}

// Store is implemented by the SQLite repository.
type Store interface {
	Create(ctx context.Context, sb Sandbox) error
	Get(ctx context.Context, id string) (Sandbox, error)
	GetByName(ctx context.Context, name string) (Sandbox, error)
	List(ctx context.Context, opts ListOptions) ([]Sandbox, string, error)
	Update(ctx context.Context, sb Sandbox) error
	Delete(ctx context.Context, id string) error
}
