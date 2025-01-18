---
hide:
  - navigation
  # - toc
---

<!-- markdownlint-disable MD041 -->

Crow CI is a _Continuous Integration & Continuous Delivery_ (CI/CD)[^1] application.
It is designed to be **lightweight**, simple to use and **fast**.

## Why self-host CI/CD?

There are many CI/CD options offered for free.
As with all hosted options which are provided for free, there are limitations and downsides.
Self-hosting provides ways to get rid of these.

For CI/CD specifically, the following limitations of public CI offerings can be tackled:

- Limited build times
- Low-performing runners
- Storing secrets on remote platforms
- Platform/Forge bound

## Why use Crow CI in the first place?

Crow CI gives you full control over the selection of nodes which will process builds.
You can store secrets on servers you own and run as many builds as you want.
In addition, Crow CI works with many different forges (GitHub, GitLab, Gitea, Forgejo, Bitbucket), allowing you to transition between these or use a consistent CI syntax across all of them.

Combined with the [Crow CI Autoscaler](configuration/autoscaler.md), you can make use of powerful CI runners on any cloud provider, giving you the full power the cloud space has to offer while keeping costs to a minimum.

## Containers at the core

Crow CI is a _container-only_ application.
This stands in contrast to other common CI/CD apps like [GitHub Actions](https://docs.github.com/en/actions), [GitLab Runner](https://docs.gitlab.com/runner/), [Jenkins](https://www.jenkins.io/), [Circle CI](https://circleci.com/) or [Drone CI](https://www.drone.io/).
While these can also execute pipelines in containers if desired, they are "host-first", i.e., by default pipelines run directly on a VM and are subsequently restricted to the OS on that VM.

Crow in contrast only uses containers and is hence operating-system agnostic, i.e. it does not matter which OS is running on the host which is executing the pipelines.

With respect to well-known CI/CD providers, Crow is mostly comparable to [Circle CI](https://circleci.com/), which is also making use of a container-only approach.

## Target audience

Crow CI is Apache-2.0 licensed, lightweight and can be used in private environments without restriction.
It has a very small footprint (< 200 MB memory consumption) in idle and can even be installed and practically used on a Raspberry Pi with 4GB (or even less, depending on what the builds will do).

Crow uses a SQLite by default and can do a lot with in during runtime.
If you are planning a bigger instance with multiple (active) users and > 20 active repos, it is recommended to configure it with a Postgres or MariaDB database for performance reasons.[^2]

## History

Crow has been forked from [Woodpecker CI](https://woodpecker-ci.org/) (v3.0.0) in January 2025 from the former Woodpecker maintainer [@pat-s](https://github.com/pat-s).
Motivation for the fork was built around the improvement of infrastructure-related processes (releases, docs, governance) and professionalizing the project ecosystem.

Woodpecker itself is a fork of Drone CI (v0.8.91) from April 2019 with the first standalone version being released on September 9th 2019.

## License

Crow CI is licensed under the Apache 2.0 license (following inheritance from Woodpecker and Drone).

<!-- markdownlint-disable -->

[^1]: [This RedHat blog post explains the concept of CI/CD in more detail.](https://www.redhat.com/en/topics/devops/what-is-ci-cd)

[^2]: This is primarily because Crow (still) stores pipeline logs in the DB. A refactoring to storing these outside the DB by default is planned but now yet implemented.

<!-- markdownlint-enable -->

## Logo

The logo was designed and kindly contributed by [Bold Crow AI](https://boldcrow.ai/).
