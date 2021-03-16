package dao

import (
	"github.com/sakura-rip/mongodbgen/sample"
	"go.mongodb.org/mongo-driver/bson"
)

func (cl *UserDaoClient) InsertItem(userId string, item string) error {
	return cl.AddToArray(userId, bson.M{"items": item})
}

func (cl *UserDaoClient) DeleteItem(userId, item string) error {
	return cl.RemoveFromArray(userId, bson.M{"items": item})
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
