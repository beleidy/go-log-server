# Go log server
This server acts as a centralised location to process your logs from different processes. Send your log item in using a POST request and chose where to store or display it. Options include, console, file and SQL database.

When storing in a database, this server allows you to always use a fixed number of connections with the database server. This avoids the possibility of your connections getting rejected if you were logging directly from your decentralised proccesses to the database.

## Prerequistes
* Go
* PostgreSQL (for enabling logging to database)

## Installing
### Get the source code
```
go get -d gitlab.com/beleidy/go-log-server
```
### Setup your configuration variables
```
cd $GOPATH/src/gitlab.com/beleidy/go-log-server
```
Edit the `.env` file. The variable names and their description below:

| **Env variable** | **type** | **description** | **default value**
|:-------------:|:-----:| ---------- |:-----:|
| logToScreen | bool | enable logging to stout | true |
| logToFile | bool | enable logging to file on the server | true |
| logToDb |  bool | enable logging to database | true |
| dbHost |  string | address of database server | localhost |
| dbPort |  uint16 | port number of the database server | 5432 |
| dbName |  string | name of the database | logs |
| dbUser |  string | username for database | postgres |
| dbPassword |  string | passwod for database | postgres |
| dbMaxConnections |  int | maximum number of connections between the server and your database | 95 |
| logFilePath |  string | path to write log file if enabled | "" |

### Run the server
```
go run .
```
Your server is now running at port `8080`

### Test with a sample log item
The server expects to recieve a `POST` request with a JSON in the body in the format
```
{
    "id": <string>,
    "time": <time ISO string>,
    "level": <int>,
    "message": <string>
}
```

so we can use [cURL](https://curl.haxx.se/docs/manual.html) to send a POST request from the command line. Open your terminal and enter
```
curl -d '{"id":"test-log", "time":"2019-04-16T08:03:46.657Z", "level":3, "message":"if you can see this, logging worked"}' \
    -H "Content-Type: application/json" -X POST http://localhost:8080
```
You should now be able to see the log item wherever you have logging enabled.

## Built with
* [Golang](https://golang.org/)
* [fasthttp](https://github.com/valyala/fasthttp)


