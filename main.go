package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TarsCloud/TarsGo/tars"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/PonyWilliam/tarsConfig/domain/model"
	"github.com/PonyWilliam/tarsConfig/domain/repository"
	"github.com/PonyWilliam/tarsConfig/tars-protocol/SanXiangBank"
)
var rp repository.IMangeConfig
func init(){
	db_data := model.GetMysqlConfig()
	//初始化数据库连接，并对数据库表进行检查
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",db_data.User,db_data.Pwd,db_data.Host,db_data.Port,db_data.Dbname)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}
	//传递db给User Repository进行维护
	rp = repository.InitConfig(db)
	_ = rp.Init()
}
func main() {
	// Get server config
	cfg := tars.GetServerConfig()

	// New servant imp
	imp := new(ManagerConfigImp)
	err := imp.Init()
	if err != nil {
		fmt.Printf("ManagerConfigImp init fail, err:(%s)\n", err)
		os.Exit(-1)
	}
	// New servant
	app := new(SanXiangBank.ManagerConfig)
	// Register Servant
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".ManagerConfigObj")

	// Run application
	tars.Run()
}
