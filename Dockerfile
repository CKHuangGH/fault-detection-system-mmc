FROM golang
RUN mkdir -p /app
# 設定工作目錄
WORKDIR /app

# 複製程式碼到容器內
COPY fault_detection.go go.mod go.sum .

RUN go mod download
# 編譯程式碼
RUN go build -o myapp

# 使用較小的映像，適用於執行
# FROM golang:1.20 AS runner

ENTRYPOINT ["./myapp"]
# # 設定工作目錄
# WORKDIR /app

# # 從 builder 階段複製編譯好的二進制檔案到 runner 階段
# COPY --from=builder /app/myapp .

# 執行程式
# CMD ["./myapp"]