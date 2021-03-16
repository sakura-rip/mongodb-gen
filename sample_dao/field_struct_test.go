package dao

import (
	"fmt"
	"github.com/sakura-rip/mongodbgen/sample"
	"testing"
)

func TestDeleteProfile(t *testing.T) {
	err := cl.DeleteProfile(uid)
	if err != nil {
		t.Error(err)
	}
}
func TestInsertOrUpdateProfile(t *testing.T) {
	err := cl.InsertOrUpdateProfile(uid, "this is new profile")
	if err != nil {
		t.Error(err)
	}
}
func TestGetUserProfile(t *testing.T) {
	profile, err := cl.GetUserProfile(uid)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v\n", profile)
}
func TestGetUserByProfile(t *testing.T) {
	profile, err := cl.GetUserByProfile(sample.Profile{Name: "test"})
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v\n", profile)
}
func TestUpdateOrInsertUserProfileName(t *testing.T) {
	err := cl.UpdateOrInsertUserProfileName(uid, "new name")
	if err != nil {
		t.Error(err)
	}
}
func TestDeleteUserProfileName(t *testing.T) {
	err := cl.DeleteUserProfileName(uid)
	if err != nil {
		t.Error(err)
	}
}
func TestGetUserProfileName(t *testing.T) {
	name, err := cl.GetUserProfileName(uid)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v\n", name)
}
func TestGetUserByProfileName(t *testing.T) {
	name, err := cl.GetUserByProfileName("test")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v\n", name)
}
