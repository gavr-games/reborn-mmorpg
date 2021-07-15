FROM ruby:3.0.2-alpine

RUN mkdir -p /usr/src/api
WORKDIR /usr/src/api

RUN apk add --update --upgrade gcc less postgresql-dev \
    git build-base xz-dev libc6-compat \
    postgresql-client linux-headers tzdata && \
    rm -rf /var/cache/apk/*

RUN gem install bundler

EXPOSE 4567
