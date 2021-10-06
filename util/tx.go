package util

import (
	"context"
	"devlog/common"
	"devlog/ent"
	"github.com/labstack/echo/v4"
)

func WithTx(c echo.Context, fn func(ctx *common.Context, tx *ent.Tx) (interface{}, error)) (interface{}, error) {
	ctx := c.(*common.Context)
	tx, err := ctx.Client().Tx(context.Background())

	if err != nil {
		return nil, err
	}

	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
			panic(v)
		}
	}()

	v, err := fn(ctx, tx)

	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return v, nil
}
