FROM golang:1.15.3-alpine3.12

ARG UID
ENV UID=$UID

RUN adduser --uid $UID --home /go gouser --disabled-password
WORKDIR /go