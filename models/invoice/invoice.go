package invoice

import (
	"invoiceApi/models"

	"gorm.io/gorm"
)


type Invoice struct {
	gorm.Model
	Title  string `json:"title"`
	Published *bool `json:"published"`
	UserId int32   `gorm:"column:userId" json:"userId"`
	Author  *models.User `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
}
