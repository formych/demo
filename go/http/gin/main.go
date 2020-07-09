package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/", Index)
	engine.GET("/who", Who)
	engine.GET("/who/:id", ID)

	if err := engine.Run(":80"); err != nil {
		log.Fatalf("server stoped, err: %s", err)
	}
}

const (
	defaultID   = "2011118120"
	defaultName = "formych"
)

// Index 首页
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"name": defaultName,
		"text": "Hello, World!",
		"time": time.Now().Format("2006-01-02 15:04:05"),
	})
}

// Who 名字
func Who(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		name = defaultName
	}
	c.JSON(http.StatusOK, map[string]string{
		"name": name,
		"text": "Hello, Wordl!",
		"time": time.Now().Format("2006-01-02 15:04:05"),
	})
}

// ID ID
func ID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		id = defaultID
	}
	c.JSON(http.StatusOK, map[string]string{
		"id":   id,
		"text": "Hello, Wordl!",
		"time": time.Now().Format("2006-01-02 15:04:05"),
	})
}
