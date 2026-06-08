#!/bin/sh

REMOTE_REPOSITORY=zot-tailscale.tail768005.ts.net:5000

podman machine start || echo "[podman already started]"

podman build . -t cff4ha:development --no-cache
podman tag cff4ha:development $REMOTE_REPOSITORY/cff4ha:development
podman push $REMOTE_REPOSITORY/cff4ha:development

kubectl rollout restart deployment cff4ha-cff4ha -n cff4ha --insecure-skip-tls-verify
