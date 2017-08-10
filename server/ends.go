package server

import (
	"fmt"
	"net/http"

	"github.com/ipfans/echo-session"
	"github.com/labstack/echo"
)

func (s *GatewayServer) userAuth(c echo.Context) (err error) {
	sess := session.Default(c)
	user := User{}
	c.Bind(&user)
	if s.authUserService.AuthUser(user.Username, user.Password) {
		sess.Set(Username, user.Username)
		sess.Save()
		return c.JSON(http.StatusOK, &DataMessage{http.StatusOK, fmt.Sprintf("auth for %s", user.Username)})
	} else {
		return c.JSON(http.StatusForbidden, &DataMessage{http.StatusForbidden, "auth failed"})
	}
	return
}

func (s *GatewayServer) userRegister(c echo.Context) (err error) {
	user := &User{}
	c.Bind(user)
	if err := s.authUserService.SaveUser(user.Username, user.Password); err == nil {
		s.AddRoleForUser(user.Username, s.DefaultRegisterRole)
		return c.JSON(http.StatusOK, &DataMessage{http.StatusOK, fmt.Sprintf("register for %s", user.Username)})
	} else {
		return c.JSON(http.StatusBadRequest, &DataMessage{http.StatusBadRequest, err.Error()})
	}
	return
}

func (s *GatewayServer) userUpdate(c echo.Context) (err error) {
	user := &User{}
	c.Bind(user)
	if s.authUserService.UpdatePassword(user.Username, user.Password, user.NewPassword) {
		return c.JSON(http.StatusOK, &DataMessage{http.StatusOK, fmt.Sprintf("password updated for %s", user.Username)})
	} else {
		return c.JSON(http.StatusBadRequest, &DataMessage{http.StatusBadRequest, "password update failed"})
	}

	return
}

func (s *GatewayServer) enforceAuth(c echo.Context) (err error) {
	p := new(Policy)
	if err = c.Bind(p); err != nil {
		return
	}
	passed := s.Enforce(p.User, p.Path, p.Method)
	return c.JSON(http.StatusOK, &SuccessMessage{http.StatusOK, passed})
}

func (s *GatewayServer) addPolicy(c echo.Context) (err error) {
	p := new(Policy)
	if err = c.Bind(p); err != nil {
		return
	}
	success := s.AddPolicy(p.User, p.Path, p.Method)
	return c.JSON(http.StatusOK, &SuccessMessage{http.StatusOK, success})
}

func (s *GatewayServer) getPolicies(c echo.Context) (err error) {
	data := s.GetPolicy()
	return c.JSON(http.StatusOK, &DataMessage{http.StatusOK, data})
}

func (s *GatewayServer) getAllAuthorities(c echo.Context) (err error) {
	data := s.GetAllObjects()
	return c.JSON(http.StatusOK, &DataMessage{http.StatusOK, data})
}

func (s *GatewayServer) getAllRoles(c echo.Context) (err error) {
	data := s.GetAllRoles()
	return c.JSON(http.StatusOK, &DataMessage{http.StatusOK, data})
}

func (s *GatewayServer) getAllUsers(c echo.Context) (err error) {
	data := s.GetAllSubjects()
	return c.JSON(http.StatusOK, &DataMessage{http.StatusOK, data})
}

func (s *GatewayServer) getAllActions(c echo.Context) (err error) {
	data := s.GetAllActions()
	return c.JSON(http.StatusOK, &DataMessage{http.StatusOK, data})
}

func (s *GatewayServer) getGroupPolicies(c echo.Context) (err error) {
	data := s.GetGroupingPolicy()
	return c.JSON(http.StatusOK, &DataMessage{http.StatusOK, data})
}

func (s *GatewayServer) addRoleToUser(c echo.Context) (err error) {
	ur := new(UserRole)
	if err = c.Bind(ur); err != nil {
		return
	}
	success := s.AddRoleForUser(ur.User, ur.Role)
	return c.JSON(http.StatusOK, &SuccessMessage{http.StatusOK, success})
}

func (s *GatewayServer) removeRoleFromUser(c echo.Context) (err error) {
	ur := new(UserRole)
	if err = c.Bind(ur); err != nil {
		return
	}
	s.DeleteRoleForUser(ur.User, ur.Role)
	return c.String(http.StatusOK, "")
}

func (s *GatewayServer) getUserRoles(c echo.Context) (err error) {
	ur := new(UserRole)
	if err = c.Bind(ur); err != nil {
		return
	}
	data := s.GetRolesForUser(ur.User)
	return c.JSON(http.StatusOK, &DataMessage{http.StatusOK, data})
}

func (s *GatewayServer) getRoleUsers(c echo.Context) (err error) {
	ur := new(UserRole)
	if err = c.Bind(ur); err != nil {
		return
	}
	data := s.GetUsersForRole(ur.Role)
	return c.JSON(http.StatusOK, &DataMessage{http.StatusOK, data})
}

func (s *GatewayServer) removePolicy(c echo.Context) (err error) {
	p := new(Policy)
	if err = c.Bind(p); err != nil {
		return
	}
	success := s.RemovePolicy(p.User, p.Path, p.Method)
	return c.JSON(http.StatusOK, &SuccessMessage{http.StatusOK, success})
}
