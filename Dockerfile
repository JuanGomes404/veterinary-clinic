FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /api 

FROM alpine:latest

WORKDIR /
COPY --from=builder /api /api

EXPOSE 8080
CMD ["/api"]
