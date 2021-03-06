migration
=========

migration is a project to create a Go framework for handling database migrations using actual Go code. It is still in early development, but the general idea is to pattern migration definition after `go test` unit test definitions. A user should be able to write a package containing functions that look like the example below, run `go generate`, and get a package that they can import elsewhere in their module that will perform the desired migrations on an `*sql.DB` instance.

Example of Intended Functionality
---------------------------------

This is very early design and hasn't been thought through completely yet. The idea is to pattern after ActiveRecord's migrations, but with some changes. In particular, a dependency system is planned that will hopefully make it easier to structure migrations in a project being worked on by multiple people simultaneously.

```go
func MigrateInit(m *migration.M) {
	m.CreateTable("example", func(t *migration.T) {
		t.AddColumn("id", migration.BigInt).PrimaryKey()
		t.AddColumn("name", migration.String).NotNull(),
		t.AddColumn("val", migration.String)
	})
}

func MigrateMakeValNotNull(m *migration.M) {
	m.Require("Init")

	m.AlterTable("example", func(t *migration.T) {
		t.Column("val").NotNull()
	})
}
```
