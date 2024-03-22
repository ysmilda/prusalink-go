FROM golang:latest AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 

# Move to working directory /build
WORKDIR /build

# Copy the go mod files separately so the build step can be cached
COPY go.mod .
COPY go.sum .

# Download the needed dependencies
RUN go mod download

COPY cli cli

# Build the application
RUN go build -o main ./cli

# Start from scratch, keeps the container as small as possible
FROM scratch

# Copy the compiled program
COPY --from=builder /build/main /

# Execute the compiled program
ENTRYPOINT [ "/main" ]
