package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func dbStatusHandler(c *gin.Context) {
    err := pingDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    } else {
        c.JSON(http.StatusOK, gin.H{"status": "DB connection is fine."})
    }
}

func dbSetupHandler(c *gin.Context) {
    err := doSetupDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    } else {
        c.JSON(http.StatusOK, gin.H{"status": "DB has been setup correctly."})
    }
}

func getLocaisHandler(c *gin.Context) {
    response := getLocations()
    if len(response) == 0 {
        c.JSON(http.StatusNoContent, "No results were found.")
    } else {
        c.JSON(http.StatusOK, response)
    }
}

func setLocalHandler(c *gin.Context) {
    var request Location
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    } else {
        err = saveNewLocation(request)
        if err != nil {
            c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
        }
        c.JSON(http.StatusOK, gin.H{"status": "Request was recorded successfully!"})
    }
}

func getLocalHandler (c *gin.Context) {
     name := c.Param("name")
     response := getLocationByName(name)
     if response.Name == "" {
         c.JSON(http.StatusNoContent, gin.H{"error": "No results were found."})
     } else {
         c.JSON(http.StatusOK, response)
     }
}

func getUsuariosHandler (c *gin.Context) {
    response := getUsers()
    if len(response) == 0 {
        c.JSON(http.StatusNoContent, "No results were found.")
    } else {
        c.JSON(http.StatusOK, response)
    }
}
