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

package schema_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/crowci/crow/v3/pipeline/frontend/yaml/linter/schema"
	"github.com/stretchr/testify/assert"
)

func TestSchema(t *testing.T) {
	t.Parallel()

	testTable := []struct {
		name     string
		testFile string
		fail     bool
	}{
		{
			name:     "Clone",
			testFile: ".crow/test-clone.yaml",
		},
		{
			name:     "Clone skip",
			testFile: ".crow/test-clone-skip.yaml",
		},
		{
			name:     "Matrix",
			testFile: ".crow/test-matrix.yaml",
		},
		{
			name:     "Multi Pipeline",
			testFile: ".crow/test-multi.yaml",
		},
		{
			name:     "Plugin",
			testFile: ".crow/test-plugin.yaml",
		},
		{
			name:     "Run on",
			testFile: ".crow/test-run-on.yaml",
		},
		{
			name:     "Service",
			testFile: ".crow/test-service.yaml",
		},
		{
			name:     "Step",
			testFile: ".crow/test-step.yaml",
		},
		{
			name:     "When",
			testFile: ".crow/test-when.yaml",
		},
		{
			name:     "Workspace",
			testFile: ".crow/test-workspace.yaml",
		},
		{
			name:     "Labels",
			testFile: ".crow/test-labels.yaml",
		},
		{
			name:     "Map and Sequence Merge",
			testFile: ".crow/test-merge-map-and-sequence.yaml",
		},
		{
			name:     "Broken Config",
			testFile: ".crow/test-broken.yaml",
			fail:     true,
		},
		{
			name:     "Array syntax",
			testFile: ".crow/test-array-syntax.yaml",
			fail:     false,
		},
		{
			name:     "Step DAG syntax",
			testFile: ".crow/test-dag.yaml",
			fail:     false,
		},
		{
			name:     "Custom backend",
			testFile: ".crow/test-custom-backend.yaml",
			fail:     false,
		},
		{
			name:     "Broken Plugin by environment",
			testFile: ".crow/test-broken-plugin.yaml",
			fail:     true,
		},
		{
			name:     "Broken Plugin by commands",
			testFile: ".crow/test-broken-plugin2.yaml",
			fail:     true,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			fi, err := os.Open(tt.testFile)
			assert.NoError(t, err, "could not open test file")
			defer fi.Close()
			configErrors, err := schema.Lint(fi)
			if tt.fail {
				if len(configErrors) == 0 {
					assert.Error(t, err, "Expected config errors but got none")
				}
			} else {
				assert.NoError(t, err, fmt.Sprintf("Validation failed: %v", configErrors))
			}
		})
	}
}
