package main

import (
	"github.com/fatih/structtag"
	"github.com/iancoleman/strcase"
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/xerrors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	supportExtension = ".go"
)

func getAllCollections(dirPath string) []Collection {
	fileSet := token.NewFileSet()

	list, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal("error occurred during opening dir: %w", err)
	}
	var cols []Collection
	for _, d := range list {
		if d.IsDir() || !strings.HasSuffix(d.Name(), supportExtension) {
			continue
		}
		fileName := filepath.Join(dirPath, d.Name())
		fileSrc, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatal("error occurred during reading file: %w", err)
		}
		if file, err := parser.ParseFile(fileSet, "", fileSrc, parser.AllErrors); err == nil {
			baseFname := strings.TrimSuffix(d.Name(), supportExtension)
			colName := strcase.ToCamel(baseFname)

			if err = isCollectionStructExist(file, colName); err != nil {
				log.Fatal(err)
			}
			col := Collection{
				Name:        colName,
				LowerName:   strcase.ToLowerCamel(colName),
				Fields:      map[string]CollectionField{},
				PackageName: file.Name.Name,
				FileName:    fileName,
			}

			for _, field := range file.Scope.Objects[colName].Decl.(*ast.TypeSpec).Type.(*ast.StructType).Fields.List {
				if len(field.Names) != 1 {
					continue
				}
				//TODO: map, struct, arrayの対応を。
				tags, _ := structtag.Parse(field.Tag.Value)
				bsonTag, _ := tags.Get("bson")
				colFil := CollectionField{
					RootName:  colName,
					FieldName: field.Names[0].Name,
					FieldType: string(fileSrc)[field.Type.Pos()-1 : field.Type.End()-1],
					LowerName: strcase.ToLowerCamel(fileName),
					BsonName:  bsonTag.Name,
				}
				col.Fields[colFil.FieldName] = colFil
			}

			cols = append(cols, col)
		}
	}
	return cols
}

//ファイル名と同じ、Collection名となる構造体の存在を確認する
func isCollectionStructExist(file *ast.File, structName string) error {
	if obj, ok := file.Scope.Objects[structName]; ok == false {
		return xerrors.Errorf("A structure with the same name as the file name is required.\n"+
			"The structure should contain a field with the tag `bson:\"_id\"`.\n"+
			"Required structure name:, %w", structName)
	} else {
		if !isStruct(obj) {
			return xerrors.Errorf(structName, " must be a structure.")
		}
		if getStructNameHaveID(obj) == "" {
			return xerrors.Errorf("Structure: ", structName, " must have a field with tag `bson:\"_id\"`")
		}
		return nil
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
