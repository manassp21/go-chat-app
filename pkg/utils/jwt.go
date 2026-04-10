package utils

import(
	"errors"
	"time"
	"go-chat-app/pkg/config"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct{
	UserID int `json:"userid"`
	UserName string `json:"username"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(userid int, username , email string) (string, error) {
	claims:=&Claims{
		UserID : userid,
		UserName : username,
		Email : email,
		RegisteredClaims : jwt.RegisteredClaims{
			ExpiresAt : jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(config.AppConfig.JWTExpirationHours))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.AppConfig.JWTSecret))

	if err!=nil{
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*Claims, error){
	claims:=&Claims{}

	token, err:=jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error){
		if _,ok:=token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.AppConfig.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}