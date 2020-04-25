package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// creates a single signed JWT token
func CreateToken(id string, exp time.Time, secret string) (string, error) {

	type Claims struct {
		ID string `json:"id"`
		jwt.StandardClaims
	}

	claims := &Claims{
		id,
		jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}

	// create token
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign token
	return t.SignedString([]byte(secret))
}

// generates an access and refresh token pair
func GetTokenPair(id string, secret string) (*TokenPair, error) {
	aexp := time.Now().Add(time.Minute * 15)
	rexp := time.Now().Add(time.Hour * 1)
	at, err := CreateToken(id, aexp, secret)
	if err != nil {
		return nil, err
	}
	rt, err := CreateToken(id, rexp, secret)
	if err != nil {
		return nil, err
	}
	res := TokenPair{
		AccessToken:  at,
		RefreshToken: rt,
	}
	return &res, nil
}

func ValidateToken(token string, secret string) (jwt.MapClaims, error) {

	type Claims struct {
		ID string `json:"id"`
		jwt.StandardClaims
	}

	// parse
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(tt *jwt.Token) (interface{}, error) {
		// check signing method
		if _, ok := tt.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	// get claims from token
	if claims, ok := t.Claims.(*Claims); ok && t.Valid {
		fmt.Printf("Claims found, ID: %s", claims.ID)
		return nil, nil
	} else {
		return nil, errors.New("Invalid token")
	}
}
