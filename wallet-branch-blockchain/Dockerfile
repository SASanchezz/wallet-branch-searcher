# Build stage
FROM golang:1.19

WORKDIR /app-go

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Run stage
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./src/main/
EXPOSE 8080

CMD ["./main"]
