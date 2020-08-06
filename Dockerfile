# BUILDER
# hadolint ignore=DL3006,DL3007
FROM golang:1.14.2 AS builder
WORKDIR /go/src/github.com/bots-house/birzzha

# install deps
COPY go.mod go.sum ./
RUN go mod download

# build project
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

# RUNNER
# hadolint ignore=DL3006,DL3007
FROM alpine:latest  
# hadolint ignore=DL3018
RUN apk --no-cache add ca-certificates

WORKDIR /root/
EXPOSE 8000
COPY --from=builder /go/src/github.com/bots-house/birzzha/birzzha /bin/
ENTRYPOINT ["/bin/birzzha"]