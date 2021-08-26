LABEL maintainer="Yosia <yosiaagustadewa@gmail.com>"

# build
FROM golang:1.16.4 AS builder

WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED='0' && go build -o login-dev main.go 

# certs
FROM alpine:latest as certs
RUN apk --update add ca-certificates


# runtime
FROM debian:buster-slim

COPY --from=builder /go/src/app/login-dev /login-dev

COPY identity_cred.json /
COPY config.json /

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENV PATH="/go/bin:${PATH}"

CMD ["/login-dev"]
