FROM golang:1.23.2 AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# https://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/products cmd/products/main.go
COPY cmd/products/configs/*-prod.yml bin/configs/

FROM alpine AS products
ARG GIT_COMMIT_HASH="undefined"
ENV GIT_COMMIT_HASH=$GIT_COMMIT_HASH
WORKDIR /products
COPY --from=builder build/bin bin/
# https://stackoverflow.com/questions/21553353/what-is-the-difference-between-cmd-and-entrypoint-in-a-dockerfile
WORKDIR /products/bin
CMD ["./products"]
