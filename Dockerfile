FROM golang:1.13 AS building

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY . /usr/src/app

RUN \
  mkdir -p build && \
  go build -o build/http_runner ./runner/http

FROM alpine:3.10

# ref. https://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
RUN \
  mkdir /lib64 && \
  ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

COPY \
  --from=building \
  /usr/src/app/build/http_runner \
  /usr/bin/http_runner

CMD "http_runner"
