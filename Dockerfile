FROM golang:1.21 AS builder

WORKDIR /app/

ADD go.mod go.sum ./
RUN go mod download && go mod verify

ADD . ./

RUN go build -o ./bin/main ./cmd/app

FROM golang:1.21 AS server

WORKDIR /app/

COPY --from=builder /app/bin/main ./
COPY --from=builder /app/internal/assets ./assets

ENV ASSETS_DIR=assets/

EXPOSE 80

CMD ["./main"]
