# STEP 1 build executable binary

FROM golang:1.22.3-alpine3.19 as builder

RUN apk update
RUN apk add git

# Create appuser on builder image
RUN adduser -D -g '' appuser

COPY ./cmd/ /tmp/app/cmd
COPY ./env/ /tmp/app/env
COPY go.* /tmp/app/
WORKDIR /tmp/app
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o /tmp/miniapi github.com/udhos/miniapi/cmd/miniapi

#
# STEP 2 build a small image from alpine
#
# curl: https://security.alpinelinux.org/srcpkg/curl
#
FROM alpine:3.19.1
COPY --from=builder /tmp/miniapi /bin/miniapi
#RUN apk add curl=8.1.2-r0 libcrypto3=3.1.0-r4 libssl3=3.1.0-r4
RUN apk update && \
    apk add curl && \
    apk upgrade libssl3 libcrypto3 busybox busybox-binsh ssl_client libcurl nghttp2-libs && \
    adduser -D -g '' user && \
    rm -rf /var/cache/apk/*
USER user
ENTRYPOINT ["/bin/miniapi"]
