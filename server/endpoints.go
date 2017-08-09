package server

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/ipfans/echo-session"
	"fmt"
)

func (s *AuthServer) userAuth(c echo.Context) (err error) {
	sess := session.Default(c)
	user := User{}
	c.Bind(&user)
	if (s.authUserService.AuthUser(user.Username, user.Password)) {
		sess.Set(Username, user.Username)
		sess.Save()
		return c.JSON(http.StatusOK, &DataMessage{http.StatusOK, fmt.Sprintf("auth for %s", user.Username)})
	} else {
		return c.JSON(http.StatusForbidden, &DataMessage{http.StatusForbidden, "auth failed"})
	}
	return
}

func (s *AuthServer) userRegister(c echo.Context) (err error) {
	user := &User{}
	c.Bind(user)
	if err := s.authUserService.SaveUser(user.Username, user.Password); err != nil {
		return c.JSON(http.StatusOK, &DataMessage{http.StatusOK, fmt.Sprintf("register for %s", user.Username)})
	} else {
		return c.JSON(http.StatusBadRequest, &DataMessage{http.StatusBadRequest, err})
	}

	return
}

func (s *AuthServer) userUpdate(c echo.Context) (err error) {
	user := &User{}
	c.Bind(user)
	if (s.authUserService.UpdatePassword(user.Username, user.Password, user.NewPassword)) {
		return c.JSON(http.StatusOK, &DataMessage{http.StatusOK, fmt.Sprintf("password updated for %s", user.Username)})
	} else {
		return c.JSON(http.StatusBadRequest, &DataMessage{http.StatusBadRequest, "password update failed"})
	}

	return
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
