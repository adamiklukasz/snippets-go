package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	g := gin.Default()
	g.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// http://127.0.0.1:8080/user/lukasz/a/b/c -> name=lukasz, rest=/a/b/c
	g.GET("/user/:name/*rest", func(c *gin.Context) {
		name := c.Param("name")
		rest := c.Param("rest")
		q1 := c.DefaultQuery("set", "15")
		q2 := c.Query("get")

		fmt.Printf("%s %s\n", q1, q2)
		fmt.Printf("%s\n", rest)
		fmt.Printf("%s\n", c.FullPath()) // route definition, not URL
		c.String(200, "Hello %s", name)
	})

	// curl -X POST 127.0.0.1:8080/post -F name=Lukasz -F msg=hello -F msg=world
	g.POST("/post", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		msgs := ctx.PostFormArray("msg")

		fmt.Printf("name=%#v\n", name)
		fmt.Printf("msg=%#v\n", msgs)
	})

	g.POST("/file", func(ctx *gin.Context) {
		file, _ := ctx.FormFile("file")

		fmt.Printf("%s\n", file.Filename)
		f, _ := file.Open()
		b, _ := ioutil.ReadAll(f)
		fmt.Printf("b=%s\n", string(b))

		ctx.String(200, "%s uploaded", file.Filename)
	})

	_ = g.Run("127.0.0.1:8080")
}
