## STEP

1. Build proxy docker image

```bash
docker build -f caddy.dockerfile -t nuttchai/caddy:1.0.0 .
```

2. Push docker image to docker hub

```bash
docker push nuttchai/caddy:1.0.0
```

3. Import docker image to docker compose file (DON'T FORGET TO USE SAME NETWORK)
