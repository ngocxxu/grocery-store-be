# Giai đoạn 1: Dependency stage
FROM golang:1.23.2 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./main.go

# Giai đoạn 2: Final stage
# FROM alpine:3.20.0
# COPY --from=builder /app/main .

# Env
ENV POSTGRES_HOST=postgres
ENV POSTGRES_DB_URL=${POSTGRES_DB_URL}
ENV POSTGRES_USER=${POSTGRES_USER}
ENV POSTGRES_PASSWORD=${POSTGRES_PASSWORD}
ENV POSTGRES_DB=${POSTGRES_DB}

EXPOSE 8030
CMD ["./main"]
