// Copyright 2018 Drone.IO Inc.
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

package session

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/crowci/crow/v3/server"
	"github.com/crowci/crow/v3/server/model"
	"github.com/crowci/crow/v3/server/store"
	"github.com/crowci/crow/v3/server/store/types"
)

func Repo(c *gin.Context) *model.Repo {
	v, ok := c.Get("repo")
	if !ok {
		return nil
	}
	r, ok := v.(*model.Repo)
	if !ok {
		return nil
	}
	r.Perm = Perm(c)
	return r
}

func SetRepo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			_store   = store.FromContext(c)
			fullName = strings.TrimLeft(c.Param("repo_full_name"), "/")
			_repoID  = c.Param("repo_id")
			user     = User(c)
		)

		var repo *model.Repo
		var err error
		if _repoID != "" {
			var repoID int64
			repoID, err = strconv.ParseInt(_repoID, 10, 64)
			if err != nil {
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
			repo, err = _store.GetRepo(repoID)
		} else {
			repo, err = _store.GetRepoName(fullName)
		}

		if repo != nil && err == nil {
			c.Set("repo", repo)
			c.Next()
			return
		}

		// debugging
		log.Debug().Err(err).Msgf("cannot find repository %s", fullName)

		if user == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if errors.Is(err, types.RecordNotExist) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		_ = c.AbortWithError(http.StatusInternalServerError, err)
	}
}

func Perm(c *gin.Context) *model.Perm {
	v, ok := c.Get("perm")
	if !ok {
		return nil
	}
	u, ok := v.(*model.Perm)
	if !ok {
		return nil
	}
	return u
}

func SetPerm() gin.HandlerFunc {
	return func(c *gin.Context) {
		_store := store.FromContext(c)
		user := User(c)
		repo := Repo(c)
		_forge, err := server.Config.Services.Manager.ForgeFromRepo(repo)
		if err != nil {
			log.Error().Err(err).Msg("Cannot get forge from repo")
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		perm := new(model.Perm)

		if user != nil {
			var err error
			perm, err = _store.PermFind(user, repo)
			if err != nil {
				log.Error().Err(err).Msgf("error fetching permission for %s %s",
					user.Login, repo.FullName)
			}
			if time.Unix(perm.Synced, 0).Add(time.Hour).Before(time.Now()) {
				_repo, err := _forge.Repo(c, user, repo.ForgeRemoteID, repo.Owner, repo.Name)
				if err == nil {
					log.Debug().Msgf("synced user permission for %s %s", user.Login, repo.FullName)
					perm = _repo.Perm
					perm.Repo = repo
					perm.RepoID = repo.ID
					perm.UserID = user.ID
					perm.Synced = time.Now().Unix()
					if err := _store.PermUpsert(perm); err != nil {
						_ = c.AbortWithError(http.StatusInternalServerError, err)
						return
					}
				}
			}
		}

		if perm == nil {
			perm = new(model.Perm)
		}

		if user != nil && user.Admin {
			perm.Pull = true
			perm.Push = true
			perm.Admin = true
		}

		if repo.Visibility == model.VisibilityPublic || (repo.Visibility == model.VisibilityInternal && user != nil) {
			perm.Pull = true
		}

		if user != nil {
			log.Debug().Msgf("%s granted %+v permission to %s",
				user.Login, perm, repo.FullName)
		} else {
			log.Debug().Msgf("guest granted %+v to %s", perm, repo.FullName)
		}

		c.Set("perm", perm)
		c.Next()
	}
}

func MustPull(c *gin.Context) {
	user := User(c)
	perm := Perm(c)

	if perm.Pull {
		c.Next()
		return
	}

	// debugging
	if user != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Debug().Msgf("user %s denied read access to %s",
			user.Login, c.Request.URL.Path)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		log.Debug().Msgf("guest denied read access to %s %s",
			c.Request.Method,
			c.Request.URL.Path,
		)
	}
}

func MustPush(c *gin.Context) {
	user := User(c)
	perm := Perm(c)

	// if the user has push access, immediately proceed
	// the middleware execution chain.
	if perm.Push {
		c.Next()
		return
	}

	// debugging
	if user != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Debug().Msgf("user %s denied write access to %s",
			user.Login, c.Request.URL.Path)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		log.Debug().Msgf("guest denied write access to %s %s",
			c.Request.Method,
			c.Request.URL.Path,
		)
	}
}
