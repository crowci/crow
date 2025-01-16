### Variables

GO_PACKAGES := `go list ./... | grep -v /vendor/`
TARGETOS := `go env GOOS`
TARGETARCH := `go env GOARCH`
BIN_SUFFIX := if TARGETOS == "windows" { ".exe" } else { "" }
DIST_DIR := "dist"
VERSION := "dev"
VERSION_NUMBER := "0.0.0"
CI_COMMIT_SHA := `git rev-parse HEAD`
TAGS := ""
STATIC_BUILD := "true"

# Conditional assignment for LDFLAGS if STATIC_BUILD is true
# FIXME: https://github.com/casey/just/issues/11

LDFLAGS := if STATIC_BUILD == "true" { "-s -w -extldflags '-static'" } else { "" }

# only used to compile server

CGO_ENABLED := "1"
HAS_GO := `hash go > /dev/null 2>&1 && echo "GO" || echo "NOGO"`
XGO_VERSION := if HAS_GO == "GO" { "go-1.23.x" } else { "" }
CGO_CFLAGS := if HAS_GO == "GO" { `go env CGO_CFLAGS` } else { "" }

### Recipes
## general

fmt:
    find . -name '*.go' -not -path './vendor/*' -exec gci write {} \;

test: test-agent test-server test-server-datastore test-cli test-lib

generate-openapi: install-tools
    go run github.com/swaggo/swag/cmd/swag fmt
    CGO_ENABLED=0 go generate cmd/server/openapi.go

install-tools:
    @hash golangci-lint > /dev/null 2>&1 || go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
    @hash gofumpt > /dev/null 2>&1 || go install mvdan.cc/gofumpt@latest
    @hash addlicense > /dev/null 2>&1 || go install github.com/google/addlicense@latest
    @hash mockery > /dev/null 2>&1 || go install github.com/vektra/mockery/v2@latest
    @hash protoc-gen-go > /dev/null 2>&1 || go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    @hash protoc-gen-go-grpc > /dev/null 2>&1 || go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

## test

test-agent:
    go test -race -cover -coverprofile agent-coverage.out -timeout 60s -tags 'test {{ TAGS }}' github.com/crowci/crow/v3/cmd/agent github.com/crowci/crow/v3/agent/...

test-server:
    go test -race -cover -coverprofile server-coverage.out -timeout 60s -tags 'test {{ TAGS }}' github.com/crowci/crow/v3/cmd/server `go list github.com/crowci/crow/v3/server/... | grep -v '/store'`

test-cli:
    go test -race -cover -coverprofile cli-coverage.out -timeout 60s -tags 'test {{ TAGS }}' github.com/crowci/crow/v3/cmd/cli github.com/crowci/crow/v3/cli/...

test-server-datastore:
    go test -timeout 300s -tags 'test {{ TAGS }}' -run TestMigrate github.com/crowci/crow/v3/server/store/...
    go test -race -timeout 100s -tags 'test {{ TAGS }}' -skip TestMigrate github.com/crowci/crow/v3/server/store/...

test-server-datastore-coverage:
    go test -race -cover -coverprofile datastore-coverage.out -timeout 300s -tags 'test {{ TAGS }}' github.com/crowci/crow/v3/server/store/...

[working-directory('web')]
test-ui:
    pnpm install --frozen-lockfile
    pnpm run lint
    pnpm run format:check
    pnpm run typecheck
    pnpm run test

test-lib:
    go test -race -cover -coverprofile coverage.out -timeout 60s -tags 'test {{ TAGS }}' `go list ./... | grep -v '/cmd\|/agent\|/cli\|/server'`

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

lint:
    golangci-lint run

[working-directory('web')]
build-ui:
    pnpm install --frozen-lockfile
    pnpm build

build-agent:
    CGO_ENABLED=0 GOOS={{ TARGETOS }} GOARCH={{ TARGETARCH }} go build -tags '{{ TAGS }}' -ldflags '{{ LDFLAGS }}' -o {{ DIST_DIR }}/crow-agent{{ BIN_SUFFIX }} github.com/crowci/crow/v3/cmd/agent

