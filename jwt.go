package main

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

const SecretKey = "SecretKey"

func main() {
	router := gin.Default()

	jwt := router.Group("/jwt")
	{
		jwt.GET("/generate", generateJwt)
		jwt.GET("/verify/:token", verifyJwt)
	}
	router.Run(":8080")
}

func generateJwt(c *gin.Context) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "iss": "www.host-that-issue-jwt.com",
        "aud": "tenant/client_id",
        "sub":"user-id",
        "jti":"unique-string",
        "iat": time.Now().Unix(),
        "nbf": time.Now().Add(time.Second * 10).Unix(),
        "exp": time.Now().Add(time.Second * 30).Unix(),
        "author": "Teerapong Chantakard",
    })

    if tokenString, err := token.SignedString([]byte(SecretKey)); err != nil {
		c.String(http.StatusOK, err.Error())
    } else {
    	c.String(http.StatusUnauthorized, tokenString)
	}
}

func verifyJwt(c *gin.Context) {
	myToken := c.Param("token")

	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
        return []byte(SecretKey), nil
    })

    if err == nil && token.Valid {
        c.JSON(http.StatusOK, token)
    } else {
        c.String(http.StatusUnauthorized, err.Error())
    }
}
