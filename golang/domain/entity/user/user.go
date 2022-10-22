package user

type User struct {
	ID       string `json:"id" db:"id" gorm:"primaryKey; column:id"`
	Name     string `json:"name,omitempty"`
	Logo     string `json:"logo,omitempty"`
	Email    string `json:"email,omitempty"`
	CretaeAt string `json:"create_at,omitempty"`
	UpdateAt string `json:"update_at,omitempty"`
}

func (User) TableName() string {
	return "users"
}
