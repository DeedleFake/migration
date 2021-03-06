package migration

import (
	"context"
	"database/sql"

	"deedles.dev/migration/internal/util"
)

// M is a type passed to Migrate functions to configure the migration.
type M struct {
	name         string
	deps         util.Set[string]
	steps        []mstep
	irreversible bool
}

func (m M) migrateUp(ctx context.Context, tx *sql.Tx, dialect Dialect) error {
	for _, step := range m.steps {
		err := step.migrateUp(ctx, tx, dialect)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m M) migrateDown(ctx context.Context, tx *sql.Tx, dialect Dialect) error {
	if m.irreversible {
		return ErrIrreversible
	}

	for _, step := range m.steps {
		err := step.migrateDown(ctx, tx, dialect)
		if err != nil {
			return err
		}
	}
	return nil
}

// Require marks other migrations as being dependencies of this one.
// In other words, the named migrations should be applied before this
// one is.
//
// Calling this function more than once is equivalent to calling it
// once with all of the same arguments.
//
// The provided migration names should be the name of the migration
// function minus the "Migrate" prefix. For example,
//
//    func MigrateFirst(m *migration.M) {}
//
//    func MigrateSecond(m *migration.M) {
//      // MigrateSecond depends on MigrateFirst.
//      m.Require("First")
//    }
func (m *M) Require(migrations ...string) {
	for _, mig := range migrations {
		m.deps.Add(mig)
	}
}

// CreateTable creates a new table using a configuration determined by
// f.
func (m *M) CreateTable(name string, f func(*T)) {
	t := T{name: name}
	m.steps = append(m.steps, &t)
	f(&t)
}

// UpDown registers separate manual steps for migrating up and down.
// The provided functions should perform opposite operations.
func (m *M) UpDown(up func(*MUp), down func(*MDown)) {
	var updown updown
	m.steps = append(m.steps, &updown)
	up(&updown.up)
	down(&updown.down)
}

// SQL registers a custom SQL statement to be executed as part of the
// migration.
//
// Warning: Use of this method causes the migration to become
// irreversible. To register reversible SQL statements, use UpDown.
func (m *M) SQL(stmt string, args ...any) {
	m.irreversible = true
	m.steps = append(m.steps, sqlstep{stmt: stmt, args: args})
}

type sqlstep struct {
	stmt string
	args []any
}

func (s sqlstep) migrateUp(ctx context.Context, tx *sql.Tx, dialect Dialect) error {
	_, err := tx.ExecContext(ctx, s.stmt, s.args...)
	return err
}

func (s sqlstep) migrateDown(ctx context.Context, tx *sql.Tx, dialect Dialect) error {
	return ErrIrreversible
}
