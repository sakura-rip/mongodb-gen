package main

type Collection struct {
	//Structの名前
	Name string
	//strcase.ToLowerCamel(Name)
	LowerName string
	//Structの各フィールドの情報
	//map[Bson名]CollectionField
	Fields map[string]CollectionField
	//対象Structの存在するファイルの名前
	FileName string

	GoModPackageName string
	PackageName      string

	IdFieldType string
}

func NewCollection() Collection {
	return Collection{
		Fields: map[string]CollectionField{},
	}
}

type CollectionField struct {
	//Collection.Name
	RootName string
	//Collection.LowerName
	RootLowerName string
	//Collection.IdFieldType
	RootIdFieldType string
	//Collection.PackageName
	PackageName string

	//フィールド名
	FieldName string
	//フィールドのType
	FieldType string
	//strings.ToLower(FieldName)
	LowerName string
	//`bson:"BsonName"`
	BsonName string
}

func NewCollectionField() CollectionField {
	return CollectionField{}
}
