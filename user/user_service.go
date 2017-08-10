package user

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type UserService struct {
	engine *xorm.Engine
}

func NewUserService(connStr string) *UserService {
	rt := &UserService{}
	engine, err := xorm.NewEngine("mysql", connStr)
	if err != nil {
		panic(err)
	}
	rt.engine = engine
	rt.engine.ShowSQL(true)
	err = engine.Sync2(new(User))
	if err != nil {
		panic(err)
	}
	return rt
}

func (us *UserService) AuthUser(username, password string) (b bool) {
	b = false
	user := &User{Username: username}
	us.engine.Get(user)
	if ComparePassword(user.Password, password) == nil {
		b = true
	}
	return
}

func (us *UserService) SaveUser(username, password string) (err error) {
	has, _ := us.engine.Get(&User{Username: username})
	if has {
		return errors.Errorf("user %s has existed", username)
	}
	_, err = us.engine.Insert(NewEncryptedUser(username, password))
	return
}

func (us *UserService) UpdatePassword(userName, oldPassword, newPassword string) (success bool) {
	success = false
	user := &User{Username: userName}
	success, _ = us.engine.Get(user)
	if !success {
		return
	}
	if ComparePassword(user.Password, oldPassword) == nil {
		rs, _ := us.engine.Exec("update `user` set `password` = ? where `username` = ?", CryptPass(newPassword), userName)
		affected, _ := rs.RowsAffected()
		if affected > 0 {
			success = true
		}
		return
	} else {
		success = false
		return
	}

}
