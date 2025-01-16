package main

import (
	"testing"

	"github.com/crowci/crow/v3/cmd/server/openapi"
	"github.com/stretchr/testify/assert"
)

func TestSetupOpenApiStaticConfig(t *testing.T) {
	setupOpenAPIStaticConfig()
	assert.Equal(t, "/api", openapi.SwaggerInfo.BasePath)
}
