package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	func helloFunc(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	}
*/
func Pinghandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
func hellohandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hi back to you",
	})
}

func main() {
	/*
		http.HandleFunc("/", helloFunc)
		log.Fatal(http.ListenAndServe(":8080", nil))
	*/
	application := gin.Default()
	application.GET("/ping", Pinghandler)
	application.GET("/hello", hellohandler)

	application.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
