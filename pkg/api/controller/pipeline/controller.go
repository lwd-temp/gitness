// Copyright 2023 Harness, Inc.
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

package pipeline

import (
	"github.com/harness/gitness/pkg/auth/authz"
	"github.com/harness/gitness/pkg/store"
	"github.com/harness/gitness/types/check"

	"github.com/jmoiron/sqlx"
)

type Controller struct {
	defaultBranch string
	db            *sqlx.DB
	uidCheck      check.PathUID
	repoStore     store.RepoStore
	triggerStore  store.TriggerStore
	authorizer    authz.Authorizer
	pipelineStore store.PipelineStore
}

func NewController(
	db *sqlx.DB,
	uidCheck check.PathUID,
	authorizer authz.Authorizer,
	repoStore store.RepoStore,
	triggerStore store.TriggerStore,
	pipelineStore store.PipelineStore,
) *Controller {
	return &Controller{
		db:            db,
		uidCheck:      uidCheck,
		repoStore:     repoStore,
		triggerStore:  triggerStore,
		authorizer:    authorizer,
		pipelineStore: pipelineStore,
	}
}