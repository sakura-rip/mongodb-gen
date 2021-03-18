// auto generated by github.com/sakura-rip tool
// Version : develop@v1
// DO NOT EDIT

package dao

import (
	"github.com/sakura-rip/mongodbgen/sample"
	"go.mongodb.org/mongo-driver/bson"
)

func (cl *UserDaoClient) DeleteItemsField(id string) error {
	return cl.DeleteAttributes(id, bson.M{"items": 1})
}

func (cl *UserDaoClient) InsertOrUpdateItemsField(id string, object string) error {
	return cl.UpdateAttribute(id, bson.M{"items": object})
}

func (cl *UserDaoClient) InsertItems(id string, object string) error {
	return cl.AddToArray(id, bson.M{"items": object})
}

func (cl *UserDaoClient) DeleteItems(id string, object string) error {
	return cl.RemoveFromArray(id, bson.M{"items": object})
}

func (cl *UserDaoClient) GetUserByItemsIn(objects ...interface{}) ([]sample.User, error) {
	return cl.Get(bson.M{"items": In(objects...)})
}

func (cl *UserDaoClient) GetUserByItemsElemMatch(elem bson.M) ([]sample.User, error) {
	return cl.Get(bson.M{"items": ElemMatch(elem)})
}

func (cl *UserDaoClient) GetUserByItemsAll(attr ...interface{}) ([]sample.User, error) {
	return cl.Get(bson.M{"items": All(attr)})
}

func (cl *UserDaoClient) GetUserByItemsSize(size int64) ([]sample.User, error) {
	return cl.Get(bson.M{"items": Size(size)})
}