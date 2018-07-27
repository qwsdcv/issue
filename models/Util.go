package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	jwt "github.com/dgrijalva/jwt-go"
)

// FormatTime formats *time.Time to string with format "2006-01-02 15:04:05"
func FormatTime(t *time.Time) (ret string) {
	return t.Format("2006-01-02 15:04:05")
}

//CreateToken will return a token.
func CreateToken() (tokenStr string, err error) {
	key := beego.AppConfig.String("private_key")
	keyBuff := []byte(key)
	claims := &jwt.StandardClaims{
		ExpiresAt: 10,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString(keyBuff)
	return
}

//Valid check the token is valid.
func Valid(tokenString string) (valid bool, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		key := beego.AppConfig.String("private_key")
		keyBuff := []byte(key)
		return keyBuff, nil
	})
	if err != nil {
		valid = false
	} else {
		valid = token.Valid
	}

	return
}
