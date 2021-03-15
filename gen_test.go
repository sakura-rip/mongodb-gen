package main

import (
	"fmt"
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

func TestGenerateCollectionFieldBase(t *testing.T) {
	targetDirName = "sample"
	genTargetDir = targetDirName + "_dao"
	cols := getAllCollections()
	for _, field := range cols[0].Fields {
		base, err := generateCollectionFieldBase(field)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(base)
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
