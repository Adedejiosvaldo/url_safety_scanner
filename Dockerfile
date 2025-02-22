# FROM golang:1.23-alpine AS builder
# WORKDIR /app
# COPY . .
# RUN go mod download
# # RUN go build -o ./main ./main.go

# RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# FROM alpine:latest AS runner
# WORKDIR /app
# COPY --from=builder /app/example-golang .
# EXPOSE 8080
# ENTRYPOINT ["./main"]



FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build with updated Go version
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

EXPOSE 8080

CMD ["./main"]
