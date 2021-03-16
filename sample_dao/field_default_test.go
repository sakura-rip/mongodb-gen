package dao

import (
	"fmt"
	"testing"
)

func TestInsertOrUpdateUserEmail(t *testing.T) {
	err := cl.InsertOrUpdateUserEmail(uid, "newEmail")
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteUserEmail(t *testing.T) {
	err := cl.DeleteUserEmail(uid)
	if err != nil {
		t.Error(err)
	}
}

func TestGetUserEmail(t *testing.T) {
	email, err := cl.GetUserEmail(uid)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v\n", email)
}

func TestGetUserByEmail(t *testing.T) {
	user, err := cl.GetUserByEmail("email")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v\n", user)
}
