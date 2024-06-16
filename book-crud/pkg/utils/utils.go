package utils

import (
	"encoding/json"
	"io"
	"log"
)

func ParseBody(raw io.ReadCloser, x any) {
	body, err := io.ReadAll(raw)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, x)
	if err != nil {
		log.Fatal(err)
	}
}
