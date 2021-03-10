package main

import (
	"golang.org/x/xerrors"
	"html/template"
	"io/ioutil"
	"os"
	"syscall"
)

func CreateTemplateFile(path, tmpPath string, data interface{}, funcMap template.FuncMap) error {
	if isFileExist(path) {
		err := os.Remove(path)
		if err != nil {
			return xerrors.Errorf("error occurred during removing file: %w", err)
		}
	}

	fp, err := os.Open(tmpPath)
	if err != nil {
		return xerrors.Errorf("error occurred during opening file:%w", err)
	}
	defer fp.Close()

	txt, err := ioutil.ReadAll(fp)
	if err != nil {
		return xerrors.Errorf("error occurred during reading file:%w", err)
	}
	tmp := template.Must(template.New("mongo-db-gen").Funcs(funcMap).Parse(string(txt)))
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		return xerrors.Errorf("error occurred during creating file:%w", err)
	}
	defer file.Close()

	err = tmp.Execute(file, data)
	if err != nil {
		return xerrors.Errorf("error occurred during executing template:%w", err)
	}
	return nil
}

func isFileExist(path string) bool {
	_, err := os.Stat(path)
	if pErr, ok := err.(*os.PathError); ok {
		if pErr.Err == syscall.ENOTDIR {
			return false
		}
	}
	if os.IsNotExist(err) {
		return false
	}

	return true
}