build-cli:
    CGO_ENABLED=0 GOOS={{ TARGETOS }} GOARCH={{ TARGETARCH }} go build -tags '{{ TAGS }}' -ldflags '{{ LDFLAGS }}' -o {{ DIST_DIR }}/crow-cli{{ BIN_SUFFIX }} github.com/crowci/crow/v3/cmd/cli

# build-server

# env PLATFORMS=linux/amd64 just build-server
build-server: build-ui vendor
    just cross-compile-server

vendor:
    go mod tidy
    go mod vendor

# build for specific platform: env PLATFORMS='linux|arm64/v8' just cross-compile-server

# build for local platform: just cross-compile-server
cross-compile-server:
    #!/usr/bin/env bash
    set -euxo pipefail

    PLATFORMS="${PLATFORMS:-{{ TARGETOS }}/{{ TARGETARCH }}}"
    IFS=';' read -ra PLATFORMS_ARRAY <<< "${PLATFORMS}"
    for platform in "${PLATFORMS_ARRAY[@]}"; do
        IFS='|' read -ra PLATFORM_PARTS <<< "$platform"
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
    tree "{{ DIST_DIR }}"

check-xgo:
    hash xgo > /dev/null 2>&1 || go install src.techknowlogick.com/xgo@latest

# this should not be called directly - use 'cross-compile-server' instead
release-server-xgo: check-xgo
    @echo "------------------"
    @echo "Building for:"
    @echo "- os: ${TARGETOS}"
    @echo "- arch (xgo): ${TARGETARCH}"
    @echo "------------------"
    CGO_CFLAGS="{{ CGO_CFLAGS }}" xgo -go {{ XGO_VERSION }} -dest {{ DIST_DIR }}/server/${TARGETOS}_${TARGETARCH} -tags 'netgo osusergo grpcnotrace {{ TAGS }}' -ldflags '-linkmode external {{ LDFLAGS }} -X github.com/crowci/crow/v3/version.Version={{ VERSION }}' -targets ${TARGETOS}/${TARGETARCH} -out crow-server -pkg cmd/server .
    # move binary into subfolder depending on target os and arch
    if [ "${XGO_IN_XGO:-0}" -eq "1" ]; then \
      echo "inside xgo image"; \
      mkdir -p {{ DIST_DIR }}/server/${TARGETOS}_${TARGETARCH}; \
      mv -vf /build/crow-server* {{ DIST_DIR }}/server/${TARGETOS}_${TARGETARCH}/crow-server{{ BIN_SUFFIX }}; \
    else \
      echo "outside xgo image"; \
      [ -f "{{ DIST_DIR }}/server/${TARGETOS}_${TARGETARCH}/crow-server{{ BIN_SUFFIX }}" ] && rm -v {{ DIST_DIR }}/server/${TARGETOS}_${TARGETARCH}/crow-server{{ BIN_SUFFIX }}; \
      mv -v {{ DIST_DIR }}/server/${TARGETOS}_${TARGETARCH}/crow-server* {{ DIST_DIR }}/server/${TARGETOS}_${TARGETARCH}/crow-server{{ BIN_SUFFIX }}; \
    fi

## agent

