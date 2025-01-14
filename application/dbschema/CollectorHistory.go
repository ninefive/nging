// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	
	"time"
)

type CollectorHistory struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*CollectorHistory
	namer   func(string) string
	connID  int
	
	Id            	uint64  	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	ParentId      	uint64  	`db:"parent_id" bson:"parent_id" comment:"父ID" json:"parent_id" xml:"parent_id"`
	PageId        	uint    	`db:"page_id" bson:"page_id" comment:"页面ID" json:"page_id" xml:"page_id"`
	PageParentId  	uint    	`db:"page_parent_id" bson:"page_parent_id" comment:"父页面ID" json:"page_parent_id" xml:"page_parent_id"`
	PageRootId    	uint    	`db:"page_root_id" bson:"page_root_id" comment:"入口页面ID" json:"page_root_id" xml:"page_root_id"`
	HasChild      	string  	`db:"has_child" bson:"has_child" comment:"是否有子级" json:"has_child" xml:"has_child"`
	Url           	string  	`db:"url" bson:"url" comment:"页面网址" json:"url" xml:"url"`
	UrlMd5        	string  	`db:"url_md5" bson:"url_md5" comment:"页面网址MD5" json:"url_md5" xml:"url_md5"`
	Title         	string  	`db:"title" bson:"title" comment:"页面标题" json:"title" xml:"title"`
	Content       	string  	`db:"content" bson:"content" comment:"页面内容MD5" json:"content" xml:"content"`
	RuleMd5       	string  	`db:"rule_md5" bson:"rule_md5" comment:"规则标识MD5" json:"rule_md5" xml:"rule_md5"`
	Data          	string  	`db:"data" bson:"data" comment:"采集到的数据" json:"data" xml:"data"`
	Created       	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Exported      	uint    	`db:"exported" bson:"exported" comment:"最近导出时间" json:"exported" xml:"exported"`
}

func (this *CollectorHistory) Trans() *factory.Transaction {
	return this.trans
}

func (this *CollectorHistory) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *CollectorHistory) SetConnID(connID int) factory.Model {
	this.connID = connID
	return this
}

func (this *CollectorHistory) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.trans)
	}
	return factory.NewModel(structName,this.connID).Use(this.trans)
}

func (this *CollectorHistory) Objects() []*CollectorHistory {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *CollectorHistory) NewObjects() *[]*CollectorHistory {
	this.objects = []*CollectorHistory{}
	return &this.objects
}

func (this *CollectorHistory) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.connID).SetTrans(this.trans).SetCollection(this.Name_()).SetModel(this)
}

func (this *CollectorHistory) SetNamer(namer func (string) string) factory.Model {
	this.namer = namer
	return this
}

func (this *CollectorHistory) Name_() string {
	if this.namer != nil {
		return this.namer("collector_history")
	}
	return factory.TableNamerGet("collector_history")(this)
}

func (this *CollectorHistory) FullName_(connID ...int) string {
	if len(connID) > 0 {
		return factory.DefaultFactory.Cluster(connID[0]).Table(this.Name_())
	}
	return factory.DefaultFactory.Cluster(this.connID).Table(this.Name_())
}

func (this *CollectorHistory) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *CollectorHistory) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *CollectorHistory) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *CollectorHistory) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *CollectorHistory) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *CollectorHistory) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.HasChild) == 0 { this.HasChild = "N" }
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

func (this *CollectorHistory) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	
	if len(this.HasChild) == 0 { this.HasChild = "N" }
	return this.Setter(mw, args...).SetSend(this).Update()
}

func (this *CollectorHistory) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *CollectorHistory) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) error {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *CollectorHistory) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) error {
	
	if val, ok := kvset["has_child"]; ok && val != nil { if v, ok := val.(string); ok && len(v) == 0 { kvset["has_child"] = "N" } }
	return this.Setter(mw, args...).SetSend(kvset).Update()
}

func (this *CollectorHistory) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		
	if len(this.HasChild) == 0 { this.HasChild = "N" }
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
	if len(this.HasChild) == 0 { this.HasChild = "N" }
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

func (this *CollectorHistory) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *CollectorHistory) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *CollectorHistory) Reset() *CollectorHistory {
	this.Id = 0
	this.ParentId = 0
	this.PageId = 0
	this.PageParentId = 0
	this.PageRootId = 0
	this.HasChild = ``
	this.Url = ``
	this.UrlMd5 = ``
	this.Title = ``
	this.Content = ``
	this.RuleMd5 = ``
	this.Data = ``
	this.Created = 0
	this.Exported = 0
	return this
}

func (this *CollectorHistory) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["ParentId"] = this.ParentId
	r["PageId"] = this.PageId
	r["PageParentId"] = this.PageParentId
	r["PageRootId"] = this.PageRootId
	r["HasChild"] = this.HasChild
	r["Url"] = this.Url
	r["UrlMd5"] = this.UrlMd5
	r["Title"] = this.Title
	r["Content"] = this.Content
	r["RuleMd5"] = this.RuleMd5
	r["Data"] = this.Data
	r["Created"] = this.Created
	r["Exported"] = this.Exported
	return r
}

