package controllers

import (
	"github.com/gin-gonic/gin"
)

// BlogStruct struct
type BlogStruct struct {
}

func (*BlogStruct) getName(c *gin.Context) {
	c.JSON(200, gin.H{"Hello": "world"})
}

//Hello test hellow
func (BlogStruct) Hello(c *gin.Context) {
	c.JSON(200, gin.H{"Hello": "world"})
}
