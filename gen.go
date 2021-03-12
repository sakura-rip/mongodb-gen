package main

func generateClientFile(packageName string, collections []Collection) error {
	return createTemplateFile(
		targetDirName+"/client.go", "./templates/client.go.tmpl",
		map[string]interface{}{
			"Collections": collections,
			"PackageName": packageName,
		}, nil,
	)
}
