package main

import (
	"testing"
)

func TestGetArea(t *testing.T) {
	trArea := triangle{
		5.0,
		7.0,
	}
	if trArea.getArea() != 17.5 {
		t.Errorf("The area of triange is not equals to 35 rather it is %v", trArea.getArea())
	}

	sqArea := square{5}

	if sqArea.getArea() != 25 {
		t.Errorf("The area of triange is not equals to 35 rather it is %v", sqArea.getArea())
	}
}
