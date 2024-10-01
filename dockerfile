FROM golang:1.23.1-alpine as build

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN go build -v -o run_server


FROM alpine:2.6
WORKDIR /app

COPY --from=build /app/templates /app/templates
COPY --from=build /app/run_server /app/run_server

EXPOSE 8080

CMD ["./run_server"]