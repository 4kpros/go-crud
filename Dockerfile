FROM golang:1.22

# Work directory
WORKDIR /app

# Installing dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copying all the files
COPY . .

# Build
RUN make install
RUN make build-all

# Exposing server port
EXPOSE 3100

# Run
CMD ["/go-api/build/main"]