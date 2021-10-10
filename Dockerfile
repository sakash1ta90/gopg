FROM golang:1.17-alpine

WORKDIR /go/src
RUN ls -la
COPY . .

RUN apk upgrade --update && \
    apk --no-cache add git && \
    go get -u github.com/cosmtrek/air && \
    go install github.com/swaggo/swag/cmd/swag@latest && \
    go build -o /go/bin/air github.com/cosmtrek/air && \
    go build -o /go/bin/swag github.com/swaggo/swag/cmd/swag && \
    adduser -D gouser && chown -R gouser /go/src
USER gouser

CMD ["air", "-c", ".air.toml"]