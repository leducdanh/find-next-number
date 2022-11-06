package user

import "time"

type User struct {
	ID       string    `json:"id" db:"id" gorm:"primaryKey; column:id" validate:"required"`
	Name     string    `json:"name,omitempty" validate:"required"`
	Email    string    `json:"email,omitempty" validate:"required,email"`
	CretaeAt time.Time `json:"create_at,omitempty"`
	UpdateAt time.Time `json:"update_at,omitempty"`
}

func (User) TableName() string {
	return "users"
}
