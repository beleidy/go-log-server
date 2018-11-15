package main

import (
	"time"
)

// Entry is a log entry
type Entry struct {
	ID      string    `json:"id"`
	Time    time.Time `json:"time"`
	Message string    `json:"message"`
}
