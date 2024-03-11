# Step 1: Modules caching
FROM golang:1.19.4 as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:1.19.4 as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /webappbot
WORKDIR /webappbot
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags migrate -o /bin/webappbot ./cmd/webappbot

# Step 3: Final
FROM scratch
COPY --from=builder /webappbot/config /config
COPY --from=builder /webappbot/migrations /migrations
COPY --from=builder /bin/webappbot /webappbot
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/app"]