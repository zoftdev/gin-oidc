package main

import (
	"ginoidc/core"
	"net/url"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
)

func errorHandler(c *gin.Context) {
	c.String(500, "ERROR...")
}

func f(c *gin.Context) {
	c.JSON(201, gin.H{
		"myname": "gin1",
	})
}

func main() {

	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// store := sessions.NewCookieStore([]byte("secret"))
	// r.Use(sessions.Sessions("mysession", store))

	iss, _ := url.Parse("http://159.138.240.104:8080/auth/realms/master")
	backendUrl, _ := url.Parse("http://localhost:8080/")
	logout, _ := url.Parse("http://localhost:8080/hi")

	param := core.InitParams{
		Router:        r,
		ClientId:      "vroom",
		ClientSecret:  "4ae1e487-462a-4c46-b117-a9759ffd2399",
		Issuer:        *iss,
		ClientUrl:     *backendUrl,
		Scopes:        []string{"openid"},
		ErrorHandler:  errorHandler,
		PostLogoutUrl: *logout,
	}

	r.Use(core.Init(param))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hi", f)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
