# Proposed PR to upstream/main

## Title
feat: add Render deployment support and wall-mode loading fixes

## Body
## Summary
This branch introduces a set of deployment and wall-display improvements on top of commit `e775890`.

At a high level, it does three things:

1. Adds **Render-oriented deployment support** and docs for running a demo deployment.
2. Fixes **template generation/tooling mismatch** in Docker builds by aligning templ invocation with the pinned module/toolchain.
3. Adds/fixes **wall mode behavior** by ensuring route registration and improving wall image loading so repeated images are avoided and loading can continue as users scroll.

## Main changes

### 1) Render deployment + demo defaults
- Adds `render.yaml` blueprint for a web service deployment.
- Updates config handling so runtime port behavior works with Render's dynamic `PORT` semantics when `KIOSK_PORT` is not explicitly set.
- Adds `README.CODEX_RENDER.md` with practical setup notes for Codex and Render deployment.

How it works:
- The service can be provisioned from blueprint and configured with `KIOSK_IMMICH_API_KEY`.
- Runtime port selection is delegated to env-based configuration compatible with Render defaults.

### 2) Docker/templ generation alignment
- Docker build pipeline is updated to a multi-stage frontend+backend build flow.
- Replaces ad-hoc templ install/generate mismatch with generation driven by the Go toolchain/module version (`go tool templ generate` in build flow).
- Includes generated frontend assets from the frontend build stage in the final backend build context.

How it works:
- Frontend assets are built in dedicated Node/pnpm stages and copied into backend build context.
- Backend compile step runs after template generation so generated files are present and version-compatible.

### 3) Wall route and image loading fixes
- Registers/supports the wall route path in router wiring.
- Adjusts wall view/template logic to use a valid dynamic image endpoint.
- Adds load behavior changes to reduce visible image repetition and support continued/infinite loading behavior as users move through the wall.

How it works:
- Image URLs are generated against dynamic endpoints rather than invalid/static forms.
- Client wall loading appends new image sets and avoids simple repeat loops, improving long-running wall displays.

## Commits included
- `ef428480` Adjust Render setup for Immich demo and add Codex build guide
- `9ff303c9` Fix Docker templ generation version mismatch
- `30214a12` Fix wall route registration and host-agnostic wall image URLs
- `e9bc1c41` Fix wall image loading to use valid dynamic image endpoint
- `cd41d837` Fix wall page image repetition and add infinite loading

## Notes for upstream reviewers
- This branch is based on an older upstream point and is currently far behind `upstream/main`; a **rebase/cherry-pick strategy is recommended** before merge.
- Suggested integration order:
  1. Render deployment/docs commits
  2. Docker/templ alignment commit
  3. Wall route/loading commits

## Validation checklist
- [ ] Build image succeeds with current toolchain versions
- [ ] `templ` generation works in CI and local Docker builds
- [ ] Wall route is reachable
- [ ] Wall image URLs resolve correctly in browser network tab
- [ ] Infinite load path in wall view does not re-emit immediate duplicates
