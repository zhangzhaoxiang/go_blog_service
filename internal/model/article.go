package model

/*
@Time    : 2021/4/15 8:19 上午
@Author  : Ziks
@Email   : zhangzhaoxiang7@163.com
@File    : article.go
@Software: GoLand
*/

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}
