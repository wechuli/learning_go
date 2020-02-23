package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	mcPostBody := map[string]interface{}{
		"this": "that",
	}

	bytes.ne

	body, _ := json.Marshal(mcPostBody)
	fmt.Println(string(body))
}
