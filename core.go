// Package core defines shared domain errors.
package core

import "errors"

// ErrNotFound reports a missing resource.
var ErrNotFound = errors.New("not found")

// ErrConflict reports a stale generation or an invalid state transition.
var ErrConflict = errors.New("conflict")
