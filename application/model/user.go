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

package model

import (
	"errors"

	"encoding/gob"

	"github.com/admpub/nging/application/dbschema"
	"github.com/admpub/nging/application/library/common"
	"github.com/admpub/nging/application/model/base"
	"github.com/webx-top/com"
	"github.com/webx-top/db"
	"github.com/webx-top/echo"
)

func init() {
	gob.Register(&dbschema.User{})
}

func NewUser(ctx echo.Context) *User {
	return &User{
		User: &dbschema.User{},
		Base: base.New(ctx),
	}
}

type User struct {
	*dbschema.User
	*base.Base
}

func (u *User) Exists(username string) (bool, error) {
	n, e := u.Param().SetArgs(db.Cond{`username`: username}).Count()
	return n > 0, e
}

func (u *User) Exists2(username string, excludeUID uint) (bool, error) {
	n, e := u.Param().SetArgs(db.And(
		db.Cond{`username`: username},
		db.Cond{`id`: db.NotEq(excludeUID)},
	)).Count()
	return n > 0, e
}

func (u *User) CheckPasswd(username string, password string) (exists bool, err error) {
	exists = true
	err = u.Get(nil, `username`, username)
	if err != nil {
		if err == db.ErrNoMoreRows {
			exists = false
		}
		return
	}
	if u.User.Disabled == `Y` {
		err = errors.New(u.T(`该用户已被禁用`))
		return
	}
	if u.User.Password != com.MakePassword(password, u.User.Salt) {
		err = errors.New(u.T(`密码不正确`))
	}
	return
}

func (u *User) check(editMode bool) (err error) {
	ctx := u.Base.Context
	if len(u.Username) == 0 {
		return ctx.E(`用户名不能为空`)
	}
	if len(u.Email) == 0 {
		return ctx.E(`Email不能为空`)
	}
	if !com.IsUsername(u.Username) {
		return errors.New(ctx.T(`用户名不能包含特殊字符(只能由字母、数字、下划线和汉字组成)`))
	}
	if !ctx.Validate(`email`, u.Email, `email`).Ok() {
		return ctx.E(`Email地址"%s"格式不正确`, u.Email)
	}
	if len(u.Mobile) > 0 && !ctx.Validate(`mobile`, u.Mobile, `mobile`).Ok() {
		return ctx.E(`手机号"%s"格式不正确`, u.Mobile)
	}
	if !editMode || ctx.Form(`modifyPwd`) == `1` {
		if len(u.Password) < 8 {
			return ctx.E(`密码不能少于8个字符`)
		}
		if ctx.Form(`confirmPwd`) != u.Password {
			return ctx.E(`密码与确认密码不一致`)
		}
	}
	if len(u.Disabled) == 0 {
		u.Disabled = `N`
	}
	if len(u.Online) == 0 {
		u.Online = `N`
	}
	var exists bool
	if editMode {
		exists, err = u.Exists2(u.Username, u.Id)
	} else {
		exists, err = u.Exists(u.Username)
	}
	if err != nil {
		return
	}
	if exists {
		err = ctx.E(`用户名已经存在`)
	}
	return
}

func (u *User) Add() (err error) {
	err = u.check(false)
	if err != nil {
		return
	}
	u.Salt = com.Salt()
	u.Password = com.MakePassword(u.Password, u.Salt)
	_, err = u.User.Add()
	if err == nil {
		if len(u.Avatar) > 0 {
			var newPath string
			newPath, err = common.MoveAvatarToUserDir(u.Base.Context, u.Avatar, `user-avatar`, uint64(u.Id))
			if err == nil && newPath != u.Avatar {
				err = u.SetField(nil, `avatar`, newPath, db.Cond{`id`: u.Id})
			}
		}
	}
	return
}

func (u *User) UpdateField(uid uint, set map[string]interface{}) (err error) {
	err = u.check(true)
	if err != nil {
		return
	}
	ctx := u.Base.Context
	if ctx.Form(`modifyPwd`) == `1` {
		u.Password = com.MakePassword(u.Password, u.Salt)
		set[`password`] = u.Password
	}
	err = u.Param().SetSend(set).SetArgs(`id`, uid).Update()
	if err == nil && len(u.Avatar) == 0 {
		err = common.RemoveAvatar(`user-avatar`, uint64(uid))
	}
	return
}

func (u *User) NeedCheckU2F(uid uint) bool {
	u2f := &dbschema.UserU2f{}
	n, _ := u2f.Count(nil, `uid`, uid)
	return n > 0
}

func (u *User) GetUserAllU2F(uid uint) ([]*dbschema.UserU2f, error) {
	u2f := &dbschema.UserU2f{}
	all := []*dbschema.UserU2f{}
	_, err := u2f.ListByOffset(&all, nil, 0, -1, `uid`, uid)
	return all, err
}

func (u *User) U2F(uid uint, typ string) (u2f *dbschema.UserU2f, err error) {
	u2f = &dbschema.UserU2f{}
	err = u2f.Get(nil, db.And(db.Cond{`uid`: uid}, db.Cond{`type`: typ}))
	return
}

func (u *User) Register(user, pass, email string) error {
	userSchema := &dbschema.User{}
	userSchema.Username = user
	userSchema.Email = email
	userSchema.Salt = com.Salt()
	userSchema.Password = com.MakePassword(pass, userSchema.Salt)
	userSchema.Disabled = `N`
	_, err := userSchema.Add()
	u.User = userSchema
	return err
}

func (u *User) SetSession(users ...*dbschema.User) {
	userCopy := u.ClearPasswordData(users...)
	u.Context.Session().Set(`user`, &userCopy)
}

func (u *User) ClearPasswordData(users ...*dbschema.User) dbschema.User {
	var user dbschema.User
	if len(users) > 0 {
		user = *(users[0])
	} else {
		user = *(u.User)
	}
	user.Password = ``
	user.Salt = ``
	user.SafePwd = ``
	return user
}

func (u *User) UnsetSession() {
	u.Context.Session().Delete(`user`)
}

func (u *User) VerifySession(users ...*dbschema.User) error {
	var user *dbschema.User
	if len(users) > 0 {
		user = users[0]
	} else {
		user, _ = u.Context.Session().Get(`user`).(*dbschema.User)
	}
	if user == nil {
		return common.ErrUserNotLoggedIn
	}
	err := u.Get(nil, db.Cond{`id`: user.Id})
	if err != nil {
		if err != db.ErrNoMoreRows {
			return err
		}
		u.UnsetSession()
		return common.ErrUserNotFound
	}
	if u.User.Updated != user.Updated {
		u.SetSession()
		u.Context.Set(`user`, user)
	}
	return nil
}
