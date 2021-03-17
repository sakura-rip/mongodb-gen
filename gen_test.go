package main

import (
	"testing"
)

func TestGenerateClientFile(t *testing.T) {
	targetDirName = "sample"
	genTargetDir = targetDirName + "_dao"
	cols := getAllCollections()
	err := generateClientFile(cols)
	if err != nil {
		t.Error(err)
	}
}

func TestGenerateCollectionBaseFile(t *testing.T) {
	targetDirName = "sample"
	genTargetDir = targetDirName + "_dao"
	cols := getAllCollections()
	for _, col := range cols {
		err := generateCollectionBaseFile(col)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestGenerateQueryFile(t *testing.T) {
	targetDirName = "sample"
	genTargetDir = targetDirName + "_dao"
	err := generateQueryFile()
	if err != nil {
		t.Error(err)
	}
}

func TestGenerateFieldDefaultFile(t *testing.T) {
	targetDirName = "sample"
	genTargetDir = targetDirName + "_dao"
	cols := getAllCollections()
	for _, col := range cols {
		for fName, field := range col.Fields {
			if fName == col.IdFieldName {
				continue
			}
			switch {
			//array
			case strings.HasPrefix(field.FieldType, "[]"):
			//map
			case strings.HasPrefix(field.FieldType, "map["):
			//default(int, str...etc
			case isKnownType(field.FieldType):
				err := generateFieldKnownTypeFile(field)
				if err != nil {
					t.Error(err)
				}
			//struct
			default:

			}
		}
	}
}
