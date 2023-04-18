# STEP 1 build executable binary

FROM golang:1.20.3-alpine3.17 as builder

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
FROM alpine:3.17.3
COPY --from=builder /tmp/miniapi /bin/miniapi
RUN apk add curl=8.0.1-r0
RUN adduser -D -g '' user
USER user
ENTRYPOINT ["/bin/miniapi"]
