package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/valyala/fasthttp"
)

var config Configuration

func main() {
	err := envconfig.Process("log", &config)
	check(err)

	go initMessage()

	dbQueue := make(chan Entry, 100)
	go dbLogger(dbQueue)

	postQueue := make(chan []byte)
	go postDecoder(postQueue, dbQueue)

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
	fmt.Println("Logging to Postgres Database: ")
	fmt.Println("Host Name: ", config.DbHost)
	fmt.Println("Port: ", config.DbPort)
	fmt.Println("Database Name: ", config.DbName)
	fmt.Println("User Name: ", config.DbUser)
	fmt.Println("Max DB connections: ", config.DbMaxConnections)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
