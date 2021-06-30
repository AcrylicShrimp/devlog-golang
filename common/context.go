package common

import (
	"context"
	"devlog/ent"
	"devlog/env"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	awsS3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
	client      *ent.Client
	admin       *ent.Admin
	awsS3Client *awsS3.Client
}

func NewContext(c echo.Context, client *ent.Client) *Context {
	config, err := awsConfig.LoadDefaultConfig(context.Background())

	if err != nil {
		panic(err)
	}

	return &Context{c, client, nil, awsS3.NewFromConfig(config, func(options *awsS3.Options) {
		options.Region = env.AWSS3Region
	})}
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

func (c *Context) AWSS3Client() *awsS3.Client {
	return c.awsS3Client
}
