package user

import (
	"encoding/json"
	"fmt"
	"github.com/Soontao/go-simple-api-gateway/types"
	"testing"
)

func TestNewUser(t *testing.T) {
	uJson, _ := json.Marshal(types.NewEncryptedUser("tetsUser", "testPass"))
	fmt.Println(string(uJson))
}
