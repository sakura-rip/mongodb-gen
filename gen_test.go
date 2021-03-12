package main

import "testing"

func TestGenerateClientFile(t *testing.T) {
	targetDirName = "."
	err := generateClientFile("sample", []Collection{})
	if err != nil {
		t.Error(err)
	}
}
