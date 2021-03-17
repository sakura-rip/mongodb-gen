// auto generated by github.com/sakura-rip tool
// Version : develop@v1
// DO NOT EDIT

package dao

import (
	"github.com/sakura-rip/mongodbgen/sample"
	"go.mongodb.org/mongo-driver/bson"
)

func (cl *UserDaoClient) InsertOrUpdateUserEmail(id string, newEmail string) error {
	return cl.UpdateAttribute(id, bson.M{"email": newEmail})
}

func (cl *UserDaoClient) DeleteUserEmail(id string) error {
	return cl.DeleteAttributes(id, bson.M{"email": 1})
}

func (cl *UserDaoClient) GetUserEmail(id string) (val string, err error) {
	user, err := cl.GetAttributes(id, "email")
	if err != nil {
		return val, err
	}
	return user.Email, err
}

func (cl *UserDaoClient) GetUserByEmail(email string) ([]sample.User, error) {
	users, err := cl.Get(bson.M{"email": email})
	if err != nil {
		return nil, err
	}
	return users, nil
}
