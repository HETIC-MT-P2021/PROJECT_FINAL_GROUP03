FROM golang:1.15-alpine

# ----- SETUP -----

# Enable Go modules
ENV GO111MODULE=on

# Set the current working with go absolute path
WORKDIR /go/src/github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/frontApi

# ----- INSTALL -----

# Copy go.mod + go.sum for install before full copy
COPY frontApi/go.mod .

COPY frontApi/go.sum .

# Download all dependencies
RUN go mod download -x

# ----- COPY + RUN -----

# Copy the source from the current directory to the container
COPY frontApi/ .

# Install 'air' live-reload go module
RUN go get -u github.com/cosmtrek/air

# Use the excutable
ENTRYPOINT ["/go/bin/air"]
