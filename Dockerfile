FROM golang:1.20-alpine AS build

WORKDIR /app
COPY fault_detection.go go.mod go.sum .
RUN go mod download
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/main .

CMD ["./main"]