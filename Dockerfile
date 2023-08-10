FROM golang:1.20 AS builder

WORKDIR /app

COPY fault_detection.go go.mod go.sum .

RUN go mod vendor

RUN go build -mod vendor -o myapp fault_detection.go

FROM alpine

WORKDIR /

COPY --from=builder /app/myapp .

CMD ["./myapp"]