In the following, different reverse proxy setups are given (in alphabetical order) to make Crow work behind a reverse proxy:

## Apache

The following modules are required:

- `proxy`
- `proxy_http`

```txt
ProxyPreserveHost On

RequestHeader set X-Forwarded-Proto "https"

ProxyPass / http://127.0.0.1:8000/
ProxyPassReverse / http://127.0.0.1:8000/
```

## Caddy

```txt
# WebUI and API
crow.example.com {
  reverse_proxy crow-server:8000
}

# expose gRPC
crow-agent.example.com {
  reverse_proxy h2c://crow-server:9000
}
```

!!!info
    The above configuration shows how to create reverse-proxies for server and agent communication. If the agent is configured to use SSL, do not forget to enable `CROW_GRPC_SECURE`.

## Nginx

```conf
server {
    listen 80;
    server_name crow.example.com;

    location / {
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header Host $http_host;

        proxy_pass http://127.0.0.1:8000;
        proxy_redirect off;
        proxy_http_version 1.1;
        proxy_buffering off;

        chunked_transfer_encoding off;
    }
}
```

!!!info
    This does not cover an SSL configuration with NGINX but only shows how to properly forward incoming requests through NGINX to Crow.

## Ngrok

Start `ngrok` using the designed Crow port, e.g. `ngrok http 8000`.
This will return a response similar to the following

Set `CROW_HOST` to the returned URL and (re)start Crow.

## Tunnelmole

Start tunnelmole using the designed Crow port, e.g. `tmole 8000`.
This will return a response similar to the following

```txt
tmole 8000
http://bvdo5f-ip-49-183-170-144.tunnelmole.net is forwarding to localhost:8000
https://bvdo5f-ip-49-183-170-144.tunnelmole.net is forwarding to localhost:8000
```

Set `CROW_HOST` to the returned URL (e.g. exx.tunnelmole.net) and (re)start Crow.

## Traefik

To install the crow server behind a Traefik load balancer, both the http and the gRPC ports must be exposed and configured.

Here is a comprehensive example, which uses `traefik` running via docker compose and applies TLS termination and automatic redirection from http to https.

```yaml
services:
  server:
    image: ghcr.io/crowci/crow-server:latest
    environment:
      # Crow settings ...

    networks:
      - dmz # externally defined network, so that traefik can connect to the server
    volumes:
      - crow-server-data:/var/lib/crow/

    deploy:
      labels:
        - traefik.enable=true

        # web server
        - traefik.http.services.crow-service.loadbalancer.server.port=8000

        - traefik.http.routers.crow-secure.rule=Host(`cd.your-domain.com`)
        - traefik.http.routers.crow-secure.tls=true
        - traefik.http.routers.crow-secure.tls.certresolver=letsencrypt
        - traefik.http.routers.crow-secure.entrypoints=web-secure
        - traefik.http.routers.crow-secure.service=crow-service

        - traefik.http.routers.crow.rule=Host(`cd.your-domain.com`)
        - traefik.http.routers.crow.entrypoints=web
        - traefik.http.routers.crow.service=crow-service

        - traefik.http.middlewares.crow-redirect.redirectscheme.scheme=https
        - traefik.http.middlewares.crow-redirect.redirectscheme.permanent=true
        - traefik.http.routers.crow.middlewares=crow-redirect@docker

        #  gRPC service
        - traefik.http.services.crow-grpc.loadbalancer.server.port=9000
        - traefik.http.services.crow-grpc.loadbalancer.server.scheme=h2c

        - traefik.http.routers.crow-grpc-secure.rule=Host(`crow-grpc.your-domain.com`)
        - traefik.http.routers.crow-grpc-secure.tls=true
        - traefik.http.routers.crow-grpc-secure.tls.certresolver=letsencrypt
        - traefik.http.routers.crow-grpc-secure.entrypoints=web-secure
        - traefik.http.routers.crow-grpc-secure.service=crow-grpc

        - traefik.http.routers.crow-grpc.rule=Host(`crow-grpc.your-domain.com`)
        - traefik.http.routers.crow-grpc.entrypoints=web
        - traefik.http.routers.crow-grpc.service=crow-grpc

        - traefik.http.middlewares.crow-grpc-redirect.redirectscheme.scheme=https
        - traefik.http.middlewares.crow-grpc-redirect.redirectscheme.permanent=true
        - traefik.http.routers.crow-grpc.middlewares=crow-grpc-redirect@docker

networks:
  dmz:
    external: true
```
