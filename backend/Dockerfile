FROM golang:1.12.6-alpine3.10 AS build

RUN apk add --no-cache git gcc libc-dev
RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/bonnetn/dare/backend

COPY Gopkg.lock .
COPY Gopkg.toml .
RUN dep ensure -vendor-only

COPY app.go .
COPY internal/ ./internal/

RUN go build .

FROM alpine:3.10

WORKDIR /dare

COPY --from=build /go/src/github.com/bonnetn/dare/backend/backend .

EXPOSE 9090

CMD ./backend
