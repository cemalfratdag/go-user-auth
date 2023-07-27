ARG GO_VERSION=1.20.5

FROM golang:${GO_VERSION}

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy
RUN go mod download
RUN go build -o main .

EXPOSE 3003