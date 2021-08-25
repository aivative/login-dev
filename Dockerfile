LABEL maintainer="Yosia <yosiaagustadewa@gmail.com>"

# build
FROM golang:1.16 as build
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED='0' && go build -v -o /app .

# runtime
FROM gcr.io/distroless/base
COPY --from=build /app /app

# NOT USED IN PRODUCTION
#
# COPY identity_cred.json .
# COPY config.json .

EXPOSE 10001

CMD ["/app"]