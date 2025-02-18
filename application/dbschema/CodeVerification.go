// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	
	"time"
)

type CodeVerification struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*CodeVerification
	namer   func(string) string
	connID  int
	
	Id         	uint64  	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Code       	string  	`db:"code" bson:"code" comment:"验证码" json:"code" xml:"code"`
	Created    	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	OwnerId    	uint64  	`db:"owner_id" bson:"owner_id" comment:"所有者ID" json:"owner_id" xml:"owner_id"`
	OwnerType  	string  	`db:"owner_type" bson:"owner_type" comment:"所有者类型" json:"owner_type" xml:"owner_type"`
	Used       	uint    	`db:"used" bson:"used" comment:"使用时间" json:"used" xml:"used"`
	Purpose    	string  	`db:"purpose" bson:"purpose" comment:"目的" json:"purpose" xml:"purpose"`
	Start      	uint    	`db:"start" bson:"start" comment:"有效时间" json:"start" xml:"start"`
	End        	uint    	`db:"end" bson:"end" comment:"失效时间" json:"end" xml:"end"`
	Disabled   	string  	`db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	SendMethod 	string  	`db:"send_method" bson:"send_method" comment:"发送方式(mobile-手机;email-邮箱)" json:"send_method" xml:"send_method"`
	SendTo     	string  	`db:"send_to" bson:"send_to" comment:"发送目标" json:"send_to" xml:"send_to"`
}

func (this *CodeVerification) Trans() *factory.Transaction {
	return this.trans
}

func (this *CodeVerification) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *CodeVerification) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *CodeVerification) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *CodeVerification) Objects() []*CodeVerification {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *CodeVerification) NewObjects() *[]*CodeVerification {
	this.objects = []*CodeVerification{}
	return &this.objects
}

func (this *CodeVerification) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *CodeVerification) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *CodeVerification) Name_() string {
	if this.namer != nil {
		return this.namer("code_verification")
	}
	return factory.TableNamerGet("code_verification")(this)
}

func (this *CodeVerification) FullName_(connID ...int) string {
	if len(connID) > 0 {
		return factory.DefaultFactory.Cluster(connID[0]).Table(this.Name_())
	}
	return factory.DefaultFactory.Cluster(this.connID).Table(this.Name_())
}

func (this *CodeVerification) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *CodeVerification) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *CodeVerification) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *CodeVerification) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *CodeVerification) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *CodeVerification) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.SendMethod) == 0 { this.SendMethod = "mobile" }
	if len(this.OwnerType) == 0 { this.OwnerType = "user" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint64(v)
		}
	}
	return
}

func (this *CodeVerification) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	
	if len(this.SendMethod) == 0 { this.SendMethod = "mobile" }
	if len(this.OwnerType) == 0 { this.OwnerType = "user" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	return this.Setter(mw, args...).SetSend(this).Update()
}

func (this *CodeVerification) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *CodeVerification) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) error {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *CodeVerification) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) error {
	
	if val, ok := kvset["send_method"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["send_method"] = "mobile" } }
	if val, ok := kvset["owner_type"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["owner_type"] = "user" } }
	if val, ok := kvset["disabled"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["disabled"] = "N" } }
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *CodeVerification) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		
	if len(this.SendMethod) == 0 { this.SendMethod = "mobile" }
	if len(this.OwnerType) == 0 { this.OwnerType = "user" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.SendMethod) == 0 { this.SendMethod = "mobile" }
	if len(this.OwnerType) == 0 { this.OwnerType = "user" }
	if len(this.Disabled) == 0 { this.Disabled = "N" }
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint64(v)
		}
	}
	return 
}

func (this *CodeVerification) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *CodeVerification) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *CodeVerification) Reset() *CodeVerification {
	this.Id = 0
	this.Code = ``
	this.Created = 0
	this.OwnerId = 0
	this.OwnerType = ``
	this.Used = 0
	this.Purpose = ``
	this.Start = 0
	this.End = 0
	this.Disabled = ``
	this.SendMethod = ``
	this.SendTo = ``
	return this
}

func (this *CodeVerification) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Code"] = this.Code
	r["Created"] = this.Created
	r["OwnerId"] = this.OwnerId
	r["OwnerType"] = this.OwnerType
	r["Used"] = this.Used
	r["Purpose"] = this.Purpose
	r["Start"] = this.Start
	r["End"] = this.End
	r["Disabled"] = this.Disabled
	r["SendMethod"] = this.SendMethod
	r["SendTo"] = this.SendTo
	return r
}

