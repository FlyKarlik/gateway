FROM golang:1.19.6-alpine3.17 as builder

RUN mkdir /app
ADD ./gateway /app
ADD ./comet /comet
ADD ./protos /protos
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM alpine AS pruduction
COPY --from=builder /app/main .
COPY --from=builder /app/cert/ca-cert.pem /cert/ca-cert.pem

EXPOSE 3000
CMD ["./main"]