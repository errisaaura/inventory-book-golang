package auth

import (
	"inventory-book-project/models"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

func HomeHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/login")
}

func LoginGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"content": "",
	})
}

func LoginPostHandler(c *gin.Context) {
	var credential models.Login
	err := c.Bind(&credential)
	if err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Uername/Password is invalid request"})
	}

	if credential.Username != models.USER || credential.Password != models.PASSWORD {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"content": "Uername/Password is invalid"})
	} else {
		// token
		claim := jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
			Issuer:    "test",
			IssuedAt:  time.Now().Unix(),
		}

		sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
		token, err := sign.SignedString([]byte(models.SECRET))
		if err != nil {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{"content": "invalid login"})
			c.Abort()
		}

		// fmt.Println(token)
		q := url.Values{}
		q.Set("auth", token)
		location := url.URL{Path: "/books", RawQuery: q.Encode()}
		c.Redirect(http.StatusMovedPermanently, location.RequestURI())

	}
}
