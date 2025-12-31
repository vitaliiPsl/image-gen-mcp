FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /bin/image-gen-mcp ./cmd/server

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /bin/image-gen-mcp /bin/image-gen-mcp

ENTRYPOINT ["/bin/image-gen-mcp"]
