package handling

import (
	"testing"
)

type emailDummyModel struct {
	x string `validation:"email"`
}

func TestEmail_empty(t *testing.T) {
	model := emailDummyModel{""}
	executeTest(t, model, 1)
}

func TestEmail_invalid(t *testing.T) {
	model := emailDummyModel{"test"}
	executeTest(t, model, 1)
}

func TestEmail_valid(t *testing.T) {
	model := emailDummyModel{"test@test.com"}
	executeTest(t, model, 0)
}
