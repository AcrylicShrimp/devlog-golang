package common

import (
	"context"
	"devlog/ent"
	"github.com/labstack/echo"
)

type Context struct {
	echo.Context
	client *ent.Client
	ctx    context.Context
}

func NewContext(c echo.Context, client *ent.Client, ctx context.Context) *Context {
	return &Context{c, client, ctx}
}

func (c *Context) Client() *ent.Client {
	return c.client
}

func (c *Context) Ctx() context.Context {
	return c.ctx
}
