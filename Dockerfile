# 使用 golang 官方的 Golang 鏡像作為基本映像
FROM golang:1.20 AS builder

# 設定工作目錄
WORKDIR /app

# 複製程式碼到容器內
COPY main.go .

# 編譯程式碼
RUN go build -o myapp

# 使用較小的映像，適用於執行
FROM golang:1.20 AS runner

# 設定工作目錄
WORKDIR /app

# 從 builder 階段複製編譯好的二進制檔案到 runner 階段
COPY --from=builder /app/myapp .

# 執行程式
CMD ["./myapp"]