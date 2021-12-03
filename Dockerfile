# syntax=docker/dockerfile:1
FROM golang:1.15-alpine

ARG APP_PATH
# ----- SETUP -----

# Enable Go modules
ENV GO111MODULE=on

# Set the current working with go absolute path
WORKDIR /go/src/github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/app

# ----- INSTALL -----

# Copy the source from the current directory to the container
COPY $APP_PATH/ .

RUN go mod tidy

# Download all dependencies
RUN go mod download -x

# Install 'air' live-reload go module
RUN go get -u github.com/cosmtrek/air

# Use the excutable
ENTRYPOINT ["/go/bin/air"]
