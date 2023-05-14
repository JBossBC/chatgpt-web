package utils

import (
	"chatgpt-web/internal/app/dao"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func MarshalJWT(user dao.User) ([]byte, error) {
	mapping := jwt.MapClaims{}
	mapping["username"] = user.Username
	mapping["level"] = user.Level
	mapping["levelDeadline"] = user.LevelDeadline
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return []byte(nil), err
	}
	mapping["exd"] = time.Now().Add(4 * time.Hour).In(location)
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, mapping)
	tokenStr, err := token.SigningString()
	if err != nil {
		return []byte(nil), nil
	}
	return []byte(tokenStr), nil
}
