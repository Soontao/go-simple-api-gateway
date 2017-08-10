package user

import (
	"testing"
)

func TestUserService_SaveUser(t *testing.T) {
	us := NewUserService("")
	us.engine.ShowSQL(true)
	username := "testUser"
	password := "testPassword"
	newPassword := "newTestPassword"

	us.engine.Exec("delete from user where username=?", username)

	err := us.SaveUser(username, password)
	if err != nil {
		t.Fatal(err)
	}
	success := us.AuthUser(username, password)
	if !success {
		t.Log("auth failed")
		t.Fail()
	}
	success = us.UpdatePassword(username, password, newPassword)
	if !success {
		t.Log("updated failed")
		t.Fail()
	}
	success = us.AuthUser(username, newPassword)
	if !success {
		t.Log("auth failed")
		t.Fail()
	}
}
