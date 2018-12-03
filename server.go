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
    
    // Base URL
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "Welcome to the DoaSanca RESTful API!",
                                    "location endpoint": "/locais",
                                    "users endpoint": "/usuarios"})
	})
    // Management routes
    debug := router.Group("/db") 
    {
        debug.GET("/status", dbStatusHandler)
        debug.GET("/setup", dbSetupHandler)
    }
    // API routes
    router.GET("/locais", getLocaisHandler)
    router.POST("/locais", setLocalHandler)
    router.GET("/locais/:name", getLocalHandler)
    router.GET("/usuarios", getUsuariosHandler)
    router.GET("/usuarios/:email", getUsuarioHandler)
    
    router.Run(":" + port)
}
