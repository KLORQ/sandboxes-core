package sandbox

import "testing"

func TestValidateTransition(t *testing.T) {
	if err := ValidateTransition(Pending, Creating); err != nil {
		t.Fatal(err)
	}
	if err := ValidateTransition(Running, Deleted); err == nil {
		t.Fatal("expected invalid transition")
	}
}
