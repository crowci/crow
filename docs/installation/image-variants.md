## Image tags

!!!info
    There is no `latest` tag on purpose to prevent accidental major version upgrades.
    Either use a Semver tag or one of the rolling major/minor version tags.

- `vX.Y.Z`: SemVer tags for specific releases, no entrypoint shell (scratch image)
  - `vX.Y`
  - `vX`
- `vX.Y.Z-alpine`: SemVer tags for specific releases, based on Alpine, rootless for Server and CLI (as of v3.0).
  - `vX.Y-alpine`
  - `vX-alpine`
- `next`: Built from the `main` branch
- `pull_<PR_ID>`: Images built from Pull Request branches.

## Image registries

Images are pushed to the GitHub Container Registry (ghcr.io) of the project.

!!!info
    If there funds available to cover a DockerHub team subscription, images will also be published to DockerHub.

- [crow-server (ghcr.io)](https://hub.docker.com/repository/docker/crowci/crow-server)

- [crow-agent (ghcr.io)](https://hub.docker.com/repository/docker/crowci/crow-agent)

- [crow-cli (ghcr.io)](https://hub.docker.com/repository/docker/crowci/crow-cli)

- [crow-autoscaler (ghcr.io)](https://hub.docker.com/repository/docker/crowci/autoscaler)
