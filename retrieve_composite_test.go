package main

import (
	"fmt"
	"testing"
)

func TestRetrieveComposite(t *testing.T) {
	result, err := RetrieveComposite()

	if err != nil {
		t.Error(err)
	}

	fmt.Printf("Result: %v", *result)
}
