package stash

import (
	"github.com/sakura-rip/mongodbgen/sample"
	"go.mongodb.org/mongo-driver/bson"
)

func (cl *UserDaoClient) DeleteTagField(userId string) error {
	return cl.DeleteAttributes(userId, bson.M{"tag": 1})
}

func (cl *UserDaoClient) InsertOrUpdateTagField(userId string, object interface{}) error {
	return cl.UpdateAttribute(userId, bson.M{"tag": object})
}

func (cl *UserDaoClient) DeleteUserTag(userId, key string) error {
	return cl.DeleteAttributes(userId, bson.M{"tag." + key: 1})
}

func (cl *UserDaoClient) InsertOrUpdateUserTag(userId, key string, tag sample.Tag) error {
	return cl.UpdateAttribute(userId, bson.M{"tag." + key: tag})
}

func (cl *UserDaoClient) UpdateUserTagId(userId, key, id string) error {
	return cl.UpdateAttribute(userId, bson.M{"tag." + key + ".id": id})
}

func (cl *UserDaoClient) DeleteUserTagId(userId, key string) error {
	return cl.DeleteAttributes(userId, bson.M{"tag." + key + ".id": 1})
}
