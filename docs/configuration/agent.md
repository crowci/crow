<!-- markdownlint-disable MD041 -->
To configure an agent, the following settings are required at a minimum:

```sh
CROW_SERVER=<connection uri> # usually 'e.g. woodpecker-server:9000'
CROW_AGENT_SECRET="<token>"
```

## Workflows per agent

!!!note
A pipeline often consists of multiple workflows.
It is recommended therefore to increase this setting to a value matching the resources of the surrounding environment (i.e. if builds are executed on a server with 4GB memory, allowing 100 workflows in parallel is likely not a good match).

By default, every agent is allowed to execute on workflow at a time.
This can be controlled via `CROW_MAX_WORKFLOWS` to increase the amount of parallel workflows an agent is allowed to process.

## Workflow filters

Agents can be configured to only process workflows of certain repositories.

```sh
CROW_AGENT_LABELS=repo=<username>/*
```

This config would restrict an agent to builds of this particular namespace (being it an org or user).
This could be used together with the deployment of the agent on a specific server, and by this isolating specific pipelines from others.
