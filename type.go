package migration

type columnType interface {
	SQL(Dialect) string
}

type stringType struct{}

func (stringType) SQL(dialect Dialect) string {
	return "TEXT"
}

type intType struct{}

func (intType) SQL(dialect Dialect) string {
	return "INT"
}
