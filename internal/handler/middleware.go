package handler

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	keyUserID  = "user_id"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
)

type TokenClaims struct {
	jwt.StandardClaims
	UId          string `json:"uid"`
	UserFullName string `json:"user_full_name"`
}

func ParseToken(accessToken string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token isn't valid")
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return nil, errors.New("token claims are not type of *tokenClaims")
	}

	return claims, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		claim, err := ParseToken(tokenString[7:])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		userID, err := strconv.Atoi(claim.UId)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization user id is not correct "})
			c.Abort()
			return
		}

		c.Set(keyUserID, userID)
		c.Next()
	}
}
