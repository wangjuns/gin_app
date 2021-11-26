FROM golang:latest

RUN mkdir /code
WORKDIR /code

ADD . /code/

RUN go build -o app .


ENTRYPOINT ["./app"]