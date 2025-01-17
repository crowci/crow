package permissions

import (
	"github.com/crowci/crow/v3/server/model"
	"github.com/crowci/crow/v3/shared/utils"
)

func NewOwnersAllowlist(owners []string) *OwnersAllowlist {
	return &OwnersAllowlist{owners: utils.SliceToBoolMap(owners)}
}

type OwnersAllowlist struct {
	owners map[string]bool
}

func (o *OwnersAllowlist) IsAllowed(repo *model.Repo) bool {
	return len(o.owners) < 1 || o.owners[repo.Owner]
}
