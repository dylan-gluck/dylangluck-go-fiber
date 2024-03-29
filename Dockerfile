# syntax=docker/dockerfile:1

FROM golang:1.22

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
ADD *.go ./

RUN mkdir ./blog
COPY blog/ ./blog/

RUN mkdir ./public
COPY public/ ./public/

RUN mkdir ./util
COPY util/ ./util/

RUN mkdir ./views
COPY views/ ./views/

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o ./web

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

# Run
CMD ["./web"]
