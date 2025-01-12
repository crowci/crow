---
hide:
  - navigation
  # - toc
---
# Introduction

Crow is a *Continuous Integration & Continuous Delivery* (CI/CD)[^1] application.
It is designed to be lightweight, simple to use and fast. 

A typical CI/CD pipeline usually includes the following steps:

1. Running tests and linters
2. Building an application
3. Deploying an application

However, these days CI/CD apps are also often used as flexible task schedulers with central logging and can be tailored to pretty much any use case.

## Containers at the core

Crow is *only* using containers to run pipelines.
This stands in contrast to other common CI/CD apps like GitHub Actions, GitLab Runner or Jenkins.
While these can also execute pipelines in containers if desired, they are "Host-first", i.e. by default they run pipelines directly on a VM and are restricted to the OS on that VM.

Crow in contrast only uses containers and is hence operating-system agnostic, i.e. it does not matter which OS is running on the host which is executing the pipelines.

With respect to well-known CI/CD providers, Crow is mostly comparable to Circle CI, which is also using a container-only approach.

## Target audience

Crow is FOSS, lightweight and shines when being self-hosted in private environments[^3].
It has a very small footprint (< 200 MB memory consumption) in idle and can even be installed and practically used on a Raspberry Pi with 4GB (or even less, depending on what the builds will do).

Crow uses a SQLite by default and can do a lot with in during runtime.
If you are planning a bigger instance with multiple (active) users and > 20 active repos, it is recommended to configure it with a Postgres or MariaDB database for performance reasons.[^2]

## History

Crow has been forked from Woodpecker CI (v3.0.0) in January 2025 from the previous Woodpecker maintainer [@pat-s](https://github.com/pat-s).
Motivation was built around improving infrastructure-related processes of the project (releases, docs, governance) which mainly found no agreement in previous discussions in Woodpecker.

Woodpecker itself is a fork of Drone CI (v0.8.91) in April 2019 with the first standalone version being release on September 9th 2019.

### License

Due to being a fork of Woodpecker and Drone, Crow is licensed under the Apache 2.0 license.

<!-- markdownlint-disable -->
[^1]: [This RedHat blog post explains the concept of CI/CD in more detail.](https://www.redhat.com/en/topics/devops/what-is-ci-cd)
[^2]: This is primarily because Crow (still) stores pipeline logs in the DB. A refactoring to storing these outside the DB by default is planned but now yet implemented.
[^3]: Crow can also be run in public which is successfully shown on Codeberg which uses Woodpecker (which Crow was forked from) as the main CI/CD engine on its platform. Doing so however requires more administration RBAC-work on the admin side (giving permissions, enforcing trusted plugins, caring about secret exposure) than running Crow within a closed private environment.
<!-- markdownlint-enable -->
