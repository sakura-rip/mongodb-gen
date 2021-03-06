// auto generated by https://github.com/sakura-rip/mongodb-gen
// Version : develop@v1
// DO NOT EDIT

package dao

import (
	"github.com/sakura-rip/mongodb-gen/sample"
	"go.mongodb.org/mongo-driver/bson"
)

func (cl *UserDaoClient) DeleteProfile(id string) error {
	return cl.DeleteAttributes(id, bson.M{"profile": 1})
}

func (cl *UserDaoClient) InsertOrUpdateProfile(id string, object interface{}) error {
	return cl.UpdateAttribute(id, bson.M{"profile": object})
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

func (cl *UserDaoClient) UpdateOrInsertUserProfileName(id string, name string) error {
	return cl.UpdateAttribute(id, bson.M{"profile.name": name})
}

func (cl *UserDaoClient) DeleteUserProfileName(id string) error {
	return cl.DeleteAttributes(id, bson.M{"profile.name": 1})
}

func (cl *UserDaoClient) GetUserProfileName(id string) (val string, err error) {
	user, err := cl.GetAttribute(id, bson.D{{"profile.name", 1}})
	if err != nil {
		return val, err
	}
	return user.Profile.Name, err
}

func (cl *UserDaoClient) GetUserByProfileName(name string) ([]sample.User, error) {
	return cl.Get(bson.M{"profile.name": name})
}
