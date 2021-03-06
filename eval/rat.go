package eval

import (
	"errors"
	"fmt"
	"math/big"
)

var ErrOnlyStrOrRat = errors.New("only str or rat may be converted to rat")

// Rat is a rational number.
type Rat struct {
	b *big.Rat
}

var _ Value = Rat{}

func (Rat) Kind() string {
	return "string"
}

func (r Rat) Eq(a interface{}) bool {
	if r == a {
		return true
	}
	r2, ok := a.(Rat)
	if !ok {
		return false
	}
	return r.b.Cmp(r2.b) == 0
}

func (r Rat) Repr(int) string {
	return "(rat " + r.String() + ")"
}

func (r Rat) String() string {
	if r.b.IsInt() {
		return r.b.Num().String()
	}
	return r.b.String()
}

// ToRat converts a Value to rat. A str can be converted to a rat if it can be
// parsed. A rat is returned as-is. Other types of values cannot be converted.
func ToRat(v Value) (Rat, error) {
	switch v := v.(type) {
	case Rat:
		return v, nil
	case String:
		r := big.Rat{}
		_, err := fmt.Sscanln(string(v), &r)
		if err != nil {
			return Rat{}, fmt.Errorf("%s cannot be parsed as rat", v.Repr(NoPretty))
		}
		return Rat{&r}, nil
	default:
		return Rat{}, ErrOnlyStrOrRat
	}
}
