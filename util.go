//Copyright (c)  2020 go-generalize
//Released under the MIT license
//https://opensource.org/licenses/mit-license.php

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
)

// ExecCommand - コマンドを実行して出力結果とエラーを返す
func ExecCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	b, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	if exitCode := cmd.ProcessState.ExitCode(); exitCode != 0 {
		return "", fmt.Errorf("failed to exec git command: (exit code: %d, output: %s)", exitCode, string(b))
	}

	return string(b), nil
}

// GetGitRootPath - カレントディレクトリが所属するGitリポジトリの直下のパスを返す
func GetGitRootPath() string {
	s, err := ExecCommand("git", "rev-parse", "--show-superproject-working-tree", "--show-toplevel")

	if err != nil {
		log.Fatalf("failed to exec git command: %v", err)
	}

	return strings.TrimRight(s, "\r\n")
}

// GetGoModPath - Gitリポジトリ直下から幅優先探索でgo.modを探す
func GetGoModPath() string {
	dirs := []string{
		GetGitRootPath(),
	}

	result := ""
	for len(dirs) != 0 && result == "" {
		d := dirs[0]
		dirs = dirs[1:]

		err := filepath.Walk(d, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				// 直下はSkipしない
				if path == d {
					return nil
				}

				if filepath.Base(path) == "node_modules" {
					return filepath.SkipDir
				}

				dirs = append(dirs, path)

				return filepath.SkipDir
			}

			if filepath.Base(path) == "go.mod" {
				result = path
			}

			return nil
		})
		if err != nil {
			return ""
		}
	}

	return result
}

// GetGoRootPath - go.modがあるフォルダを返す
func GetGoRootPath() string {
	return filepath.Dir(GetGoModPath())
}

// GetGoRootPackageName - Goのルートパッケージ名をgo.modから取得する
func GetGoRootPackageName() (string, error) {
	goModPath := GetGoModPath()
	d, err := ioutil.ReadFile(goModPath)
	if err != nil {
		return "", err
	}

	f, err := modfile.Parse("", d, nil)
	if err != nil {
		return "", err
	}

	if len(f.Module.Mod.Path) == 0 {
		return "", fmt.Errorf("package name was not found")
	}

	return f.Module.Mod.Path, nil
}

func isKnownType(str string) bool {
	switch str {
	case "bool", "byte", "complex128", "complex64", "float32", "float64", "int", "int16", "int32", "int64", "int8", "rune", "string", "uint", "uint16", "uint32", "uint64", "uint8", "uintptr", "interface{}":
		return true
	default:
		return false
	}
}

func getArrayType(field CollectionField) string {
	ft := field.FieldType[2:]
	switch {
	case isKnownType(ft):
		return ft
	default:
		return "interface{}"
	}
}

func getMapKeyType(field CollectionField) string {
	//map[]string
	start := strings.Index(field.FieldType, "[")
	end := strings.Index(field.FieldType, "]")
	key := field.FieldType[start-1 : end-1]
	switch {
	case isKnownType(key):
		return key
	default:
		return "interface{}"
	}
}

func getMapValueType(field CollectionField) string {
	end := strings.Index(field.FieldType, "]")
	key := field.FieldType[end-1 : len(field.FieldType)-1]
	switch {
	case isKnownType(key):
		return key
	//	TODO:valueがmapやStructの時の対応
	default:
		return "interface{}"
	}
}

func sprints(base string, val ...interface{}) string {
	return fmt.Sprintf(base+strings.Repeat(".%s", len(val)), val...)
}

func isStructFieldValue(field string) bool {
	switch {
	case isKnownType(field):
		return false
	case strings.HasPrefix(field, "["):
		return false
	case strings.HasPrefix(field, "map["):
		return false
	default:
		return true
	}
}
