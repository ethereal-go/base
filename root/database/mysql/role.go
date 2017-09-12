package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/ethereal-go/base/root/database"
)

type RoleService struct {
	DB *gorm.DB
}

// User returns a user for a given id.
func (s RoleService) Roles() (roles []*database.Role, err error) {
	s.DB.Find(&roles)
	return
}

func  (s RoleService) Role(id int) (role *database.Role, err error)  {
	s.DB.First(role, id)
	s.DB.Model(&role).Related(&database.Role{})
	return
}

func (s *RoleService) CreateRole(u *database.Role) {

}

func (s *RoleService) DeleteRole(id int) {

}
