// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devlog/ent/predicate"
	"devlog/ent/unsavedpostattachment"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UnsavedPostAttachmentDelete is the builder for deleting a UnsavedPostAttachment entity.
type UnsavedPostAttachmentDelete struct {
	config
	hooks    []Hook
	mutation *UnsavedPostAttachmentMutation
}

// Where adds a new predicate to the UnsavedPostAttachmentDelete builder.
func (upad *UnsavedPostAttachmentDelete) Where(ps ...predicate.UnsavedPostAttachment) *UnsavedPostAttachmentDelete {
	upad.mutation.predicates = append(upad.mutation.predicates, ps...)
	return upad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (upad *UnsavedPostAttachmentDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(upad.hooks) == 0 {
		affected, err = upad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UnsavedPostAttachmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			upad.mutation = mutation
			affected, err = upad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(upad.hooks) - 1; i >= 0; i-- {
			mut = upad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, upad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (upad *UnsavedPostAttachmentDelete) ExecX(ctx context.Context) int {
	n, err := upad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (upad *UnsavedPostAttachmentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: unsavedpostattachment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: unsavedpostattachment.FieldID,
			},
		},
	}
	if ps := upad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, upad.driver, _spec)
}

// UnsavedPostAttachmentDeleteOne is the builder for deleting a single UnsavedPostAttachment entity.
type UnsavedPostAttachmentDeleteOne struct {
	upad *UnsavedPostAttachmentDelete
}

// Exec executes the deletion query.
func (upado *UnsavedPostAttachmentDeleteOne) Exec(ctx context.Context) error {
	n, err := upado.upad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{unsavedpostattachment.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (upado *UnsavedPostAttachmentDeleteOne) ExecX(ctx context.Context) {
	upado.upad.ExecX(ctx)
}
