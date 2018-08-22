FROM golang:alpine AS build-back
COPY . /go/src/github.com/jlevesy/readstack
WORKDIR /go/src/github.com/jlevesy/readstack
RUN apk --update add make git && \
  go get -u github.com/golang/dep/cmd/dep && \
  make static_build

FROM scratch
ENV READSTACK_WEB_ASSETS_PATH=/readstack/web
ENV READSTACK_LISTEN_PORT=8080

COPY --from=build-back /go/src/github.com/jlevesy/readstack/dist/server /readstack/server

EXPOSE 8080
CMD ["/readstack/server"]
