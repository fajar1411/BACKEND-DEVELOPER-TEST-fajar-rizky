package middlewares

import (
	"test/config"

	"github.com/golang-jwt/jwt"
)

func InitJWTS(c *config.AppConfig) {
	key = c.JWT_SECRET
}

func ExtractToken(t interface{}) int {
	user := t.(*jwt.Token)
	userId := -1
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		switch claims["userID"].(type) {
		case float64:
			userId = int(claims["userID"].(float64))
		case int:
			userId = claims["userID"].(int)
		}
	}
	return userId
}

func GenerateJWT(id int) (string, interface{}) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	useToken, _ := token.SignedString([]byte(config.GetConfig().JWT_SECRET))
	return useToken, token
}
