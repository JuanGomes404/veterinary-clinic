FROM golang:1.23-alpine


WORKDIR /app
COPY .  /app

RUN CGO_ENABLED=0 GOOS=linux go build -o /api 

EXPOSE 8080
CMD ["/app"]
