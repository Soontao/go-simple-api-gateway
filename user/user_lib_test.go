package user

import (
	"encoding/json"
	"fmt"
	"testing"
	"github.com/Soontao/go-simple-api-gateway/types"
)

func TestNewUser(t *testing.T) {
	uJson, _ := json.Marshal(types.NewEncryptedUser("tetsUser", "testPass"))
	fmt.Println(string(uJson))
}
