package handling

import (
	"testing"
)

type emailDummyModel struct {
	x string `validation:"email"`
}

func TestEmail_invalidEmpty(t *testing.T) {
	model := emailDummyModel{""}
	executeTest(t, model, 1)
}

func TestEmail_invalidString(t *testing.T) {
	model := emailDummyModel{"test"}
	executeTest(t, model, 1)
}

func TestEmail_valid(t *testing.T) {
	model := emailDummyModel{"test@test.com"}
	executeTest(t, model, 0)
}
