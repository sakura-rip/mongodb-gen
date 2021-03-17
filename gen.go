package main

import "html/template"

func generateClientFile(collections []Collection) error {
	return createTemplateFile(
		genTargetDir+"/client_gen.go", "./templates/client.go.tmpl",
		map[string]interface{}{
			"Collections": collections,
		}, nil,
	)
}

func generateCollectionBaseFile(col Collection) error {
	return createTemplateFile(
		genTargetDir+"/"+col.Name+"CollectionBase_gen.go",
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

func generateFieldKnownTypeFile(field CollectionField) error {
	return createTemplateFile(
		genTargetDir+"/"+field.Root.Name+field.FieldName+"Field_gen.go",
		"./templates/field_knownType.go.tmpl",
		field, nil)
}

func generateFieldArrayTypeFile(field CollectionField) error {
	return createTemplateFile(
		genTargetDir+"/"+field.Root.Name+field.FieldName+"Field_gen.go",
		"./templates/field_arrayType.go.tmpl",
		field, template.FuncMap{"getArrayType": getArrayType})
}
