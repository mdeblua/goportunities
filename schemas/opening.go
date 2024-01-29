package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int
}

type OpeningResponse struct {
	Id        uint      `json:"id"`
	createdAt time.Time `json:"created_at"`
	updatedAt time.Time `json:"updated_at"`
	deletedAt time.Time `json:"deleted_at, omitEmpty"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Remote    bool      `json:"remote"`
	Link      string    `json:"link"`
	Salary    int       `json:"salary"`
}
