// Copyright 2023 Woodpecker Authors
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
	"github.com/crowci/crow/v3/cmd/server/openapi"
	"github.com/crowci/crow/v3/version"
)

// Generate docs/openapi.json via:
//go:generate go run github.com/swaggo/swag/cmd/swag init -g cmd/server/openapi.go --outputTypes go -output openapi -d ../../
//go:generate go run openapi_json_gen.go openapi.go
//go:generate go run github.com/getkin/kin-openapi/cmd/validate ../../docs/openapi.json

// setupOpenAPIStaticConfig initializes static content (version) for the OpenAPI config.
//
//	@title			Crow CI API
//	@description	Crow CI is a lightweight, community-driven CI application for self-hosted environments.
//	@description	To get a personal access token (PAT) for authentication, please log in
//	@description	and go to you personal profile page, by clicking the user icon at the top right.
//	@BasePath		/api
//	@contact.name	Crow CI
//	@contact.url	https://crowci.dev/
func setupOpenAPIStaticConfig() {
	openapi.SwaggerInfo.Version = version.String()
}
