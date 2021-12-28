/*
 * Copyright (c) 2021 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package modfetch

import (
	"path/filepath"

	"golang.org/x/mod/module"

	"github.com/goplus/gop/env"
)

// -----------------------------------------------------------------------------

var (
	GOMODCACHE = env.GOMODCACHE()
)

func DownloadCachePath(mod module.Version) (string, error) {
	encPath, err := module.EscapePath(mod.Path)
	if err != nil {
		return "", err
	}
	return filepath.Join(GOMODCACHE, "cache/download", encPath, "@v", mod.Version+".zip"), nil
}

func ModCachePath(mod module.Version) (string, error) {
	encPath, err := module.EscapePath(mod.Path)
	if err != nil {
		return "", err
	}
	return filepath.Join(GOMODCACHE, encPath+"@"+mod.Version), nil
}

// -----------------------------------------------------------------------------
