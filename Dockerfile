# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

# GCC needed for sqlite3
RUN apk add build-base

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./
ADD api ./api
ADD common ./common
ADD domain ./domain
ADD infrastructure ./infrastructure
ADD sqlite ./sqlite

RUN go build -o /cf-proposal

EXPOSE 9000

CMD [ "/cf-proposal" ]