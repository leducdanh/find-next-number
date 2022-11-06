package user

import (
	"errors"
	"time"

	g_ "test-github/domain/repository/gorm"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"

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

	users := make([]*user.User, 0)

	if s.db.Find(&users).Error != nil {
		return nil, errors.New("Datacenter not found")
	}

	return users, nil
}

func (s *UserOrm) PostUser(c echo.Context) (interface{}, error) {
	var user user.User
	err := c.Bind(&user)
	if err != nil {
		return nil, err
	}
	user.CretaeAt = time.Now()
	user.UpdateAt = time.Now()

	result := s.db.Create(&user) // pass pointer of data to Create
	if result.Error != nil {
		return nil, errors.New("Create fail")
	}

	return result, nil
}
