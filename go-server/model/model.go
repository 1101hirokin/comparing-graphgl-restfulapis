package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id        uint      `json:"id" gorm:"unique" gorm:"AUTO_INCREMENT"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Bio       string    `json:"bio"`
	UrlAvatar string    `json:"url_avatar"`
}

type Measurement struct {
	MemoryUsageChange uint64  `json:"memory_usage_change"`
	ProcessTime       float64 `json:"process_time"`
}

type Pagination struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type UsersResponse struct {
	Success     bool         `json:"success"`
	Message     string       `json:"message"`
	Users       *[]User      `json:"users"`
	Pagination  *Pagination  `json:"pagination"`
	Measurement *Measurement `json:"measurement"`
}
type UserResponse struct {
	Success     bool         `json:"success"`
	Message     string       `json:"message"`
	User        *User        `json:"user"`
	Measurement *Measurement `json:"measurement"`
}
type GraphQLResponse struct {
	Success     bool         `json:"success"`
	Message     string       `json:"message"`
	Result      interface{}  `json:"result"`
	Measurement *Measurement `json:"measurement"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	return db
}
