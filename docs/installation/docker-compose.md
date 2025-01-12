# docker-compose

This exemplary docker-compose setup shows the deployment of a Crow instance connected to GitHub (`CROW_GITHUB=true`).
If you are using another forge, please change this including the respective secret settings.

It creates persistent volumes for the server and agent config directories.
The bundled SQLite DB is stored in `/var/lib/crow` and is the most important part to be persisted as it holds all users and repository information.

The server uses the default port 8000 and gets exposed to the host here, so Crow can be accessed through this port on the host or by a reverse proxy sitting in front of it.

```yaml
services:
  crow-server:
    image: ghcr.io/crowci/crow-server:latest
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

  agent:
    image: ghcr.io/crowci/crow-agent:latest
    restart: always
    depends_on:
      - crow-server
    volumes:
      - crow-agent:/etc/crow
      - /var/run/docker.sock:/var/run/docker.sock
    environment:
      - CROW_SERVER=crow-server:9000
      - CROW_AGENT_SECRET=${CROW_AGENT_SECRET}
```
