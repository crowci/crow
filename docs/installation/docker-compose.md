<!-- markdownlint-disable MD041 -->
This exemplary docker-compose setup shows the deployment of a Crow instance connected to GitHub (`CROW_GITHUB=true`).
If you are using another forge, please change this including the respective secret settings.

It creates persistent volumes for the server and agent config directories.
The bundled SQLite DB is stored in `/var/lib/crow` and is the most important part to be persisted as it holds all users and repository information.

The server uses the default port 8000 and gets exposed to the host here, so Crow can be accessed through this port on the host or by a reverse proxy sitting in front of it.

!!! information

    The deployment of `crow-autoscaler` is optional.
    Only a brief skeleton is added below for demonstration purposes.
    Please check [the Autoscaler documentation](../configuration/autoscaler.md) and decide whether or not the Autoscaler is a good fit for your environment.
    It also contains the information how to configure the Autoscaler in detail.

```yaml
services:
  crow-server:
    image: ghcr.io/crowci/crow-server:<tag>
    ports:
      - 8000:8000
    volumes:
      - crow-server:/var/lib/crow/
    environment:
      - CROW_OPEN=true
      - CROW_HOST=${CROW_HOST}
      - CROW_GITHUB=true
      - CROW_GITHUB_CLIENT=${CROW_GITHUB_CLIENT}
      - CROW_GITHUB_SECRET=${CROW_GITHUB_SECRET}
      - CROW_AGENT_SECRET=${CROW_AGENT_SECRET}

  crow-agent:
    image: ghcr.io/crowci/crow-agent:<tag>
    restart: always
    depends_on:
      - crow-server
    volumes:
      - crow-agent:/etc/crow
      - /var/run/docker.sock:/var/run/docker.sock
    environment:
      - CROW_SERVER=crow-server:9000
      - CROW_AGENT_SECRET=${CROW_AGENT_SECRET}

  # OPTIONAL
  crow-autoscaler:
    image: ghcr.io/crowci/crow-autoscaler:<tag>
    restart: always
    depends_on:
      - crow-server
    environment:
      - CROW_SERVER=https://your-crow-server.com
      - CROW_TOKEN=${CROW_TOKEN}
      - CROW_MIN_AGENTS=0
      - CROW_MAX_AGENTS=3
      - CROW_WORKFLOWS_PER_AGENT=2
      - CROW_GRPC_ADDR=https://grpc.your-crow-server.com
      - CROW_GRPC_SECURE=true
      - CROW_PROVIDER=hetznercloud
      - CROW_HETZNERCLOUD_API_TOKEN=${CROW_HETZNERCLOUD_API_TOKEN}
```
