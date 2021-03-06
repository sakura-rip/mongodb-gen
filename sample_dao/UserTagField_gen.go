// auto generated by https://github.com/sakura-rip/mongodb-gen
// Version : develop@v1
// DO NOT EDIT

package dao

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func (cl *UserDaoClient) DeleteTagField(id string) error {
	return cl.DeleteAttributes(id, bson.M{"tag": 1})
}

func (cl *UserDaoClient) InsertOrUpdateTagField(id string, object interface{}) error {
	return cl.UpdateAttribute(id, bson.M{"tag": object})
}

func (cl *UserDaoClient) DeleteUserTag(id string, key interface{}) error {
	return cl.DeleteAttributes(id, bson.M{fmt.Sprintf("tag.%s", key): 1})
}

func (cl *UserDaoClient) InsertOrUpdateUserTag(id string, key interface{}, value interface{}) error {
	return cl.UpdateAttribute(id, bson.M{fmt.Sprintf("tag.%s", key): value})
}
