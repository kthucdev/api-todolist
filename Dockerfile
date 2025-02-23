# Stage 1: Build ứng dụng Go
FROM golang:1.23 AS builder

WORKDIR /app

# Copy module files và download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy toàn bộ source code
COPY . .

# Build file binary từ cmd/api/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/api

# Kiểm tra file sau khi build
RUN ls -lah /app

# Stage 2: Tạo một image nhẹ để chạy app
FROM scratch

WORKDIR /root/

# Copy binary từ builder stage
COPY --from=builder /app/main .

# Mở cổng 8080
EXPOSE 8080

# Chạy ứng dụng
CMD ["./main"]
