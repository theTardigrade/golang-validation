package validation

import "testing"

type dummyValidationModel struct {
	a string   `validation:"required,minlen=4,maxlen=8"`
	b string   `validation:"required,email,minlen=6,maxlen=5000" name:"Email address"`
	c []string `validation:"minlen=4,maxlen=8"`
	d int      `validation:"min=5,max=5000,divisible=5"`
	e uint     `validation:"min=25,max=5000,divisible=25"`
	f float64  `validation:"min=0.5,max=5.5"`
	g float32  `validation:"min=0.5,max=5.5"`
	h *int     `validation:"required"`
	i int64    `validation:"min=1,max=9999999,divisible=3"`
}

func BenchmarkValidation(b *testing.B) {
	model := dummyValidationModel{}

	for n := 0; n < b.N; n++ {
		if _, _, err := Validate(model); err != nil {
			//panic(err)
		}
	}
}
