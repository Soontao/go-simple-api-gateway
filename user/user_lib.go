package user

import (
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

func CryptPass(pass string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(bytes)
}

func ComparePassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// User db model
type User struct {
	UUID      string    `json:"uuid" xorm:"'uuid' pk"`
	Username  string    `json:"username" xorm:"unique"`
	Password  string    `json:"-"`
	CreatedAt time.Time `xorm:"created" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated" json:"updated_at"`
}

// NewEncryptedUser Entity
func NewEncryptedUser(username, password string) (u *User) {
	u = new(User)
	u.Username = username
	u.Password = CryptPass(password)
	u.UUID = strings.Replace(uuid.NewV4().String(), "-", "", -1)
	return
}
