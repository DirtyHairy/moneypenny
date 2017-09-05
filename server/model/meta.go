package model

type Meta struct {
	Id      uint64 `db:"id,primarykey"`
	Version uint64 `db:"version,notnull"`
}
