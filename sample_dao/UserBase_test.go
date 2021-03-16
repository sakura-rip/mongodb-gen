package dao

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sakura-rip/mongodbgen/sample"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

var (
	cl         *UserDaoClient
	uid        string
	SampleUser *sample.User
)

func newUuid() string {
	uid, _ := uuid.NewUUID()
	return uid.String()
}

func newUser(uuid string) *sample.User {
	return &sample.User{
		Id:    uuid,
		Email: "email",
		Profile: sample.Profile{
			Name: "test",
		},
		Profiles: []sample.Profile{{Name: "test2"}},
		Items:    []string{"item1", "item2", "item3"},
		Hoge:     1245,
		Fuck:     4589,
		Tag: map[string]sample.Tag{
			"tag1": {Description: "this is item1"},
			"tag2": {Description: "this is item2"},
		},
	}
}

func init() {
	cl = NewUserDaoClient()
	uid = newUuid()
	SampleUser = newUser(uid)
	cl.InsertOne(SampleUser)
}

func TestInsertOne(t *testing.T) {
	one, err := cl.InsertOne(SampleUser)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v\n", one)
}

func TestInsertMany(t *testing.T) {
	one, err := cl.InsertMany([]interface{}{SampleUser})
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v\n", one)
}

func TestDeleteOne(t *testing.T) {
	TestInsertMany(t)
	one, err := cl.DeleteOne(uid)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v\n", one)
}

func TestUpdateAttribute(t *testing.T) {
	TestInsertOne(t)
	err := cl.UpdateAttribute(uid, bson.M{"email": "new email"})
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteAttributes(t *testing.T) {
	TestInsertOne(t)
	err := cl.DeleteAttributes(uid, bson.M{"email": 1, "hoge": 1})
	if err != nil {
		t.Error(err)
	}
}

func TestAddToArray(t *testing.T) {
	TestInsertOne(t)
	err := cl.AddToArray(uid, bson.M{"items": "newitem"})
	if err != nil {
		t.Error(err)
	}
}

func TestUserDaoClient_RemoveFromArray(t *testing.T) {
	TestInsertOne(t)
	err := cl.RemoveFromArray(uid, bson.M{"items": "item1"})
	if err != nil {
		t.Error(err)
	}
}

func TestUserDaoClient_Get(t *testing.T) {
	TestInsertOne(t)
	get, err := cl.Get(And(bson.M{"email": "email"}))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v\n", get)
}

func TestUserDaoClient_GetAll(t *testing.T) {
	TestInsertOne(t)
	all, err := cl.GetAll()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v\n", all)
}

func TestUserDaoClient_GetAttribute(t *testing.T) {
	TestInsertOne(t)
	attribute, err := cl.GetAttributes(uid, "hoge")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v\n", attribute)
}
