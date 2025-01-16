package permissions

import (
	"testing"

	"github.com/crowci/crow/v3/server/model"
	"github.com/stretchr/testify/assert"
)

func TestOwnersAllowlist(t *testing.T) {
	ol := NewOwnersAllowlist([]string{"woodpecker-ci"})
	assert.True(t, ol.IsAllowed(&model.Repo{Owner: "woodpecker-ci"}))
	assert.False(t, ol.IsAllowed(&model.Repo{Owner: "not-woodpecker-ci"}))
	empty := NewOwnersAllowlist([]string{})
	assert.True(t, empty.IsAllowed(&model.Repo{Owner: "woodpecker-ci"}))
	assert.True(t, empty.IsAllowed(&model.Repo{Owner: "not-woodpecker-ci"}))
}
