FROM golang:latest

RUN mkdir /code
WORKDIR /code

RUN go get -u github.com/go-gonic/gin


ENTRYPOINT ["go", "run", "app.go"]