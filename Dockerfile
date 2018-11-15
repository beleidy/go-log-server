FROM golang:1.11 as builder

WORKDIR /go/src/github.com/beleidy/simple-log-server
ADD . .

ENV CGO_ENABLED=0
ENV GOOS=linux


RUN go get -d -v ./...
RUN go build -ldflags '-w -s' -a -installsuffix cgo

FROM scratch
COPY --from=builder /go/src/github.com/beleidy/simple-log-server/simple-log-server ./simple-log-server

ENV LOGTOSCREEN=false
ENV LOGTOFILE=false
ENV LOGTODB=true
ENV DBHOST=localhost
ENV DBPORT=5432
ENV DBNAME=logs
ENV DBUSER=postgres
ENV DBPASSWORD=postgres
ENV DBMAXCONNECTIONS=95
ENV LOGFILEPATH=""

EXPOSE 8080

CMD [ "./simple-log-server" ]



