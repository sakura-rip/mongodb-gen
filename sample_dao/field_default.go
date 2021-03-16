package dao

import (
	"github.com/sakura-rip/mongodbgen/sample"
	"go.mongodb.org/mongo-driver/bson"
)

func (cl *UserDaoClient) InsertOrUpdateUserEmail(id, newEmail string) error {
	return cl.UpdateAttribute(id, bson.M{"email": newEmail})
}

func (cl *UserDaoClient) DeleteUserEmail(id string) error {
	return cl.DeleteAttributes(id, bson.M{"email": 1})
}

func (cl *UserDaoClient) GetUserEmail(id string) (string, error) {
	user, err := cl.GetAttributes(id, "email")
	if err != nil {
		return "", err
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
