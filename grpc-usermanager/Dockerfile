# syntax=docker/dockerfile:1

FROM golang:1.20

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy server files
COPY server/main.go ./
RUN mkdir proto
COPY proto/*.go ./proto/

# Build go image
RUN go build -o /server
EXPOSE 50051
CMD [ "/server" ]
