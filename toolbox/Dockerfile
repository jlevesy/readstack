FROM golang:alpine

ARG UID
ARG GID

RUN apk add --update \
    git \
    bash \
    nodejs \
    yarn \
    make \
    g++ \
    postgresql-client && \
    go get -u -d github.com/mattes/migrate/cli github.com/lib/pq && \
    go build -tags 'postgres' -o /usr/local/bin/migrate github.com/mattes/migrate/cli && \
    go get -u github.com/golang/dep/cmd/dep && \
    yarn global add web-ext && \
    adduser -D -u ${UID} readstack-developer && \
    chown -R readstack-developer /go

USER $UID:$GID
