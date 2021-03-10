package main

type Collection struct {
	//Structの名前
	Name string
	//strings.ToLower(Name)
	LowerName string
	//Structの各フィールドの情報
	//map[Bson名]CollectionField
	Fields map[string]CollectionField
	//パッケージ名
	PackageName string
	//対象Structの存在するファイルの名前
	FileName string
}

func NewCollection() Collection {
	return Collection{
		Fields: map[string]CollectionField{},
	}
}

type CollectionField struct {
	//親の名前(Collection.Name)
	RootName string
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
