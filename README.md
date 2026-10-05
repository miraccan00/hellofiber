# HelloFiber

HelloFiber Go Fiber API

The product team's app in the *GitOps in Production* series on [miraccanyilmaz.me](https://miraccanyilmaz.me):
Argo CD deploys it through [product-helloapi-gitops](https://github.com/miraccan00/product-helloapi-gitops),
the platform side is [platform-gitops](https://github.com/miraccan00/platform-gitops), and each article's
runnable lab is a folder in [blog-wiki](https://github.com/miraccan00/blog-wiki).

| Article | Branch used |
|---|---|
| Argo CD in HA, Explained by Breaking It · [EN](https://miraccanyilmaz.me/en/blog/argocd-ha-app-of-apps/) · [TR](https://miraccanyilmaz.me/blog/argocd-ha-app-of-apps/) | `blog-04` |
| Argo CD SSO Integration: OIDC and RBAC with ZITADEL · [EN](https://miraccanyilmaz.me/en/blog/argocd-sso-zitadel/) · [TR](https://miraccanyilmaz.me/blog/argocd-sso-zitadel/) | `blog-04` (unchanged) |
| Moving Secrets into Vault: From base64 in Git to Vault and ESO Without Downtime · [EN](https://miraccanyilmaz.me/en/blog/vault-eso-secret-migration/) · [TR](https://miraccanyilmaz.me/blog/vault-eso-secret-migration/) | `blog-04` (unchanged) |

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