// Copyright 2024 Woodpecker Authors
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

package migration

import (
	"src.techknowlogick.com/xormigrate"
	"xorm.io/xorm"
)

var removeRepoScm = xormigrate.Migration{
	ID: "remove-repo-scm",
	MigrateSession: func(sess *xorm.Session) (err error) {
		type repos struct {
			SCMKind string `xorm:"varchar(50) 'scm'"`
		}

		// ensure columns to drop exist
		if err := sess.Sync(new(repos)); err != nil {
			return err
		}

		return dropTableColumns(sess, "repos", "scm")
	},
}
