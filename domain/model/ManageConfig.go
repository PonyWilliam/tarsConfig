package model

import "gorm.io/plugin/soft_delete"
type RequireConfig struct{
	ID int64 `gorm:"primary_key;auto_increment;not_null" json:"id"`
	Name string `json:"name"` //配置名称
	RequireLevel int64 `json:"requirelevel"` //要求用户等级
	RequireWork int64 `json:"requirework"` //1->学生、在业 2->在业 3->均可
	RequireAge int64 `json:"requireage"` //对年龄是否有要求
	RequirePromise bool `json:"requirepromise"` //对失信是否有要求

	Created  int64 `gorm:"autoCreateTime"`//创建时间
	Updated  int64 `gorm:"autoUpdateTime"`//修改时间
	DeletedAt soft_delete.DeletedAt//删除时间，数据库存档一段时间，可以做个定时脚本定时删除
	
}
