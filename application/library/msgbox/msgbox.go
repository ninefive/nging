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

package msgbox

import (
	"fmt"
	"os"
	"strings"

	"github.com/admpub/go-pretty/table"
	"github.com/admpub/go-pretty/text"
)

func Error(title, content string, args ...interface{}) {
	Render(title, content, `error`, args...)
}

func Success(title, content string, args ...interface{}) {
	Render(title, content, `success`, args...)
}

func Info(title, content string, args ...interface{}) {
	Render(title, content, `info`, args...)
}

func Warn(title, content string, args ...interface{}) {
	Render(title, content, `warn`, args...)
}

func Debug(title, content string, args ...interface{}) {
	Render(title, content, `debug`, args...)
}

func Render(title, content, typ string, args ...interface{}) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{title})
	t.AppendRow([]interface{}{""})
	if len(args) > 0 {
		content = fmt.Sprintf(content, args...)
	}
	for _, row := range strings.Split(content, "\n") {
		t.AppendRow([]interface{}{row})
	}
	t.AppendRow([]interface{}{""})
	t.AppendFooter(table.Row{"Powered by webx.top"})
	t.SetStyle(table.StyleColoredRedWhiteOnBlack)
	headerColor := text.Colors{text.BgRed, text.FgWhite, text.Bold}
	switch typ {
	case `success`:
		headerColor[0] = text.BgGreen
	case `info`:
		headerColor[0] = text.BgBlue
	case `warn`:
		headerColor[0] = text.BgHiYellow
		headerColor[1] = text.FgHiRed
	case `debug`:
		headerColor[0] = text.BgHiBlack
	}
	t.Style().Color.Header = headerColor
	t.Style().Color.Footer = text.Colors{text.BgWhite, text.FgBlack, text.Italic}
	t.Style().Color.Row = text.Colors{text.BgWhite, text.FgBlack}
	t.Style().Color.RowAlternate = text.Colors{text.BgWhite, text.FgBlack}
	t.SetAlignHeader([]text.Align{text.AlignCenter})
	t.SetAlignFooter([]text.Align{text.AlignRight})
	t.Render()
}
