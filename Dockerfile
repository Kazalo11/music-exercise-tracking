#build stage
FROM golang:1.22.7-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/app


#run stage
FROM debian:bullseye-slim
COPY --from=builder /app/server /server
EXPOSE 8080
CMD ["/server"]