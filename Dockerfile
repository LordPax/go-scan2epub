# Stage 1: Build the application
FROM golang:1.23-alpine AS builder

ARG VERSION
ARG URL=https://github.com/LordPax/go-scan2epub

WORKDIR /app

RUN apk add --no-cache git && \
    git clone --branch ${VERSION} ${URL} . && \
    go mod download && \
    go build -o scan2epub

# Stage 2: Create the final lightweight image
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /

COPY --from=builder /app/scan2epub /scan2epub

RUN mkdir books

VOLUME ["/root/.config/scan2epub", "/books"]

CMD ["/scan2epub", "inter"]
