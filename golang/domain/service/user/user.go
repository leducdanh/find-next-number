package user

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"test-github/domain/orm/user"
	"test-github/util"
)

type UserService struct {
	orm  *user.UserOrm
	resp *util.Response
}

func (s *UserService) NewUserOrm() {
	s.resp = &util.Response{}
	s.orm = &user.UserOrm{}
	s.orm.NewUserOrm()
}

func (s *UserService) ListAll(c echo.Context) error {
	societies, _ := s.orm.GetAll()
	s.resp.Status = "OK"
	s.resp.Data = societies
	return c.JSON(http.StatusOK, s.resp)
}
