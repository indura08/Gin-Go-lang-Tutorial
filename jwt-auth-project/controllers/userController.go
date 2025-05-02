package controllers

import (
	"jwt-auth/initializers"
	"jwt-auth/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	//get the email/pass off req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	//hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to hash password",
		})

		return
	}

	//create user
	newUser := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&newUser)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error happned while creating the user, check again",
		})
	}

	//respond
	c.JSON(http.StatusOK, gin.H{})

}

func Login(c *gin.Context) {
	//get the email and password form the re body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// Look up reqested user
	var user models.User
	//initializers.DB.Where("email = ?", body.Email).First(&user)
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password , please check again",
		})

		return
	}

	//compare sent in password with user password hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
	}

	//generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
		"email": user.Email,
	})

	//sign and get the completed encoded token
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error happened while generating the token",
		})

		return
	}

	//send it to the client
	//or set the token into a cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*24, "", "", false, true) //me method eke parameters tika blgnna mouse hover krla method ek gwat , ehmai tutorial eke uneth

	c.JSON(200, gin.H{
		"token": tokenString,
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("User")
	c.JSON(http.StatusOK, gin.H{
		"message": &user,
	})
}
