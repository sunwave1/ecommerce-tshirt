package database

import (
	"database/sql"
	"fmt"
	"strings"
)

type Builder struct {
	table   string
	columns []string
	wheres  []string
	orWhere []string
	updates []string
	limit   int
	raw     string
}

func (builder *Builder) getColums() string {

	if len(builder.columns) <= 0 {
		return "*"
	}

	return strings.Join(builder.columns, ", ")
}

func (builder *Builder) getWheres() string {
	return strings.Join(builder.wheres, " AND ")
}

func (builder *Builder) getOrWheres() string {
	return strings.Join(builder.orWhere, " OR ")
}

func (builder *Builder) getUpdates() string {
	return strings.Join(builder.updates, ", ")
}

func (builder *Builder) getCondition() string {

	wheres := builder.getWheres()

	var condition string

	if len(wheres) > 0 {
		condition += " WHERE "
	}

	if len(builder.getOrWheres()) > 0 {
		condition += fmt.Sprintf(" OR %s", builder.getOrWheres())
	}

	if builder.limit > 0 {
		condition += fmt.Sprintf(" LIMIT %d", builder.limit)
	}

	return condition
}

func (builder *Builder) Limit(limit int) *Builder {

	builder.limit = limit

	return builder
}

func (builder *Builder) Where(where string) *Builder {

	builder.wheres = append(builder.wheres, where)

	return builder
}

func (builder *Builder) OrWhere(orWhere string) *Builder {

	builder.orWhere = append(builder.orWhere, orWhere)

	return builder
}

func (builder *Builder) Table(table string) *Builder {

	builder.table = table

	return builder
}

func (builder *Builder) Columns(columns []string) *Builder {

	builder.columns = columns

	return builder
}

func (builder *Builder) Update(updates []string) *Builder {

	builder.updates = updates

	return builder
}

func (builder *Builder) GetSql() string {

	if len(builder.raw) > 0 {
		return builder.raw
	}

	return fmt.Sprintf("SELECT %s FROM %s %s", builder.getColums(), builder.table, builder.getCondition())
}

func (builder *Builder) Dump() *Builder {

	fmt.Println(builder.GetSql())

	return builder
}

func (builder *Builder) Raw(raw string) *Builder {

	builder.raw = raw

	return builder
}

func (builder *Builder) Get(callback func(db *sql.DB, row *sql.Rows, columns []string) error) error {

	db := GetConnection()

	rows, err := db.Query(builder.GetSql())

	if err != nil {
		return err
	}

	columns, erroColumn := rows.Columns()

	if erroColumn != nil {
		return erroColumn
	}

	for rows.Next() {
		if errorRow := callback(db, rows, columns); errorRow != nil {
			return errorRow
		}
	}

	return nil
}
