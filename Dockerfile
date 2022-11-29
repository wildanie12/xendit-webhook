FROM golang:1.18-alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/wildanie12/xendit-webhook

COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/xendit-webhook

FROM alpine

COPY --from=builder /go/bin/xendit-webhook /go/bin/xendit-webhook

