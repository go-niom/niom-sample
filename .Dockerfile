
# build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
ENV GOPATH /go
WORKDIR /go/src
COPY . /go/src/app-build
RUN cd /go/src/app-build && env GOOS=linux CGO_ENABLED=0 go build .

#final stage
FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*
WORKDIR /app
COPY --from=builder /go/src/app-build/niom-sample /app
COPY .env /app

LABEL Name=niom-sample Version=0.0.1
EXPOSE 4000
CMD ["./niom-sample"]
