FROM golang:1.22-alpine AS build
RUN apk add --no-cache git gcc musl-dev
RUN adduser -D -g '' appuser
USER appuser
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY --chown=appuser:appuser . .
RUN go build -o taskly-cli ./cmd/taskly


FROM alpine:3.18 AS production
RUN apk add --no-cache ca-certificates
RUN adduser -D -g '' appuser
USER appuser
WORKDIR /app
COPY --from=build /app/taskly-cli .
EXPOSE 3000
CMD ["./taskly-cli"]