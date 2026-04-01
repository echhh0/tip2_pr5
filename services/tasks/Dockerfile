FROM golang:1.26-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/tasks ./services/tasks/cmd/tasks

FROM alpine:3.20
RUN adduser -D -H appuser

USER appuser
WORKDIR /app

COPY --from=builder /bin/tasks /app/tasks

EXPOSE 8082

ENTRYPOINT ["/app/tasks"]