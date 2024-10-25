# Giai đoạn 1: Dependency stage
FROM golang:1.23.2 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./main.go

# Giai đoạn 2: Final stage
FROM alpine:3.20.0
COPY --from=builder /app/main .

# Thiết lập các biến môi trường
ENV POSTGRES_HOST=postgres
ENV POSTGRES_DB_URL=${POSTGRES_DB_URL}
ENV POSTGRES_USER=${POSTGRES_USER}
ENV POSTGRES_PASSWORD=${POSTGRES_PASSWORD}
ENV POSTGRES_DB=${POSTGRES_DB}

# Thiết lập quyền thực thi cho file chính
RUN chmod +x /main

# Mở cổng 8030
EXPOSE 8030

# Khởi động ứng dụng
CMD ["/main"]
