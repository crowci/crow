package docker

import (
	backend "github.com/crowci/crow/v3/pipeline/backend/types"
	"github.com/go-viper/mapstructure/v2"
)

// BackendOptions defines all the advanced options for the docker backend.
type BackendOptions struct {
	User string `mapstructure:"user"`
}

func parseBackendOptions(step *backend.Step) (BackendOptions, error) {
	var result BackendOptions
	if step == nil || step.BackendOptions == nil {
		return result, nil
	}
	err := mapstructure.WeakDecode(step.BackendOptions[EngineName], &result)
	return result, err
}
