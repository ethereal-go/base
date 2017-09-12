package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/ethereal-go/base/root/database"

)

type UserService struct {
	DB *gorm.DB
}

// User returns a user for a given id.
func (s UserService) Users() (users []*database.User, err error) {
	s.DB.Find(&users)
	return
}

func  (s UserService) User(id int) (user *database.User, err error)  {
	s.DB.First(user, id)
	s.DB.Model(&user).Related(&database.Role{})
	return
}

func (s *UserService) CreateUser(u *database.User) {

}

func (s *UserService) DeleteUser(id int) {

}
