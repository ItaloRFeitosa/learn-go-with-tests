FROM golang:latest

WORKDIR /app/src

# ENV GOPATH=/go/src
# ENV GOROOT=/go/root

COPY ./ ./ 