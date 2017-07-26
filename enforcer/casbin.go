package enforcer

import (
	"github.com/casbin/casbin"
	"github.com/casbin/xorm-adapter"
)

var enforcer *casbin.Enforcer;

func SingletonCasbinEnforcer(connStr string) (*casbin.Enforcer) {
	if (enforcer == nil) {
		Adapter := xormadapter.NewAdapter("mysql", connStr)
		enforcer = casbin.NewEnforcer("enforcer/casbin.conf", Adapter)
	}
	return enforcer
}
