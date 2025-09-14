FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o gateway ./gateway/cmd/gateway

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/gateway .

EXPOSE 8000
CMD ["./gateway"]
