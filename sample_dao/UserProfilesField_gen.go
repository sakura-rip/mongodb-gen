// auto generated by github.com/sakura-rip tool
// Version : develop@v1
// DO NOT EDIT

package dao

import (
	"github.com/sakura-rip/mongodbgen/sample"
	"go.mongodb.org/mongo-driver/bson"
)

func (cl *UserDaoClient) DeleteProfilesField(id string) error {
	return cl.DeleteAttributes(id, bson.M{"profiles": 1})
}

func (cl *UserDaoClient) InsertOrUpdateProfilesField(id string, object interface{}) error {
	return cl.UpdateAttribute(id, bson.M{"profiles": object})
}

func (cl *UserDaoClient) InsertProfiles(id string, object interface{}) error {
	return cl.AddToArray(id, bson.M{"profiles": object})
}

func (cl *UserDaoClient) DeleteProfiles(id string, object interface{}) error {
	return cl.RemoveFromArray(id, bson.M{"profiles": object})
}

func (cl *UserDaoClient) GetUserByProfilesIn(objects ...interface{}) ([]sample.User, error) {
	return cl.Get(bson.M{"profiles": In(objects...)})
}

func (cl *UserDaoClient) GetUserByProfilesElemMatch(elem bson.M) ([]sample.User, error) {
	return cl.Get(bson.M{"profiles": ElemMatch(elem)})
}

func (cl *UserDaoClient) GetUserByProfilesAll(attr ...interface{}) ([]sample.User, error) {
	return cl.Get(bson.M{"profiles": All(attr)})
}

func (cl *UserDaoClient) GetUserByProfilesSize(size int64) ([]sample.User, error) {
	return cl.Get(bson.M{"profiles": Size(size)})
}
