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
package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"github.com/admpub/confl"
	"github.com/admpub/log"

	//"github.com/admpub/license_gen/lib"

	"github.com/admpub/nging/application/library/caddy"
	"github.com/admpub/nging/application/library/ftp"
	"github.com/admpub/securecookie"
	"github.com/webx-top/codec"
	"github.com/webx-top/com"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/middleware/language"
)

func NewConfig() *Config {
	c := &Config{}
	c.ConfigInDB = NewConfigInDB(c)
	return c
}

type Config struct {
	DB DB `json:"db"`

	Sys struct {
		VhostsfileDir          string            `json:"vhostsfileDir"`
		AllowIP                []string          `json:"allowIP"`
		SSLAuto                bool              `json:"sslAuto"`
		SSLEmail               string            `json:"sslEmail"`
		SSLHosts               []string          `json:"sslHosts"`
		SSLCacheDir            string            `json:"sslCacheDir"`
		SSLKeyFile             string            `json:"sslKeyFile"`
		SSLCertFile            string            `json:"sslCertFile"`
		EditableFileExtensions map[string]string `json:"editableFileExtensions"`
		EditableFileMaxSize    string            `json:"editableFileMaxSize"`
		EditableFileMaxBytes   int64             `json:"editableFileMaxBytes"`
		ErrorPages             map[int]string    `json:"errorPages"`
		CmdTimeout             string            `json:"cmdTimeout"`
		CmdTimeoutDuration     time.Duration     `json:"-"`
		ShowExpirationTime     int64             `json:"showExpirationTime"` //显示过期时间：0为始终显示；大于0为距离剩余到期时间多少秒的时候显示；小于0为不显示
	} `json:"sys"`

	Cron struct {
		PoolSize int    `json:"poolSize"`
		Template string `json:"template"` //发信模板
	} `json:"cron"`

	Cookie struct {
		Domain   string `json:"domain"`
		MaxAge   int    `json:"maxAge"`
		Path     string `json:"path"`
		HttpOnly bool   `json:"httpOnly"`
		HashKey  string `json:"hashKey"`
		BlockKey string `json:"blockKey"`
	} `json:"cookie"`

	Caddy    caddy.Config    `json:"caddy"`
	FTP      ftp.Config      `json:"ftp"`
	Language language.Config `json:"language"`
	Download struct {
		SavePath string `json:"savePath"`
	} `json:"download"`
	//License lib.LicenseData `json:"license,omitempty"`

	*ConfigInDB `json:"-"`

	connectedDB bool
}

// ConnectedDB 数据库是否已连接，如果没有连接则自动连接
func (c *Config) ConnectedDB(autoConn ...bool) bool {
	if c.connectedDB {
		return c.connectedDB
	}
	n := len(autoConn)
	if n == 0 || (n > 0 && autoConn[0]) {
		err := c.connectDB()
		if err != nil {
			log.Error(err)
		}
	}
	return c.connectedDB
}

func (c *Config) connectDB() error {
	err := ConnectDB(c)
	if err != nil {
		return err
	}
	c.connectedDB = true
	return nil
}

func (c *Config) SetDebug(on bool) *Config {
	c.ConfigInDB.SetDebug(on)
	return c
}

func (c *Config) Codec() codec.Codec {
	return codec.Default
}

func (c *Config) Encode(raw string, keys ...string) string {
	var key string
	if len(keys) > 0 && len(keys[0]) > 0 {
		key = com.Md5(keys[0])
	} else {
		key = c.Cookie.HashKey
	}
	return codec.Default.Encode(raw, key)
}

func (c *Config) Decode(encrypted string, keys ...string) string {
	if len(encrypted) == 0 {
		return ``
	}
	var key string
	if len(keys) > 0 && len(keys[0]) > 0 {
		key = com.Md5(keys[0])
	} else {
		key = c.Cookie.HashKey
	}
	return codec.Default.Decode(encrypted, key)
}

func (c *Config) InitSecretKey() *Config {
	c.Cookie.BlockKey = c.GenerateRandomKey()
	c.Cookie.HashKey = c.GenerateRandomKey()
	return c
}

func (c *Config) GenerateRandomKey() string {
	return com.Md5(string(securecookie.GenerateRandomKey(32)))
}

func (c *Config) Reload(newConfig *Config) error {
	engines := []string{}
	if !reflect.DeepEqual(newConfig.Caddy, c.Caddy) {
		engines = append(engines, `caddy`)
	}
	if !reflect.DeepEqual(newConfig.FTP, c.FTP) {
		engines = append(engines, `ftp`)
	}
	DefaultCLIConfig.Reload(engines...)
	return nil
}

func (c *Config) AsDefault() {
	echo.Set(`DefaultConfig`, c)
	c.ConfigInDB.Init()
	DefaultConfig = c
}

func (c *Config) SaveToFile() error {
	b, err := confl.Marshal(c)
	if err == nil {
		err = os.MkdirAll(filepath.Dir(DefaultCLIConfig.Conf), os.ModePerm)
		if err == nil {
			/*
				_, e := os.Stat(DefaultCLIConfig.Conf + `.sample`)
				if e != nil {
					if os.IsNotExist(e) {
						old, err := ioutil.ReadFile(DefaultCLIConfig.Conf)
						if err == nil {
							err = ioutil.WriteFile(DefaultCLIConfig.Conf+`.sample`, old, os.ModePerm)
						}
						if err != nil {
							return err
						}
					}
				}
			*/
			err = ioutil.WriteFile(DefaultCLIConfig.Conf, b, os.ModePerm)
		}
	}
	return err
}
