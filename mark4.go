package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func main() {
// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "Hello",
// 		})
// 	})
// 	r.Run()
// }

// func main() {
// 	r := gin.Default()
// 	r.GET("/getFunctionallity", getFunc)
// 	r.Run()
// }
// func getFunc(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Hello Virtual World",
// 	})
// }

//parameters in path ---> with String
//that is coming with String value
// func main() {
// 	router := gin.Default()
// 	//make one handler will match /user/john but will not match
// 	router.GET("/user/:name", func(c *gin.Context) {
// 		name := c.Param("name")
// 		c.String(http.StatusOK, "Hello %s", name)
// 	})
// 	router.Run(":8080")
// }

// func main() {
// 	gin.SetMode(gin.ReleaseMode)
// 	router := gin.Default()
// 	router.GET("/user/:name/*action", func(c *gin.Context) {
// 		name := c.Param("name")
// 		action := c.Param("action")
// 		message := name + "is" + action
// 		c.String(http.StatusOK, message)
// 	})
// 	router.Run(":8080")
// }

// func main() {
// 	router := gin.Default()
// 	router.POST("/user/:name/*action", func(c *gin.Context) {
// 		b := c.FullPath() == "/user/:name/*action"
// 		c.String(http.StatusOK, "%t", b)
// 	})
// 	router.Run(":8080")
// }

func main() {
	router := gin.Default()
	router.GET("/user/groups", func(c *gin.Context) {
		c.String(http.StatusOK, "The available groups are [...]")
	})
	router.Run(":8080")
}
