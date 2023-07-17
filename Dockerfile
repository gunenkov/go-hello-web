FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /myapp

FROM alpine:3.18.2

WORKDIR /

COPY --from=builder /myapp /myapp

EXPOSE 8080

ENV GIN_MODE=release

ENTRYPOINT ["/myapp"]