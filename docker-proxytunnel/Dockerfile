FROM alpine:latest AS builder
LABEL maintainer="Fabio"

RUN apk add --no-cache \
    build-base \
    git \
    wget \
    tar \
    openssl \
    openssl-dev \
    asciidoc \
    xmlto

RUN git clone https://github.com/proxytunnel/proxytunnel.git && \
    cd proxytunnel && \
    make && \
    make install && \
    cd .. && \
    rm -rf proxytunnel

COPY ./entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

FROM alpine:latest AS final
LABEL maintainer="Fabio"

COPY --from=builder /usr/local/bin/proxytunnel /usr/local/bin/proxytunnel
COPY --from=builder /entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]