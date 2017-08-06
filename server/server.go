package server

import (
	"github.com/labstack/echo"
	_ "github.com/go-sql-driver/mysql"
	"github.com/Soontao/go-simple-api-gateway/enforcer"
	"github.com/casbin/casbin"
	"github.com/labstack/echo/middleware"
	"net/url"
	"github.com/labstack/gommon/log"
	casbinMw "github.com/Soontao/echo-contrib/casbin"
)

type AuthServer struct {
	*echo.Echo
	*casbin.Enforcer
	resourceHost *url.URL
}

func NewAuthServer(connStr string, resourceHostStr string) (s *AuthServer) {
	resourceHost, err := url.Parse(resourceHostStr)
	if err != nil {
		log.Fatal(err)
	}
	s = &AuthServer{echo.New(), enforcer.NewCasbinEnforcer(connStr), resourceHost}
	s.Use(NewCoockieSession())
	s.mountAuthenticateEndpoints()
	s.mountAuthorizationEndPoints()
	s.mountReverseProxy()
	s.HideBanner = true
	s.LoadPolicy()
	return
}

func (s *AuthServer) mountReverseProxy() {
	s.Group("/").Use(casbinMw.Middleware(s.Enforcer), middleware.Proxy(&middleware.RoundRobinBalancer{
		Targets: []*middleware.ProxyTarget{
			&middleware.ProxyTarget{
				URL: s.resourceHost,
			},
		},
	}))
}

func (s *AuthServer) mountAuthenticateEndpoints() {
	api := s.Group("/_/auth/api")
	api.Any("/auth", s.authAPI)
}

func (s *AuthServer) mountAuthorizationEndPoints() {
	api := s.Group("/_/gateway/api")
	policy := api.Group("/policy")
	policy.GET("/", s.getPolicies).Name = "Get All Policies"
	policy.GET("/group", s.getGroupPolicies).Name = "Get Group Policies"
	policy.GET("/authorities", s.getAllAuthorities)
	policy.GET("/methods", s.getAllActions)
	policy.POST("/enforce", s.enforceAuth).Name = "Find Some Authority"
	policy.PUT("/", s.addPolicy).Name = "Add Policy"
	policy.DELETE("/", s.removePolicy).Name = "Remove Authority"
	role := api.Group("/role")
	role.GET("/", s.getAllRoles)
	role.PUT("/", s.addRoleToUser).Name = "Add Role To User"
	role.DELETE("/", s.removeRoleFromUser).Name = "Remove Role From User"
	role.GET("/users", s.getRoleUsers).Name = "Get Users of a Role"
	user := api.Group("/user")
	user.GET("/", s.getAllUsers)
	user.GET("/role", s.getUserRoles).Name = "Get Roles of a User"
}
