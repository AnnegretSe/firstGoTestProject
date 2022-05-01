FROM golang:1.18

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY src/go.mod ./
RUN go mod download

COPY src/main.go ./

RUN go build -o /hello_world_http

# Make sure to expose the port the HTTP server is using
EXPOSE 8080

# Run the app binary when we run the container
ENTRYPOINT ["/hello_world_http"]
