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
		c.String(http.StatusOK, "Hello user!")
	})

    router.GET("/locais/:name", func(c *gin.Context) {
        name := c.Param("name")
		c.String(http.StatusOK, "Página do local %s!", name)
	})
   
    db, err := setupDB()
    router.GET("/db/status", func(c *gin.Context) {
        err = pingDB(db)
        
        if err != nil {
            c.String(http.StatusInternalServerError, "%q", err)
        } else {
            c.String(http.StatusOK, "DB connection is fine.")
        }
    })
    
    router.Run(":" + port)
}
