// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	
	"time"
)

type FtpUser struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*FtpUser
	namer   func(string) string
	connID  int
	
	Id          	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"" json:"id" xml:"id"`
	Username    	string  	`db:"username" bson:"username" comment:"用户名" json:"username" xml:"username"`
	Password    	string  	`db:"password" bson:"password" comment:"密码" json:"password" xml:"password"`
	Banned      	string  	`db:"banned" bson:"banned" comment:"是否禁止连接" json:"banned" xml:"banned"`
	Directory   	string  	`db:"directory" bson:"directory" comment:"授权目录(一行一个) " json:"directory" xml:"directory"`
	IpWhitelist 	string  	`db:"ip_whitelist" bson:"ip_whitelist" comment:"IP白名单(一行一个) " json:"ip_whitelist" xml:"ip_whitelist"`
	IpBlacklist 	string  	`db:"ip_blacklist" bson:"ip_blacklist" comment:"IP黑名单(一行一个) " json:"ip_blacklist" xml:"ip_blacklist"`
	Created     	uint    	`db:"created" bson:"created" comment:"创建时间 " json:"created" xml:"created"`
	Updated     	uint    	`db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
	GroupId     	uint    	`db:"group_id" bson:"group_id" comment:"用户组" json:"group_id" xml:"group_id"`
}

func (this *FtpUser) Trans() *factory.Transaction {
	return this.trans
}

func (this *FtpUser) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *FtpUser) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *FtpUser) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *FtpUser) Objects() []*FtpUser {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *FtpUser) NewObjects() *[]*FtpUser {
	this.objects = []*FtpUser{}
	return &this.objects
}

func (this *FtpUser) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *FtpUser) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *FtpUser) Name_() string {
	if this.namer != nil {
		return this.namer("ftp_user")
	}
	return factory.TableNamerGet("ftp_user")(this)
}

func (this *FtpUser) FullName_(connID ...int) string {
	if len(connID) > 0 {
		return factory.DefaultFactory.Cluster(connID[0]).Table(this.Name_())
	}
	return factory.DefaultFactory.Cluster(this.connID).Table(this.Name_())
}

func (this *FtpUser) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *FtpUser) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *FtpUser) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *FtpUser) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *FtpUser) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *FtpUser) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Banned) == 0 { this.Banned = "N" }
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	return
}

func (this *FtpUser) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	this.Updated = uint(time.Now().Unix())
	if len(this.Banned) == 0 { this.Banned = "N" }
	return this.Setter(mw, args...).SetSend(this).Update()
}

func (this *FtpUser) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *FtpUser) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) error {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *FtpUser) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) error {
	
	if val, ok := kvset["banned"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["banned"] = "N" } }
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *FtpUser) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		this.Updated = uint(time.Now().Unix())
	if len(this.Banned) == 0 { this.Banned = "N" }
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Banned) == 0 { this.Banned = "N" }
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	return 
}

func (this *FtpUser) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *FtpUser) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *FtpUser) Reset() *FtpUser {
	this.Id = 0
	this.Username = ``
	this.Password = ``
	this.Banned = ``
	this.Directory = ``
	this.IpWhitelist = ``
	this.IpBlacklist = ``
	this.Created = 0
	this.Updated = 0
	this.GroupId = 0
	return this
}

func (this *FtpUser) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Username"] = this.Username
	r["Password"] = this.Password
	r["Banned"] = this.Banned
	r["Directory"] = this.Directory
	r["IpWhitelist"] = this.IpWhitelist
	r["IpBlacklist"] = this.IpBlacklist
	r["Created"] = this.Created
	r["Updated"] = this.Updated
	r["GroupId"] = this.GroupId
	return r
}

