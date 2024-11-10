FROM golang:1.23 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o stresstest cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/stresstest .

RUN apk add --no-cache bash

ENTRYPOINT ["./stresstest"]