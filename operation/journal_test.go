package operation

import (
	"context"
	"errors"
	"testing"

	core "github.com/KLORQ/sandboxes-core"
)

type fakeStore struct {
	items map[string]Operation
}

func newFakeStore() *fakeStore {
	return &fakeStore{items: make(map[string]Operation)}
}

func (s *fakeStore) Create(_ context.Context, op Operation) error {
	if _, ok := s.items[op.ID]; ok {
		return core.ErrConflict
	}
	s.items[op.ID] = op
	return nil
}

func (s *fakeStore) Get(_ context.Context, id string) (Operation, error) {
	op, ok := s.items[id]
	if !ok {
		return Operation{}, core.ErrNotFound
	}
	return op, nil
}

func (s *fakeStore) ListBySandbox(_ context.Context, sandboxID string) ([]Operation, error) {
	var out []Operation
	for _, op := range s.items {
		if op.SandboxID == sandboxID {
			out = append(out, op)
		}
	}
	return out, nil
}

func (s *fakeStore) Update(_ context.Context, op Operation) error {
	if _, ok := s.items[op.ID]; !ok {
		return core.ErrNotFound
	}
	s.items[op.ID] = op
	return nil
}

func TestJournalLifecycle(t *testing.T) {
	ctx := context.Background()
	store := newFakeStore()
	journal := NewJournal(store)

	op := Operation{ID: "op-1", SandboxID: "sb-1", Kind: KindCreate}
	if err := journal.Begin(ctx, op); err != nil {
		t.Fatal(err)
	}
	if err := journal.MarkRunning(ctx, op.ID); err != nil {
		t.Fatal(err)
	}
	if err := journal.Complete(ctx, op.ID); err != nil {
		t.Fatal(err)
	}

	got, err := store.Get(ctx, op.ID)
	if err != nil {
		t.Fatal(err)
	}
	if got.State != Done {
		t.Fatalf("state = %s", got.State)
	}
}

func TestJournalRejectsTerminalTransition(t *testing.T) {
	ctx := context.Background()
	store := newFakeStore()
	journal := NewJournal(store)

	op := Operation{ID: "op-1", SandboxID: "sb-1", Kind: KindCreate}
	if err := journal.Begin(ctx, op); err != nil {
		t.Fatal(err)
	}
	if err := journal.Fail(ctx, op.ID, "boom"); err != nil {
		t.Fatal(err)
	}
	if err := journal.Complete(ctx, op.ID); !errors.Is(err, core.ErrConflict) {
		t.Fatalf("expected conflict, got %v", err)
	}
}
