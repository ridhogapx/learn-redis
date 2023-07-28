package main

import "testing"

func TestComposite(t *testing.T) {
	err := Composite()

	if err != nil {
		t.Error(err)
	}
}
