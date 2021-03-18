// auto generated by https://github.com/sakura-rip/mongodb-gen
// Version : develop@v1
// DO NOT EDIT

package dao

import (
	"github.com/sakura-rip/mongodbgen/sample"
	"go.mongodb.org/mongo-driver/bson"
)

func (cl *GroupDaoClient) DeleteMembersField(id string) error {
	return cl.DeleteAttributes(id, bson.M{"members": 1})
}

func (cl *GroupDaoClient) InsertOrUpdateMembersField(id string, object string) error {
	return cl.UpdateAttribute(id, bson.M{"members": object})
}

func (cl *GroupDaoClient) InsertMembers(id string, object string) error {
	return cl.AddToArray(id, bson.M{"members": object})
}

func (cl *GroupDaoClient) DeleteMembers(id string, object string) error {
	return cl.RemoveFromArray(id, bson.M{"members": object})
}

func (cl *GroupDaoClient) GetGroupByMembersIn(objects ...interface{}) ([]sample.Group, error) {
	return cl.Get(bson.M{"members": In(objects...)})
}

func (cl *GroupDaoClient) GetGroupByMembersElemMatch(elem bson.M) ([]sample.Group, error) {
	return cl.Get(bson.M{"members": ElemMatch(elem)})
}

func (cl *GroupDaoClient) GetGroupByMembersAll(attr ...interface{}) ([]sample.Group, error) {
	return cl.Get(bson.M{"members": All(attr)})
}

func (cl *GroupDaoClient) GetGroupByMembersSize(size int64) ([]sample.Group, error) {
	return cl.Get(bson.M{"members": Size(size)})
}
