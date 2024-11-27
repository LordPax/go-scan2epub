# Stage 1: Build the application
FROM --platform=$BUILDPLATFORM golang:1.23-alpine AS builder

ARG VERSION
ARG URL=https://github.com/LordPax/go-scan2epub
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app

RUN mkdir /books
RUN apk add --no-cache git && \
    git clone --branch ${VERSION} ${URL} . && \
    go mod download && \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o scan2epub

# Stage 2: Create the final lightweight image
FROM alpine:latest

WORKDIR /

COPY --from=builder /app/scan2epub /scan2epub
COPY --from=builder /books /books
VOLUME ["/root/.config/scan2epub", "/books"]

CMD ["/scan2epub", "inter"]
