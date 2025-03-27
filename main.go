package main

import (
  "log"
  "net/http"
  "os"
  "github.com/gin-gonic/gin"
)

func main() {
  environment := os.Getenv("ENV")
  if environment == "PRODUCTION" {
    gin.SetMode(gin.ReleaseMode)
  }

  router := gin.Default()
  router.GET("/", func(ctx *gin.Context) {
   ctx.JSON(http.StatusOK, gin.H{"message": "Hello, World"})
  })


  if err := router.Run(":8080"); err != nil {
   log.Fatal("Server failed: ", err)
  }
}
