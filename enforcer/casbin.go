package enforcer

import (
	"github.com/casbin/casbin"
	"github.com/casbin/xorm-adapter"
)

func NewCasbinEnforcer(connStr string) *casbin.Enforcer {
	Adapter := xormadapter.NewAdapter("mysql", connStr, true)
	enforcer := casbin.NewEnforcer(casbin.NewModel(CasbinConf), Adapter)
	return enforcer
}
