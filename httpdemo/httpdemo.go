package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

const KeyRequestId = "requestId"

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	// 添加middleware
	r.Use(func(c *gin.Context) {
		s := time.Now()
		// path response code, log latency
		logger.Info("incoming request",
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)))
		c.Next()
	}, func(c *gin.Context) {
		c.Set(KeyRequestId, rand.Int())
		c.Next()
	})
	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		if rid, exists := c.Get(KeyRequestId); exists {
			h[KeyRequestId] = rid
		}
		c.JSON(200, h)
	})
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, "hello")
	})
	r.Run()
}
