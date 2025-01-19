// Copyright 2022 Woodpecker Authors
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

package pipeline

import (
	"context"
	"strconv"

	"github.com/crowci/crow/v3/cli/common"
	"github.com/crowci/crow/v3/cli/internal"
	crow "github.com/crowci/crow/v3/crow-go/crow"
	"github.com/urfave/cli/v3"
)

var pipelineShowCmd = &cli.Command{
	Name:      "show",
	Usage:     "show pipeline information",
	ArgsUsage: "<repo-id|repo-full-name> [pipeline]",
	Action:    pipelineShow,
	Flags:     common.OutputFlags("table"),
}

func pipelineShow(ctx context.Context, c *cli.Command) error {
	repoIDOrFullName := c.Args().First()
	client, err := internal.NewClient(ctx, c)
	if err != nil {
		return err
	}
	repoID, err := internal.ParseRepo(client, repoIDOrFullName)
	if err != nil {
		return err
	}
	pipelineArg := c.Args().Get(1)

	var number int64
	if pipelineArg == "last" || len(pipelineArg) == 0 {
		// Fetch the pipeline number from the last pipeline
		pipeline, err := client.PipelineLast(repoID, crow.PipelineLastOptions{})
		if err != nil {
			return err
		}
		number = pipeline.Number
	} else {
		number, err = strconv.ParseInt(pipelineArg, 10, 64)
		if err != nil {
			return err
		}
	}

	pipeline, err := client.Pipeline(repoID, number)
	if err != nil {
		return err
	}

	return pipelineOutput(c, []*crow.Pipeline{pipeline})
}
