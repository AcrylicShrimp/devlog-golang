package common

import (
	"devlog/ent"
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
	client *ent.Client
	admin  *ent.Admin
}

func NewContext(c echo.Context, client *ent.Client) *Context {
	return &Context{c, client, nil}
}

func (c *Context) Client() *ent.Client {
	return c.client
}

func (c *Context) Admin() *ent.Admin {
	return c.admin
}

func (c *Context) SetAdmin(admin *ent.Admin) {
	c.admin = admin
}
