#builder
FROM golang:1.23 AS builder

WORKDIR /build

ENV CGO_ENABLED=0
ENV GOCACHE=/go-cache
ENV GOMODCACHE=/gomod-cache

COPY go.mod go.sum ./
COPY cmd/ cmd/
COPY internal/ internal/

RUN --mount=type=cache,target=/gomod-cache --mount=type=cache,target=/go-cache\
  go build -o app ./cmd/...

#final
FROM scratch

COPY --from=builder /build/app /app
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/app"]
