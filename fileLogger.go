package main

import (
	"bufio"
	"os"
)

func fileLogger(fileQueue chan Entry) {
	fd, err := os.OpenFile(config.LogFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	check(err)
	defer fd.Close()

	fw := bufio.NewWriter(fd)

	for {
		e := <-fileQueue
		fw.WriteString(e.String())
		fw.Flush()
	}
}
