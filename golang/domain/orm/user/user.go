package user

import (
	"errors"

	g_ "test-github/domain/repository/gorm"

	"github.com/jinzhu/gorm"

	"test-github/domain/entity/user"
)

// db, _ := gorm.NewDb()

type UserOrm struct {
	db *gorm.DB
}

func (s *UserOrm) NewUserOrm() {
	g := g_.Gorm{}
	s.db = g.NewDb()[0]
	s.db.AutoMigrate(&user.User{})
}

func (s *UserOrm) GetAll() ([]*user.User, error) {

	societies := make([]*user.User, 0)

	if s.db.Find(&societies).Error != nil {
		return nil, errors.New("Datacenter not found")
	}

	return societies, nil
}
