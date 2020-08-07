# build static binary
FROM golang:1.14.6-alpine3.11 as builder 

RUN apk --no-cache add  \
    tzdata \
    zip \
    ca-certificates \
    git 

WORKDIR /usr/share/zoneinfo
RUN zip -q -r -0 /zoneinfo.zip .

WORKDIR /go/src/github.com/bots-house/birzzha

# download dependencies 
COPY go.mod go.sum ./
RUN go mod download 

COPY . .

ARG REVISION

# compile 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
      -ldflags="-w -s -extldflags \"-static\" -X \"main.revision=${REVISION}\"" -a \
      -o /bin/birzzha .

# run 
FROM scratch

ENV ZONEINFO /zoneinfo.zip

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /zoneinfo.zip /
COPY --from=builder /bin/birzzha /bin/birzzha

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD [ "/bin/birzzha", "-health" ]

EXPOSE 8000

ENTRYPOINT [ "/bin/birzzha" ]
