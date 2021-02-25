package handling

import (
	"testing"

	"github.com/theTardigrade/golang-validation/internal/data"
)

type maxlenStringDummyModel struct {
	x string `validation:"maxlen=4"`
}

func TestMaxlen_stringInvalid(t *testing.T) {
	model := maxlenStringDummyModel{}
	datum := maxlenDatum{}

	for _, s := range []string{"these", "tests", "abcdefghijklmnopqrstuvwxyz"} {
		model.x = s

		executeTest(t, model, func(m *data.Main, t *data.Tag) (s []string) {
			s = append(s, datum.FailureMessage(m, t))
			return
		})
	}
}

func TestMaxlen_stringValid(t *testing.T) {
	model := maxlenStringDummyModel{}

	for _, s := range []string{"", "a", "at", "the", "this"} {
		model.x = s

		executeTest(t, model, nil)
	}
}

type maxlenSliceDummyModel struct {
	x []string `validation:"maxlen=4"`
}

func TestMaxlen_sliceInvalid(t *testing.T) {
	model := maxlenSliceDummyModel{}
	datum := maxlenDatum{}

	for _, s := range [][]string{
		[]string{"a", "b", "c", "d", "e"},
		[]string{"0", "1", "2", "3", "4", "5", "6"},
	} {
		model.x = s

		executeTest(t, model, func(m *data.Main, t *data.Tag) (s []string) {
			s = append(s, datum.FailureMessage(m, t))
			return
		})
	}
}

func TestMaxlen_sliceValid(t *testing.T) {
	model := maxlenSliceDummyModel{}

	for _, s := range [][]string{
		[]string{},
		[]string{"a"},
		[]string{"the", "this", "these"},
	} {
		model.x = s

		executeTest(t, model, nil)
	}
}
