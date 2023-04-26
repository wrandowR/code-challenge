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


ENTRYPOINT  ["/app/processor"]