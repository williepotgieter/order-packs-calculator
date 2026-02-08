FROM golang:1.25.7-trixie AS builder

WORKDIR /app

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -trimpath -ldflags="-s -w" -o packapp cmd/app/main.go

FROM alpine:3.23

WORKDIR /app

RUN addgroup -S appuser \
    && adduser -S -G appuser -H -s /sbin/nologin appuser

COPY --from=builder --chown=appuser:appuser /app/packapp /app/packapp

USER appuser

EXPOSE $API_PORT

ENTRYPOINT ["/app/packapp"]