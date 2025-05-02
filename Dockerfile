FROM golang:1.24.2-alpine3.21 AS gobuilder
WORKDIR /src
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . ./
RUN --mount=type=cache,target=/root/.cache \
    --mount=type=cache,target=/go/pkg/mod \
    go build -o mcp-husqvarna-automower *.go

FROM alpine:3.21
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=gobuilder /src/mcp-husqvarna-automower /app/
ENTRYPOINT ["/app/mcp-husqvarna-automower"]