package main

import (
	"fmt"

	"github.com/ahasunos/lingolab/go/ginandgorm/controllers"
	"github.com/ahasunos/lingolab/go/ginandgorm/initializers"
	"github.com/ahasunos/lingolab/go/ginandgorm/models"
	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Homepage live: Where APIs flow smoother than Michael's awkward office moments!",
	})
}

func hostAndServe() {

	// func gin.Default() *gin.Engine
	// Default returns an Engine instance with the Logger and Recovery middleware already attached
	r := gin.Default()

	// func (*gin.RouterGroup).GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
	// GET is a shortcut for router.Handle("GET", path, handlers).
	r.GET("/", home) // This could be specified as: r.Handle("GET", "/", home)
	r.GET("/getCharacters", controllers.GetCharacters)
	r.POST("/addCharacters", controllers.AddCharacters)

	// func (*gin.Engine).Run(addr ...string) (err error)
	// Run attaches the router to a http.Server and starts listening and serving HTTP requests.
	// It is a shortcut for http.ListenAndServe(addr, router)
	// Note: this method will block the calling goroutine indefinitely unless an error happens.
	r.Run() // listen and serve on 0.0.0.0:8080, since address is not specified!
}

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDB()
}

func main() {
	// Migrate the schema
	initializers.DB.AutoMigrate(&models.Character{})
	fmt.Println("Starting server...")
	hostAndServe()
}
