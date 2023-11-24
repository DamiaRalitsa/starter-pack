package commands

import (
	"gorm.io/gorm"
)

type Users struct {
	ID        int    `gorm:"primary key;autoincrement" json:"id"`
	Name      string `json:"name"`
	Title     string `json:"title"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at`
}

func MigrateUser(db *gorm.DB) error {
	err := db.AutoMigrate(&Users{})
	return err
}
