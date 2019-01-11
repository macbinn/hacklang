package builtin

import (
	"database/sql"
	"fmt"
	"github.com/macbinn/hacklang/value"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strings"
)

type Table struct {
	db *sql.DB
	name string
}

func (t *Table) all() value.Object {
	rows, err := t.db.Query("select * from " + t.name)
	if err != nil {
		log.Print(err)
		return nil
	}
	types, err := rows.ColumnTypes()
	if err != nil {
		log.Print(err)
		return nil
	}
	list := NewEmptyList()
	for rows.Next() {
		m := NewEmptyMap()
		var dst []interface{}
		for _, t := range types {
			//log.Printf("%s %s", t.Name(), t.DatabaseTypeName())
			// todo: check more types
			if t.DatabaseTypeName() == "INTEGER" {
				s := NewNumber(0)
				m.Val[t.Name()] = s
				dst = append(dst, &s.Int)
			} else {
				s := NewString("")
				m.Val[t.Name()] = s
				dst = append(dst, &s.S)
			}
		}
		err := rows.Scan(dst...)
		if err != nil {
			log.Print(err)
			return nil
		}
		list.L.PushBack(m)
	}
	return list
}

func (t *Table) new(args ...value.Object) value.Object {
	m := args[0].(*Map).Val
	var cols []string
	var vals []string
	var values []interface{}
	for name, val := range m {
		cols = append(cols, name)
		vals = append(vals, "?")
		values = append(values, Convert(val))
	}
	colstr := strings.Join(cols, ", ")
	valstr := strings.Join(vals, ", ")
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", t.name, colstr, valstr)
	result, err := t.db.Exec(query, values...)
	if err != nil {
		return nil
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil
	}
	m["id"] = NewNumber(int(id))
	return nil
}

func (t *Table) del(args ...value.Object) value.Object {
	id := args[0].(*Number).Int
	result, err := t.db.Exec("DELETE FROM "+t.name+" WHERE `id`=?", id)
	if err != nil {
		return nil
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return nil
	}
	if affected == 1 {
		return NewBool(true)
	}
	return nil
}

func (t *Table) Get(name string) value.Object {
	switch name {
	case "all":
		return NewFunction("sql.Table.all", func(args ...value.Object) value.Object {
			return t.all()
		})
	case "new":
		return NewFunction("sql.Table.new", func(args ...value.Object) value.Object {
			return t.new(args...)
		})
	case "del":
		return NewFunction("sql.Table.del", func(args ...value.Object) value.Object {
			return t.del(args...)
		})
	}
	return nil
}

func (t *Table) Repr() string {
	return fmt.Sprintf("<sql.Table %s>", t.name)
}

type DB struct {
	db *sql.DB
	tables *Map
}

func (d *DB) Get(name string) value.Object {
	table := d.tables.Get(name)
	if table == nil {
		// check table exists
		table = &Table{
			db: d.db,
			name: name,
		}
		d.tables.Val[name] = table
	}
	return table
}

func (d *DB) Repr() string {
	return "<sql.DB>"
}

func NewDB(db *sql.DB) *DB {
	return &DB{
		db: db,
		tables: NewEmptyMap(),
	}
}

// sql.db(`sqlite3`, `db.sqlite3`)
func newDB(args ...value.Object) value.Object {
	driver := args[0].(*String).S
	source := args[1].(*String).S
	db, err := sql.Open(driver, source)
	if err != nil {
		return nil
	}
	return NewDB(db)
}

func init() {
	GlobalScope.Register("sql", NewMap(map[string]value.Object{
		"db": NewFunction("sql.db", newDB),
	}))
}

