FROM golang:1.17.1-alpine as builder
WORKDIR /build

COPY . /build/
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app ./cmd/main.go

# generate clean, final image for end users
FROM alpine:3.11.3
COPY --from=builder /build/app .
COPY --from=builder /build/configs configs

# executable
ENTRYPOINT [ "/app" ]