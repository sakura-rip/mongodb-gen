package stash

import (
	"github.com/sakura-rip/mongodbgen/sample"
	"testing"
)

func TestDeleteTagField(t *testing.T) {
	err := cl.DeleteTagField(uid)
	if err != nil {
		t.Error(err)
	}
}
func TestInsertOrUpdateTagField(t *testing.T) {
	err := cl.InsertOrUpdateTagField(uid, sample.Tag{
		Id:          "newtag",
		Description: "",
	})
	if err != nil {
		t.Error(err)
	}
}
func TestDeleteUserTag(t *testing.T) {
	err := cl.DeleteUserTag(uid, "tag1")
	if err != nil {
		t.Error(err)
	}
}
func TestInsertOrUpdateUserTag(t *testing.T) {
	err := cl.InsertOrUpdateUserTag(uid, "newtag2", sample.Tag{
		Id:          "newtag2",
		Description: "this is new tag 2",
	})
	if err != nil {
		t.Error(err)
	}
}
func TestUpdateUserTagId(t *testing.T) {
	err := cl.UpdateUserTagId(uid, "tag1", "newtagid1")
	if err != nil {
		t.Error(err)
	}
}
func TestDeleteUserTagId(t *testing.T) {
	err := cl.DeleteUserTagId(uid, "tag1")
	if err != nil {
		t.Error(err)
	}
}
