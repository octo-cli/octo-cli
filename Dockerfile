# Build
FROM golang:1.11-alpine AS go-github-cli-build-env

RUN apk add --no-cache git musl-dev gcc

WORKDIR /go/src/github.com/go-github-cli/go-github-cli
COPY . .

RUN script/bootstrap
RUN script/build

# Package
FROM alpine

COPY --from=go-github-cli-build-env /go/src/github.com/go-github-cli/go-github-cli/bin/go-github-cli /bin/go-github-cli
RUN apk add --no-cache ca-certificates jq
ENTRYPOINT ["go-github-cli"]
