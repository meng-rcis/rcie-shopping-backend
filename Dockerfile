FROM golang:1.19-alpine as builder

# Set Environment Variable
ENV HOME /app
ENV CGO_ENABLED 0
ENV GOOS linux

# Set the Current Working Directory inside the container
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build app
RUN go build -a -installsuffix cgo -o ./main ./cmd/api

# Multi-stages build
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary & env file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Set Application Environment Path
ENV ENV_PATH .env

# Expose application port
EXPOSE 8080

# Run the application
CMD [ "./main" ]