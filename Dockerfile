# Multi-stage build setup (https://docs.docker.com/develop/develop-images/multistage-build/)

# Stage 1 (to create a "build" image, ~850MB)
FROM golang:1.17.2 AS builder
RUN go version

COPY . /go/src/github.com/marcusramberg/plexpod
WORKDIR /go/src/github.com/marcusramberg/plexpod

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

FROM alpine:3.14.2
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/marcusramberg/plexpod/app .

EXPOSE 8123
ENTRYPOINT ["./app"]
