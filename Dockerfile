FROM golang:1.12-alpine

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app/go-sample-app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/go-sample-app .


# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./out/go-sample-app"]

# # Giai đoạn 1: Dependency stage
# FROM golang:1.23.2 AS builder
# # Env
# ARG POSTGRES_DB_URL
# ARG POSTGRES_USER 
# ARG POSTGRES_PASSWORD
# ARG POSTGRES_DB

# ENV POSTGRES_HOST=postgres
# ENV POSTGRES_DB_URL=${POSTGRES_DB_URL}
# ENV POSTGRES_USER=${POSTGRES_USER}
# ENV POSTGRES_PASSWORD=${POSTGRES_PASSWORD}
# ENV POSTGRES_DB=${POSTGRES_DB}

# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download
# COPY . .
# RUN go build -o main ./main.go

# # Giai đoạn 2: Final stage
# FROM alpine:3.20.0
# RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*
# WORKDIR /app
# COPY --from=builder /app/main .
# COPY .env /app

# EXPOSE 8030
# CMD ["./main"]
