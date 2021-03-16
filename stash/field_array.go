package stash

import (
	"github.com/sakura-rip/mongodbgen/sample"
	"go.mongodb.org/mongo-driver/bson"
)

func (cl *UserDaoClient) DeleteItemField(id string) error {
	return cl.DeleteAttributes(id, bson.M{"items": 1})
}

func (cl *UserDaoClient) InsertOrUpdateItemField(id string, object interface{}) error {
	return cl.UpdateAttribute(id, bson.M{"items": object})
}

func (cl *UserDaoClient) InsertItem(id string, item string) error {
	return cl.AddToArray(id, bson.M{"items": item})
}

func (cl *UserDaoClient) DeleteItem(id, item string) error {
	return cl.RemoveFromArray(id, bson.M{"items": item})
}

func (cl *UserDaoClient) GetUserByItemIn(item ...interface{}) ([]sample.User, error) {
	return cl.Get(bson.M{"items": In(item...)})
}

func (cl *UserDaoClient) GetUserByItemElemMatch(elem bson.M) ([]sample.User, error) {
	return cl.Get(bson.M{"items": ElemMatch(elem)})
}

func (cl *UserDaoClient) GetUserByItemAll(attr ...interface{}) ([]sample.User, error) {
	return cl.Get(bson.M{"items": All(attr)})
}

func (cl *UserDaoClient) GetUserByItemSize(size int64) ([]sample.User, error) {
	return cl.Get(bson.M{"items": Size(size)})
}
