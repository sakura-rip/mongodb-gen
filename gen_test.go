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
