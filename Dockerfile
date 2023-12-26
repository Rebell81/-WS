FROM golang:1.21.5-alpine3.19 AS builder
WORKDIR /build
COPY  . .
RUN go mod download
RUN go build -o /cws main.go

FROM alpine:3.19
WORKDIR /
COPY --from=builder /cws /

CMD ["/cws"]