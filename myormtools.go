package myORMTools

import (
	"database/sql"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"reflect"
)

import _ "github.com/go-sql-driver/mysql"

type Id int64

type Fetchable[G any] interface {
	Get(id Id) G
}

func FromRows[T any](rows *sql.Rows) []T {
	var ds = []T{}
	for rows.Next() {
		d, err := ScanInto[T](rows)
		if err != nil {
			continue
		}
		ds = append(ds, *d)
	}
	return ds
}

const Y = "y"
const N = "n"

func QueryColumns[Model any]() []string {
	var zero [0]Model
	cols := []string{}
	el := reflect.TypeOf(zero).Elem()
	for i := 0; i < el.NumField(); i++ {
		f := el.Field(i)
		tag := f.Tag.Get("db")
		if tag != "" {
			cols = append(cols, tag)
		}
	}
	return cols
}

func PointerArray[T any](p *T) []any {
	arr := []any{}
	el := reflect.TypeOf(p).Elem()
	val := reflect.ValueOf(p).Elem()
	for i := 0; i < el.NumField(); i++ {
		f := el.Field(i)
		if f.Tag.Get("db") != "" {
			fs := val.Field(i)
			if fs.CanAddr() {
				a := fs.Addr()
				if fs.CanInterface() {
					arr = append(arr, a.Interface())
				}
			}
		}
	}
	return arr
}

func ScanInto[T any](s interface{}) (*T, error) {
	var c T
	var err error
	scanArgs := PointerArray(&c)
	switch t := s.(type) {
	case *sql.Row:
		err = t.Scan(scanArgs...)
	case *sql.Rows:
		err = t.Scan(scanArgs...)
	default:
		err = fmt.Errorf("tipo invalido de scaneable")
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return &c, err
}

func BaseQuery[T any](table string) sq.SelectBuilder {
	cols := QueryColumns[T]()
	return sq.Select(cols...).From(table)
}

type FilteringOptions struct {
	Pagination *PaginationData
	SqlFilter  *map[string]interface{}
}
