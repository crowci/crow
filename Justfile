### Variables
GO_PACKAGES := `go list ./... | grep -v /vendor/`

TARGETOS := `go env GOOS`
TARGETARCH := `go env GOARCH`

BIN_SUFFIX := if TARGETOS == "windows" { ".exe" } else { "" }

DIST_DIR := "dist"

VERSION := "next"
VERSION_NUMBER := "0.0.0"
CI_COMMIT_SHA := `git rev-parse HEAD`

TAGS := ""
LDFLAGS := ""

HAS_GO := `hash go > /dev/null 2>&1 && echo "GO" || echo "NOGO"`

XGO_VERSION := if HAS_GO == "GO" { "go-1.23.x" } else { "" }
CGO_CFLAGS := if HAS_GO == "GO" { `go env CGO_CFLAGS` } else { "" }

### Recipes

## docs

docs-venv-fish:
    python3 -m venv venv
    fish -c 'source venv/bin/activate.fish && pip3 install -r docs/requirements.txt'

docs-mike ARG:
    fish -c 'source venv/bin/activate.fish && mike {{ ARG }}'

docs-serve:
    fish -c 'source venv/bin/activate.fish && mkdocs --version && mkdocs serve'

docs-deploy:
    fish -c 'source venv/bin/activate.fish && mkdocs --version && mkdocs gh-deploy'

## git

cherry-pick COMMIT:
    git fetch crow && git cherry-pick {{ COMMIT }}

## build

[working-directory: 'web']
build-ui:
    pnpm install --frozen-lockfile
    pnpm build

build-agent:
    CGO_ENABLED=0 GOOS={{TARGETOS}} GOARCH={{TARGETARCH}} go build -tags '{{TAGS}}' -ldflags '{{LDFLAGS}}' -o {{DIST_DIR}}/crow-agent{{BIN_SUFFIX}} go.woodpecker-ci.org/woodpecker/v3/cmd/agent

build-cli: ## Build cli
	CGO_ENABLED=0 GOOS={{TARGETOS}} GOARCH={{TARGETARCH}} go build -tags '{{TAGS}}' -ldflags '{{LDFLAGS}}' -o {{DIST_DIR}}/crow-cli{BIN_SUFFIX} go.woodpecker-ci.org/woodpecker/v3/cmd/cli

# build-server
# env PLATFORMS=linux/amd64 just build-server
build-server: build-ui vendor
    just cross-compile-server

vendor:
    go mod tidy
    go mod vendor

# build for specific platform: `env PLATFORMS='linux|amd64' just cross-compile-server`
cross-compile-server:
    #!/usr/bin/env bash
    set -euxo pipefail

    IFS=',' read -ra PLATFORMS_ARRAY <<< "${PLATFORMS}"
    for platform in "${PLATFORMS_ARRAY[@]}"; do
        IFS='/' read -ra PLATFORM_PARTS <<< "$platform"
        TARGETOS="${PLATFORM_PARTS[0]}"
        if [ -n "${PLATFORM_PARTS[1]:-}" ]; then
            TARGETARCH_XGO="${PLATFORM_PARTS[1]//arm64\/v8/arm64}"
            TARGETARCH_XGO="${TARGETARCH_XGO//arm\/v7/arm-7}"
            TARGETARCH="${PLATFORM_PARTS[1]//arm64\/v8/arm64}"
            TARGETARCH="${TARGETARCH//arm\/v7/arm}"
        else
            TARGETARCH_XGO=""
            TARGETARCH=""
        fi
        env TARGETOS=$TARGETOS TARGETARCH=$TARGETARCH TARGETARCH_XGO=$TARGETARCH_XGO just release-server-xgo || exit 1
    done
    tree "{{DIST_DIR}}"

# this should not be called standalone - use 'cross-compile-server' instead
release-server-xgo:
    #!/usr/bin/env bash
    set -euxo pipefail

    echo "Building for:"
    echo "os: ${TARGETOS}"
    echo "arch orgi: ${TARGETARCH}"
    echo "arch (xgo): ${TARGETARCH}"
    # build via xgo
    CGO_CFLAGS="{{CGO_CFLAGS}}" xgo -go {{XGO_VERSION}} -dest {{DIST_DIR}}/server/${TARGETOS}_${TARGETARCH} -tags 'netgo osusergo grpcnotrace {{TAGS}}' -ldflags '-linkmode external {{LDFLAGS}}' -targets ${TARGETOS}/${TARGETARCH} -out crow-server -pkg cmd/server .
    # move binary into subfolder depending on target os and arch
    if [ "$${XGO_IN_XGO:-0}" -eq "1" ]; then \
      echo "inside xgo image"; \
      mkdir -p {{DIST_DIR}}/server/${TARGETOS}_${TARGETARCH}; \
      mv -vf /build/crow-server* {{DIST_DIR}}/server/${TARGETOS}_${TARGETARCH}/crow-server{{BIN_SUFFIX}}; \
    else \
      echo "outside xgo image"; \
      [ -f "{{DIST_DIR}}/server/${TARGETOS}_${TARGETARCH}/crow-server{{BIN_SUFFIX}}" ] && rm -v {{DIST_DIR}}/server/${TARGETOS}_${TARGETARCH}/crow-server{{BIN_SUFFIX}}; \
      mv -v {{DIST_DIR}}/server/${TARGETOS}_${TARGETARCH}/crow-server* {{DIST_DIR}}/server/${TARGETOS}_${TARGETARCH}/crow-server{{BIN_SUFFIX}}; \
    fi

## images

# env PLATFORMS='linux/amd64,linux/arm64' just image-server-alpine
image-server-alpine:
    just cross-compile-server
    echo $GITHUB_PKGS_TOKEN | docker login ghcr.io -u crowci-bot --password-stdin
    docker buildx build --platform $PLATFORMS -t ghcr.io/crowci/crow-server:dev-alpine -f docker/Dockerfile.server.alpine.multiarch.rootless --push .

# env PLATFORMS='linux/amd64,linux/arm64' just image-server
image-server:
    just cross-compile-server
    echo $GITHUB_PKGS_TOKEN | docker login ghcr.io -u crowci-bot --password-stdin
    docker buildx build --platform $PLATFORMS -t ghcr.io/crowci/crow-server:dev -f docker/Dockerfile.server.multiarch.rootless --push .

# env PLATFORMS='linux/amd64,linux/arm64' just image-agent
image-agent:
    echo $GITHUB_PKGS_TOKEN | docker login ghcr.io -u crowci-bot --password-stdin
    docker buildx build --platform $PLATFORMS -t ghcr.io/crowci/crow-agent:dev -f docker/Dockerfile.agent.multiarch --push .

# env PLATFORMS='linux/amd64,linux/arm64' just image-agent-alpine
image-agent-alpine:
    echo $GITHUB_PKGS_TOKEN | docker login ghcr.io -u crowci-bot --password-stdin
    docker buildx build --platform $PLATFORMS -t ghcr.io/crowci/crow-agent:dev-alpine -f docker/Dockerfile.agent.alpine.multiarch --push .