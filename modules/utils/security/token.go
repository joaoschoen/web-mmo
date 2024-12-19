package security

import (
	"os"
	"time"

	"web-mmo/modules/api/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type CustomClaims struct {
	SessionID string `json:"session_id"`
	jwt.RegisteredClaims
}

func GenerateToken(user model.UserData, session_id string) (string, error) {
	//TODO
	// currently using secret, change to RSA key at least
	uuid := uuid.New().String()

	TOKEN_SECRET := os.Getenv("TOKEN_SECRET")
	println(uuid)

	// Token Claims
	claims := CustomClaims{
		SessionID: uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Token expiration
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(TOKEN_SECRET))
	println(signedToken)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

//TODO
// func ValidateToken(token string){
// 	jwt.NewValidator()
// }
