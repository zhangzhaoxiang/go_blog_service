package model

/*
@Time    : 2021/4/15 8:23 上午
@Author  : Ziks
@Email   : zhangzhaoxiang7@163.com
@File    : article_tag.go
@Software: GoLand
*/

type ArticleTag struct {
	*Model
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
