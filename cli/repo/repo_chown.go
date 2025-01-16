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

package repo

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/crowci/crow/v3/cli/internal"
)

var repoChownCmd = &cli.Command{
	Name:      "chown",
	Usage:     "assume ownership of a repository",
	ArgsUsage: "<repo-id|repo-full-name>",
	Action:    repoChown,
}

func repoChown(ctx context.Context, c *cli.Command) error {
	repoIDOrFullName := c.Args().First()
	client, err := internal.NewClient(ctx, c)
	if err != nil {
		return err
	}
	repoID, err := internal.ParseRepo(client, repoIDOrFullName)
	if err != nil {
		return err
	}

	repo, err := client.RepoChown(repoID)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully assumed ownership of repository %s\n", repo.FullName)
	return nil
}
