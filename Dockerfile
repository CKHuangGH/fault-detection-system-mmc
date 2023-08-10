FROM golang:1.20 AS builder

RUN mkdir -p /app

WORKDIR /app

COPY fault_detection.go go.mod go.sum .

RUN go mod download

RUN go build -o myapp

FROM alpine

RUN mkdir -p /app

COPY --from=builder /app/myapp /app/myapp

WORKDIR /app 

ENTRYPOINT ["./myapp"]