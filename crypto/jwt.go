package crypto

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func NewJwtToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userId,
		"nbf": time.Now().Add(time.Second * 3600).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	secretKey := os.Getenv("SECRET_KEY")
	tokenString, err := token.SignedString(secretKey)
	return tokenString, err
}

func ParseValidJwtToken(tokenString string) error {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	secretKey := os.Getenv("SECRET_KEY")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Println(claims["id"], claims["nbf"])
	} else {
		log.Println(err)
	}
	return err
}
