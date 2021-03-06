// auto generated by https://github.com/sakura-rip/mongodb-gen
// Version : develop@v1
// DO NOT EDIT

package dao

import (
	"github.com/sakura-rip/mongodb-gen/sample"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (cl *NewDaoClient) DeleteOne(id string) (int64, error) {
	result, err := cl.Col.DeleteOne(cl.ctx, bson.M{"_id": id})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, err
}

func (cl *NewDaoClient) DeleteMany(attr bson.M) (int64, error) {
	result, err := cl.Col.DeleteMany(cl.ctx, attr)
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, err
}

func (cl *NewDaoClient) InsertOne(new *sample.New) (string, error) {
	result, err := cl.Col.InsertOne(cl.ctx, new)
	if err != nil {
		return "", err
	}
	if oid, ok := result.InsertedID.(string); ok {
		return oid, nil
	}
	return "", nil
}

func (cl *NewDaoClient) InsertMany(objs []interface{}) ([]string, error) {
	result, err := cl.Col.InsertMany(cl.ctx, objs)
	if err != nil {
		return nil, err
	}
	var strIds []string
	for _, id := range result.InsertedIDs {
		strIds = append(strIds, id.(string))
		return strIds, nil
	}
	return strIds, nil
}

func (cl *NewDaoClient) UpdateAttribute(id string, attrToUpdate bson.M) error {
	_, err := cl.Col.UpdateOne(cl.ctx, bson.M{"_id": id}, Set(attrToUpdate))
	return err
}

func (cl *NewDaoClient) DeleteAttributes(id string, attrToDelete bson.M) error {
	_, err := cl.Col.UpdateOne(cl.ctx, bson.M{"_id": id}, Unset(attrToDelete))
	return err
}

func (cl *NewDaoClient) AddToArray(id string, object bson.M) error {
	_, err := cl.Col.UpdateOne(cl.ctx, bson.M{"_id": id}, AddToSet(object))
	return err
}

func (cl *NewDaoClient) RemoveFromArray(id string, object bson.M) error {
	_, err := cl.Col.UpdateOne(cl.ctx, bson.M{"_id": id}, Pull(object))
	return err
}

func (cl *NewDaoClient) Get(filter bson.M) ([]sample.New, error) {
	res, err := cl.Col.Find(
		cl.ctx, filter)
	if err != nil {
		return nil, err
	}
	var news []sample.New
	if err := res.All(cl.ctx, &news); err != nil {
		return nil, err
	}
	return news, err
}

func (cl *NewDaoClient) GetNew(id string) (sample.New, error) {
	res := cl.Col.FindOne(
		cl.ctx,
		bson.D{{"_id", id}},
	)
	var new sample.New
	if err := res.Decode(&new); err != nil {
		return new, err
	}

	return new, nil
}

func (cl *NewDaoClient) GetAll() ([]sample.New, error) {
	return cl.Get(bson.M{})
}

func (cl *NewDaoClient) GetAttribute(id string, attributes bson.D) (sample.New, error) {
	res := cl.Col.FindOne(
		cl.ctx,
		bson.D{{"_id", id}},
		options.FindOne().SetProjection(attributes),
	)
	var new sample.New
	if err := res.Decode(&new); err != nil {
		return new, err
	}
	return new, nil
}

func (cl *NewDaoClient) GetAttributes(id string, attributes ...string) (sample.New, error) {
	var projections = bson.D{}
	for _, attribute := range attributes {
		projections = append(projections, bson.E{Key: attribute, Value: 1})
	}
	return cl.GetAttribute(id, projections)
}
