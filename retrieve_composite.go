package main

func RetrieveComposite() (*string, error) {
	result, err := client.Get("manga_5").Result()

	if err != nil {
		return nil, err
	}

	return &result, nil
}
