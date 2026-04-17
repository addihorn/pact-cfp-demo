FROM golang:1.23-alpine AS builder
WORKDIR /src

COPY apps/api/go.mod ./apps/api/go.mod
WORKDIR /src/apps/api
RUN go mod download

COPY apps/api .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/api ./cmd/api

FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /out/api /app/api
EXPOSE 8080
ENTRYPOINT ["/app/api"]