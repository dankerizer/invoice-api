package invoice

import (
	"gorm.io/gorm"
)


type InvoiceDetail struct {
	gorm.Model
	InvoiceId int32   `gorm:"column:invoiceId" json:"invoiceId"`
	Key	string `json:"key"`
	Value string `json:"value"`
}
