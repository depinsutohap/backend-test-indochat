package main

import (
	"log"
	"net/http"
	"os"
	
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/russross/blackfriday"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		 yourHtmlString := "<html><body>I am cached HTML!<br>Hi<br></body></html>"

		//Write your 200 header status (or other status codes, but only WriteHeader once)
		c.Writer.WriteHeader(http.StatusOK)
		//Convert your cached html string to byte array
		c.Writer.Write([]byte(yourHtmlString))
	})

	router.Run(":" + port)
}
