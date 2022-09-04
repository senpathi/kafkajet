FROM golang:1.18.5-alpine3.16 AS builder

WORKDIR /src
COPY . .
RUN go build -ldflags="-s -w" -o ./kafka-jet *.go

FROM alpine:3.16.0

WORKDIR /opt
COPY --from=builder /src/kafka-jet /opt

ENTRYPOINT ["sh","-c","./kafka-jet"]