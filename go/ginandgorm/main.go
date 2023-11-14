package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func home(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "ping",
		"greet":   "hello",
	})
}

func requestHandler() {

	// func gin.Default() *gin.Engine
	// Default returns an Engine instance with the Logger and Recovery middleware already attached
	r := gin.Default()

	// func (*gin.RouterGroup).GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
	// GET is a shortcut for router.Handle("GET", path, handlers).
	r.GET("/", home) // This could be specified as: r.Handle("GET", "/", home)

	// func (*gin.Engine).Run(addr ...string) (err error)
	// Run attaches the router to a http.Server and starts listening and serving HTTP requests.
	// It is a shortcut for http.ListenAndServe(addr, router)
	// Note: this method will block the calling goroutine indefinitely unless an error happens.
	r.Run() // listen and serve on 0.0.0.0:8080, since address is not specified!
}

func main() {
	fmt.Println("Starting server...")
	requestHandler()
}
