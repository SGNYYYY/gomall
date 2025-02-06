package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserId int32  `json:"userid"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type JwtResp struct {
	UserId int32  `json:"userid"`
	Role   string `json:"role"`
	Token  string `json:"token"`
}

func GenerateToken(user_id int32, role string) (string, error) {
	claims := &Claims{
		UserId: user_id,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(720 * time.Minute)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                        // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                        // 生效时间
			Issuer:    "gomall-demo",                                         // 签发者
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func VerifyToken(tokenString string) (*JwtResp, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil || !token.Valid {
		return &JwtResp{
			Token: "",
		}, nil
	}
	newToken, err := GenerateToken(claims.UserId, claims.Role)
	if err != nil {
		return nil, err
	}
	resp := &JwtResp{
		UserId: claims.UserId,
		Role:   claims.Role,
		Token:  newToken,
	}
	return resp, nil
}