release-agent:
    GOOS=linux   GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '{{ LDFLAGS }}' -tags 'grpcnotrace {{ TAGS }}' -o {{ DIST_DIR }}/agent/linux_amd64/crow-agent       github.com/crowci/crow/v3/cmd/agent
    GOOS=linux   GOARCH=arm64 CGO_ENABLED=0 go build -ldflags '{{ LDFLAGS }}' -tags 'grpcnotrace {{ TAGS }}' -o {{ DIST_DIR }}/agent/linux_arm64/crow-agent       github.com/crowci/crow/v3/cmd/agent
    GOOS=linux   GOARCH=arm   CGO_ENABLED=0 go build -ldflags '{{ LDFLAGS }}' -tags 'grpcnotrace {{ TAGS }}' -o {{ DIST_DIR }}/agent/linux_arm/crow-agent         github.com/crowci/crow/v3/cmd/agent
    GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '{{ LDFLAGS }}' -tags 'grpcnotrace {{ TAGS }}' -o {{ DIST_DIR }}/agent/windows_amd64/crow-agent.exe github.com/crowci/crow/v3/cmd/agent
    GOOS=darwin  GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '{{ LDFLAGS }}' -tags 'grpcnotrace {{ TAGS }}' -o {{ DIST_DIR }}/agent/darwin_amd64/crow-agent      github.com/crowci/crow/v3/cmd/agent
    GOOS=darwin  GOARCH=arm64 CGO_ENABLED=0 go build -ldflags '{{ LDFLAGS }}' -tags 'grpcnotrace {{ TAGS }}' -o {{ DIST_DIR }}/agent/darwin_arm64/crow-agent      github.com/crowci/crow/v3/cmd/agent
    # tar binary files
    tar -cvzf {{ DIST_DIR }}/crow-agent_linux_amd64.tar.gz   -C {{ DIST_DIR }}/agent/linux_amd64   crow-agent
    tar -cvzf {{ DIST_DIR }}/crow-agent_linux_arm64.tar.gz   -C {{ DIST_DIR }}/agent/linux_arm64   crow-agent
    tar -cvzf {{ DIST_DIR }}/crow-agent_linux_arm.tar.gz     -C {{ DIST_DIR }}/agent/linux_arm     crow-agent
    tar -cvzf {{ DIST_DIR }}/crow-agent_darwin_amd64.tar.gz  -C {{ DIST_DIR }}/agent/darwin_amd64  crow-agent
    tar -cvzf {{ DIST_DIR }}/crow-agent_darwin_arm64.tar.gz  -C {{ DIST_DIR }}/agent/darwin_arm64  crow-agent
    # zip binary files
    rm -f  {{ DIST_DIR }}/crow-agent_windows_amd64.zip
    zip -j {{ DIST_DIR }}/crow-agent_windows_amd64.zip          {{ DIST_DIR }}/agent/windows_amd64/crow-agent.exe

## cli

release-cli:
    GOOS=linux   GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '{{ LDFLAGS }}' -o {{ DIST_DIR }}/cli/linux_amd64/crow-cli       github.com/crowci/crow/v3/cmd/cli
    GOOS=linux   GOARCH=arm64 CGO_ENABLED=0 go build -ldflags '{{ LDFLAGS }}' -o {{ DIST_DIR }}/cli/linux_arm64/crow-cli       github.com/crowci/crow/v3/cmd/cli
    GOOS=linux   GOARCH=arm   CGO_ENABLED=0 go build -ldflags '{{ LDFLAGS }}' -o {{ DIST_DIR }}/cli/linux_arm/crow-cli         github.com/crowci/crow/v3/cmd/cli
    GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '{{ LDFLAGS }}' -o {{ DIST_DIR }}/cli/windows_amd64/crow-cli.exe github.com/crowci/crow/v3/cmd/cli
    GOOS=darwin  GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '{{ LDFLAGS }}' -o {{ DIST_DIR }}/cli/darwin_amd64/crow-cli      github.com/crowci/crow/v3/cmd/cli
    GOOS=darwin  GOARCH=arm64 CGO_ENABLED=0 go build -ldflags '{{ LDFLAGS }}' -o {{ DIST_DIR }}/cli/darwin_arm64/crow-cli      github.com/crowci/crow/v3/cmd/cli
    # tar binary files
    tar -cvzf {{ DIST_DIR }}/crow-cli_linux_amd64.tar.gz   -C {{ DIST_DIR }}/cli/linux_amd64   crow-cli
    tar -cvzf {{ DIST_DIR }}/crow-cli_linux_arm64.tar.gz   -C {{ DIST_DIR }}/cli/linux_arm64   crow-cli
    tar -cvzf {{ DIST_DIR }}/crow-cli_linux_arm.tar.gz     -C {{ DIST_DIR }}/cli/linux_arm     crow-cli
    tar -cvzf {{ DIST_DIR }}/crow-cli_darwin_amd64.tar.gz  -C {{ DIST_DIR }}/cli/darwin_amd64  crow-cli
    tar -cvzf {{ DIST_DIR }}/crow-cli_darwin_arm64.tar.gz  -C {{ DIST_DIR }}/cli/darwin_arm64  crow-cli
    # zip binary files
    rm -f  {{ DIST_DIR }}/crow-cli_windows_amd64.zip
    zip -j {{ DIST_DIR }}/crow-cli_windows_amd64.zip          {{ DIST_DIR }}/cli/windows_amd64/crow-cli.exe

