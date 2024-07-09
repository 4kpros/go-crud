FROM golang:1.22

# Work directory
WORKDIR /go-api

# Installing dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copying all the files
COPY . .

# Build
RUN make build-all

# Exposing server port
EXPOSE 3100

# Run
CMD ["/build/main"]