package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/bobopylabepolhk/ypshortener/pkg/urlutils"
	"github.com/labstack/echo/v4"
)

var UserIDCookie string = "ypshortener_user_id"
var idLength int = 12

func ValidateUserID(userID string, secret string) bool {
	hash := hmac.New(sha256.New, []byte(secret))
	binaryToken := []byte(userID)[:idLength]
	hash.Write(binaryToken)
	s := hash.Sum(nil)

	return userID[idLength:] == base64.StdEncoding.EncodeToString(s)
}

func GenerateUserID(secret string) string {
	token := urlutils.CreateRandomToken(idLength)
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(token))
	return fmt.Sprintf("%v%v", token, base64.StdEncoding.EncodeToString(hash.Sum(nil)))
}

func GetUserID(ctx echo.Context) string {
	return ctx.Get(UserIDCookie).(string)
}
