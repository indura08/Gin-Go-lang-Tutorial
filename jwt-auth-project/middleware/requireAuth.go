package middleware

import (
	"fmt"
	"jwt-auth/initializers"
	"jwt-auth/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	fmt.Println("In middleware") //this is for testing purposes

	//get the cookie off the request body
	tokenString, err := c.Cookie("Authorization")
	fmt.Println("TOKEN:", tokenString)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//Decode / validate it

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//fmt.Println(claims["foo"], claims["nbf"])

		//check expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//find the user with token sub
		var user models.User
		initializers.DB.First(&user, claims["sub"]) //methna user kiynne hambena user wa save rnna hdna variable eke location ekai

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//Attach to rq
		c.Set("User", user)

		//Continue
		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
