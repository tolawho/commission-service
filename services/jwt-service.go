package services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

// JWTService is a contract of what jwtService can do
type JWTService interface {
	GenerateToken(userID uint) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	*jwt.RegisteredClaims
	UserID uint `json:"user_id"`
}

type jwtService struct {
	secretKey string
	issuer    string
}

// NewJWTService method is creates a new instance of JWTService
func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	return os.Getenv("JWT_SECRET")
}

func (j *jwtService) GenerateToken(UserID uint) string {
	// Get the token instance with the Signing method
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	// Choose an expiration time. Shorter the better
	exp := time.Now().Add(time.Hour * 24)
	// Add your claims
	token.Claims = &jwtCustomClaim{
		&jwt.RegisteredClaims{
			// Set the exp and claims. sub is usually the userID
			ExpiresAt: jwt.NewNumericDate(exp),
		},
		UserID,
	}
	// Sign the token with your secret key
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
	} else {
		fmt.Println(err)
	}
	return token, err
}
