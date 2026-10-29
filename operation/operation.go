// Package operation defines durable operation records.
package operation

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
}
