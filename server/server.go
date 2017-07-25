package server

import (
	"github.com/labstack/echo"
	_ "github.com/go-sql-driver/mysql"
	"github.com/casbin/casbin"
	"github.com/casbin/xorm-adapter"
)

type GateWayServer struct {
	e *echo.Echo
	c *casbin.Enforcer
}

func NewGateWayServer(mysqlConnStr string) (s *GateWayServer) {
	a := xormadapter.NewAdapter("mysql", mysqlConnStr)
	s = &GateWayServer{e: echo.New(), c: casbin.NewEnforcer("cabin.conf", a)}
	s.mountEndPointsTo(s.e)
	return
}

func (gate *GateWayServer) mountEndPointsTo(e *echo.Echo) {
	e.GET("/auth", gate.findAuth)
	e.PUT("/auth", gate.addAuth)
	e.DELETE("/auth", gate.delAuth)
}

func (gate *GateWayServer) findAuth(c echo.Context) (err error) { return }

func (gate *GateWayServer) addAuth(c echo.Context) (err error) { return }

func (gate *GateWayServer) delAuth(c echo.Context) (err error) { return }
