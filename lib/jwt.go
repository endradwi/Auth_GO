package lib

import (
	"time"

	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
)

var JWT_SECRET []byte = []byte(GetMD5hash("SECRET_KEY"))

func GeneratedToken(payload any) string {

	sig, _ := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: JWT_SECRET}, (&jose.SignerOptions{}).WithType("JWT"))
	baseinfo := jwt.Claims{
		IssuedAt: jwt.NewNumericDate(time.Now()),
	}

	token, _ := jwt.Signed(sig).Claims(baseinfo).Claims(payload).Serialize()
	return token
}
