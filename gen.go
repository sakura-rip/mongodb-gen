package main

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

func generateQueryFile() error {
	return createTemplateFile(
		genTargetDir+"/"+"query.go",
		"./templates/query.go.tmpl",
		nil, nil)
}
