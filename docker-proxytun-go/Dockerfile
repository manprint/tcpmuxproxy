FROM debian:bookworm-slim
LABEL maintainer="Fabio"

RUN apt update && \
    apt install -y curl && \
    ARCH=$(uname -m) && \
    if [ "$ARCH" = "x86_64" ]; then \
        curl -L -o /bin/proxy https://raw.githubusercontent.com/manprint/tcpmuxproxy/refs/heads/main/builds/proxy-linux-amd64; \
    elif [ "$ARCH" = "aarch64" ]; then \
        curl -L -o /bin/proxy https://raw.githubusercontent.com/manprint/tcpmuxproxy/refs/heads/main/builds/proxy-linux-arm64; \
    else \
        echo "Unsupported architecture: $ARCH" && exit 1; \
    fi && \
    chmod +x /bin/proxy && \
    rm -rf /var/lib/apt/lists/*

CMD ["/bin/proxy"]