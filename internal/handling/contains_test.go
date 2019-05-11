package handling

import (
	"testing"

	"github.com/theTardigrade/validation/internal/data"
)

type containsStringDummyModel struct {
	x string `validation:"contains=x"`
}

func TestContains_stringInvalid(t *testing.T) {
	model := containsStringDummyModel{}
	datum := containsDatum{}

	for _, v := range []string{"a", "X", "test"} {
		model.x = v

		executeTest(t, model, func(m *data.Main, t *data.Tag) (s []string) {
			s = append(s, datum.FailureMessage(m, t))
			return
		})
	}
}

func TestContains_stringValid(t *testing.T) {
	model := containsStringDummyModel{}

	for _, v := range []string{"", "x", "latex", "express"} {
		model.x = v

		executeTest(t, model, nil)
	}
}

type containsRequiredStringDummyModel struct {
	x string `validation:"required,contains=x"`
}

func TestContainsRequired_stringInvalid(t *testing.T) {
	model := containsRequiredStringDummyModel{}
	cDatum := containsDatum{}
	rDatum := requiredDatum{}

	for _, v := range []string{"", "a", "alphabetic"} {
		model.x = v

		executeTest(t, model, func(m *data.Main, t *data.Tag) (s []string) {
			switch t.Key {
			case "contains":
				s = append(s, cDatum.FailureMessage(m, t))
			case "required":
				if len(v) == 0 {
					s = append(s, rDatum.FailureMessage(m, t))
				}
			}
			return
		})
	}
}

func TestContainsRequired_stringValid(t *testing.T) {
	model := containsRequiredStringDummyModel{}

	for _, v := range []string{"x", "expressionistic", "xylophone"} {
		model.x = v

		executeTest(t, model, nil)
	}
}
