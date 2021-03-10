package main

type Collection struct {
	Name      string
	LowerName string
	Fields    []CollectionField

	PackageName string
}

type CollectionField struct {
	RootName  string
	FieldName string
	FieldType string
	LowerName string
	BsonName  string
}
