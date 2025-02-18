/*
   Nging is a toolbox for webmasters
   Copyright (C) 2018-present  Wenhui Shen <swh@admpub.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package cmd

import (
	stdLog "log"
	"os"
	"path/filepath"

	"github.com/admpub/log"
	"github.com/admpub/nging/application/cmd/event"
	"github.com/admpub/nging/application/library/config"
	"github.com/admpub/nging/application/library/license"
	"github.com/spf13/cobra"
	"github.com/webx-top/echo"
)

func Init() {
	config.Version.Number = echo.String(`VERSION`)
	config.Version.Label = echo.String(`LABEL`)
	if !event.Bindata {
		license.SkipLicenseCheck = true
		event.Develop = true
	}
	buildTime := echo.String(`BUILD_TIME`)
	if len(buildTime) == 0 {
		gitFile := filepath.Join(echo.Wd(), `.git/index`)
		f, err := os.Stat(gitFile)
		if err == nil {
			buildTime = f.ModTime().Format(`20060102150405`)
		}
		echo.Set(`BUILD_TIME`, buildTime)
	}
	config.Version.BuildTime = buildTime
	stdLog.SetOutput(log.Writer(log.LevelInfo))
	event.Licensed = license.Ok(``)
	event.MustLicensed = echo.Bool(`MUST_LICENSED`)
	config.Version.Licensed = event.Licensed
	config.Version.CommitID = echo.String(`COMMIT`)
	config.Version.DBSchema = echo.Float64(`SCHEMA_VER`)
	rootCmd.Short = event.Short
}

func Add(cmds ...*cobra.Command) {
	rootCmd.AddCommand(cmds...)
}
