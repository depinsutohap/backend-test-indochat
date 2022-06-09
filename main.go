package main

import (
	"log"
	"net/http"
	"os"
	"fmt"
	
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
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
		yourHtmlString := ""
		number := 99
		for i := 1; i < 11; i++ {
			yourHtmlString += fmt.Sprintf("%d * %d = %d\n", i, number, number * i)
		  }

		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write([]byte(yourHtmlString))
	})

	router.Run(":" + port)
}
