FROM golang:alpine AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/gorpc-template
COPY . .
RUN go get -d -v
RUN go build -o /gorpc-template

FROM scratch
COPY --from=builder /gorpc-template /gorpc-template
ENTRYPOINT ["/gorpc-template"]