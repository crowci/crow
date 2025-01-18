---
hide:
  - toc
---

# Server

| Name                                                | Description                                                                                                                              | Value     |
| --------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | --------- |
| `WOODPECKER_BACKEND_K8S_NAMESPACE`                  | The k8s namespace to execute pipelines in                                                                                                | `crow`    |
| `WOODPECKER_BACKEND_K8S_POD_ANNOTATIONS`            | Additional annotations to apply to worker Pods.<br> Must be a YAML object, e.g.<br> `{"example.com/test-annotation":"test-value"}`.      |           |
| `WOODPECKER_BACKEND_K8S_POD_LABELS_ALLOW_FROM_STEP` | Determines if Pod annotations can be defined from a step's<br> backend options.                                                          | `false`   |
| `WOODPECKER_BACKEND_K8S_POD_LABELS`                 | Additional labels to apply to worker Pods.<br> Must be a YAML object,<br> e.g. `{"example.com/test-label":"test-value"}`                 |           |
| `WOODPECKER_BACKEND_K8S_POD_NODE_SELECTOR`          | Additional node selector to apply to worker pods.<br> Must be a YAML object, e.g.<br> `{"topology.kubernetes.io/region":"eu-central-1"}` |           |
| `WOODPECKER_BACKEND_K8S_PULL_SECRET_NAMES`          | Secret names to pull images from private repositories.                                                                                   |           |
| `WOODPECKER_BACKEND_K8S_SECCTX_NONROOT`             | Whether containers must be run as a non-root user.                                                                                       | `false`   |
| `WOODPECKER_BACKEND_K8S_STORAGE_CLASS`              | The storage class to use for the temporary pipeline volume.                                                                              |           |
| `WOODPECKER_BACKEND_K8S_STORAGE_RWX`                | Whether a RWX should be used for the temporary pipeline volume.<br> If false, RWO is used instead.                                       | `true`    |
| `WOODPECKER_BACKEND_K8S_STORAGE_CLASS`              | The storage class to use for the pipeline volume.                                                                                        |           |
| `WOODPECKER_BACKEND_K8S_VOLUME_SIZE`                | The volume size of the temporary pipeline volume.                                                                                        | `10G`     |
| `WOODPECKER_BACKEND_LOCAL_TEMP_DIR`                 | Directory in which pipelines are executed                                                                                                | `$TMPDIR` |
| `WOODPECKER_LOG_LEVEL`                              | Logging level. Possible values are `trace`, `debug`, `info`, <br>`warn`, `error`, `fatal`, `panic`, and `disabled`.                      |           |
| `WOODPECKER_LOG_FILE`                               | Output destination for logs. <br>`stdout` and `stderr` can be used as special keywords.                                                  | `stderr`  |

## Agent

| Name                             | Description                                                                                                                                                                                                                                                                              | Value            |
| -------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------- |
| `WOODPECKER_AGENT_CONFIG_FILE`   | Filepath containing agent config, e.g. `/etc/woodpecker/agent.conf`                                                                                                                                                                                                                      |                  |
| `WOODPECKER_AGENT_LABELS`        | Configures custom labels for the agent, to enable workflow filtering.<br> Accepts a list of key-value pairs like `key=value,second-key=*`.<br> Agents provide three additional labels <br> `platform=os/arch`, `hostname=my-agent` and `repo=*` <br> which can be overwritten if needed. |                  |
| `WOODPECKER_AGENT_SECRET_FILE`   | Filepath containing agent secret, e.g. `/etc/woodpecker/agent-secret.conf`                                                                                                                                                                                                               |                  |
| `WOODPECKER_AGENT_SECRET`        | A shared secret used by server and agents to authenticate communication.<br> A secret can be generated via `openssl rand -hex 32`.                                                                                                                                                       |                  |
| `WOODPECKER_BACKEND`             | Woodpecker backend to use. <br>Possible values are `auto-detect`, `docker`, `local` or `kubernetes`.                                                                                                                                                                                     | `auto-detect`    |
| `WOODPECKER_CONNECT_RETRY_COUNT` | number of times agent retries to connect to the server.                                                                                                                                                                                                                                  | `5`              |
| `WOODPECKER_CONNECT_RETRY_DELAY` | delay between agent connection retries to the server.                                                                                                                                                                                                                                    | `2s`             |
| `WOODPECKER_DEBUG_NOCOLOR`       | Disable colored debug output.                                                                                                                                                                                                                                                            | `true`           |
| `WOODPECKER_DEBUG_PRETTY`        | Enable pretty-printed debug output.                                                                                                                                                                                                                                                      | `false`          |
| `WOODPECKER_GRPC_SECURE`         | Whether the connection to `WOODPECKER_SERVER` should be made via SSL.                                                                                                                                                                                                                    | `false`          |
| `WOODPECKER_GRPC_VERIFY`         | Whether the `grpc` server certificate should be verified. <br>Only valid when `WOODPECKER_GRPC_SECURE=true`.                                                                                                                                                                             | `true`           |
| `WOODPECKER_HEALTHCHECK_ADDR`    | Healthcheck endpoint address.                                                                                                                                                                                                                                                            | `:3000`          |
| `WOODPECKER_HEALTHCHECK`         | Enable healthcheck endpoint.                                                                                                                                                                                                                                                             | `true`           |
| `WOODPECKER_HOSTNAME`            | Agent hostname                                                                                                                                                                                                                                                                           |                  |
| `WOODPECKER_KEEPALIVE_TIME`      | With no activity after this duration, the agent pings the server <br> to check if the transport is still alive.                                                                                                                                                                          |                  |
| `WOODPECKER_KEEPALIVE_TIMEOUT`   | After pinging for a keepalive check, the agent waits for this duration <br>before closing unresponsive connections.                                                                                                                                                                      | `20s`            |
| `WOODPECKER_LOG_LEVEL`           | Logging level. Possible values are `trace`, `debug`, `info`, <br>`warn`, `error`, `fatal`, `panic`, and `disabled`.                                                                                                                                                                      |                  |
| `WOODPECKER_MAX_WORKFLOWS`       | Number of parallel workflows.                                                                                                                                                                                                                                                            | `1`              |
| `WOODPECKER_SERVER`              | gRPC address of the server.                                                                                                                                                                                                                                                              | `localhost:9000` |
| `WOODPECKER_USERNAME`            | gRPC username.                                                                                                                                                                                                                                                                           | `x-oauth-basic`  |
| `WOODPECKER_LOG_FILE`            | Output destination for logs. `stdout` and `stderr` can be used as special keywords.                                                                                                                                                                                                      | `stderr`         |