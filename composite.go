package main

import "encoding/json"

// Should be json
type Manga struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func Composite() error {
	json, err := json.Marshal(Manga{
		Title:  "Attack on Titan",
		Author: "Hajime",
	})

	if err != nil {
		return err
	}

	err = client.Set("manga_5", json, 0).Err()

	if err != nil {
		return err
	}

	return nil
}
