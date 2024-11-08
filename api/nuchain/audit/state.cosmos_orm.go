// Code generated by protoc-gen-go-cosmos-orm. DO NOT EDIT.

package audit

import (
	context "context"
	ormlist "cosmossdk.io/orm/model/ormlist"
	ormtable "cosmossdk.io/orm/model/ormtable"
	ormerrors "cosmossdk.io/orm/types/ormerrors"
)

type EntryTable interface {
	Insert(ctx context.Context, entry *Entry) error
	Update(ctx context.Context, entry *Entry) error
	Save(ctx context.Context, entry *Entry) error
	Delete(ctx context.Context, entry *Entry) error
	Has(ctx context.Context, datatype string, id string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, datatype string, id string) (*Entry, error)
	List(ctx context.Context, prefixKey EntryIndexKey, opts ...ormlist.Option) (EntryIterator, error)
	ListRange(ctx context.Context, from, to EntryIndexKey, opts ...ormlist.Option) (EntryIterator, error)
	DeleteBy(ctx context.Context, prefixKey EntryIndexKey) error
	DeleteRange(ctx context.Context, from, to EntryIndexKey) error

	doNotImplement()
}

type EntryIterator struct {
	ormtable.Iterator
}

func (i EntryIterator) Value() (*Entry, error) {
	var entry Entry
	err := i.UnmarshalMessage(&entry)
	return &entry, err
}

type EntryIndexKey interface {
	id() uint32
	values() []interface{}
	entryIndexKey()
}

// primary key starting index..
type EntryPrimaryKey = EntryDatatypeIdIndexKey

type EntryDatatypeIdIndexKey struct {
	vs []interface{}
}

func (x EntryDatatypeIdIndexKey) id() uint32            { return 0 }
func (x EntryDatatypeIdIndexKey) values() []interface{} { return x.vs }
func (x EntryDatatypeIdIndexKey) entryIndexKey()        {}

func (this EntryDatatypeIdIndexKey) WithDatatype(datatype string) EntryDatatypeIdIndexKey {
	this.vs = []interface{}{datatype}
	return this
}

func (this EntryDatatypeIdIndexKey) WithDatatypeId(datatype string, id string) EntryDatatypeIdIndexKey {
	this.vs = []interface{}{datatype, id}
	return this
}

type entryTable struct {
	table ormtable.Table
}

func (this entryTable) Insert(ctx context.Context, entry *Entry) error {
	return this.table.Insert(ctx, entry)
}

func (this entryTable) Update(ctx context.Context, entry *Entry) error {
	return this.table.Update(ctx, entry)
}

func (this entryTable) Save(ctx context.Context, entry *Entry) error {
	return this.table.Save(ctx, entry)
}

func (this entryTable) Delete(ctx context.Context, entry *Entry) error {
	return this.table.Delete(ctx, entry)
}

func (this entryTable) Has(ctx context.Context, datatype string, id string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, datatype, id)
}

func (this entryTable) Get(ctx context.Context, datatype string, id string) (*Entry, error) {
	var entry Entry
	found, err := this.table.PrimaryKey().Get(ctx, &entry, datatype, id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &entry, nil
}

func (this entryTable) List(ctx context.Context, prefixKey EntryIndexKey, opts ...ormlist.Option) (EntryIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return EntryIterator{it}, err
}

func (this entryTable) ListRange(ctx context.Context, from, to EntryIndexKey, opts ...ormlist.Option) (EntryIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return EntryIterator{it}, err
}

func (this entryTable) DeleteBy(ctx context.Context, prefixKey EntryIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this entryTable) DeleteRange(ctx context.Context, from, to EntryIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this entryTable) doNotImplement() {}

var _ EntryTable = entryTable{}

func NewEntryTable(db ormtable.Schema) (EntryTable, error) {
	table := db.GetTable(&Entry{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Entry{}).ProtoReflect().Descriptor().FullName()))
	}
	return entryTable{table}, nil
}

type StateStore interface {
	EntryTable() EntryTable

	doNotImplement()
}

type stateStore struct {
	entry EntryTable
}

func (x stateStore) EntryTable() EntryTable {
	return x.entry
}

func (stateStore) doNotImplement() {}

var _ StateStore = stateStore{}

func NewStateStore(db ormtable.Schema) (StateStore, error) {
	entryTable, err := NewEntryTable(db)
	if err != nil {
		return nil, err
	}

	return stateStore{
		entryTable,
	}, nil
}
