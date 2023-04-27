FROM golang:alpine AS build
ENV GOPROXY=https://proxy.golang.org
WORKDIR /app


COPY go.mod go.sum /app/
RUN go mod download
COPY  . /app/
RUN go build -o processor

FROM alpine 

WORKDIR /app

COPY --from=build /app/processor /app/
COPY ./infrastructure/datastore/migrations /app/code-challenge/infrastructure/datastore/migrations
COPY ./templates /app/code-challenge/templates


ENTRYPOINT  ["/app/processor"]