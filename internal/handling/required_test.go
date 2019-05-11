package handling

import (
	"testing"

	"github.com/theTardigrade/validation/internal/data"
)

type requiredStringDummyModel struct {
	x string `validation:"required"`
}

func TestRequired_stringInvalid(t *testing.T) {
	model := requiredStringDummyModel{}
	datum := requiredDatum{}

	executeTest(t, model, func(m *data.Main, t *data.Tag) (s []string) {
		s = append(s, datum.FailureMessage(m, t))
		return
	})
}

func TestRequired_stringValid(t *testing.T) {
	model := requiredStringDummyModel{}

	for _, s := range []string{" ", "-", "x", "the", "test"} {
		model.x = s

		executeTest(t, model, nil)
	}
}

type requiredPointerDummyModel struct {
	x *string   `validation:"required"`
	y **string  `validation:"required"`
	z ***string `validation:"required"`
}

func TestRequired_pointerInvalid(t *testing.T) {
	model := requiredPointerDummyModel{}
	datum := requiredDatum{}

	executeTest(t, model, func(m *data.Main, t *data.Tag) (s []string) {
		s = append(s, datum.FailureMessage(m, t))
		return
	})
}

func TestRequired_pointerValid(t *testing.T) {
	model := requiredPointerDummyModel{}

	one, two, three := "", "2", "three"

	for _, s := range []*string{&one, &two, &*&three} {
		model.x = s
		model.y = &s
		model.z = &model.y

		executeTest(t, model, nil)
	}
}
