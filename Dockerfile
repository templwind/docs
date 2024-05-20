# Development stage with Air for live reloading
FROM golang:1.22-alpine as dev

# Install system dependencies
RUN apk add --no-cache bash git curl gcc musl-dev

# Install Air for live reloading
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

RUN go install github.com/a-h/templ/cmd/templ@latest

# Set the Working Directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go app (optional here, primarily for testing builds)
RUN CGO_ENABLED=1 go build -o /go/bin/app

# Command to run Air for development with live reloading
CMD ["air", "-c", ".air.toml"]

# Production stage using Distroless for minimal footprint
FROM gcr.io/distroless/static-debian11 as prod

# Copy the compiled binary from the development stage
COPY --from=dev /go/bin/app /

# Set the container's binary as the entrypoint
CMD ["/app"]
