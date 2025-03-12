package main

import (
    "log"
    "netmonitor/internal/delivery/http"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Файл .env не найден!")
    }

    router := gin.Default()
    http.RegisterRoutes(router)

    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Ошибка запуска сервера: %v", err)
    }
}
