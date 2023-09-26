package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestRetrieveComposite(t *testing.T) {
	result, err := RetrieveComposite()

	// Case: Should not returning error
	if err != nil {
		t.Error(err)
	}

	// Case: Should returning struct and not error
	var manga Manga
	errUnmarshal := json.Unmarshal([]byte(*result), &manga)

	if errUnmarshal != nil {
		t.Error(errUnmarshal)
	}

	// Case: should returning Struct value as i expected
	if manga.Title != "Attack on Titan" {
		t.Error("Test is not as expected!")
	}

	if manga.Author != "Hajime" {
		t.Error("Test is not as expected")
	}

  fmt.Println("Manga:", manga.Title)
}
