package main

import (
	"github.com/fatih/structtag"
	"github.com/iancoleman/strcase"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"golang.org/x/xerrors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	supportExtension = ".go"
)

func getAllCollections(targetDirName string) []Collection {
	fileSet := token.NewFileSet()
	list, err := os.ReadDir(targetDirName)
	if err != nil {
		log.Fatal("error occurred during opening dir: %w", err)
	}
	var cols []Collection
	for _, d := range list {
		if d.IsDir() || !strings.HasSuffix(d.Name(), supportExtension) {
			continue
		}
		fileName := filepath.Join(targetDirName, d.Name())
		fileSrc, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatal("error occurred during reading file: %w", err)
		}
		if file, err := parser.ParseFile(fileSet, "", fileSrc, parser.AllErrors); err == nil {
			baseFname := strings.TrimSuffix(d.Name(), supportExtension)
			colName := strcase.ToCamel(baseFname)

			idField, err := checkCollectionStructExist(file, colName)
			if err != nil {
				log.Fatal(err)
			}
			col := Collection{
				Name:        colName,
				LowerName:   strcase.ToLowerCamel(colName),
				Fields:      map[string]CollectionField{},
				FileName:    fileName,
				PackageName: file.Name.Name,
				IdFieldName: idField.Names[0].Name,
				IdFieldType: types.ExprString(idField.Type),
			}

			for _, field := range file.Scope.Objects[colName].Decl.(*ast.TypeSpec).Type.(*ast.StructType).Fields.List {
				if len(field.Names) != 1 {
					continue
				}
				tags, _ := structtag.Parse(strings.Trim(field.Tag.Value, "`"))
				bsonTag, _ := tags.Get("bson")
				colFil := CollectionField{
					Root:      col,
					FieldName: field.Names[0].Name,
					FieldType: types.ExprString(field.Type),
					LowerName: strcase.ToLowerCamel(field.Names[0].Name),
					BsonName:  bsonTag.Name,
				}
				colFil.StructFields = getCollectionField(field, &colFil)
				col.Fields[colFil.FieldName] = colFil
			}

			cols = append(cols, col)
		}
	}
	return cols
}

func getCollectionField(field *ast.Field, col *CollectionField) []StructField {
	var fields []StructField
	if isStructField(field) {
		for _, field := range field.Type.(*ast.Ident).Obj.Decl.(*ast.TypeSpec).Type.(*ast.StructType).Fields.List {
			fields = append(fields, parseStructField(col, field, nil)...)
		}
	}
	return fields
}

func parseStructField(col *CollectionField, field *ast.Field, baseColFil *StructField) []StructField {
	var fields []StructField
	tags, _ := structtag.Parse(strings.Trim(field.Tag.Value, "`"))
	bsonTag, _ := tags.Get("bson")
	fName := field.Names[0].Name
	strField := StructField{
		Name:      fName,
		FieldType: types.ExprString(field.Type),
		LowerName: strcase.ToLowerCamel(fName),
	}
	if baseColFil == nil {
		strField.LocationBson = col.BsonName + "." + bsonTag.Name
		strField.Location = col.FieldName + "." + fName
		strField.FieldLocation = col.FieldName + fName
	} else {
		strField.LocationBson = baseColFil.LocationBson + "." + bsonTag.Name
		strField.Location = baseColFil.Location + "." + fName
		strField.FieldLocation = baseColFil.FieldLocation + fName
	}
	if isStructFieldValue(strField.FieldType) {
		strField.FieldType = col.Root.PackageName + "." + strField.FieldType
	}
	fields = append(fields, strField)
	if !isStructField(field) {
		return fields
	}
	for _, b := range field.Type.(*ast.Ident).Obj.Decl.(*ast.TypeSpec).Type.(*ast.StructType).Fields.List {
		fields = append(fields, parseStructField(col, b, &strField)...)
	}
	return fields
}

//ファイル名と同じ、Collection名となる構造体の存在を確認する
func checkCollectionStructExist(file *ast.File, structName string) (*ast.Field, error) {
	if obj, ok := file.Scope.Objects[structName]; ok == false {
		return nil, xerrors.Errorf("A structure with the same name as the file name is required.\n"+
			"The structure should contain a field with the tag `bson:\"_id\"`.\n"+
			"Required structure name:, %w", structName)
	} else {
		if !isStruct(obj) {
			return nil, xerrors.Errorf("%w must be a structure.", structName)
		}
		if field := getStructFieldHaveID(obj); field == nil {
			return nil, xerrors.Errorf("Structure: %w %w", structName, " must have one field with tag `bson:\"_id\"`")
		} else {
			return field, nil
		}
	}
}

func getStructFieldHaveID(obj *ast.Object) *ast.Field {
	objStruct := obj.Decl.(*ast.TypeSpec).Type.(*ast.StructType)
	var fields []*ast.Field
	for _, field := range objStruct.Fields.List {
		//fieldName := field.Names[0].Name
		if field.Tag == nil {
			log.Fatal("all struct field must have bson tag")
		}
		tags, err := structtag.Parse(strings.Trim(field.Tag.Value, "`"))
		if err != nil {
			log.Fatal("error occurred during parsing field tag:", err)
		}
		if val, _ := tags.Get("bson"); val.Name == "_id" {
			fields = append(fields, field)
		}
	}
	if len(fields) == 1 {
		return fields[0]
	}
	return nil
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

func isStructField(field *ast.Field) bool {
	iden, ok := field.Type.(*ast.Ident)
	if !ok {
		return false
	}
	if iden.Obj == nil {
		return false
	}
	ts, ok := iden.Obj.Decl.(*ast.TypeSpec)
	if !ok {
		return false
	}
	_, ok = ts.Type.(*ast.StructType)
	return ok
}
