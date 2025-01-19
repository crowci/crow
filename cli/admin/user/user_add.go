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

package user

import (
	"context"
	"fmt"

	"github.com/crowci/crow/v3/cli/internal"
	crow "github.com/crowci/crow/v3/crow-go/crow"
	"github.com/urfave/cli/v3"
)

var userAddCmd = &cli.Command{
	Name:      "add",
	Usage:     "add a user",
	ArgsUsage: "<username>",
	Action:    userAdd,
}

func userAdd(ctx context.Context, c *cli.Command) error {
	login := c.Args().First()

	client, err := internal.NewClient(ctx, c)
	if err != nil {
		return err
	}

	user, err := client.UserPost(&crow.User{Login: login})
	if err != nil {
		return err
	}
	fmt.Printf("Successfully added user %s\n", user.Login)
	return nil
}
