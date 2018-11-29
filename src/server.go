package main

import (
    "os"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    port := os.Getenv("PORT")
    
    if port == "" {
        log.Fatal("$PORT must be set")
    }

    router := gin.New()
    router.Use(gin.Logger())
    
    router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello user!", nil)
	})

    router.GET("/:name", func(c *gin.Context) {
        name := c.Param("name")
		c.String(http.StatusOK, "Hello %s!", name)
	})

    router.Run(":" + port)
}
