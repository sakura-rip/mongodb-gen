package main

func generateClientFile(collections []Collection) error {
	return createTemplateFile(
		GenDirName+"/client.go", "./templates/client.go.tmpl",
		map[string][]Collection{
			"Collections": collections,
		}, nil,
	)
}
