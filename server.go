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
        log.Printf("$PORT not set, defaulting PORT to 8080.\n")
        port = "8080"
    }

    router := gin.Default()
    db, err := getDB()

    if err != nil {
        log.Fatalf("Failed to get database connection: %q", err)
    }
    defer db.Close()
    
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "Welcome to the DoaSanca RESTful API!"})
	})

    // Management routes
    router.GET("/db/status", func(c *gin.Context) {
        err = pingDB()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        } else {
            c.JSON(http.StatusOK, gin.H{"status": "DB connection is fine."})
        }
    })
    router.GET("/db/setup", func(c *gin.Context) {
        err = setupDB()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        } else {
            c.JSON(http.StatusOK, gin.H{"status": "DB has been setup correctly."})
        }
    })
    
    // API routes
    router.GET("/locais", func(c *gin.Context) {
        response := getLocations()
        if response.Name == "" {
            c.JSON(http.StatusNoContent, "No results were found.")
        } else {
            c.JSON(http.StatusOK, response)
        }
    })
    router.POST("/locais/", func(c *gin.Context) {
        var request Location
        if err = c.ShouldBindJSON(&request); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        } else {
            c.JSON(http.StatusOK, gin.H{"status": "Request was recorded successfully!"})
        }
    })
    router.GET("/locais/:name", func(c *gin.Context) {
        name := c.Param("name")
        response := getLocationByName(name)
		if response.Name == "" {
            c.JSON(http.StatusNoContent, gin.H{"error": "No results were found."})
        } else {
            c.JSON(http.StatusOK, response)
        }
	})

    router.Run(":" + port)
}
