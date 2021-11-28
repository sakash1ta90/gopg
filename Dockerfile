FROM golang:1.17-alpine as build

WORKDIR /go/src
COPY . .

RUN set -eux && \
    apk upgrade --update && \
    apk --no-cache add git alpine-sdk build-base && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air && \
    go install github.com/swaggo/swag/cmd/swag@latest

ENV GO111MODULE=on

RUN set -eux && \
  go build -gcflags "all=-N -l" -o app ./main.go

FROM alpine:3

WORKDIR /app

COPY --from=build /go/src/app .

RUN set -x && \
  addgroup go && \
  adduser -D -G go go && \
  chown -R go:go /app/app

CMD ["./app"]