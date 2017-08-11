package server

import (
	"github.com/Soontao/go-simple-api-gateway/key"
	"github.com/ipfans/echo-session"
	"github.com/labstack/echo"
)

// BasicAuthSessionMw is used for reading basic auth header and save username if it passed
func (s *GatewayServer) BasicAuthSessionMw(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		username, password, exist := c.Request().BasicAuth()
		sess := session.Default(c)
		if exist && username != "" && s.authUserService.AuthUser(username, password) {
			sess.Set(key.KEY_Username, username)
		} else {
			sess.Delete(key.KEY_Username)
		}
		sess.Save()
		return next(c)
	}
}
