package http

import (
    "github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
    api := r.Group("/api")
    {
        api.GET("/health", healthCheck)
    }
}

func healthCheck(c *gin.Context) {
    c.JSON(200, gin.H{"status": "OK"})
}
