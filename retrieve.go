package main

func Retrieve() (*string, error) {
	val, err := client.Get("name").Result()

	if err != nil {
		return nil, err
	}

	return &val, nil
}
