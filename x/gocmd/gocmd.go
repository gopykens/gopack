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

package gocmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/goplus/gop/x/gopenv"
	"github.com/goplus/mod/env"
)

type GopEnv = env.Gop

type Config struct {
	Gop   *GopEnv
	Flags []string
	Run   func(cmd *exec.Cmd) error
}

// -----------------------------------------------------------------------------

func doWithArgs(op string, conf *Config, args ...string) (err error) {
	if conf == nil {
		conf = new(Config)
	}
	exargs := make([]string, 1, 16)
	exargs[0] = op
	exargs = appendLdflags(exargs, conf.Gop)
	exargs = append(exargs, conf.Flags...)
	exargs = append(exargs, args...)
	cmd := exec.Command("go", exargs...)
	run := conf.Run
	if run == nil {
		run = runCmd
	}
	return run(cmd)
}

func runCmd(cmd *exec.Cmd) (err error) {
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// -----------------------------------------------------------------------------

const (
	ldFlagVersion   = "-X \"github.com/goplus/gop/env.buildVersion=%s\""
	ldFlagBuildDate = "-X \"github.com/goplus/gop/env.buildDate=%s\""
	ldFlagGopRoot   = "-X \"github.com/goplus/gop/env.defaultGopRoot=%s\""
)

const (
	ldFlagAll = ldFlagVersion + " " + ldFlagBuildDate + " " + ldFlagGopRoot
)

func loadFlags(env *GopEnv) string {
	return fmt.Sprintf(ldFlagAll, env.Version, env.BuildDate, env.Root)
}

func appendLdflags(exargs []string, env *GopEnv) []string {
	if env == nil {
		env = gopenv.Get()
	}
	return append(exargs, "-ldflags", loadFlags(env))
}

// -----------------------------------------------------------------------------
