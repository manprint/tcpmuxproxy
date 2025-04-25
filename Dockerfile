FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN ARCH=$(uname -m) && \
    if [ "$ARCH" = "x86_64" ]; then \
        GOOS=linux GOARCH=amd64 go build -o proxy; \
    elif [ "$ARCH" = "aarch64" ]; then \
        GOOS=linux GOARCH=arm64 go build -o proxy; \
    else \
        echo "Unsupported architecture: $ARCH" && exit 1; \
    fi

FROM scratch AS final

COPY --from=builder /app/proxy /proxy

CMD ["/proxy"]



