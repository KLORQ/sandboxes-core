package operation

import "testing"

func TestValidateTransition(t *testing.T) {
	if err := ValidateTransition(Pending, Running); err != nil {
		t.Fatal(err)
	}
	if err := ValidateTransition(Done, Running); err == nil {
		t.Fatal("expected terminal transition to fail")
	}
}

func TestOperationValidate(t *testing.T) {
	op := Operation{ID: "op-1", SandboxID: "sb-1", Kind: KindCreate, State: Pending}
	if err := op.Validate(); err != nil {
		t.Fatal(err)
	}
}
