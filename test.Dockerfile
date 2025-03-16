FROM golang:1.23 as golang

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go test -c -o /usr/local/bin/test ./...

FROM ubuntu:22.04
COPY --from=golang /usr/local/bin/test /usr/local/bin/test

CMD ["test"]
