# Codex + Render Setup Notes

This file documents the **extra setup** required in this repository for:

1. Building/running successfully in Codex environments.
2. Deploying correctly on Render.

## 1) Codex build setup

Some Codex environments do not have generated template output ready, and may also miss frontend build artifacts if they were not built recently.

Use the following sequence from the repo root:

```bash
# Install templ codegen CLI (required by backend templates)
go install github.com/a-h/templ/cmd/templ@latest

# Generate Go template files
templ generate

# Build frontend assets
cd frontend
pnpm install
pnpm build
cd ..

# Run tests
go test ./...
```

If `go test ./...` fails with missing packages under `internal/templates/...`, it usually means `templ generate` has not been run in the current environment.

## 2) Render setup

This repository contains a `render.yaml` blueprint configured for Immich's public demo host:

- `KIOSK_IMMICH_URL` defaults to `https://demo.immich.app`.
- `KIOSK_IMMICH_API_KEY` is required and must be set in Render.

### Deploy steps

1. In Render, create a **Blueprint** service from this repository.
2. Open the generated `immich-kiosk-demo` service.
3. Set `KIOSK_IMMICH_API_KEY` in the service environment settings.
4. Trigger a deploy.

### Getting a demo API key

You still need a valid API key from the Immich demo instance account you are using.

General flow in Immich:

1. Sign in to the Immich instance.
2. Go to account settings.
3. Create a new API key.
4. Copy the key into Render as `KIOSK_IMMICH_API_KEY`.

## 3) Port behavior on Render

Render injects a dynamic `PORT` variable for web services. This project supports that automatically (when `KIOSK_PORT` is not explicitly set), so no additional port mapping is required.
