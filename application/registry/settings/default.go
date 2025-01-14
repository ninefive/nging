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

package settings

import (
	"strings"

	"github.com/admpub/log"
	"github.com/admpub/nging/application/dbschema"
	"github.com/webx-top/db"
	"github.com/webx-top/echo"
)

var configDefaults = map[string]map[string]*dbschema.Config{
	`base`: {
		`apiKey`: {
			Key:         `apiKey`,
			Label:       `API密钥`,
			Description: ``,
			Value:       ``,
			Group:       `base`,
			Type:        `text`,
			Sort:        0,
			Disabled:    `N`,
		},
		`debug`: {
			Key:         `debug`,
			Label:       `调试模式`,
			Description: ``,
			Value:       `1`,
			Group:       `base`,
			Type:        `text`,
			Sort:        0,
			Disabled:    `N`,
		},
	},
	`smtp`: {
		`username`: {
			Key:         `username`,
			Label:       `登录名`,
			Description: ``,
			Value:       ``,
			Group:       `smtp`,
			Type:        `text`,
			Sort:        0,
			Disabled:    `N`,
		},
		`password`: {
			Key:         `password`,
			Label:       `密码`,
			Description: ``,
			Value:       ``,
			Group:       `smtp`,
			Type:        `text`,
			Sort:        0,
			Disabled:    `N`,
			Encrypted:   `Y`,
		},
		`host`: {
			Key:         `host`,
			Label:       `服务器`,
			Description: ``,
			Value:       `smtp.exmail.qq.com`,
			Group:       `smtp`,
			Type:        `text`,
			Sort:        0,
			Disabled:    `N`,
		},
		`port`: {
			Key:         `port`,
			Label:       `端口`,
			Description: ``,
			Value:       `465`,
			Group:       `smtp`,
			Type:        `text`,
			Sort:        0,
			Disabled:    `N`,
		},
		`secure`: {
			Key:         `secure`,
			Label:       `认证方式`,
			Description: ``,
			Value:       `SSL`,
			Group:       `smtp`,
			Type:        `text`,
			Sort:        0,
			Disabled:    `N`,
		},
		`identity`: {
			Key:         `identity`,
			Label:       `身份`,
			Description: ``,
			Value:       ``,
			Group:       `smtp`,
			Type:        `text`,
			Sort:        0,
			Disabled:    `N`,
		},
		`timeout`: {
			Key:         `timeout`,
			Label:       `超时时间`,
			Description: ``,
			Value:       ``,
			Group:       `smtp`,
			Type:        `text`,
			Sort:        0,
			Disabled:    `N`,
		},
		`engine`: {
			Key:         `engine`,
			Label:       `发送引擎`,
			Description: ``,
			Value:       `mail`,
			Group:       `smtp`,
			Type:        `text`,
			Sort:        0,
			Disabled:    `N`,
		},
		`from`: {
			Key:         `from`,
			Label:       `发信人地址`,
			Description: ``,
			Value:       ``,
			Group:       `smtp`,
			Type:        `text`,
			Sort:        0,
			Disabled:    `N`,
		},
		`queueSize`: {
			Key:         `queueSize`,
			Label:       `并发数量`,
			Description: ``,
			Value:       `10`,
			Group:       `smtp`,
			Type:        `text`,
			Sort:        0,
			Disabled:    `N`,
		},
	},
	`log`: {
		`saveFile`: {
			Key:         `saveFile`,
			Label:       `保存路径`,
			Description: ``,
			Value:       ``,
			Group:       `log`,
			Type:        `text`,
			Sort:        0,
			Disabled:    `N`,
		},
		`fileMaxBytes`: {
			Key:         `fileMaxBytes`,
			Label:       `日志文件尺寸`,
			Description: ``,
			Value:       ``,
			Group:       `log`,
			Type:        `text`,
			Sort:        0,
			Disabled:    `N`,
		},
		`targets`: {
			Key:         `targets`,
			Label:       `输出`,
			Description: ``,
			Value:       `console`,
			Group:       `log`,
			Type:        `list`,
			Sort:        0,
			Disabled:    `N`,
		},
		`colorable`: {
			Key:         `colorable`,
			Label:       `彩色日志`,
			Description: ``,
			Value:       `0`,
			Group:       `log`,
			Type:        `text`,
			Sort:        0,
			Disabled:    `N`,
		},
	},
}

func AddDefaultConfig(group string, configs map[string]*dbschema.Config) {
	if strings.Contains(group, `.`) {
		panic(`Group name is not allowed to contain ".": ` + group)
	}
	if _, y := configDefaults[group]; !y {
		configDefaults[group] = map[string]*dbschema.Config{}
	}
	for key, conf := range configs {
		if conf.Group != group {
			conf.Group = group
		}
		configDefaults[group][key] = conf
	}
}

func GetDefaultConfig(group string) map[string]*dbschema.Config {
	r, _ := configDefaults[group]
	return r
}

func GetDefaultConfigOk(group string) (map[string]*dbschema.Config, bool) {
	r, y := configDefaults[group]
	return r, y
}

func ConfigDefaultsAsStore() echo.H {
	return configAsStore(configDefaults)
}

func ConfigDefaults() map[string]map[string]*dbschema.Config {
	return configDefaults
}

func Init() error {
	log.Debug(`Initialize the configuration data in the database table`)
	m := &dbschema.Config{}
	_, err := m.ListByOffset(nil, func(r db.Result) db.Result {
		return r.Group(`group`)
	}, 0, -1)
	if err != nil {
		return err
	}
	existsList := m.Objects()
	existsIndex := map[string]int{}
	for index, row := range existsList {
		existsIndex[row.Group] = index
	}
	for _, setting := range Settings() {
		group := setting.Group
		if _, ok := existsIndex[group]; ok {
			continue
		}
		gs, ok := GetDefaultConfigOk(group)
		if !ok {
			continue
		}
		for _, conf := range gs {
			_, err = conf.Add()
			if err != nil {
				return err
			}
		}
	}
	return err
}

// ConfigAsStore {Group:{Key:ValueObject}}
func ConfigAsStore() echo.H {
	r := echo.H{}
	m := &dbschema.Config{}
	m.ListByOffset(nil, nil, 0, -1, db.Cond{`disabled`: `N`})
	for _, row := range m.Objects() {
		decoder := GetDecoder(row.Group)
		res, err := DecodeConfigValue(row, decoder)
		if err != nil {
			log.Error(err)
			continue
		}
		value := res.Get(`ValueObject`, row.Value)
		if _, y := r[row.Group]; !y {
			r[row.Group] = echo.H{row.Key: value}
		} else {
			r.Store(row.Group).Set(row.Key, value)
		}
	}
	return r
}

// {Group:{Key:ValueObject}}
func configAsStore(configList map[string]map[string]*dbschema.Config) echo.H {
	r := echo.H{}
	for group, configs := range configList {
		v := echo.H{}
		decoder := GetDecoder(group)
		for key, row := range configs {
			if row.Disabled == `Y` {
				continue
			}
			res, err := DecodeConfigValue(row, decoder)
			if err != nil {
				log.Error(err)
				continue
			}
			value := res.Get(`ValueObject`, row.Value)
			v.Set(key, value)
		}
		r.Set(group, v)
	}
	return r
}
