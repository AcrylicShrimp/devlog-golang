// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devlog/ent/admin"
	"devlog/ent/adminsession"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AdminSessionCreate is the builder for creating a AdminSession entity.
type AdminSessionCreate struct {
	config
	mutation *AdminSessionMutation
	hooks    []Hook
}

// SetToken sets the "token" field.
func (asc *AdminSessionCreate) SetToken(s string) *AdminSessionCreate {
	asc.mutation.SetToken(s)
	return asc
}

// SetExpiresAt sets the "expires_at" field.
func (asc *AdminSessionCreate) SetExpiresAt(t time.Time) *AdminSessionCreate {
	asc.mutation.SetExpiresAt(t)
	return asc
}

// SetCreatedAt sets the "created_at" field.
func (asc *AdminSessionCreate) SetCreatedAt(t time.Time) *AdminSessionCreate {
	asc.mutation.SetCreatedAt(t)
	return asc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (asc *AdminSessionCreate) SetNillableCreatedAt(t *time.Time) *AdminSessionCreate {
	if t != nil {
		asc.SetCreatedAt(*t)
	}
	return asc
}

// SetUserID sets the "user" edge to the Admin entity by ID.
func (asc *AdminSessionCreate) SetUserID(id int) *AdminSessionCreate {
	asc.mutation.SetUserID(id)
	return asc
}

// SetUser sets the "user" edge to the Admin entity.
func (asc *AdminSessionCreate) SetUser(a *Admin) *AdminSessionCreate {
	return asc.SetUserID(a.ID)
}

// Mutation returns the AdminSessionMutation object of the builder.
func (asc *AdminSessionCreate) Mutation() *AdminSessionMutation {
	return asc.mutation
}

// Save creates the AdminSession in the database.
func (asc *AdminSessionCreate) Save(ctx context.Context) (*AdminSession, error) {
	var (
		err  error
		node *AdminSession
	)
	asc.defaults()
	if len(asc.hooks) == 0 {
		if err = asc.check(); err != nil {
			return nil, err
		}
		node, err = asc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdminSessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = asc.check(); err != nil {
				return nil, err
			}
			asc.mutation = mutation
			if node, err = asc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(asc.hooks) - 1; i >= 0; i-- {
			if asc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = asc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, asc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (asc *AdminSessionCreate) SaveX(ctx context.Context) *AdminSession {
	v, err := asc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (asc *AdminSessionCreate) Exec(ctx context.Context) error {
	_, err := asc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asc *AdminSessionCreate) ExecX(ctx context.Context) {
	if err := asc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (asc *AdminSessionCreate) defaults() {
	if _, ok := asc.mutation.CreatedAt(); !ok {
		v := adminsession.DefaultCreatedAt()
		asc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (asc *AdminSessionCreate) check() error {
	if _, ok := asc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`ent: missing required field "token"`)}
	}
	if v, ok := asc.mutation.Token(); ok {
		if err := adminsession.TokenValidator(v); err != nil {
			return &ValidationError{Name: "token", err: fmt.Errorf(`ent: validator failed for field "token": %w`, err)}
		}
	}
	if _, ok := asc.mutation.ExpiresAt(); !ok {
		return &ValidationError{Name: "expires_at", err: errors.New(`ent: missing required field "expires_at"`)}
	}
	if _, ok := asc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := asc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New("ent: missing required edge \"user\"")}
	}
	return nil
}

func (asc *AdminSessionCreate) sqlSave(ctx context.Context) (*AdminSession, error) {
	_node, _spec := asc.createSpec()
	if err := sqlgraph.CreateNode(ctx, asc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (asc *AdminSessionCreate) createSpec() (*AdminSession, *sqlgraph.CreateSpec) {
	var (
		_node = &AdminSession{config: asc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: adminsession.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: adminsession.FieldID,
			},
		}
	)
	if value, ok := asc.mutation.Token(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: adminsession.FieldToken,
		})
		_node.Token = value
	}
	if value, ok := asc.mutation.ExpiresAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: adminsession.FieldExpiresAt,
		})
		_node.ExpiresAt = value
	}
	if value, ok := asc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: adminsession.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := asc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   adminsession.UserTable,
			Columns: []string{adminsession.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: admin.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.admin_sessions = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AdminSessionCreateBulk is the builder for creating many AdminSession entities in bulk.
type AdminSessionCreateBulk struct {
	config
	builders []*AdminSessionCreate
}

// Save creates the AdminSession entities in the database.
func (ascb *AdminSessionCreateBulk) Save(ctx context.Context) ([]*AdminSession, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ascb.builders))
	nodes := make([]*AdminSession, len(ascb.builders))
	mutators := make([]Mutator, len(ascb.builders))
	for i := range ascb.builders {
		func(i int, root context.Context) {
			builder := ascb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AdminSessionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ascb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ascb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ascb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ascb *AdminSessionCreateBulk) SaveX(ctx context.Context) []*AdminSession {
	v, err := ascb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ascb *AdminSessionCreateBulk) Exec(ctx context.Context) error {
	_, err := ascb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ascb *AdminSessionCreateBulk) ExecX(ctx context.Context) {
	if err := ascb.Exec(ctx); err != nil {
		panic(err)
	}
}
