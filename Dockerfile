# FROM golang:1.20 AS builder

# RUN mkdir -p /app

# WORKDIR /app

# COPY fault_detection.go go.mod go.sum .

# RUN go mod download

# RUN go build -o myapp

# FROM alpine

# WORKDIR /

# COPY --from=builder /myapp .

# ENTRYPOINT ["./myapp"]

# 階段 1：編譯 Go 應用程序
FROM golang:1.20-alpine AS build

WORKDIR /app
COPY fault_detection.go go.mod go.sum .
RUN go mod download

RUN go build -o main .

# 階段 2：生成最小的運行時映像
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/main .

# 刪除不必要的依賴
RUN apk --no-cache add ca-certificates
RUN apk del .build-deps

CMD ["./main"]