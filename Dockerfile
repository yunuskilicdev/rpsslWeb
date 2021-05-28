FROM golang:1.16-alpine

RUN apk add --no-cache git
WORKDIR src/github.com/yunuskilicdev/rpssl
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/rpssl-app .

# This container exposes port 8080 to the ouside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./out/rpssl-app"]