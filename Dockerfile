# BUILDER
FROM golang:1.14.2 AS builder
WORKDIR /go/src/github.com/bots-house/birzzha

# install deps
ADD go.mod go.sum ./
RUN go mod download

# build project
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

# RUNNER
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/
EXPOSE 8000
COPY --from=builder /go/src/github.com/bots-house/birzzha/birzzha /bin/
ENTRYPOINT ["/bin/birzzha"]