package server

import (
	"github.com/labstack/echo"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"github.com/Soontao/go-simple-api-gateway/enforcer"
	"github.com/casbin/casbin"
	"github.com/labstack/echo/middleware"
	"net/url"
	"github.com/labstack/gommon/log"
	casbinMw "github.com/labstack/echo-contrib/casbin"
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
	s = &AuthServer{echo.New(), enforcer.SingletonCasbinEnforcer(connStr), resourceHost}
	// Basic Auth Here
	s.Use(casbinMw.Middleware(s.Enforcer))
	s.mountEndPoints();
	s.mountReverseProxy();
	s.HideBanner = true
	s.LoadPolicy()
	return
}

func (s *AuthServer) mountReverseProxy() {
	s.Group("/").Use(middleware.Proxy(&middleware.RoundRobinBalancer{
		Targets: []*middleware.ProxyTarget{
			&middleware.ProxyTarget{
				URL: s.resourceHost,
			},
		},
	}))
}

func (s *AuthServer) mountEndPoints() {
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

func (s *AuthServer) enforceAuth(c echo.Context) (err error) {
	p := new(Policy)
	if err = c.Bind(p); err != nil {
		return
	}
	passed := s.Enforce(p.User, p.Path, p.Method)
	return c.JSON(http.StatusOK, &SuccessMessage{http.StatusOK, passed})
}

func (s *AuthServer) addPolicy(c echo.Context) (err error) {
	p := new(Policy)
	if err = c.Bind(p); err != nil {
		return
	}
	success := s.AddPolicy(p.User, p.Path, p.Method)
	return c.JSON(http.StatusOK, &SuccessMessage{http.StatusOK, success})
}

func (s *AuthServer) getPolicies(c echo.Context) (err error) {
	data := s.GetPolicy()
	return c.JSON(http.StatusOK, &DataMessage{http.StatusOK, data})
}

func (s *AuthServer) getAllAuthorities(c echo.Context) (err error) {
	data := s.GetAllObjects()
	return c.JSON(http.StatusOK, &DataMessage{http.StatusOK, data})
}

func (s *AuthServer) getAllRoles(c echo.Context) (err error) {
	data := s.GetAllRoles()
	return c.JSON(http.StatusOK, &DataMessage{http.StatusOK, data})
}

func (s *AuthServer) getAllUsers(c echo.Context) (err error) {
	data := s.GetAllSubjects()
	return c.JSON(http.StatusOK, &DataMessage{http.StatusOK, data})
}

func (s *AuthServer) getAllActions(c echo.Context) (err error) {
	data := s.GetAllActions()
	return c.JSON(http.StatusOK, &DataMessage{http.StatusOK, data})
}

func (s *AuthServer) getGroupPolicies(c echo.Context) (err error) {
	data := s.GetGroupingPolicy()
	return c.JSON(http.StatusOK, &DataMessage{http.StatusOK, data})
}

func (s *AuthServer) addRoleToUser(c echo.Context) (err error) {
	ur := new(UserRole)
	if err = c.Bind(ur); err != nil {
		return
	}
	success := s.AddRoleForUser(ur.User, ur.Role)
	return c.JSON(http.StatusOK, &SuccessMessage{http.StatusOK, success})
}

func (s *AuthServer) removeRoleFromUser(c echo.Context) (err error) {
	ur := new(UserRole)
	if err = c.Bind(ur); err != nil {
		return
	}
	s.DeleteRoleForUser(ur.User, ur.Role)
	return c.String(http.StatusOK, "")
}

func (s *AuthServer) getUserRoles(c echo.Context) (err error) {
	ur := new(UserRole)
	if err = c.Bind(ur); err != nil {
		return
	}
	data := s.GetRolesForUser(ur.User)
	return c.JSON(http.StatusOK, &DataMessage{http.StatusOK, data})
}

func (s *AuthServer) getRoleUsers(c echo.Context) (err error) {
	ur := new(UserRole)
	if err = c.Bind(ur); err != nil {
		return
	}
	data := s.GetUsersForRole(ur.Role)
	return c.JSON(http.StatusOK, &DataMessage{http.StatusOK, data})
}

func (s *AuthServer) removePolicy(c echo.Context) (err error) {
	p := new(Policy)
	if err = c.Bind(p); err != nil {
		return
	}
	success := s.RemovePolicy(p.User, p.Path, p.Method)
	return c.JSON(http.StatusOK, &SuccessMessage{http.StatusOK, success})
}