# Build tar archive
build-tarball:
    mkdir -p {{ DIST_DIR }} && tar chzvf {{ DIST_DIR }}/crow-src.tar.gz \
      --exclude="*.exe" \
      --exclude="./.pnpm-store" \
      --exclude="node_modules" \
      --exclude="./dist" \
      --exclude="./data" \
      --exclude="./build" \
      --exclude="./.git" \
      .

# # Create checksums for all release files
release-checksums:
    (cd {{ DIST_DIR }}/; sha256sum *.* > checksums.txt)

## images
# platforms must be handed over via this syntax for the underlying cross-compile-server step which applies some string splitting on a list of items

# env PLATFORMS='linux|amd64;linux|arm64' just image-server
image-server:
    just cross-compile-server
    echo $GITHUB_PKGS_TOKEN | docker login ghcr.io -u crowci-bot --password-stdin
    FIXED_PLATFORMS=$(echo $PLATFORMS | sed "s/|/\//g; s/;/,/g") && docker buildx build --platform $FIXED_PLATFORMS -t ghcr.io/crowci/crow-server:dev -f docker/Dockerfile.server.multiarch.rootless --push .

# env PLATFORMS='linux/amd64,linux/arm64' just image-server-alpine
image-server-alpine:
    just cross-compile-server
    echo $GITHUB_PKGS_TOKEN | docker login ghcr.io -u crowci-bot --password-stdin
    FIXED_PLATFORMS=$(echo $PLATFORMS | sed "s/|/\//g; s/;/,/g") && docker buildx build --platform $PLATFORMS -t ghcr.io/crowci/crow-server:dev-alpine -f docker/Dockerfile.server.alpine.multiarch.rootless --push .

# env PLATFORMS='linux/amd64,linux/arm64' just image-agent
image-agent:
    echo $GITHUB_PKGS_TOKEN | docker login ghcr.io -u crowci-bot --password-stdin
    docker buildx build --platform $PLATFORMS -t ghcr.io/crowci/crow-agent:dev -f docker/Dockerfile.agent.multiarch --push .

# env PLATFORMS='linux/amd64,linux/arm64' just image-agent-alpine
image-agent-alpine:
    echo $GITHUB_PKGS_TOKEN | docker login ghcr.io -u crowci-bot --password-stdin
    docker buildx build --platform $PLATFORMS -t ghcr.io/crowci/crow-agent:dev-alpine -f docker/Dockerfile.agent.alpine.multiarch --push .

# env PLATFORMS='linux/amd64,linux/arm64' just image-cli
image-cli:
    echo $GITHUB_PKGS_TOKEN | docker login ghcr.io -u crowci-bot --password-stdin
    docker buildx build --platform $PLATFORMS -t ghcr.io/crowci/crow-cli:dev -f docker/Dockerfile.cli.multiarch.rootless --push .

# env PLATFORMS='linux/amd64,linux/arm64' just image-cli-alpine
image-cli-alpine:
    echo $GITHUB_PKGS_TOKEN | docker login ghcr.io -u crowci-bot --password-stdin
    docker buildx build --platform $PLATFORMS -t ghcr.io/crowci/crow-cli:dev-alpine -f docker/Dockerfile.cli.alpine.multiarch.rootless --push .
