# Build stage
FROM golang:1.23.2 AS builder
WORKDIR /app
COPY . .
RUN go build -o main ./main.go

# Final stage
FROM alpine:3.20.0
COPY --from=builder /app/main .
EXPOSE 8030
ENV POSTGRES_HOST=postgres
ENV POSTGRES_DB_URL=${POSTGRES_DB_URL}
ENV POSTGRES_USER=${POSTGRES_USER}
ENV POSTGRES_PASSWORD=${POSTGRES_PASSWORD}
ENV POSTGRES_DB=${POSTGRES_DB}
RUN chmod +x ./main
CMD ["./main"]