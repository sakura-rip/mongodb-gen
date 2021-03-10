package main

type Collection struct {
	Name      string
	LowerName string
	Fields    []CollectionField

	PackageName string
}

type CollectionField struct {
	Name     string
	BsonName string
}
