package factory

import (
	"context"
	"database/sql"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/sqlbuilder"
)

var DefaultFactory = New()

func Default() *Factory {
	return DefaultFactory
}

func Debug() bool {
	return DefaultFactory.Debug()
}

func CountCluster() int {
	return DefaultFactory.CountCluster()
}

func SetDebug(on bool) *Factory {
	DefaultFactory.SetDebug(on)
	return DefaultFactory
}

func SetCacher(cc Cacher) *Factory {
	DefaultFactory.SetCacher(cc)
	return DefaultFactory
}

func AddDB(databases ...db.Database) *Factory {
	DefaultFactory.AddDB(databases...)
	return DefaultFactory
}

func AddSlaveDB(databases ...db.Database) *Factory {
	DefaultFactory.AddSlaveDB(databases...)
	return DefaultFactory
}

func AddDBToCluster(index int, databases ...db.Database) *Factory {
	DefaultFactory.AddDBToCluster(index, databases...)
	return DefaultFactory
}

func AddSlaveDBToCluster(index int, databases ...db.Database) *Factory {
	DefaultFactory.AddSlaveDBToCluster(index, databases...)
	return DefaultFactory
}

func SetCluster(index int, cluster *Cluster) *Factory {
	DefaultFactory.SetCluster(index, cluster)
	return DefaultFactory
}

func AddCluster(clusters ...*Cluster) *Factory {
	DefaultFactory.AddCluster(clusters...)
	return DefaultFactory
}

func GetCluster(index int) *Cluster {
	return DefaultFactory.GetCluster(index)
}

func Tx(param *Param, ctxa ...context.Context) error {
	return DefaultFactory.Tx(param, ctxa...)
}

func NewTx(ctx context.Context, args ...int) (*Transaction, error) {
	return DefaultFactory.NewTx(ctx, args...)
}

func CloseAll() {
	DefaultFactory.CloseAll()
}

func Result(param *Param) db.Result {
	return DefaultFactory.Result(param)
}

func Driver(param *Param) interface{} {
	return DefaultFactory.Driver(param)
}

func DB(param *Param) *sql.DB {
	return DefaultFactory.DB(param)
}

// ================================
// API
// ================================

// Read ==========================

// Query query SQL. sqlRows is an *sql.Rows object, so you can use Scan() on it
// err = sqlRows.Scan(&a, &b, ...)
func Query(param *Param) (*sql.Rows, error) {
	return DefaultFactory.Query(param)
}

// QueryTo query SQL. mapping fields into a struct
func QueryTo(param *Param) (sqlbuilder.Iterator, error) {
	return DefaultFactory.QueryTo(param)
}

// QueryRow query SQL
func QueryRow(param *Param) *sql.Row {
	return DefaultFactory.QueryRow(param)
}

func SelectAll(param *Param) error {
	return DefaultFactory.SelectAll(param)
}

func SelectOne(param *Param) error {
	return DefaultFactory.SelectOne(param)
}

func SelectCount(param *Param) (int64, error) {
	return DefaultFactory.SelectCount(param)
}

func SelectList(param *Param) (func() int64, error) {
	return DefaultFactory.SelectList(param)
}

func Select(param *Param) sqlbuilder.Selector {
	return DefaultFactory.Select(param)
}

func All(param *Param) error {
	return DefaultFactory.All(param)
}

func List(param *Param) (func() int64, error) {
	return DefaultFactory.List(param)
}

func One(param *Param) error {
	return DefaultFactory.One(param)
}

func Count(param *Param) (int64, error) {
	return DefaultFactory.Count(param)
}

// Write ==========================

// Exec execute SQL
func Exec(param *Param) (sql.Result, error) {
	return DefaultFactory.Exec(param)
}

func Insert(param *Param) (interface{}, error) {
	return DefaultFactory.Insert(param)
}

func Update(param *Param) error {
	return DefaultFactory.Update(param)
}

func Upsert(param *Param, beforeUpsert ...func()) (interface{}, error) {
	return DefaultFactory.Upsert(param, beforeUpsert...)
}

func Delete(param *Param) error {
	return DefaultFactory.Delete(param)
}
