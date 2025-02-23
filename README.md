# API Todo List

## Giới thiệu

Dự án `api-todolist` là một API đơn giản cho ứng dụng quản lý công việc (To-do list), được xây dựng bằng Golang và sử dụng Gin làm framework web. Dự án được đóng gói bằng Docker và triển khai trên Render.

## Cài đặt và Chạy Ứng Dụng

### 1. Cài đặt Docker

Trước tiên, cài đặt Docker trên máy tính của mình (phù hợp với hệ điều hành của máy)

### 2. Clone repository (copy repository về máy)

```bash
git clone https://github.com/kthucdev/api-todolist.git
cd api-todolist
```

### 3. Xây dựng Docker Image

```bash
docker build -t api-todolist .
```

### 4. Chạy ứng dụng trong Docker Container

```bash
docker run -d -p 8080:8080 api-todolist
```
Lệnh này chạy container trong chế độ nền (`-d`) và ánh xạ cổng 8080.

## Triển khai lên Render

### 1. Đăng ký và Đăng nhập Render
- Truy cập [https://render.com](https://render.com) và tạo tài khoản.
- Chọn "New" > "Web Service".

### 2. Kết nối với GitHub
- Kết nối Render với tài khoản GitHub của bạn.
- Chọn kho lưu trữ `api-todolist`.

### 3. Cấu hình triển khai
- **Branch:** `main`
- **Build Command:** `go build -tags netgo -ldflags '-s -w' -o app ./cmd/api`
- **Start Command:** `./app`

### 4. Triển khai
Nhấn "Create Web Service" để Render tự động xây dựng và chạy ứng dụng.