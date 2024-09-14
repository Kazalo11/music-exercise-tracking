#build stage
FROM golang:1.23.1-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/app

#run stage
FROM alpine
WORKDIR /
COPY --from=build /app/server /server
EXPOSE 8080
CMD ["./server"]

