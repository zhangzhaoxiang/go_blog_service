package model

/*
@Time    : 2021/4/15 7:43 上午
@Author  : Ziks
@Email   : zhangzhaoxiang7@163.com
@File    : model.go
@Software: GoLand
*/

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"create_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}
