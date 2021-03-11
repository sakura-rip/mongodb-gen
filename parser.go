package main

import (
	"github.com/fatih/structtag"
	"github.com/iancoleman/strcase"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	supportExtension = ".go"
)

func findAllStructs(dirPath string) {
	fileSet := token.NewFileSet()

	list, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal("error occurred during opening dir: %w", err)
	}
	for _, d := range list {
		if d.IsDir() || !strings.HasSuffix(d.Name(), supportExtension) {
			continue
		}
		filename := filepath.Join(dirPath, d.Name())
		if file, err := parser.ParseFile(fileSet, filename, nil, parser.AllErrors); err == nil {
			baseFname := strings.TrimSuffix(d.Name(), supportExtension)

			//ファイル名と同じ、Collection名となる構造体を取得する
			if val, ok := file.Scope.Objects[strcase.ToCamel(baseFname)]; ok == false {
				log.Fatal("A structure with the same name as the file name is required.\n"+
					"The structure should contain a field with the tag `bson:\"_id\"`.\n"+
					"Required structure name:", strcase.ToCamel(baseFname))
			} else {
				if !isStruct(val) {
					log.Fatal(strcase.ToCamel(baseFname), " must be a structure.")
				}
				if getStructNameHaveID(val) == "" {
					log.Fatal("Structure: ", strcase.ToCamel(baseFname), " must have a field with tag `bson:\"_id\"`")
				}
			}

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
