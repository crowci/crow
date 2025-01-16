// Copyright 2023 Woodpecker Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package secret

import (
	"context"
	"os"
	"strings"

	"github.com/crowci/crow/v3/cli/common"
	"github.com/crowci/crow/v3/cli/internal"
	"github.com/crowci/crow/v3/woodpecker-go/woodpecker"
	"github.com/urfave/cli/v3"
)

var secretCreateCmd = &cli.Command{
	Name:      "add",
	Usage:     "add a secret",
	ArgsUsage: "[repo-id|repo-full-name]",
	Action:    secretCreate,
	Flags: []cli.Flag{
		common.RepoFlag,
		&cli.StringFlag{
			Name:  "name",
			Usage: "secret name",
		},
		&cli.StringFlag{
			Name:  "value",
			Usage: "secret value",
		},
		&cli.StringSliceFlag{
			Name:  "event",
			Usage: "limit secret to these events",
		},
		&cli.StringSliceFlag{
			Name:  "image",
			Usage: "limit secret to these images",
		},
	},
}

func secretCreate(ctx context.Context, c *cli.Command) error {
	client, err := internal.NewClient(ctx, c)
	if err != nil {
		return err
	}

	secret := &woodpecker.Secret{
		Name:   strings.ToLower(c.String("name")),
		Value:  c.String("value"),
		Images: c.StringSlice("image"),
		Events: c.StringSlice("event"),
	}
	if len(secret.Events) == 0 {
		secret.Events = defaultSecretEvents
	}
	if strings.HasPrefix(secret.Value, "@") {
		path := strings.TrimPrefix(secret.Value, "@")
		out, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		secret.Value = string(out)
	}

	repoID, err := parseTargetArgs(client, c)
	if err != nil {
		return err
	}

	_, err = client.SecretCreate(repoID, secret)
	return err
}

var defaultSecretEvents = []string{
	woodpecker.EventPush,
	woodpecker.EventTag,
	woodpecker.EventRelease,
	woodpecker.EventDeploy,
}
