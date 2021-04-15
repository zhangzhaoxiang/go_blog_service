package model

/*
@Time    : 2021/4/15 7:49 上午
@Author  : Ziks
@Email   : zhangzhaoxiang7@163.com
@File    : tag.go
@Software: GoLand
*/

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}
