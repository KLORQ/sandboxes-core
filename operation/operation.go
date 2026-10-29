// Package operation defines durable operation records.
package operation

import (
	"fmt"
	"time"
)

// Kind identifies the operation type.
type Kind string

const (
	KindCreate Kind = "create"
	KindStart  Kind = "start"
	KindStop   Kind = "stop"
	KindDelete Kind = "delete"
)

// State is the operation state.
type State string

const (
	Pending State = "pending"
	Running State = "running"
	Done    State = "done"
	Failed  State = "failed"
)

// Operation is a durable, idempotent unit of work.
type Operation struct {
	ID             string
	SandboxID      string
	Kind           Kind
	State          State
	Generation     int64
	IdempotencyKey string
	Error          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// Validate checks the operation invariants.
func (o Operation) Validate() error {
	if o.ID == "" {
		return fmt.Errorf("operation id is required")
	}
	if o.SandboxID == "" {
		return fmt.Errorf("sandbox id is required")
	}
	if o.Kind == "" {
		return fmt.Errorf("operation kind is required")
	}
	if o.State == "" {
		return fmt.Errorf("operation state is required")
	}
	return nil
}

// ValidateTransition checks a durable operation transition.
func ValidateTransition(from, to State) error {
	if from == to {
		return nil
	}
	allowed := map[State]map[State]bool{
		Pending: {Running: true, Failed: true},
		Running: {Done: true, Failed: true},
		Failed:  {Running: true},
		Done:    {},
	}
	if allowed[from][to] {
		return nil
	}
	return fmt.Errorf("invalid operation transition %s -> %s", from, to)
}
