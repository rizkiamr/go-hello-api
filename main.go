package main

import (
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
)

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
}

func main() {
	server := gin.Default()
	// This makes it so each ip can only make 5 requests per second
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Minute,
		Limit: 5,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})
	server.GET("/", mw, func(c *gin.Context) {
		c.String(200, "demo jeparadev")
	})
	server.GET("/ping", mw, func(c *gin.Context) {
		c.String(200, "pong")
	})
	server.GET("/hello", mw, func(c *gin.Context) {
		c.String(200, "hello world")
	})
	server.GET("/version", mw, func(c *gin.Context) {
		c.String(200, "1.0")
	})
	server.Run(":8080")
}
