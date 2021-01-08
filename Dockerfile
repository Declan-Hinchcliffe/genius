FROM golang:latest

# set env vars we need
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# set our working directory
WORKDIR /code/genius

# copy go mod and go sum files
COPY go.mod go.sum ./

# download dependencies with go mod
RUN go mod download

# add dlv for debugging
RUN go get github.com/go-delve/delve/cmd/dlv

# copy the code from current dir into work dir of container
COPY . .

# build the binary
RUN go build -o go/bin/api ./cmd/genius/

# expose port to outside world
EXPOSE 9000

# run the binary
CMD go/bin/api

