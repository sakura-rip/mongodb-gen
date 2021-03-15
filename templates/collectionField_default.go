package templates

const (
	DefaultField = `
//#############################
//[[.FieldName]] handlers
//Type: [[.FieldType]]
//#############################


func (cl *[[.RootName]]DaoClient) Update[[.RootName]][[.FieldName]](id, [[.FieldName]] [[.FieldType]]) error {
	return cl.UpdateAttribute(id, bson.D{{"[[.BsonName]]", [[.FieldName]]}})
}

func (cl *[[.RootName]]DaoClient) Delete[[.RootName]][[.FieldName]](id [[.IdFieldType]]) error {
	return cl.DeleteAttributes(id, "[[.BsonName]]")
}

func (cl *[[.RootName]]DaoClient) Get[[.RootName]][[.FieldName]](id [[.IdFieldType]]) ([[.FieldType]], error) {
	[[.LowerName]], err := cl.GetAttributes(id, "[[.BsonName]]")
	if err != nil {
		return "", err
	}
	return [[.LowerName]].Name, err
}

func (cl [[.RootName]]DaoClient) Get[[.RootName]]sBy[[.FieldName]]([[.LowerName]] [[.FieldType]]) ([][[.PackageName]].[[.RootName]], error) {
	[[.LowerName]]s, err := cl.Get(bson.M{"[[.BsonName]]": [[.LowerName]]})
	if err != nil {
		return nil, err
	}
	return [[.LowerName]]s, nil
}
`
)
