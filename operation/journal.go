package operation

import (
	"context"
	"fmt"
	"time"

	core "github.com/KLORQ/sandboxes-core"
)

// Journal persists operation transitions.
type Journal struct {
	store Store
	now   func() time.Time
}

// NewJournal creates a journal backed by store.
func NewJournal(store Store) *Journal {
	return &Journal{store: store, now: time.Now}
}

// Begin creates a pending operation.
func (j *Journal) Begin(ctx context.Context, op Operation) error {
	if op.State == "" {
		op.State = Pending
	}
	if err := op.Validate(); err != nil {
		return err
	}
	if op.CreatedAt.IsZero() {
		op.CreatedAt = j.now()
	}
	op.UpdatedAt = op.CreatedAt
	return j.store.Create(ctx, op)
}

// MarkRunning moves a pending operation to running.
func (j *Journal) MarkRunning(ctx context.Context, id string) error {
	return j.transition(ctx, id, Running, "")
}

// Complete moves a running operation to done.
func (j *Journal) Complete(ctx context.Context, id string) error {
	return j.transition(ctx, id, Done, "")
}

// Fail moves an operation to failed and records the error.
func (j *Journal) Fail(ctx context.Context, id, message string) error {
	return j.transition(ctx, id, Failed, message)
}

func (j *Journal) transition(ctx context.Context, id string, to State, message string) error {
	op, err := j.store.Get(ctx, id)
	if err != nil {
		return err
	}
	if err := ValidateTransition(op.State, to); err != nil {
		return fmt.Errorf("%w: %w", core.ErrConflict, err)
	}
	op.State = to
	op.Error = message
	op.UpdatedAt = j.now()
	return j.store.Update(ctx, op)
}
