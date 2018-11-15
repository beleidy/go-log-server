package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/valyala/fasthttp"
)

var config Configuration

func main() {
	err := envconfig.Process("", &config)
	check(err)

	go initMessage()

	var loggingChannels []chan Entry

	if config.LogToScreen {
		screenQueue := make(chan Entry, 100)
		loggingChannels = append(loggingChannels, screenQueue)
		go screenLogger(screenQueue)
	}

	if config.LogToFile {
		fileQueue := make(chan Entry, 100)
		loggingChannels = append(loggingChannels, fileQueue)
		go fileLogger(fileQueue)
	}

	if config.LogToDb {
		dbQueue := make(chan Entry, 100)
		loggingChannels = append(loggingChannels, dbQueue)
		go dbLogger(dbQueue)
	}

	postQueue := make(chan []byte)
	go postDecoder(postQueue, loggingChannels)

	m := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/":
			switch string(ctx.Method()) {
			case "GET":
				getHandlerFunc(ctx)
			case "POST":
				postHandlerFunc(ctx, postQueue)
			}
		}
	}
	fasthttp.ListenAndServe(":8080", m)
}

func getHandlerFunc(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "You have reached the central logging server")
}

func postHandlerFunc(ctx *fasthttp.RequestCtx, postQueue chan []byte) {
	postQueue <- ctx.PostBody()
	fmt.Fprint(ctx, "ACK")
}

func initMessage() {
	fmt.Println("Fast Log Server started:")

	fmt.Println("Logging to Screen: ", config.LogToScreen)

	fmt.Println("Logging to File: ", config.LogToFile)
	if config.LogToFile {
		fmt.Println("Log file path: ", config.LogFilePath)
	}

	fmt.Println("Logging to Postgres Database: ", config.LogToDb)
	if config.LogToDb {
		fmt.Println("Host Name: ", config.DbHost)
		fmt.Println("Port: ", config.DbPort)
		fmt.Println("Database Name: ", config.DbName)
		fmt.Println("User Name: ", config.DbUser)
		fmt.Println("Max DB connections: ", config.DbMaxConnections)

	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
