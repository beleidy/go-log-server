package main

import (
	"bytes"
	"encoding/json"
)

func postDecoder(postQueue chan []byte,
	dbQueue chan Entry) {

	for {
		postBody := <-postQueue
		decoder := json.NewDecoder(bytes.NewReader(postBody))
		var entry Entry
		err := decoder.Decode(&entry)
		check(err)
		dbQueue <- entry

	}
}
