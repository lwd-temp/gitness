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

package repo

import (
	"net/http"

	"github.com/harness/gitness/pkg/api/controller/repo"
	"github.com/harness/gitness/pkg/api/render"
	"github.com/harness/gitness/pkg/api/request"
)

// HandleBlame returns the git blame output for a file.
func HandleBlame(repoCtrl *repo.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		session, _ := request.AuthSessionFrom(ctx)

		repoRef, err := request.GetRepoRefFromPath(r)
		if err != nil {
			render.TranslatedUserError(w, err)
			return
		}

		path := request.GetOptionalRemainderFromPath(r)

		// line_from is optional, skipped if set to 0
		lineFrom, err := request.QueryParamAsPositiveInt64OrDefault(r, request.QueryParamLineFrom, 0)
		if err != nil {
			render.TranslatedUserError(w, err)
			return
		}

		// line_to is optional, skipped if set to 0
		lineTo, err := request.QueryParamAsPositiveInt64OrDefault(r, request.QueryParamLineTo, 0)
		if err != nil {
			render.TranslatedUserError(w, err)
			return
		}

		gitRef := request.GetGitRefFromQueryOrDefault(r, "")

		stream, err := repoCtrl.Blame(ctx, session, repoRef, gitRef, path, int(lineFrom), int(lineTo))
		if err != nil {
			render.TranslatedUserError(w, err)
			return
		}

		render.JSONArrayDynamic(ctx, w, stream)
	}
}