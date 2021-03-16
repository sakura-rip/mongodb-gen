package dao

import (
	"github.com/sakura-rip/mongodbgen/sample"
	"go.mongodb.org/mongo-driver/bson"
)

func (cl *UserDaoClient) DeleteProfile(userId string) error {
	return cl.DeleteAttributes(userId, bson.M{"profile": 1})
}

func (cl *UserDaoClient) InsertOrUpdateProfile(userId string, object interface{}) error {
	return cl.UpdateAttribute(userId, bson.M{"profile": object})
}

func (cl *UserDaoClient) GetUserProfile(id string) (sample.Profile, error) {
	attribute, err := cl.GetAttribute(id, bson.D{{"profile", 1}})
	if err != nil {
		return sample.Profile{}, err
	}
	return attribute.Profile, err
}

func (cl *UserDaoClient) GetUserByProfile(profile sample.Profile) ([]sample.User, error) {
	return cl.Get(bson.M{"profile": profile})
}

func (cl *UserDaoClient) UpdateOrInsertUserProfileName(userId, name string) error {
	return cl.UpdateAttribute(userId, bson.M{"profile.name": name})
}

func (cl *UserDaoClient) DeleteUserProfileName(userId string) error {
	return cl.DeleteAttributes(userId, bson.M{"profile.name": 1})
}

func (cl *UserDaoClient) GetUserProfileName(id string) (string, error) {
	attribute, err := cl.GetAttribute(id, bson.D{{"profile.name", 1}})
	if err != nil {
		return "", err
	}
	return attribute.Profile.Name, err
}

func (cl *UserDaoClient) GetUserByProfileName(name string) ([]sample.User, error) {
	return cl.Get(bson.M{"profile.name": name})
}
