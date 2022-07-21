FROM golang:1.18.3-alpine

RUN mkdir /go/src/jwt-tutorial
COPY . /go/src/jwt-tutorial
WORKDIR /go/src/jwt-tutorial

RUN go install

RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

CMD ["air"]