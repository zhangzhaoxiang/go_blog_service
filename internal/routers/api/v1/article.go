package v1

import "github.com/gin-gonic/gin"

/*
@Time    : 2021/4/15 8:35 上午
@Author  : Ziks
@Email   : zhangzhaoxiang7@163.com
@File    : article.go
@Software: GoLand
*/

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context)    {}
func (a Article) List(c *gin.Context)   {}
func (a Article) Create(c *gin.Context) {}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}
