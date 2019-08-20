FROM golang

ENV GO111MODULE=on
COPY . /go/src/github.com/rkorpalski/Hello
WORKDIR /go/src/github.com/rkorpalski/Hello

RUN go get ./...
RUN go build
CMD Hello