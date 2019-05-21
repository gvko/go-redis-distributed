package main

import (
  "github.com/gin-gonic/gin"
  "os"
)

func main() {
  os.Setenv("PORT", "8080")
  server := gin.Default()

  server.GET("/ping", func(ctx *gin.Context) {
    ctx.JSON(200, gin.H{
      "message": "pong",
    })
  })

  server.Run() // listen and serve on localhost:PORT (env var PORT)
}
