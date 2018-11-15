package main

import (
	"bytes"
	"fmt"
	"time"
)

// Entry is a log entry
type Entry struct {
	ID      string    `json:"id"`
	Time    time.Time `json:"time"`
	Level   int       `json:"level"`
	Message string    `json:"message"`
}

func (e Entry) String() string {
	return fmt.Sprintf("%s\t%v\t%d\t%s\n", e.ID, e.Time, e.Level, e.Message)
}

// Entries is an array of Entry structs
type Entries []Entry

func (es Entries) String() string {
	var b bytes.Buffer
	for _, e := range es {
		b.WriteString(e.String())
	}
	return b.String()

}
