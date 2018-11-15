package main

import (
	"fmt"
)

func screenLogger(screenQueue chan Entry) {

	for {
		fmt.Print(<-screenQueue)
	}
}
