package main

import (
	"github.com/fatih/structtag"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"
)

func findAllStructs(path string) {
	fileSet := token.NewFileSet()
	packages, err := parser.ParseDir(fileSet, path, nil, parser.AllErrors)
	if err != nil {
		log.Fatal("err on parse package: ", err)
	}
	var collections []Collection
	for _, pkg := range packages {
		collection := NewCollection()
		var rootName string
		for fName, file := range pkg.Files {
			//Field
			for _, obj := range file.Scope.Objects {
				if !isStruct(obj) {
					continue
				}
				field := NewCollectionField()
				field.RootName = ""
				field.FieldName = ""
				field.FieldType = ""
				field.LowerName = ""
				field.BsonName = ""

			}
			collection.FileName = fName
			collection.PackageName = pkg.Name
			collection.Name = rootName
			collection.LowerName = strings.ToLower(rootName)
			collections = append(collections, collection)
		}
	}
}

func getStructNameHaveID(obj *ast.Object) string {
	objStruct := obj.Decl.(*ast.TypeSpec).Type.(*ast.StructType)
	for _, field := range objStruct.Fields.List {
		//fieldName := field.Names[0].Name
		tags, _ := structtag.Parse(field.Tag.Value)
		if val, _ := tags.Get("bson"); val.Name == "_id" {
			return obj.Decl.(*ast.TypeSpec).Name.Name
		}
	}
	return ""
}

func isStruct(obj *ast.Object) bool {
	if obj.Kind != ast.Typ {
		return false
	}
	typeSpec, ok := obj.Decl.(*ast.TypeSpec)
	if !ok {
		return false
	}
	if _, ok = typeSpec.Type.(*ast.StructType); !ok {
		return false
	}
	return true
}
