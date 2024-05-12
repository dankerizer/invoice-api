package middleware

import (
	"errors"
	"invoiceApi/helper"
	"invoiceApi/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"

	"github.com/golang-jwt/jwt"
)

type ResponseToken struct {
	Token   string         `json:"token,omitempty"`
	Expired int64          `json:"expired,omitempty"`
	User    *models.User `json:"user,omitempty"`

}

func GenerateJWTToken(user *models.User) (ResponseToken, error) {
jwtSecret := os.Getenv("JWT_SECRET")

	// claims := jwt.NewWithClaims(jwt.SigningMethodHS256, pkg.JwtToken{
	// 	UserId:    strconv.Itoa(int(user.ID)),
	// 	ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	// 	User:      user,
	// })

	expiredTime := time.Now().Add(time.Hour * 24).Unix() //1 day

	claims := jwt.MapClaims{
		"user": user,
		"iat":      time.Now().Unix(),
		"ExpiredAt":    expiredTime,
	}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// token, err := claims.SignedString([]byte(jwtSecret))
	t, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		log.Printf("Error generate jwt %v", err)
		return ResponseToken{}, &helper.RequestError{
			StatusCode: fiber.StatusInternalServerError,
			Err:        errors.New("failed generate jwt"),
		}
	}

reponse := ResponseToken{
		Token:   t,
		Expired:expiredTime,

	}

	return reponse, nil
}
