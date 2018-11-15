package main

import (
	"bytes"
	"encoding/json"
)

func postDecoder(postQueue chan []byte,
	loggingChannels []chan Entry) {

	for {
		postBody := <-postQueue

		decoder := json.NewDecoder(bytes.NewReader(postBody))
		var entry Entry
		err := decoder.Decode(&entry)
		check(err)

		for _, ch := range loggingChannels {
			ch <- entry
		}
	}
}
