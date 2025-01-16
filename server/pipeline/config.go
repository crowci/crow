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
	forge_types "github.com/crowci/crow/v3/server/forge/types"
	"github.com/crowci/crow/v3/server/model"
	"github.com/crowci/crow/v3/server/pipeline/stepbuilder"
	"github.com/crowci/crow/v3/server/store"
)

func findOrPersistPipelineConfig(store store.Store, currentPipeline *model.Pipeline, forgeYamlConfig *forge_types.FileMeta) (*model.Config, error) {
	return store.ConfigPersist(&model.Config{
		RepoID: currentPipeline.RepoID,
		Name:   stepbuilder.SanitizePath(forgeYamlConfig.Name),
		Data:   forgeYamlConfig.Data,
	})
}
