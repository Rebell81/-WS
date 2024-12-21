FROM golang:alpine3.21 AS builder
WORKDIR /build
COPY  . .
RUN go mod download
RUN go build -o /cws main.go

FROM alpine:3.21
WORKDIR /
COPY --from=builder /cws /

CMD ["/cws"]