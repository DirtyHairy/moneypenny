package persistence

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/dirtyhairy/moneypenny/model"
	"github.com/go-gorp/gorp"
	_ "github.com/mattn/go-sqlite3"
)

type provider struct {
	db    *sql.DB
	dbMap *gorp.DbMap
	meta  *model.Meta
}

func FromDbConnection(db *sql.DB) (pp Provider, err error) {
	p := provider{
		db: db,
	}

	p.dbMap = &gorp.DbMap{
		Db:            p.db,
		Dialect:       gorp.SqliteDialect{},
		TypeConverter: typeConverter{},
	}

	if err = p.init(); err != nil {
		return
	}

	pp = &p
	return
}

func FromSqlite(dbPath string) (pp Provider, err error) {
	var db *sql.DB
	if db, err = sql.Open("sqlite3", dbPath); err != nil {
		return
	}

	pp, err = FromDbConnection(db)
	return
}

func (p *provider) init() (err error) {
	p.dbMap.AddTableWithName(model.Meta{}, "meta")
	p.dbMap.AddTableWithName(model.User{}, "user")

	if err = p.dbMap.CreateTablesIfNotExists(); err != nil {
		return
	}

	if sqlError := p.dbMap.SelectOne(&p.meta, "SELECT * FROM meta"); sqlError != nil {
		if sqlError != sql.ErrNoRows {
			err = sqlError
			return
		}

		p.meta = &model.Meta{Version: SCHEMA_VERSION}

		if err = p.dbMap.Insert(p.meta); err != nil {
			return
		}
	}

	if p.meta.Version < SCHEMA_VERSION {
		err = errors.New(fmt.Sprintf(
			"DB version is %d, expected %d: database was written by an old version of moneypenny, please migrate", p.meta.Version,
			SCHEMA_VERSION,
		))
		return
	}

	if p.meta.Version > SCHEMA_VERSION {
		err = errors.New(fmt.Sprintf(
			"DB version is %d, expected %d: database was written by a newer version of moneypenny, please update", p.meta.Version,
			SCHEMA_VERSION,
		))
		return
	}

	return
}

func (p *provider) GetMeta() (meta *model.Meta, err error) {
	meta = p.meta
	return
}

func (p *provider) Close() error {
	return p.db.Close()
}
