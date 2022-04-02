package main

import (
	"context"

	"github.com/PonyWilliam/tarsConfig/domain/model"
	"github.com/PonyWilliam/tarsConfig/tars-protocol/SanXiangBank"
)

// ManagerConfigImp servant implementation
type ManagerConfigImp struct {
}

// Init servant init
func (imp *ManagerConfigImp) Init() error {
	//initialize servant here:
	//...
	return nil
}

// Destroy servant destory
func (imp *ManagerConfigImp) Destroy() {
	//destroy servant here:
	//...
}
func (imp *ManagerConfigImp) CreateConfig(ctx context.Context,cfg *SanXiangBank.Config) (bool, error) {
	temp := ImpToModel(*cfg)
	err := rp.CreateConfig(&temp)
	return err == nil,err
}
func (imp *ManagerConfigImp) FindAllConfig(ctx context.Context,res *[]SanXiangBank.Config) (bool, error) {
	err,cfgs := rp.FindAllConfig()
	var r []SanXiangBank.Config
	for _,v := range cfgs{
		r = append(r, ModelToImp(v))
	}
	res = &r
	return err == nil, err
}
func (imp *ManagerConfigImp) FindConfigByID(ctx context.Context, id int32,result *SanXiangBank.Config) (bool, error) {
	//Doing something in your function
	//...
	err,cfg := rp.FindConfigByID(int64(id))
	if err != nil{
		return false,err
	}
	r := ModelToImp(cfg)
	result = &r
	return err == nil,err
}
func (imp *ManagerConfigImp) UpdateConfig(ctx context.Context,cfg *SanXiangBank.Config) (bool, error) {
	mcfg := ImpToModel(*cfg)
	err := rp.UpdateConfig(mcfg)
	return err == nil,err
}
func (imp *ManagerConfigImp) DelConfigByID(ctx context.Context, id int32) (bool, error) {
	err := rp.DelConfigByID(int64(id))
	return err == nil,err
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