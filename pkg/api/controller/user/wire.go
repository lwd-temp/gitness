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

package user

import (
	"github.com/harness/gitness/pkg/auth/authz"
	"github.com/harness/gitness/pkg/store"
	"github.com/harness/gitness/types/check"

	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

// WireSet provides a wire set for this package.
var WireSet = wire.NewSet(
	ProvideController,
)

func ProvideController(
	db *sqlx.DB,
	principalUIDCheck check.PrincipalUID,
	authorizer authz.Authorizer,
	principalStore store.PrincipalStore,
	tokenStore store.TokenStore,
	membershipStore store.MembershipStore,
) *Controller {
	return NewController(
		db,
		principalUIDCheck,
		authorizer,
		principalStore,
		tokenStore,
		membershipStore)
}