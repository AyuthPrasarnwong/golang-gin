package jwt

import (

	//"encoding/json"
	"fmt"
	"io/ioutil"
	//"reflect"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	//"github.com/dgrijalva/jwt-go"
	jwt "github.com/dgrijalva/jwt-go"

	//"github.com/EDDYCJY/go-gin-example/pkg/e"
	//"github.com/EDDYCJY/go-gin-example/pkg/util"
)

const (
	pubKeyPath = "storage/oauth-public.key" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)


// JWT is jwt middlewares
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {


		authString := c.GetHeader("Authorization")

		kv := strings.Split(authString, " ")

		if len(kv) != 2 || kv[0] != "Bearer" {

			c.JSON(http.StatusUnauthorized, gin.H{
				"StatusCode": 401,
				"Message":  "Token not provided.",
			})

			c.Abort()
			return
		}

		tokenString := kv[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 必要的验证 RS256
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {

				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			verifyBytes, err := ioutil.ReadFile(pubKeyPath)

			if err != nil {
				return nil, err
				//log.Fatal(err)
			}

			verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)

			if err != nil {

				return nil, err
			}

			return verifyKey, nil
		})

		if err != nil {

			c.JSON(http.StatusUnauthorized, gin.H{
				"StatusCode": 401,
				"Message":   err.Error(),
			})

			c.Abort()
			return

		}

		if token != nil {
			if token.Valid {

				claims, _ := token.Claims.(jwt.MapClaims)
				//var user string = claims["username"].(string)

				//fmt.Println("TypeOf", reflect.TypeOf(claims["permissions"]))
				fmt.Println("permissions", claims["permissions"])

				c.Next()

			} else if ve, ok := err.(*jwt.ValidationError); ok {

				if ve.Errors&jwt.ValidationErrorMalformed != 0 {

					fmt.Println("Token is invalid.")

				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {

					fmt.Println("Provided token is expired.")

				} else {
					//"Couldn't handle this token:"

				}
			} else {
				//"Couldn't handle this token:"

			}
		}

		c.Next()
	}
}
