package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func ValidateSession(c *gin.Context) bool {
	cookie, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "session expired, please login again"})
			return false
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while getting cookie"})
		return false
	}

	token, err := ValidateJWT(cookie)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, signature invalid"})
			return false
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while validating token"})
		return false
	}

	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, invalid token"})
		return false
	}
	return true
}

func GenerateJWT(userid string) (string, error, time.Time) {
	// Increase token expiration time to 1 hour
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &Claims{
		Username: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
			Issuer:    "tasky",
			Audience:  "tasky-users",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SECRET_KEY))

	return tokenString, err, expirationTime
}

func ValidateJWT(token string) (*jwt.Token, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(SECRET_KEY), nil
	})
	return tkn, err
}

func RefreshToken(c *gin.Context) (bool, error, time.Time) {
	token, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			return true, nil, time.Time{}
		}
		return true, err, time.Time{}
	}

	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return true, nil, time.Time{}
		}
		return false, err, time.Time{}
	}
	if !tkn.Valid || time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		return true, nil, time.Unix(claims.ExpiresAt, 0)
	}
	return false, nil, time.Unix(claims.ExpiresAt, 0)
}
