FROM golang:1.18-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code.
COPY  . .

# Build
RUN go build ./main.go

# This is for documentation purposes only.
# To actually open the port, runtime parameters
# must be supplied to the docker command.
EXPOSE 9001

WORKDIR /

# Run
ENTRYPOINT ["/app/main"]