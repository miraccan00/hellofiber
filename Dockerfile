# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.25 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
ARG VERSION=dev
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.version=${VERSION}" -o /app/app ./cmd/main.go

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/static-debian12:nonroot AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/app /app

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app"]
