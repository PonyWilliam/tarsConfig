package main

import (
	"fmt"
	"log"

	"github.com/TarsCloud/TarsGo/tars"

	"github.com/PonyWilliam/tarsConfig/domain/model"
	"github.com/PonyWilliam/tarsConfig/tars-protocol/SanXiangBank"
)

func main() {
	comm := tars.NewCommunicator()
	obj := fmt.Sprintf("SanXiangBank.Config.ManagerConfigObj@tcp -h 127.0.0.1 -p 10015 -t 60000")
	app := new(SanXiangBank.ManagerConfig)
	comm.StringToProxy(obj, app)
	//test 区域
	cfg := model.RequireConfig{}
	cfg.Name = "默认规则"
	cfg.RequireAge = 18
	cfg.RequireLevel = 3
	cfg.RequirePromise = true
	cfg.RequireWork = 2
	res := ModelToImp(cfg)
	ret,err := app.CreateConfig(&res)
	fmt.Println(ret)
	if err != nil{
		fmt.Print(111)
		log.Fatal(err)
	}
}

func ImpToModel(cfg SanXiangBank.Config)model.RequireConfig{
	mcfg := model.RequireConfig{}
	mcfg.Name = cfg.Name
	mcfg.RequireAge = int64(cfg.Rage)
	mcfg.RequireLevel = int64(cfg.Rlevel)
	mcfg.RequireWork = int64(cfg.Rwork)
	mcfg.RequirePromise = cfg.Rpromise
	return mcfg
}
func ModelToImp(mcfg model.RequireConfig)SanXiangBank.Config{
	cfg := SanXiangBank.Config{}
	cfg.Name = mcfg.Name
	cfg.Rage = int32(mcfg.RequireAge)
	cfg.Rlevel = int32(mcfg.RequireLevel)
	cfg.Rwork = int32(mcfg.RequireWork)
	cfg.Rpromise = mcfg.RequirePromise
	return cfg
}