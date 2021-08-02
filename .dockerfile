FROM golang:latest

WORKDIR /go/src/github.com/binhosemcrause/webapi-with-go

# ADD * /app
# ADD main.go /app
# ADD webapi-with-go /app
# ADD main /app

ADD . /go/src/github.com/binhosemcrause/webapi-with-go

RUN go get github.com/gin-gonic/gin
RUN go get -u github.com/lib/pq
# RUN go get github.com/binhosemcrause/webapi-with-go
RUN go install github.com/binhosemcrause/webapi-with-go

ENV PORT=5000

EXPOSE $PORT
HEALTHCHECK CMD curl --fail http://localhost:$PORT/books || exit 1

# RUN ["chmod", "777", "/go/src/github.com/binhosemcrause/webapi-with-go/main.go"]

# RUN go build -o main .

ENTRYPOINT /go/bin/apigo
CMD ["/go/bin/apigo"]

