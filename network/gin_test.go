package network

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"testing"
)

func TestGin(t *testing.T) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/user/:name/*rest", func(c *gin.Context) {
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
	r.POST("/post", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		msgs := ctx.PostFormArray("msg")

		fmt.Printf("name=%#v\n", name)
		fmt.Printf("msg=%#v\n", msgs)
	})

	// curl -X POST 127.0.0.1:8080/file -F "file=@testfile.txt" -H "Content-Type: multipart/form-data"
	r.POST("/file", func(ctx *gin.Context) {
		file, _ := ctx.FormFile("file")

		fmt.Printf("%s\n", file.Filename)
		f, _ := file.Open()
		b, _ := ioutil.ReadAll(f)
		fmt.Printf("b=%s\n", string(b))

		ctx.String(200, "%s uploaded", file.Filename)
	})

	_ = r.Run()

	//gin.Logger
	//gin.Recovery
	//gin.LoggerWithFormatter
}

type Login struct {
	Name     string `form:"name"`
	Password string `form:"password"`
	Age      int    `form:"age"`
}

func TestGinBinding(t *testing.T) {
	r := gin.Default()

	//curl -X POST 127.0.0.1:8080/add?name=Lukasz\&age=34
	r.POST("/add", func(c *gin.Context) {
		login := Login{
			Age: 24,
		}

		err := c.ShouldBindQuery(&login)
		fmt.Printf("login=%#v\n", login)
		fmt.Printf("err=%#v\n", err)

		c.String(200, "OK")
	})

	r.Run()
}
