package dao

import (
	"fmt"
	"testing"
)

func TestInsertItem(t *testing.T) {
	err := cl.InsertItem(uid, "newItem")
	if err != nil {
		t.Error(err)
	}
}
func TestDeleteItem(t *testing.T) {
	err := cl.DeleteItem(uid, "item1")
	if err != nil {
		t.Error(err)
	}
}
func TestGetUserByItemIn(t *testing.T) {
	user, err := cl.GetUserByItemIn("item3")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v\n", user)
}

func TestGetUserByItemAll(t *testing.T) {
	user, err := cl.GetUserByItemAll("item1", "item2", "item3")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v\n", user)
}
func TestGetUserByItemSize(t *testing.T) {
	user, err := cl.GetUserByItemSize(3)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v\n", user)
}
