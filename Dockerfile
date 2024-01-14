FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./

ENV GOARCH=arm64

EXPOSE 3000

RUN go build -o MyFitnessPal-Grafana .

CMD [ "/MyFitnessPal-Grafana" ]
