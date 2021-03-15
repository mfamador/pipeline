FROM golang:1.16 as builder

# Enable modules support
ENV GO111MODULE=on

RUN mkdir -p /build
WORKDIR /build

COPY go.mod go.sum ./

COPY . .

RUN go mod download

ENV CGO_ENABLED=0
# Compile static binary
RUN go build -ldflags "-s -w" ./cmd/app

##################################################
# Final image
FROM alpine

RUN mkdir -p /app/config
WORKDIR /app

# Copy pipeline config
COPY --from=builder /build/config/pipeline.yaml config/pipeline.yaml
# Copy the binary from the previous stage
COPY --from=builder /build/pipeline pipeline

CMD ["/app/pipeline", "-c", "/app/config/pipeline.yaml"]
