// Package sandbox defines the sandbox aggregate and persistence port.
package sandbox

import (
	"context"
	"time"
)

// Sandbox is the durable sandbox record.
type Sandbox struct {
	ID         string
	Name       string
	State      State
	Generation int64
	Image      string
	Workspace  string
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
