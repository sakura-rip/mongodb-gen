package main

import "github.com/sakura-rip/mongodbgen/templates"

func generateClientFile(collections []Collection) error {
	return createTemplateFile(
		genTargetDir+"/client.go", "./templates/client.go.tmpl",
		map[string]interface{}{
			"Collections": collections,
		}, nil,
	)
}

func generateCollectionBaseFile(col Collection) error {
	return createTemplateFile(
		genTargetDir+"/"+col.Name+"Base.go",
		"./templates/CollectionBase.go.tmpl",
		col, nil,
	)
}

func generateCollectionFieldBase(field CollectionField) (string, error) {
	return ExecuteTemplateInStr(templates.DefaultField, field, nil)
}
