package main

import "testing"

func TestGenerateClientFile(t *testing.T) {
	GenDirName = "."
	err := generateClientFile("sample", []Collection{})
	if err != nil {
		t.Error(err)
	}
}
