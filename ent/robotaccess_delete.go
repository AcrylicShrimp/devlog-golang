// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devlog/ent/predicate"
	"devlog/ent/robotaccess"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RobotAccessDelete is the builder for deleting a RobotAccess entity.
type RobotAccessDelete struct {
	config
	hooks    []Hook
	mutation *RobotAccessMutation
}

// Where appends a list predicates to the RobotAccessDelete builder.
func (rad *RobotAccessDelete) Where(ps ...predicate.RobotAccess) *RobotAccessDelete {
	rad.mutation.Where(ps...)
	return rad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rad *RobotAccessDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rad.hooks) == 0 {
		affected, err = rad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RobotAccessMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rad.mutation = mutation
			affected, err = rad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rad.hooks) - 1; i >= 0; i-- {
			if rad.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (rad *RobotAccessDelete) ExecX(ctx context.Context) int {
	n, err := rad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rad *RobotAccessDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: robotaccess.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: robotaccess.FieldID,
			},
		},
	}
	if ps := rad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, rad.driver, _spec)
}

// RobotAccessDeleteOne is the builder for deleting a single RobotAccess entity.
type RobotAccessDeleteOne struct {
	rad *RobotAccessDelete
}

// Exec executes the deletion query.
func (rado *RobotAccessDeleteOne) Exec(ctx context.Context) error {
	n, err := rado.rad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{robotaccess.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rado *RobotAccessDeleteOne) ExecX(ctx context.Context) {
	rado.rad.ExecX(ctx)
}