FROM golang:1.22

# Work directory
WORKDIR /go-api

# Installing dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copying all the files
COPY . .

# Setup and start postgreSQL

# Setup and start Redis

# Setup and start Memcache

# Build
RUN make build-all

# Exposing server port
EXPOSE 3100

# Run
CMD ["/build/main"]