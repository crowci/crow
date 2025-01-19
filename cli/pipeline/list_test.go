package pipeline

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/urfave/cli/v3"

	"github.com/crowci/crow/v3/crow-go/crow"
	"github.com/crowci/crow/v3/crow-go/crow/mocks"
)

func TestPipelineList(t *testing.T) {
	testtases := []struct {
		name        string
		repoID      int64
		repoErr     error
		pipelines   []*crow.Pipeline
		pipelineErr error
		args        []string
		expected    []*crow.Pipeline
		wantErr     error
	}{
		{
			name:   "success",
			repoID: 1,
			pipelines: []*crow.Pipeline{
				{ID: 1, Branch: "main", Event: "push", Status: "success"},
				{ID: 2, Branch: "develop", Event: "pull_request", Status: "running"},
				{ID: 3, Branch: "main", Event: "push", Status: "failure"},
			},
			args: []string{"ls", "repo/name"},
			expected: []*crow.Pipeline{
				{ID: 1, Branch: "main", Event: "push", Status: "success"},
				{ID: 2, Branch: "develop", Event: "pull_request", Status: "running"},
				{ID: 3, Branch: "main", Event: "push", Status: "failure"},
			},
		},
		{
			name:   "limit results",
			repoID: 1,
			pipelines: []*crow.Pipeline{
				{ID: 1, Branch: "main", Event: "push", Status: "success"},
				{ID: 2, Branch: "develop", Event: "pull_request", Status: "running"},
				{ID: 3, Branch: "main", Event: "push", Status: "failure"},
			},
			args: []string{"ls", "--limit", "2", "repo/name"},
			expected: []*crow.Pipeline{
				{ID: 1, Branch: "main", Event: "push", Status: "success"},
				{ID: 2, Branch: "develop", Event: "pull_request", Status: "running"},
			},
		},
		{
			name:        "pipeline list error",
			repoID:      1,
			pipelineErr: errors.New("pipeline error"),
			args:        []string{"ls", "repo/name"},
			wantErr:     errors.New("pipeline error"),
		},
	}

	for _, tt := range testtases {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := mocks.NewClient(t)
			mockClient.On("PipelineList", mock.Anything, mock.Anything).Return(func(_ int64, opt crow.PipelineListOptions) ([]*crow.Pipeline, error) {
				if tt.pipelineErr != nil {
					return nil, tt.pipelineErr
				}
				if opt.Page == 1 {
					return tt.pipelines, nil
				}
				return []*crow.Pipeline{}, nil
			}).Maybe()
			mockClient.On("RepoLookup", mock.Anything).Return(&crow.Repo{ID: tt.repoID}, nil)

			command := buildPipelineListCmd()
			command.Writer = io.Discard
			command.Action = func(_ context.Context, c *cli.Command) error {
				pipelines, err := pipelineList(c, mockClient)
				if tt.wantErr != nil {
					assert.EqualError(t, err, tt.wantErr.Error())
					return nil
				}

				assert.NoError(t, err)
				assert.EqualValues(t, tt.expected, pipelines)

				return nil
			}

			_ = command.Run(context.Background(), tt.args)
		})
	}
}
