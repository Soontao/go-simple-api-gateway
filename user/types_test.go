package user

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewUser(t *testing.T) {
	uJson, _ := json.Marshal(NewEncryptedUser("tetsUser", "testPass"))
	fmt.Println(string(uJson))
}
