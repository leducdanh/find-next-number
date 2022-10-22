package gorm

import (
	"test-github/domain/repository"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"fmt"
)

type Gorm struct {
	c *repository.Config
}

func (g *Gorm) NewDb() []*gorm.DB {
	g.c = &repository.Config{}

	c, _ := g.c.NewConfig()

	arrayConnections := make([]*gorm.DB, 0)

	var db *gorm.DB

	var err error

	if c.Database != "" {
		psqlConnStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			c.UserName,
			c.Password,
			c.Host,
			c.DB_Port,
			c.Database,
		)

		db, err = gorm.Open("postgres", psqlConnStr)

		if err != nil {
			fmt.Printf("err postgres = %v \n", err)
		}

		arrayConnections = append(arrayConnections, db)

		db = &gorm.DB{}
	}
	return arrayConnections
}
