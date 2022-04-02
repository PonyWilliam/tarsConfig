package repository

import (
	"github.com/PonyWilliam/tarsConfig/domain/model"
	"gorm.io/gorm"
)

var db *gorm.DB

func init(){

}

type IMangeConfig interface{
	Init() error
	CreateConfig(config *model.RequireConfig)error
	FindAllConfig()(error,[]model.RequireConfig)
	FindConfigByID(id int64)(error,model.RequireConfig)
	UpdateConfig(cfg model.RequireConfig)(error)
	DelConfigByID(id int64)error
}

type ManageConfig struct{
	Db *gorm.DB
}
func InitConfig(db *gorm.DB)IMangeConfig{
	u := &ManageConfig{Db: db}
	return u
}
func(m ManageConfig)Init()error{
	if(!m.Db.Migrator().HasTable(&model.RequireConfig{})){
		m.Db.Migrator().CreateTable(&model.RequireConfig{})
	}
	return nil
}

func(m ManageConfig)CreateConfig(config *model.RequireConfig)error{
	res := m.Db.Omit("id").Create(config)
	return res.Error
}
func(m ManageConfig)FindAllConfig()(error,[]model.RequireConfig){
	var configs []model.RequireConfig
	res := m.Db.Find(configs)
	return res.Error,configs
}
func(m ManageConfig)FindConfigByID(id int64)(error,model.RequireConfig){
	var cfg model.RequireConfig
	res := m.Db.Where("id = ?",id).First(cfg)
	return res.Error,cfg
}
func(m ManageConfig)UpdateConfig(cfg model.RequireConfig)(error){
	if(cfg.RequirePromise == false){
		m.Db.Model(cfg).Update("requirepromise",false)
	}
	res := m.Db.Model(&cfg).Updates(cfg)
	return res.Error
}
func(m ManageConfig)DelConfigByID(id int64)error{
	res := m.Db.Delete(&model.RequireConfig{},id)
	return res.Error
}
