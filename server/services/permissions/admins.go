package permissions

import (
	"github.com/crowci/crow/v3/server/model"
	"github.com/crowci/crow/v3/shared/utils"
)

func NewAdmins(admins []string) *Admins {
	return &Admins{admins: utils.SliceToBoolMap(admins)}
}

type Admins struct {
	admins map[string]bool
}

func (a *Admins) IsAdmin(user *model.User) bool {
	return a.admins[user.Login]
}
