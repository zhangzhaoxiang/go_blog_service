package v1

import "github.com/gin-gonic/gin"

/*
@Time    : 2021/4/15 8:34 上午
@Author  : Ziks
@Email   : zhangzhaoxiang7@163.com
@File    : tag.go
@Software: GoLand
*/

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context)    {}
func (t Tag) List(c *gin.Context)   {}
func (t Tag) Create(c *gin.Context) {}
func (t Tag) Update(c *gin.Context) {}
func (t Tag) Delete(c *gin.Context) {}
