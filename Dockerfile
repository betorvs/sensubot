FROM golang:1.16.0-alpine3.12 AS golang

ARG LOC=/builds/go/src/github.com/betorvs/sensubot/
RUN apk add --no-cache git
RUN mkdir -p $LOC
ENV GOPATH /go
COPY . $LOC
ENV CGO_ENABLED 0
RUN cd $LOC && TESTRUN=true go test ./... && go build

FROM alpine:3.12
ARG LOC=/builds/go/src/github.com/betorvs/sensubot
WORKDIR /
VOLUME /tmp
RUN apk add --no-cache ca-certificates
RUN update-ca-certificates
RUN mkdir -p /app
RUN addgroup -g 1000 -S app && \
    adduser -u 1000 -G app -S -D -h /app app && \
    chmod 755 /app
COPY --from=golang $LOC/sensubot /app

EXPOSE 9090
RUN chmod +x /app/sensubot
WORKDIR /app    
USER app
CMD ["/app/sensubot"]
