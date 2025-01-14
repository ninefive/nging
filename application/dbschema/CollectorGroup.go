// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	
	"time"
)

type CollectorGroup struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*CollectorGroup
	namer   func(string) string
	connID  int
	
	Id         	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Uid        	uint    	`db:"uid" bson:"uid" comment:"用户ID" json:"uid" xml:"uid"`
	Name       	string  	`db:"name" bson:"name" comment:"组名" json:"name" xml:"name"`
	Type       	string  	`db:"type" bson:"type" comment:"类型(page-页面规则组;export-导出规则组)" json:"type" xml:"type"`
	Description	string  	`db:"description" bson:"description" comment:"说明" json:"description" xml:"description"`
	Created    	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
}

func (this *CollectorGroup) Trans() *factory.Transaction {
	return this.trans
}

func (this *CollectorGroup) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *CollectorGroup) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *CollectorGroup) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *CollectorGroup) Objects() []*CollectorGroup {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *CollectorGroup) NewObjects() *[]*CollectorGroup {
	this.objects = []*CollectorGroup{}
	return &this.objects
}

func (this *CollectorGroup) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *CollectorGroup) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *CollectorGroup) Name_() string {
	if this.namer != nil {
		return this.namer("collector_group")
	}
	return factory.TableNamerGet("collector_group")(this)
}

func (this *CollectorGroup) FullName_(connID ...int) string {
	if len(connID) > 0 {
		return factory.DefaultFactory.Cluster(connID[0]).Table(this.Name_())
	}
	return factory.DefaultFactory.Cluster(this.connID).Table(this.Name_())
}

func (this *CollectorGroup) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *CollectorGroup) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *CollectorGroup) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *CollectorGroup) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *CollectorGroup) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *CollectorGroup) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Type) == 0 { this.Type = "page" }
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

func (this *CollectorGroup) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	
	if len(this.Type) == 0 { this.Type = "page" }
	return this.Setter(mw, args...).SetSend(this).Update()
}

func (this *CollectorGroup) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *CollectorGroup) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) error {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *CollectorGroup) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) error {
	
	if val, ok := kvset["type"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["type"] = "page" } }
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *CollectorGroup) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		
	if len(this.Type) == 0 { this.Type = "page" }
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.Type) == 0 { this.Type = "page" }
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

func (this *CollectorGroup) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *CollectorGroup) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *CollectorGroup) Reset() *CollectorGroup {
	this.Id = 0
	this.Uid = 0
	this.Name = ``
	this.Type = ``
	this.Description = ``
	this.Created = 0
	return this
}

func (this *CollectorGroup) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Uid"] = this.Uid
	r["Name"] = this.Name
	r["Type"] = this.Type
	r["Description"] = this.Description
	r["Created"] = this.Created
	return r
}

