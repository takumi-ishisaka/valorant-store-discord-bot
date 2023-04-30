FROM golang:1.20.3

ENV APP_BASE_PATH = /APP_BASE_PATH
RUN apt-get update \
    && apt-get install -y --no-install-recommends

WORKDIR /app