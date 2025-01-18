// Copyright 2023 Woodpecker Authors
// Copyright 2019 Laszlo Fogas
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"
	"time"

	"github.com/crowci/crow/v3/shared/constant"
	"github.com/crowci/crow/v3/shared/logger"
	"github.com/urfave/cli/v3"
)

var flags = append([]cli.Flag{
	&cli.BoolFlag{
		Sources: cli.EnvVars("WOODPECKER_DATABASE_LOG", "WOODPECKER_LOG_XORM", "CROW_DATABASE_LOG", "CROW_LOG_XORM"),
		Name:    "db-log",
		Aliases: []string{"log-xorm"}, // TODO: remove in v4.0.0
		Usage:   "enable logging in database engine (currently xorm)",
	},
	&cli.BoolFlag{
		Sources: cli.EnvVars("WOODPECKER_DATABASE_LOG_SQL", "WOODPECKER_LOG_XORM_SQL", "CROW_DATABASE_LOG_SQL", "CROW_LOG_XORM_SQL"),
		Name:    "db-log-sql",
		Aliases: []string{"log-xorm-sql"}, // TODO: remove in v4.0.0
		Usage:   "enable logging of sql commands",
	},
	&cli.IntFlag{
		Sources: cli.EnvVars("WOODPECKER_DATABASE_MAX_CONNECTIONS", "CROW_DATABASE_MAX_CONNECTIONS"),
		Name:    "db-max-open-connections",
		Usage:   "max connections xorm is allowed create",
		Value:   100,
	},
	&cli.IntFlag{
		Sources: cli.EnvVars("WOODPECKER_DATABASE_IDLE_CONNECTIONS", "CROW_DATABASE_IDLE_CONNECTIONS"),
		Name:    "db-max-idle-connections",
		Usage:   "amount of connections xorm will hold open",
		Value:   2,
	},
	&cli.DurationFlag{
		Sources: cli.EnvVars("WOODPECKER_DATABASE_CONNECTION_TIMEOUT", "CROW_DATABASE_CONNECTION_TIMEOUT"),
		Name:    "db-max-connection-timeout",
		Usage:   "time an active connection is allowed to stay open",
		Value:   3 * time.Second,
	},
	&cli.StringFlag{
		Sources: cli.EnvVars("WOODPECKER_HOST", "CROW_HOST"),
		Name:    "server-host",
		Usage:   "server fully qualified url. Format: <scheme>://<host>[/<prefix path>]",
	},
	&cli.StringFlag{
		Sources: cli.EnvVars("WOODPECKER_SERVER_ADDR", "CROW_SERVER_ADDR"),
		Name:    "server-addr",
		Usage:   "server address",
		Value:   ":8000",
	},
	&cli.StringFlag{
		Sources: cli.EnvVars("WOODPECKER_SERVER_ADDR_TLS", "CROW_SERVER_ADDR_TLS"),
		Name:    "server-addr-tls",
		Usage:   "port https with tls (:443)",
		Value:   ":443",
	},
	&cli.StringFlag{
		Sources: cli.EnvVars("WOODPECKER_SERVER_CERT", "CROW_SERVER_CERT"),
		Name:    "server-cert",
		Usage:   "server ssl cert path",
	},
	&cli.StringFlag{
		Sources: cli.EnvVars("WOODPECKER_SERVER_KEY", "CROW_SERVER_KEY"),
		Name:    "server-key",
		Usage:   "server ssl key path",
	},
	&cli.StringFlag{
		Sources: cli.EnvVars("WOODPECKER_CUSTOM_CSS_FILE", "CROW_CUSTOM_CSS_FILE"),
		Name:    "custom-css-file",
		Usage:   "file path for the server to serve a custom .CSS file, used for customizing the UI",
	},
	&cli.StringFlag{
		Sources: cli.EnvVars("WOODPECKER_CUSTOM_JS_FILE", "CROW_CUSTOM_JS_FILE"),
		Name:    "custom-js-file",
		Usage:   "file path for the server to serve a custom .JS file, used for customizing the UI",
	},
	&cli.StringFlag{
		Sources: cli.EnvVars("WOODPECKER_GRPC_ADDR", "CROW_GRPC_ADDR"),
		Name:    "grpc-addr",
		Usage:   "grpc address",
		Value:   ":9000",
	},
	&cli.StringFlag{
		Sources: cli.NewValueSourceChain(
			cli.File(os.Getenv("WOODPECKER_GRPC_SECRET_FILE")),
			cli.EnvVar("WOODPECKER_GRPC_SECRET"),
			cli.File(os.Getenv("CROW_GRPC_SECRET_FILE")),
			cli.EnvVar("CROW_GRPC_SECRET")),
		Name:  "grpc-secret",
		Usage: "grpc jwt secret",
		Value: "secret",
		Config: cli.StringConfig{
			TrimSpace: true,
		},
	},
	&cli.StringFlag{
		Sources: cli.EnvVars("WOODPECKER_METRICS_SERVER_ADDR", "CROW_METRICS_SERVER_ADDR"),
		Name:    "metrics-server-addr",
		Usage:   "metrics server address",
		Value:   "",
	},
	&cli.StringSliceFlag{
		Sources: cli.EnvVars("WOODPECKER_ADMIN", "CROW_ADMIN"),
		Name:    "admin",
		Usage:   "list of admin users",
	},
	&cli.StringSliceFlag{
		Sources: cli.EnvVars("WOODPECKER_ORGS", "CROW_ORGS"),
		Name:    "orgs",
		Usage:   "list of approved organizations",
	},
	&cli.StringSliceFlag{
		Sources: cli.EnvVars("WOODPECKER_REPO_OWNERS", "CROW_REPO_OWNERS"),
		Name:    "repo-owners",
		Usage:   "Repositories by those owners will be allowed to be used in woodpecker",
	},
	&cli.BoolFlag{
		Sources: cli.EnvVars("WOODPECKER_OPEN", "CROW_OPEN"),
		Name:    "open",
		Usage:   "enable open user registration",
	},
	&cli.BoolFlag{
		Sources: cli.EnvVars("WOODPECKER_AUTHENTICATE_PUBLIC_REPOS", "CROW_AUTHENTICATE_PUBLIC_REPOS"),
		Name:    "authenticate-public-repos",
		Usage:   "Always use authentication to clone repositories even if they are public. Needed if the SCM requires to always authenticate as used by many companies.",
	},
	&cli.StringSliceFlag{
		Sources: cli.EnvVars("WOODPECKER_DEFAULT_CANCEL_PREVIOUS_PIPELINE_EVENTS", "CROW_DEFAULT_CANCEL_PREVIOUS_PIPELINE_EVENTS"),
		Name:    "default-cancel-previous-pipeline-events",
		Usage:   "List of event names that will be canceled when a new pipeline for the same context (tag, branch) is created.",
		Value:   []string{"push", "pull_request"},
	},
	&cli.StringFlag{
		Sources: cli.EnvVars("WOODPECKER_DEFAULT_CLONE_PLUGIN", "WOODPECKER_DEFAULT_CLONE_IMAGE", "CROW_DEFAULT_CLONE_PLUGIN", "CROW_DEFAULT_CLONE_IMAGE"),
		Name:    "default-clone-plugin",
		Aliases: []string{"default-clone-image"},
		Usage:   "The default docker image to be used when cloning the repo",
		Value:   constant.DefaultClonePlugin,
	},
	&cli.IntFlag{
		Sources: cli.EnvVars("WOODPECKER_DEFAULT_PIPELINE_TIMEOUT", "CROW_DEFAULT_PIPELINE_TIMEOUT"),
		Name:    "default-pipeline-timeout",
		Usage:   "The default time in minutes for a repo in minutes before a pipeline gets killed",
		Value:   60,
	},
	&cli.IntFlag{
		Sources: cli.EnvVars("WOODPECKER_MAX_PIPELINE_TIMEOUT", "CROW_MAX_PIPELINE_TIMEOUT"),
		Name:    "max-pipeline-timeout",
		Usage:   "The maximum time in minutes you can set in the repo settings before a pipeline gets killed",
		Value:   120,
	},
	&cli.StringSliceFlag{
		Sources: cli.EnvVars("WOODPECKER_DEFAULT_WORKFLOW_LABELS", "CROW_DEFAULT_WORKFLOW_LABELS"),
		Name:    "default-workflow-labels",
		Usage:   "The default label filter to set for workflows that has no label filter set. By default workflows will be allowed to run on any agent, if not specified in the workflow.",
	},
	&cli.DurationFlag{
		Sources: cli.EnvVars("WOODPECKER_SESSION_EXPIRES", "CROW_SESSION_EXPIRES"),
		Name:    "session-expires",
		Usage:   "session expiration time",
		Value:   time.Hour * 72,
	},
	&cli.StringSliceFlag{
		Sources: cli.EnvVars("WOODPECKER_PLUGINS_PRIVILEGED", "CROW_PLUGINS_PRIVILEGED"),
		Name:    "plugins-privileged",
		Usage:   "Allow plugins to run in privileged mode, if environment variable is defined but empty there will be none",
	},
	&cli.StringSliceFlag{
		Sources: cli.EnvVars("WOODPECKER_PLUGINS_TRUSTED_CLONE", "CROW_PLUGINS_TRUSTED_CLONE"),
		Name:    "plugins-trusted-clone",
		Usage:   "Plugins which are trusted to handle Git credentials in clone steps",
		Value:   constant.TrustedClonePlugins,
	},
	&cli.StringSliceFlag{
		Sources: cli.EnvVars("WOODPECKER_VOLUME", "CROW_VOLUME"),
		Name:    "volume",
	},
	&cli.StringFlag{
		Sources: cli.EnvVars("WOODPECKER_DOCKER_CONFIG", "CROW_DOCKER_CONFIG"),
		Name:    "docker-config",
	},
	&cli.StringSliceFlag{
		Sources: cli.EnvVars("WOODPECKER_ENVIRONMENT", "CROW_ENVIRONMENT"),
		Name:    "environment",
	},
	&cli.StringSliceFlag{
		Sources: cli.EnvVars("WOODPECKER_NETWORK", "CROW_NETWORK"),
		Name:    "network",
	},
	&cli.StringFlag{
		Sources: cli.NewValueSourceChain(
			cli.File(os.Getenv("WOODPECKER_AGENT_SECRET_FILE")),
			cli.EnvVar("WOODPECKER_AGENT_SECRET"),
			cli.File(os.Getenv("CROW_AGENT_SECRET_FILE")),
			cli.EnvVar("CROW_AGENT_SECRET")),
		Name:  "agent-secret",
		Usage: "server-agent shared password",
		Config: cli.StringConfig{
			TrimSpace: true,
		},
	},
	&cli.BoolFlag{
		Sources: cli.EnvVars("WOODPECKER_DISABLE_USER_AGENT_REGISTRATION", "CROW_DISABLE_USER_AGENT_REGISTRATION"),
		Name:    "disable-user-agent-registration",
		Usage:   "Disable user registered agents",
	},
	&cli.DurationFlag{
		Sources: cli.EnvVars("WOODPECKER_KEEPALIVE_MIN_TIME", "CROW_KEEPALIVE_MIN_TIME"),
		Name:    "keepalive-min-time",
		Usage:   "server-side enforcement policy on the minimum amount of time a client should wait before sending a keepalive ping.",
	},
	&cli.StringFlag{
		Sources: cli.EnvVars("WOODPECKER_CONFIG_SERVICE_ENDPOINT", "CROW_CONFIG_SERVICE_ENDPOINT"),
		Name:    "config-service-endpoint",
		Usage:   "url used for calling configuration service endpoint",
	},
	&cli.StringFlag{
		Sources: cli.EnvVars("WOODPECKER_DATABASE_DRIVER", "CROW_DATABASE_DRIVER"),
		Name:    "db-driver",
		Aliases: []string{"driver"}, // TODO: remove in v4.0.0
		Usage:   "database driver",
		Value:   "sqlite3",
	},
	&cli.StringFlag{
		Sources: cli.NewValueSourceChain(
			cli.File(os.Getenv("WOODPECKER_DATABASE_DATASOURCE_FILE")),
			cli.EnvVar("WOODPECKER_DATABASE_DATASOURCE"),
			cli.File(os.Getenv("CROW_DATABASE_DATASOURCE_FILE")),
			cli.EnvVar("CROW_DATABASE_DATASOURCE")),
		Name:    "db-datasource",
		Aliases: []string{"datasource"}, // TODO: remove in v4.0.0
		Usage:   "database driver configuration string",
		Value:   datasourceDefaultValue(),
		Config: cli.StringConfig{
			TrimSpace: true,
		},
	},
	&cli.StringFlag{
		Sources: cli.NewValueSourceChain(
			cli.File(os.Getenv("WOODPECKER_PROMETHEUS_AUTH_TOKEN_FILE")),
			cli.EnvVar("WOODPECKER_PROMETHEUS_AUTH_TOKEN"),
			cli.File(os.Getenv("CROW_PROMETHEUS_AUTH_TOKEN_FILE")),
			cli.EnvVar("CROW_PROMETHEUS_AUTH_TOKEN")),
		Name:  "prometheus-auth-token",
		Usage: "token to secure prometheus metrics endpoint",
		Config: cli.StringConfig{
			TrimSpace: true,
		},
	},
	&cli.StringFlag{
		Sources: cli.EnvVars("WOODPECKER_STATUS_CONTEXT", "WOODPECKER_GITHUB_CONTEXT", "WOODPECKER_GITEA_CONTEXT", "CROW_STATUS_CONTEXT", "CROW_GITHUB_CONTEXT", "CROW_GITEA_CONTEXT"),
		Name:    "status-context",
		Usage:   "status context prefix",
		Value:   "ci/crow",
	},
	&cli.StringFlag{
		Sources: cli.EnvVars("WOODPECKER_STATUS_CONTEXT_FORMAT", "CROW_STATUS_CONTEXT_FORMAT"),
		Name:    "status-context-format",
		Usage:   "status context format",
		Value:   "{{ .context }}/{{ .event }}/{{ .workflow }}{{if not (eq .axis_id 0)}}/{{.axis_id}}{{end}}",
	},
	&cli.BoolFlag{
		Sources: cli.EnvVars("WOODPECKER_MIGRATIONS_ALLOW_LONG", "CROW_MIGRATIONS_ALLOW_LONG"),
		Name:    "migrations-allow-long",
		Value:   false,
	},
	&cli.BoolFlag{
		Sources: cli.EnvVars("WOODPECKER_ENABLE_SWAGGER", "CROW_ENABLE_SWAGGER"),
		Name:    "enable-swagger",
		Value:   true,
	},
	&cli.BoolFlag{
		Sources: cli.EnvVars("WOODPECKER_DISABLE_VERSION_CHECK", "CROW_DISABLE_VERSION_CHECK"),
		Usage:   "Disable version check in admin web ui.",
		Name:    "skip-version-check",
	},
	&cli.StringFlag{
		Sources: cli.EnvVars("WOODPECKER_LOG_STORE", "CROW_LOG_STORE"),
		Name:    "log-store",
		Usage:   "log store to use ('database' or 'file')",
		Value:   "database",
	},
	&cli.StringFlag{
		Sources: cli.EnvVars("WOODPECKER_LOG_STORE_FILE_PATH", "CROW_LOG_STORE_FILE_PATH"),
		Name:    "log-store-file-path",
		Usage:   "directory used for file based log storage",
	},
	//
	// backend options for pipeline compiler
	//
	&cli.StringFlag{
		Sources: cli.EnvVars("WOODPECKER_BACKEND_NO_PROXY", "NO_PROXY", "no_proxy", "CROW_BACKEND_NO_PROXY"),
		Usage:   "if set, pass the environment variable down as \"NO_PROXY\" to steps",
		Name:    "backend-no-proxy",
	},
	&cli.StringFlag{
		Sources: cli.EnvVars("WOODPECKER_BACKEND_HTTP_PROXY", "HTTP_PROXY", "http_proxy", "CROW_BACKEND_HTTP_PROXY"),
		Usage:   "if set, pass the environment variable down as \"HTTP_PROXY\" to steps",
		Name:    "backend-http-proxy",
	},
	&cli.StringFlag{
		Sources: cli.EnvVars("WOODPECKER_BACKEND_HTTPS_PROXY", "HTTPS_PROXY", "https_proxy", "CROW_BACKEND_HTTPS_PROXY"),
		Usage:   "if set, pass the environment variable down as \"HTTPS_PROXY\" to steps",
		Name:    "backend-https-proxy",
	},
	//
	// resource limit parameters
	//
	&cli.DurationFlag{
		Sources: cli.EnvVars("WOODPECKER_FORGE_TIMEOUT", "CROW_FORGE_TIMEOUT"),
		Name:    "forge-timeout",
		Usage:   "how many seconds before timeout when fetching the Woodpecker configuration from a Forge",
		Value:   time.Second * 5,
	},
	&cli.UintFlag{
		Sources: cli.EnvVars("WOODPECKER_FORGE_RETRY", "CROW_FORGE_RETRY"),
		Name:    "forge-retry",
		Usage:   "How many retries of fetching the Woodpecker configuration from a forge are done before we fail",
		Value:   3,
	},
	//
	// generic forge settings
	//
	&cli.StringFlag{
		Name:    "forge-url",
		Usage:   "url of the forge",
		Sources: cli.EnvVars("WOODPECKER_FORGE_URL", "WOODPECKER_GITHUB_URL", "WOODPECKER_GITLAB_URL", "WOODPECKER_GITEA_URL", "WOODPECKER_FORGEJO_URL", "WOODPECKER_BITBUCKET_URL", "WOODPECKER_BITBUCKET_DC_URL", "CROW_FORGE_URL", "CROW_GITHUB_URL", "CROW_GITLAB_URL", "CROW_GITEA_URL", "CROW_FORGEJO_URL", "CROW_BITBUCKET_URL", "CROW_BITBUCKET_DC_URL"),
	},
	&cli.StringFlag{
		Sources: cli.NewValueSourceChain(
			cli.File(getFirstNonEmptyEnvVar(
				"WOODPECKER_FORGE_CLIENT_FILE",
				"WOODPECKER_GITHUB_CLIENT_FILE",
				"WOODPECKER_GITLAB_CLIENT_FILE",
				"WOODPECKER_GITEA_CLIENT_FILE",
				"WOODPECKER_FORGEJO_CLIENT_FILE",
				"WOODPECKER_BITBUCKET_CLIENT_FILE",
				"WOODPECKER_BITBUCKET_DC_CLIENT_ID_FILE",
				"CROW_FORGE_CLIENT_FILE",
				"CROW_GITHUB_CLIENT_FILE",
				"CROW_GITLAB_CLIENT_FILE",
				"CROW_GITEA_CLIENT_FILE",
				"CROW_FORGEJO_CLIENT_FILE",
				"CROW_BITBUCKET_CLIENT_FILE",
				"CROW_BITBUCKET_DC_CLIENT_ID_FILE")),
			cli.EnvVar("WOODPECKER_FORGE_CLIENT"),
			cli.EnvVar("WOODPECKER_GITHUB_CLIENT"),
			cli.EnvVar("WOODPECKER_GITLAB_CLIENT"),
			cli.EnvVar("WOODPECKER_GITEA_CLIENT"),
			cli.EnvVar("WOODPECKER_FORGEJO_CLIENT"),
			cli.EnvVar("WOODPECKER_BITBUCKET_CLIENT"),
			cli.EnvVar("WOODPECKER_BITBUCKET_DC_CLIENT_ID"),
			cli.EnvVar("CROW_FORGE_CLIENT"),
			cli.EnvVar("CROW_GITHUB_CLIENT"),
			cli.EnvVar("CROW_GITLAB_CLIENT"),
			cli.EnvVar("CROW_GITEA_CLIENT"),
			cli.EnvVar("CROW_FORGEJO_CLIENT"),
			cli.EnvVar("CROW_BITBUCKET_CLIENT"),
			cli.EnvVar("CROW_BITBUCKET_DC_CLIENT_ID")),
		Name:  "forge-oauth-client",
		Usage: "oauth2 client id",
		Config: cli.StringConfig{
			TrimSpace: true,
		},
	},
	&cli.StringFlag{
		Sources: cli.NewValueSourceChain(
			cli.File(getFirstNonEmptyEnvVar(
				"WOODPECKER_FORGE_SECRET_FILE",
				"WOODPECKER_GITHUB_SECRET_FILE",
				"WOODPECKER_GITLAB_SECRET_FILE",
				"WOODPECKER_GITEA_SECRET_FILE",
				"WOODPECKER_FORGEJO_SECRET_FILE",
				"WOODPECKER_BITBUCKET_SECRET_FILE",
				"WOODPECKER_BITBUCKET_DC_CLIENT_SECRET_FILE",
				"CROW_FORGE_SECRET_FILE",
				"CROW_GITHUB_SECRET_FILE",
				"CROW_GITLAB_SECRET_FILE",
				"CROW_GITEA_SECRET_FILE",
				"CROW_FORGEJO_SECRET_FILE",
				"CROW_BITBUCKET_SECRET_FILE",
				"CROW_BITBUCKET_DC_CLIENT_SECRET_FILE")),
			cli.EnvVar("WOODPECKER_FORGE_SECRET"),
			cli.EnvVar("WOODPECKER_GITHUB_SECRET"),
			cli.EnvVar("WOODPECKER_GITLAB_SECRET"),
			cli.EnvVar("WOODPECKER_GITEA_SECRET"),
			cli.EnvVar("WOODPECKER_FORGEJO_SECRET"),
			cli.EnvVar("WOODPECKER_BITBUCKET_SECRET"),
			cli.EnvVar("WOODPECKER_BITBUCKET_DC_CLIENT_SECRET"),
			cli.EnvVar("CROW_FORGE_SECRET"),
			cli.EnvVar("CROW_GITHUB_SECRET"),
			cli.EnvVar("CROW_GITLAB_SECRET"),
			cli.EnvVar("CROW_GITEA_SECRET"),
			cli.EnvVar("CROW_FORGEJO_SECRET"),
			cli.EnvVar("CROW_BITBUCKET_SECRET"),
			cli.EnvVar("CROW_BITBUCKET_DC_CLIENT_SECRET")),
		Name:  "forge-oauth-secret",
		Usage: "oauth2 client secret",
		Config: cli.StringConfig{
			TrimSpace: true,
		},
	},
	&cli.BoolFlag{
		Name:  "forge-skip-verify",
		Usage: "skip ssl verification",
		Sources: cli.EnvVars(
			"WOODPECKER_FORGE_SKIP_VERIFY",
			"WOODPECKER_GITHUB_SKIP_VERIFY",
			"WOODPECKER_GITLAB_SKIP_VERIFY",
			"WOODPECKER_GITEA_SKIP_VERIFY",
			"WOODPECKER_FORGEJO_SKIP_VERIFY",
			"WOODPECKER_BITBUCKET_SKIP_VERIFY",
			"CROW_FORGE_SKIP_VERIFY",
			"CROW_GITHUB_SKIP_VERIFY",
			"CROW_GITLAB_SKIP_VERIFY",
			"CROW_GITEA_SKIP_VERIFY",
			"CROW_FORGEJO_SKIP_VERIFY",
			"CROW_BITBUCKET_SKIP_VERIFY"),
	},
	&cli.StringFlag{
		Sources: cli.EnvVars("WOODPECKER_EXPERT_FORGE_OAUTH_HOST", "CROW_EXPERT_FORGE_OAUTH_HOST"),
		Name:    "forge-oauth-host",
		Usage:   "!!!for experts!!! fully qualified public forge url. Use it if your forge url WOODPECKER_FORGE_URL or WOODPECKER_GITEA_URL, ... isn't a public url. Format: <scheme>://<host>[/<prefix path>]",
	},
	//
	// Addon
	//
	&cli.StringFlag{
		Sources: cli.EnvVars("WOODPECKER_ADDON_FORGE", "CROW_ADDON_FORGE"),
		Name:    "addon-forge",
		Usage:   "path to forge addon executable",
	},
	//
	// GitHub
	//
	&cli.BoolFlag{
		Sources: cli.EnvVars("WOODPECKER_GITHUB", "CROW_GITHUB"),
		Name:    "github",
		Usage:   "github driver is enabled",
	},
	&cli.BoolFlag{
		Sources: cli.EnvVars("WOODPECKER_GITHUB_MERGE_REF", "CROW_GITHUB_MERGE_REF"),
		Name:    "github-merge-ref",
		Usage:   "github pull requests use merge ref",
		Value:   true,
	},
	&cli.BoolFlag{
		Sources: cli.EnvVars("WOODPECKER_GITHUB_PUBLIC_ONLY", "CROW_GITHUB_PUBLIC_ONLY"),
		Name:    "github-public-only",
		Usage:   "github tokens should only get access to public repos",
		Value:   false,
	},
	//
	// Gitea
	//
	&cli.BoolFlag{
		Sources: cli.EnvVars("WOODPECKER_GITEA", "CROW_GITEA"),
		Name:    "gitea",
		Usage:   "gitea driver is enabled",
	},
	//
	// Forgejo
	//
	&cli.BoolFlag{
		Sources: cli.EnvVars("WOODPECKER_FORGEJO", "CROW_FORGEJO"),
		Name:    "forgejo",
		Usage:   "forgejo driver is enabled",
	},
	//
	// Bitbucket
	//
	&cli.BoolFlag{
		Sources: cli.EnvVars("WOODPECKER_BITBUCKET", "CROW_BITBUCKET"),
		Name:    "bitbucket",
		Usage:   "bitbucket driver is enabled",
	},
	//
	// Gitlab
	//
	&cli.BoolFlag{
		Sources: cli.EnvVars("WOODPECKER_GITLAB", "CROW_GITLAB"),
		Name:    "gitlab",
		Usage:   "gitlab driver is enabled",
	},
	//
	// Bitbucket DataCenter/Server (previously Stash)
	//
	&cli.BoolFlag{
		Sources: cli.EnvVars("WOODPECKER_BITBUCKET_DC", "CROW_BITBUCKET_DC"),
		Name:    "bitbucket-dc",
		Usage:   "Bitbucket DataCenter/Server driver is enabled",
	},
	&cli.StringFlag{
		Sources: cli.NewValueSourceChain(
			cli.File(os.Getenv("WOODPECKER_BITBUCKET_DC_GIT_USERNAME_FILE")),
			cli.EnvVar("WOODPECKER_BITBUCKET_DC_GIT_USERNAME"),
			cli.File(os.Getenv("CROW_BITBUCKET_DC_GIT_USERNAME_FILE")),
			cli.EnvVar("CROW_BITBUCKET_DC_GIT_USERNAME")),
		Name:  "bitbucket-dc-git-username",
		Usage: "Bitbucket DataCenter/Server service account username",
		Config: cli.StringConfig{
			TrimSpace: true,
		},
	},
	&cli.StringFlag{
		Sources: cli.NewValueSourceChain(
			cli.File(os.Getenv("WOODPECKER_BITBUCKET_DC_GIT_PASSWORD_FILE")),
			cli.EnvVar("WOODPECKER_BITBUCKET_DC_GIT_PASSWORD"),
			cli.File(os.Getenv("CROW_BITBUCKET_DC_GIT_PASSWORD_FILE")),
			cli.EnvVar("CROW_BITBUCKET_DC_GIT_PASSWORD")),
		Name:  "bitbucket-dc-git-password",
		Usage: "Bitbucket DataCenter/Server service account password",
		Config: cli.StringConfig{
			TrimSpace: true,
		},
	},
	//
	// development flags
	//
	&cli.StringFlag{
		Sources: cli.EnvVars("WOODPECKER_DEV_WWW_PROXY", "CROW_DEV_WWW_PROXY"),
		Name:    "www-proxy",
		Usage:   "serve the website by using a proxy (used for development)",
		Hidden:  true,
	},
	//
	// expert flags
	//
	&cli.StringFlag{
		Sources: cli.EnvVars("WOODPECKER_EXPERT_WEBHOOK_HOST", "CROW_EXPERT_WEBHOOK_HOST"),
		Name:    "server-webhook-host",
		Usage:   "!!!for experts!!! fully qualified woodpecker server url called by forge's webhooks. Format: <scheme>://<host>[/<prefix path>]",
	},
	//
	// secrets encryption in DB
	//
	&cli.StringFlag{
		Sources: cli.NewValueSourceChain(
			cli.File(os.Getenv("WOODPECKER_ENCRYPTION_KEY_FILE")),
			cli.EnvVar("WOODPECKER_ENCRYPTION_KEY"),
			cli.File(os.Getenv("CROW_ENCRYPTION_KEY_FILE")),
			cli.EnvVar("CROW_ENCRYPTION_KEY")),
		Name:  "encryption-raw-key",
		Usage: "Raw encryption key",
		Config: cli.StringConfig{
			TrimSpace: true,
		},
	},
	&cli.StringFlag{
		Sources: cli.EnvVars("WOODPECKER_ENCRYPTION_TINK_KEYSET_FILE", "CROW_ENCRYPTION_TINK_KEYSET_FILE"),
		Name:    "encryption-tink-keyset",
		Usage:   "Google tink AEAD-compatible keyset file to encrypt secrets in DB",
	},
	&cli.BoolFlag{
		Sources: cli.EnvVars("WOODPECKER_ENCRYPTION_DISABLE", "CROW_ENCRYPTION_DISABLE"),
		Name:    "encryption-disable-flag",
		Usage:   "Flag to decrypt all encrypted data and disable encryption on server",
	},
}, logger.GlobalLoggerFlags...)

// If woodpecker is running inside a container the default value for
// the datasource is different from running outside a container.
func datasourceDefaultValue() string {
	found := getFirstNonEmptyEnvVar("WOODPECKER_IN_CONTAINER", "CROW_IN_CONTAINER")
	if found != "" {
		return "/var/lib/woodpecker/woodpecker.sqlite"
	}
	return "woodpecker.sqlite"
}

func getFirstNonEmptyEnvVar(envVars ...string) string {
	for _, envVar := range envVars {
		val := os.Getenv(envVar)
		if val != "" {
			return val
		}
	}
	return ""
}
