FROM golang:1.24-bookworm as base

WORKDIR /build

# Copy the go.mod and go.sum files to the /build directory
COPY go.mod ./

# Install dependencies
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the application
RUN go build -o proto.proxy

# Document the port that may need to be published
EXPOSE 8000

# Start the application
CMD ["/build/proto.proxy", "-host=0.0.0.0", "-port=8000"]

