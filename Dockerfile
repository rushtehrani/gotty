# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.9

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/yudai/gotty

# Build the gotty command inside the container.
RUN go install github.com/yudai/gotty

# Run the gotty command by default when the container starts.
ENTRYPOINT /go/bin/gotty

# Document that the service listens on port 8080.
EXPOSE 8080