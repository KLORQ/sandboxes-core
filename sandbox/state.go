// Package sandbox defines sandbox lifecycle states.
package sandbox

import "fmt"

// State is the desired or observed sandbox state.
type State string

const (
	Pending  State = "pending"
	Creating State = "creating"
	Running  State = "running"
	Stopping State = "stopping"
	Stopped  State = "stopped"
	Deleting State = "deleting"
	Deleted  State = "deleted"
	Failed   State = "failed"
	Unknown  State = "unknown"
)

// ValidateTransition checks a lifecycle transition.
func ValidateTransition(from, to State) error {
	allowed := map[State]map[State]bool{
		Pending:  {Creating: true, Failed: true},
		Creating: {Running: true, Failed: true},
		Running:  {Stopping: true, Failed: true},
		Stopping: {Stopped: true, Failed: true},
		Stopped:  {Creating: true, Deleting: true, Failed: true},
		Failed:   {Creating: true, Deleting: true},
		Deleting: {Deleted: true, Failed: true},
		Unknown:  {Creating: true, Running: true, Stopping: true, Deleted: true, Failed: true},
	}
	if from == to {
		return nil
	}
	if allowed[from][to] {
		return nil
	}
	return fmt.Errorf("invalid transition %s -> %s", from, to)
}
