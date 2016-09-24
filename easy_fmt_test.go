package easy_fmt

import (
	"fmt"
	"testing"
)

type SStruct struct {
	A *int
}
type TStruct struct {
	A int
	B string
	C *int
	D *string
	E SStruct
	F *SStruct
	G []int
	H []*int
	I []*SStruct
}

func TestStructToString(t *testing.T) {
	a := 5
	b := "abc"
	c := 6
	d := "def"
	fa := 1
	X := new(TStruct)
	X.F = new(SStruct)
	X.F.A = &fa
	X.A = a
	X.B = b
	X.C = &c
	X.D = &d
	X.G = []int{1, 2, 3}
	X.H = []*int{&a, &c}
	X.I = []*SStruct{X.F}
	fmt.Println(EasyFormat(X))
}
