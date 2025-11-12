# BASE IMAGE
FROM golang:1.25-alpine

# DEPENDENCIES
RUN apk add --no-cache gcc musl-dev sqlite

# WORK DIRECTORY
WORKDIR /app

# DEPENDENCY FILES
COPY go.mod go.sum ./
RUN go mod download

# SOURCE CODE
COPY . .

# BUILD
RUN CGO_ENABLED=1 go build -o rentroom .

# NETWORK
EXPOSE 8080

# RUNTIME
CMD ["./rentroom"]