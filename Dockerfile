FROM golang:alpine AS build-back
COPY . /go/src/github.com/jlevesy/readstack
WORKDIR /go/src/github.com/jlevesy/readstack
RUN apk --update add make git && \
  go get -u github.com/golang/dep/cmd/dep && \
  make static_build

FROM node:alpine AS build-web
COPY ./web /home/node/app
WORKDIR /home/node/app
RUN yarn install && yarn build

FROM scratch
ENV READSTACK_WEB_ASSETS_PATH=/readstack/web \
  READSTACK_LISTEN_PORT=8080

COPY --from=build-back /go/src/github.com/jlevesy/readstack/dist/server /readstack/server
COPY --from=build-web /home/node/app/dist /readstack/web

EXPOSE 8080
CMD ["/readstack/server"]
