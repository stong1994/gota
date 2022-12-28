package series

import (
	"fmt"
	"math"
)

type CustomElement interface {
	String() string
	Sort() int
	Value() interface{}
}

type customElement struct {
	e   CustomElement
	nan bool
}

// force customElement struct to implement Element interface
var _ Element = (*customElement)(nil)

func (e *customElement) Set(value interface{}) {
	e.nan = false
	switch val := value.(type) {
	case string:
		if val == "NaN" {
			e.nan = true
			return
		}
	case CustomElement:
		if value.(CustomElement) == nil {
			e.nan = true
		} else {
			e.e = value.(CustomElement)
		}

	default:
		e.nan = true
		return
	}
}

func (e customElement) Copy() Element {
	if e.IsNA() {
		return &customElement{nil, true}
	}
	return &customElement{e.e, false}
}

func (e customElement) IsNA() bool {
	return e.nan
}

func (e customElement) Type() Type {
	return Custom
}

func (e customElement) Val() ElementValue {
	if e.IsNA() {
		return nil
	}
	return e.e.Value()
}

func (e customElement) String() string {
	if e.IsNA() {
		return "NaN"
	}
	return e.e.String()
}

func (e customElement) Int() (int, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to int")
	}
	return 0, nil
}

func (e customElement) Float() float64 {
	if e.IsNA() {
		return math.NaN()
	}
	return 0.0
}

func (e customElement) Bool() (bool, error) {
	if e.IsNA() {
		return false, fmt.Errorf("can't convert NaN to bool")
	}
	return false, nil
}

func (e customElement) Eq(elem Element) bool {
	b, err := elem.Bool()
	if err != nil || e.IsNA() {
		return false
	}
	return e.e.Value() == b
}

func (e customElement) Neq(elem Element) bool {
	return e.e.Value() != elem.(CustomElement).Value()
}

func (e customElement) Less(elem Element) bool {
	if e.IsNA() {
		return false
	}
	return e.e.Sort() < elem.(CustomElement).Sort()
}

func (e customElement) LessEq(elem Element) bool {
	if e.IsNA() {
		return false
	}
	return e.e.Sort() <= elem.(CustomElement).Sort()
}

func (e customElement) Greater(elem Element) bool {
	if e.IsNA() {
		return false
	}
	return e.e.Sort() > elem.(CustomElement).Sort()
}

func (e customElement) GreaterEq(elem Element) bool {
	if e.IsNA() {
		return false
	}
	return e.e.Sort() >= elem.(CustomElement).Sort()
}
