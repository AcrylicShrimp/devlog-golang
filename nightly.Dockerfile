FROM golang:alpine AS build

WORKDIR /tmp/build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

#RUN go test
RUN mkdir out
RUN go build -ldflags '-w -s' -o ./out/ .
RUN go build -ldflags '-w -s' -o ./out/ ./cmd/new-admin

FROM alpine:latest

WORKDIR /srv/app

COPY --from=build /tmp/build/out/* .

EXPOSE 8000

ENTRYPOINT ./devlog
