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
	d := []byte(userID)
	hash.Write(d[:idLength])

	return hmac.Equal([]byte(userID[idLength:]), hash.Sum(nil))
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
