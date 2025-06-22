FROM docker.io/library/golang:1.24.4 AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o weather-api ./cmd

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /src .

ENV REDIS_PASSWORD=
ENV REDIS_ADDR=
ENV REDIS_DB=
ENV WEATHER_VISUALCROSSING_APIKEY=

EXPOSE 8080

CMD ["./weather-api"]
