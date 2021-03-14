package main

import (
	"bytes"
	"golang.org/x/xerrors"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
)

func createTemplateFile(path, tmpPath string, data interface{}, funcMap template.FuncMap) error {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return xerrors.Errorf("error occurred during mkdirall: %w", err)
	}
	if isFileExist(path) {
		err := os.Remove(path)
		if err != nil {
			return xerrors.Errorf("error occurred during removing file: %w", err)
		}
	}

	fp, err := os.Open(tmpPath)
	if err != nil {
		return xerrors.Errorf("error occurred during opening file: %w", err)
	}
	defer fp.Close()

	txt, err := ioutil.ReadAll(fp)
	if err != nil {
		return xerrors.Errorf("error occurred during reading file: %w", err)
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		return xerrors.Errorf("error occurred during creating file: %w", err)
	}
	defer file.Close()

	return ExecuteTemplate(string(txt), file, data, funcMap, path)
}

func ExecuteTemplate(txt string, file io.Writer, data interface{}, funcMap template.FuncMap, path string) error {
	if funcMap == nil {
		funcMap = make(map[string]interface{})
		funcMap["version"] = getVersion
	}
	newTmp := template.New("mongo-db-gen")
	newTmp.Delims("[[", "]]")
	tmp := template.Must(newTmp.Funcs(funcMap).Parse(txt))
	err := tmp.Execute(file, data)
	if err != nil {
		return xerrors.Errorf("error occurred during executing template: %w", err)
	}

	if path == "" {
		return nil
	}
	_, err = ExecCommand("goimports", "-w", path)
	if err != nil {
		return xerrors.Errorf("error occurred during executing goimports: %w", err)
	}

	return nil
}

func ExecuteTemplateInStr(txt string, data interface{}, funcMap template.FuncMap) (string, error) {
	var doc bytes.Buffer
	err := ExecuteTemplate(txt, &doc, data, funcMap, "")
	if err != nil {
		return "", err
	}
	return doc.String(), nil
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
