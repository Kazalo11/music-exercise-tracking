#build stage
FROM golang:1.23.1-bullseye AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN  go build -o server ./cmd/app

#run stage
FROM alpine:latest
COPY --from=build /app/server /server
EXPOSE 8080
CMD ["/server"]

