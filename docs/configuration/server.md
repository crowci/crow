## Forge & User configuration

Crow can only be used in combination with a "Forge" (i.e. Git Provider).
The following ones are currently supported:

- GitHub
- GitLab
- Gitea
- Forgejo
- Bitbucket Datacenter

```yaml
CROW_FORGE=github # gitlab,gitea,forgejo,bitbucket
```

As Crow does not have its own user registry, users are provided from the linked forge via OAuth2.
**There is no way to register users manually.**

Registration is closed by default (`CROW_OPEN=false`).
If registration is open (`CROW_OPEN=true`) then every user with an account at the configured forge can login and by this register an account.

!!! warning
    It is highly recommended to restrict login to a specific subset of users, e.g. members of an organization. Otherwise, there is no control about who can register and create repositories to run pipelines. This is also recommended for _private_ Crow instances.
    Org-based restrictions can be done by setting `CROW_ORGS`:

    ```yaml
    CROW_ORGS=org1
    ```

    This would limit the use of Crow to member of this organization, allowing administrative control about users.

## Admin users

```yaml
CROW_ADMIN=johnDoe,janeSmith
```

Besides setting this env var, admins can also be promoted manually in the UI by existing admins (`Settings -> Users -> Edit User`).

## Database

### SQLite

Crow uses a SQLite database stored under `/var/lib/CROW/`.

<!-- FIXME: add link to section -->

When using the docker-compose installation, make sure to persist (and backup) the volume serving this directory.

### Postgres

Postgres >= 11 is currently supported.

```yaml
CROW_DATABASE_DRIVER=postgres
CROW_DATABASE_DATASOURCE=postgres://<user>:<pw>@0.0.0.0:5432/<dbname>?sslmode=disable
```

### MySQL/MariaDB

