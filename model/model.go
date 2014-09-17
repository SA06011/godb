package model

// this model save database's structure
// the field represents related field in table
type Field struct {
	Id        int // unique
	FieldName string
	IsNull    bool
	IsPk      bool
	Type      int
	IsForeign bool
	IsIndex   bool
	IsUnique  bool
}

type Table struct {
	Id        int
	TableName string
	Fields    []Field
}

type Database struct {
	Id     int
	DbName string
	Tables []Table
}
