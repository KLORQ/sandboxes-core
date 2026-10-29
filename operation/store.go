package operation

import "context"

// Store is implemented by the SQLite operation journal.
type Store interface {
	Create(ctx context.Context, op Operation) error
	Get(ctx context.Context, id string) (Operation, error)
	ListBySandbox(ctx context.Context, sandboxID string) ([]Operation, error)
	Update(ctx context.Context, op Operation) error
}
