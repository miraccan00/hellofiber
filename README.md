# HelloFiber

HelloFiber Go Fiber API

## Quickstart
```bash
go run cmd/main.go
```

Then open:
```
http://localhost:8080/hello
```

Expected response:
```
Hello, World!
```

With `MESSAGE` and `APP_ENV` set (the Helm chart in
[product-helloapi-gitops](https://github.com/miraccan00/product-helloapi-gitops) sets both):
```
Hello from helloapi [DEV] (env=development, version=sha-abc1234)
```
`GET /healthz` returns `ok` (readiness/liveness probe).

## Image

Every push publishes `ghcr.io/miraccan00/hellofiber:sha-<7 chars>` for amd64 and arm64;
`latest` follows `main`.

## Docker
Build and run locally:
```bash
docker build --build-arg VERSION=local -t hellofiber .
docker run --rm -p 8080:8080 hellofiber
```