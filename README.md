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

## Docker
Build and run locally:
```bash
docker build -t hellofiber .
docker run --rm -p 8080:8080 hellofiber
```