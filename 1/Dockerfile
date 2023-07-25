# build stage
FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o health .

# final stage
FROM alpine:latest
RUN apk --no-cache add curl
WORKDIR /app
COPY --from=builder /app/health .
EXPOSE 8080
CMD ["./health"]
HEALTHCHECK --start-period=5s --retries=3 CMD curl --fail http://localhost:8080/health