The minimum version of MySQL/MariaDB required is determined by the `go-sql-driver/mysql` - see [its README](https://github.com/go-sql-driver/mysql#requirements) for more information.

```yaml
CROW_DATABASE_DRIVER=mysql
CROW_DATABASE_DATASOURCE=<user>:<pw>@tcp(0.0.0.0:3306)/<dbname>?parseTime=true
```

### Migration

Crow automatically handles database migrations, i.e. when the schema changes in a new version, the changes are applied during version upgrade.
This also means that reverting to an old version after such a migration is not supported.

## Backend

The backend specifies how and where builds are executed.
There are two production-grade backends:

- `docker`
- `kubernetes`

In addition, the `local` backends exists which allows executing pipelines on the local machine that runs the agent directly.

!!! warning
the `local` backend is aimed at local debugging and development and should not be used in production.

!!! info
Podman is not supported as a backend and cannot be used as a backend engine for the "docker" backend.
This backend needs a dedicated integration using the podman-specific golang libraries.
Contributions welcome!

### Docker

The `docker` backend can be seen as the "default" backend.
It executes each step inside a separate container started on the agent.

#### Backend Options

The following `backend_options` are available:

```yaml
steps:
  - name: test
  [...]
    backend_options:
      docker:
        # same as 'docker run --user'
        user:
```

#### Housekeeping

Crow will not automatically clean up images from the Docker installation on the host.
Doing so needs extra configuration and should be included as a regular system maintenance task.

Here are a few helper commands related to image cleanup:

**Removing all unused images**:

```sh
docker image rm $(docker images --filter "dangling=true" -q --no-trunc)
```

**Removing dangling/unused Crow volumes**:

```sh
docker volume rm $(docker volume ls --filter name=^crow_* --filter dangling=true -q)
```

**Pruning all images older than 30 days if the underlying file system partition has reached a certain size**:

```sh
# prune when disk usage > 300GB
df -k | grep -E '/var/lib/docker$' | awk '{if ($3 > 300000000) system("docker image prune -f -a --filter until=720h")}'
```

### Kubernetes

The Kubernetes backend executes steps inside standalone Pods.
For every pipeline, a temporary PVC is created for the lifetime of the pipeline to transfer files between steps.

Every step is executed in a distinct pod.
This allows every step to be run on a different node, ensuring maximum flexibility with respect to resource distribution.

!!!info
There are plans to add a "one pod per pipeline" setting which executes all steps as containers of a single pod.
This has the advantage of not requiring a RWX volume for file persistence between pods but would at the same time reduce the scheduling flexibility (as it would require the pod to be scheduled on a node that can take the step with the higher resource need).

### Private Registries

Images from private registries require authentication.
This is commonly done by providing a `imagePullSecret` in Kubernetes.

To do so, the env var `CROW_BACKEND_K8S_PULL_SECRET_NAMES` can be used.
It references **existing Kubernetes secrets** which contain registry pull secrets and allows the `agent` to make use of these in a specific namespace (defined via `CROW_BACKEND_K8S_NAMESPACE`) during pipeline creation.

The Kubernetes secret must be of type `kubernetes.io/dockerconfigjson` and look like:

```json
.dockerconfigjson: '{"auths":{"https://index.docker.io/v1/":{"auth":"<auth>","password":"<pw>","username":"<username>"}}}'
```

!!!tip
The secret value can be obtained by issuing a local login to the registry and then extracting the created config value from `~/.docker/config.json`.

### Backend options

The Kubernetes backend allows specifying requests and limits on a per-step basic, most commonly for CPU and memory.

!!!note
Adding a `resources` definition to all steps is essential for the Kubernetes backend to ensure efficient scheduling.

#### General

```yaml
steps:
  - name: 'Kubernetes step'
    image: alpine
    commands:
      - echo "Hello world"
    backend_options:
      kubernetes:
        resources:
          requests:
            memory: 200Mi
            cpu: 100m
          limits:
            memory: 400Mi
            cpu: 1000m
```

The following other common pod settings in Kubernetes are supported in `backend_options` using the same syntax as in Kubernetes:

- nodeSelector
- volumes
- tolerations
- securityContext
- annotations
- labels

If your unfamiliar with some, please check the corresponding upstream documentation in Kubernetes.

#### Noteworthy options and adaptations

In the following, special/uncommon settings are explained in more detail.

- [Limit Ranges](https://kubernetes.io/docs/concepts/policy/limit-range/) can be used to set the limits by per-namespace basis.

- `runtimeClassName` specifies the name of the `RuntimeClass` which will be used to run the Pod.
  If `runtimeClassName` is not set, the default `RuntimeHandler` will be used.
  See the [Kubernetes documentation on runtime classes](https://kubernetes.io/docs/concepts/containers/runtime-class/) for more information.

- To allow annotations and labels, the env var `crow_BACKEND_K8S_POD_ANNOTATIONS_ALLOW_FROM_STEP` and/or `crow_BACKEND_K8S_POD_LABELS_ALLOW_FROM_STEP` must be set to `true`.

- Users of [crio-o](https://github.com/cri-o/cri-o) need to configure the workspace for all pipelines centrally in order ensure correct execution (see [this issue](https://github.com/crow-ci/crow/issues/2510)):

  ```yaml
  workspace:
  base: '/crow'
  path: '/'
  ```

### Local

!!!danger
    The local backend executes pipelines on the local system without any isolation.

!!!note
    This backend does not support 'services'.

This backend should only be used in private environments where the code and pipeline can be trusted.
It should not be used for public instances where unknown users can add new repositories and execute code.
The agent should not run as a privileged user (root).

The backend will use a random directory in `$TMPDIR` to store the cloned code and execute commands.

In order to use this backend, a binary of the agent must be built and made executable on the respective server.

The backend can be used by setting

```yaml
CROW_BACKEND=local
```

The entrypoint of the `image:` definition in the pipeline is used to declare the shell, such as `bash` or `fish`.

Plugins are supported and executed as expected.
In the context of the `local` backend, plugins are effectively executable binaries, which can be located using their name if available in `$PATH`, or through an absolute path:

```yaml
steps:
  - name: build
    image: /usr/bin/tree
```

## Metrics

### Endpoint

Crow exposes a Prometheus-compatible `/metrics` endpoint protected by the environment variable `CROW_PROMETHEUS_AUTH_TOKEN`.

```yaml
global:
  scrape_interval: 60s

scrape_configs:
  - job_name: 'crow'
    bearer_token: dummyToken...

    static_configs:
      - targets: ['crow.domain.com']
```

### Authorization

To access the endpoint, a user-level API token must be created by an admin and handed over in the Prometheus configuration file as a bearer token:

```diff
 global:
   scrape_interval: 60s

 scrape_configs:
   - job_name: 'crow'
+    bearer_token: <token>

     static_configs:
        - targets: ['crow.domain.com']
```

Alternatively, the token can be read from a file:

```diff
 global:
   scrape_interval: 60s

 scrape_configs:
   - job_name: 'crow'
+    bearer_token_file: /etc/secrets/crow-monitoring-token

     static_configs:
        - targets: ['crow.domain.com']
```

### Reference

The following metrics are exposed:

```
# HELP crow_pipeline_count Pipeline count.
# TYPE crow_pipeline_count counter
crow_pipeline_count{branch="main",pipeline="total",repo="crow-ci/crow",status="success"} 3
crow_pipeline_count{branch="dev",pipeline="total",repo="crow-ci/crow",status="success"} 3
# HELP crow_pipeline_time Build time.
# TYPE crow_pipeline_time gauge
crow_pipeline_time{branch="main",pipeline="total",repo="crow-ci/crow",status="success"} 116
crow_pipeline_time{branch="dev",pipeline="total",repo="crow-ci/crow",status="success"} 155
# HELP crow_pipeline_total_count Total number of builds.
# TYPE crow_pipeline_total_count gauge
crow_pipeline_total_count 1025
# HELP crow_pending_steps Total number of pending pipeline steps.
# TYPE crow_pending_steps gauge
crow_pending_steps 0
# HELP crow_repo_count Total number of repos.
# TYPE crow_repo_count gauge
crow_repo_count 9
# HELP crow_running_steps Total number of running pipeline steps.
# TYPE crow_running_steps gauge
crow_running_steps 0
# HELP crow_user_count Total number of users.
# TYPE crow_user_count gauge
crow_user_count 1
# HELP crow_waiting_steps Total number of pipeline waiting on deps.
# TYPE crow_waiting_steps gauge
crow_waiting_steps 0
# HELP crow_worker_count Total number of workers.
# TYPE crow_worker_count gauge
crow_worker_count 4
```

## SSL

!!!tip
    Using a [reverse proxy](proxy.md) instead of configuring SSL within Crow provides more flexibility and security and adheres more to "best practices" for deploying applications. 

To enable Crow to terminate TLS directly, the following settings are available:

```sh
CROW_SERVER_CERT=/etc/certs/crow.example.com/server.crt
CROW_SERVER_KEY=/etc/certs/crow.example.com/server.key
```

SSL support is provided using the [ListenAndServeTLS](https://pkg.go.dev/net/http#ListenAndServeTLS) function from the Go standard library.

### Container configuration

In addition to the ports shown in [the docker-compose installation](docker-compose.md), port 443 must be exposed:

```yaml
 services:
   crow-server:
     [...]
     ports:
+      - 80:80
+      - 443:443
       - 9000:9000
```

Additionally, the certificate and key must be mounted and referenced:

```yaml
 services:
   crow-server:
     environment:
+      - CROW_SERVER_CERT=/etc/certs/crow.example.com/server.crt
+      - CROW_SERVER_KEY=/etc/certs/crow.example.com/server.key
     volumes:
+      - /etc/certs/crow.example.com/server.crt:/etc/certs/crow.example.com/server.crt
+      - /etc/certs/crow.example.com/server.key:/etc/certs/crow.example.com/server.key
```

## Logging

The following env vars apply to logging configuration:

```sh
# default 'info'
WOODPECKER_LOG_LEVEL
# default 'stderr'
WOODPECKER_LOG_FILE
# default 'false'
WOODPECKER_DATABASE_LOG
# default 'false'
WOODPECKER_DATABASE_LOG_SQL
# default 'database'
WOODPECKER_LOG_STORE
# not set
WOODPECKER_LOG_STORE_FILE_PATH
```

### Server & Agent

By default, Crow streams the server and agent output to `stderr` and does not persist it.

!!!note
    There is a difference between Server/Agent logs and Pipeline logs.
    The former are the logs describing the application runtime itself, the latter are the logs from the executed pipelines.

Setting `WOODPECKER_LOG_FILE` alongside with `WOODPECKER_LOG_STORE_FILE_PATH` enables file-based logging.

If `WOODPECKER_DATABASE_LOG=true` is set, logs are written into the configured database.

!!! warning
    Database logging might quickly increase the size of the DB, depending on the chosen log level.
    It is recommended to use file-based logging with automatic log-rotation (not configured automatically).

### Pipeline logs

Pipeline execution logs are stored by default alongside a pipeline configuration in the configured database.
Depending on the amount of pipelines and their output, this can fill up the database over time.

An alternative is store logs in an external file which can possibly be auto-rotated:

```sh
WOODPECKER_LOG_STORE=file
WOODPECKER_LOG_STORE_FILE_PATH=<path>
```

!!! info
    Support for external S3-based logging is planned.

## External configuration API

### Introduction

To provide additional management and preprocessing capabilities for pipeline configurations, Crow supports an HTTP API which can be enabled to call an external config service.

Before the run or restart of any pipeline Crow will make a POST request to an external HTTP API sending the current repository, build information and all current config files retrieved from the repository.
The external API can then send back new pipeline configurations that will be used immediately or respond with HTTP 204 to tell the system to use the existing configuration.

Every request sent by Crow is signed using a http-signature by a private key (ed25519) generated on the first start of the Crow server.
You can get the public key for the verification of the http-signature from http(s)://your-woodpecker-server/api/signature/public-key.

A simplistic example configuration service can be found [here](https://github.com/woodpecker-ci/example-config-service).

!!! warning
    The external config service must be trusted as it is receiving secret information about the repository and pipeline and has the ability to change pipeline configs that could potentially execute malicious tasks.

### Configuration

```sh
CROW_CONFIG_SERVICE_ENDPOINT=https://example.com/ciconfig
```

Exemplary request from Crow:

```json
{
  "repo": {
    "id": 100,
    "uid": "",
    "user_id": 0,
    "namespace": "",
    "name": "woodpecker-test-pipe",
    "slug": "",
    "scm": "git",
    "git_http_url": "",
    "git_ssh_url": "",
    "link": "",
    "default_branch": "",
    "private": true,
    "visibility": "private",
    "active": true,
    "config": "",
    "trusted": false,
    "protected": false,
    "ignore_forks": false,
    "ignore_pulls": false,
    "cancel_pulls": false,
    "timeout": 60,
    "counter": 0,
    "synced": 0,
    "created": 0,
    "updated": 0,
    "version": 0
  },
  "pipeline": {
    "author": "myUser",
    "author_avatar": "https://myforge.com/avatars/d6b3f7787a685fcdf2a44e2c685c7e03",
    "author_email": "my@email.com",
    "branch": "main",
    "changed_files": ["some-file-name.txt"],
    "commit": "2fff90f8d288a4640e90f05049fe30e61a14fd50",
    "created_at": 0,
    "deploy_to": "",
    "enqueued_at": 0,
    "error": "",
    "event": "push",
    "finished_at": 0,
    "id": 0,
    "link_url": "https://myforge.com/myUser/woodpecker-testpipe/commit/2fff90f8d288a4640e90f05049fe30e61a14fd50",
    "message": "test old config\n",
    "number": 0,
    "parent": 0,
    "ref": "refs/heads/main",
    "refspec": "",
    "clone_url": "",
    "reviewed_at": 0,
    "reviewed_by": "",
    "sender": "myUser",
    "signed": false,
    "started_at": 0,
    "status": "",
    "timestamp": 1645962783,
    "title": "",
    "updated_at": 0,
    "verified": false
  },
  "netrc": {
    "machine": "https://example.com",
    "login": "user",
    "password": "password"
  }
}
```

Exemplary response:

```json
{
  "configs": [
    {
      "name": "central-override",
      "data": "steps:\n  - name: backend\n    image: alpine\n    commands:\n      - echo \"Hello there from ConfigAPI\"\n"
    }
  ]
}
```
