package migration

import (
	"context"
	"database/sql"
)

// T provides methods to configure modifications to a database table.
type T struct {
	name  string
	steps []mstep
}

func (t T) migrateUp(ctx context.Context, tx *sql.Tx, dialect Dialect) error {
	panic("Not implemented.")
}

func (t T) migrateDown(ctx context.Context, tx *sql.Tx, dialect Dialect) error {
	panic("Not implemented.")
}

// Name returns the name of the table being modified.
func (t T) Name() string {
	return t.name
}

// String adds a string column to the table.
func (t *T) String(name string) *Column[string] {
	return addColumn[stringType, string](t, name)
}

// Int adds an integer column to the table.
func (t *T) Int(name string) *Column[int] {
	return addColumn[intType, int](t, name)
}

// Index adds an index to the table that applies to the named columns.
func (t *T) Index(names ...string) *Index {
	i := Index{names: names}
	t.steps = append(t.steps, &i)
	return &i
}

// Column represents the configuration of a column in a table.
type Column[T any] struct {
	name string
	t    columnType
	d    *T
	null bool
}

func (c Column[T]) migrateUp(ctx context.Context, tx *sql.Tx, dialect Dialect) error {
	panic("Not implemented.")
}

func (c Column[T]) migrateDown(ctx context.Context, tx *sql.Tx, dialect Dialect) error {
	panic("Not implemented.")
}

// NoDefault disables the default value of the column.
func (c *Column[T]) NoDefault() *Column[T] {
	c.d = nil
	return c
}

// Default sets the default value of the column.
func (c *Column[T]) Default(d T) *Column[T] {
	c.d = &d
	return c
}

//func (c *Column[T]) DefaultFunc() *Column[T] {
//	panic("Not implemented.")
//}

// Null configures whether or not NULL values are allowed in the
// column.
func (c *Column[T]) Null(allow bool) *Column[T] {
	c.null = allow
	return c
}

// Index represents the configuration of an index in a table.
type Index struct {
	names  []string
	unique bool
}

func (i Index) migrateUp(ctx context.Context, tx *sql.Tx, dialect Dialect) error {
	panic("Not implemented.")
}

func (i Index) migrateDown(ctx context.Context, tx *sql.Tx, dialect Dialect) error {
	panic("Not implemented.")
}

// Unique sets whether or not the column enforces uniqueness.
func (i *Index) Unique(unique bool) *Index {
	i.unique = unique
	return i
}

// addColumn adds a column of a specified type to a table.
func addColumn[C columnType, V any](t *T, name string) *Column[V] {
	var ct C
	c := Column[V]{name: name, t: ct}
	t.steps = append(t.steps, &c)
	return &c
}
