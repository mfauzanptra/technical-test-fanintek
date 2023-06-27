package helper

import (
	"fmt"
	"testAPI/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func ExtractToken(t interface{}) (int, string) {
	user := t.(*jwt.Token)
	userId := -1
	npp := ""
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		switch claims["userID"].(type) {
		case float64:
			userId = int(claims["userID"].(float64))
			npp = fmt.Sprint(claims["npp"])
		case int:
			userId = claims["userID"].(int)
			npp = fmt.Sprint(claims["npp"])
		}
		return int(userId), npp
	}
	return -1, ""
}
func GenerateJWT(id int, npp string) (string, interface{}) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = id
	claims["npp"] = npp
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	useToken, _ := token.SignedString([]byte(config.JWT_KEY))
	return useToken, token
}
