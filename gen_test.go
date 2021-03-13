package main

import "testing"

func TestGenerateClientFile(t *testing.T) {
	targetDirName = "."
	err := generateClientFile("sample", []Collection{})
	if err != nil {
		t.Error(err)
	}
}

func TestGenerateCollectionBaseFile(t *testing.T) {
	targetDirName = "."
	cols := getAllCollections("sample")
	err := generateCollectionBaseFile(cols[0])
	if err != nil {
		t.Error(err)
	}
}
