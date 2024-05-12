package pkg

import (
	"invoiceApi/models"
)

type JwtToken struct {
	UserId    string               `json:"userId,omitempty"`
	ExpiresAt int64                `json:"expired,omitempty"`
	User      *models.User 			`json:"user,omitempty"`
}

// Valid implements jwt.Claims.
func (j JwtToken) Valid() error {
	panic("unimplemented")
}
