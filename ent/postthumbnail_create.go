// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devlog/ent/post"
	"devlog/ent/postthumbnail"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// PostThumbnailCreate is the builder for creating a PostThumbnail entity.
type PostThumbnailCreate struct {
	config
	mutation *PostThumbnailMutation
	hooks    []Hook
}

// SetWidth sets the "width" field.
func (ptc *PostThumbnailCreate) SetWidth(u uint32) *PostThumbnailCreate {
	ptc.mutation.SetWidth(u)
	return ptc
}

// SetHeight sets the "height" field.
func (ptc *PostThumbnailCreate) SetHeight(u uint32) *PostThumbnailCreate {
	ptc.mutation.SetHeight(u)
	return ptc
}

// SetHash sets the "hash" field.
func (ptc *PostThumbnailCreate) SetHash(s string) *PostThumbnailCreate {
	ptc.mutation.SetHash(s)
	return ptc
}

// SetURL sets the "url" field.
func (ptc *PostThumbnailCreate) SetURL(s string) *PostThumbnailCreate {
	ptc.mutation.SetURL(s)
	return ptc
}

// SetCreatedAt sets the "created_at" field.
func (ptc *PostThumbnailCreate) SetCreatedAt(t time.Time) *PostThumbnailCreate {
	ptc.mutation.SetCreatedAt(t)
	return ptc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ptc *PostThumbnailCreate) SetNillableCreatedAt(t *time.Time) *PostThumbnailCreate {
	if t != nil {
		ptc.SetCreatedAt(*t)
	}
	return ptc
}

// SetPostID sets the "post" edge to the Post entity by ID.
func (ptc *PostThumbnailCreate) SetPostID(id int) *PostThumbnailCreate {
	ptc.mutation.SetPostID(id)
	return ptc
}

// SetPost sets the "post" edge to the Post entity.
func (ptc *PostThumbnailCreate) SetPost(p *Post) *PostThumbnailCreate {
	return ptc.SetPostID(p.ID)
}

// Mutation returns the PostThumbnailMutation object of the builder.
func (ptc *PostThumbnailCreate) Mutation() *PostThumbnailMutation {
	return ptc.mutation
}

// Save creates the PostThumbnail in the database.
func (ptc *PostThumbnailCreate) Save(ctx context.Context) (*PostThumbnail, error) {
	var (
		err  error
		node *PostThumbnail
	)
	ptc.defaults()
	if len(ptc.hooks) == 0 {
		if err = ptc.check(); err != nil {
			return nil, err
		}
		node, err = ptc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PostThumbnailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ptc.check(); err != nil {
				return nil, err
			}
			ptc.mutation = mutation
			node, err = ptc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ptc.hooks) - 1; i >= 0; i-- {
			mut = ptc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ptc *PostThumbnailCreate) SaveX(ctx context.Context) *PostThumbnail {
	v, err := ptc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (ptc *PostThumbnailCreate) defaults() {
	if _, ok := ptc.mutation.CreatedAt(); !ok {
		v := postthumbnail.DefaultCreatedAt()
		ptc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptc *PostThumbnailCreate) check() error {
	if _, ok := ptc.mutation.Width(); !ok {
		return &ValidationError{Name: "width", err: errors.New("ent: missing required field \"width\"")}
	}
	if _, ok := ptc.mutation.Height(); !ok {
		return &ValidationError{Name: "height", err: errors.New("ent: missing required field \"height\"")}
	}
	if _, ok := ptc.mutation.Hash(); !ok {
		return &ValidationError{Name: "hash", err: errors.New("ent: missing required field \"hash\"")}
	}
	if v, ok := ptc.mutation.Hash(); ok {
		if err := postthumbnail.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf("ent: validator failed for field \"hash\": %w", err)}
		}
	}
	if _, ok := ptc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New("ent: missing required field \"url\"")}
	}
	if v, ok := ptc.mutation.URL(); ok {
		if err := postthumbnail.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf("ent: validator failed for field \"url\": %w", err)}
		}
	}
	if _, ok := ptc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := ptc.mutation.PostID(); !ok {
		return &ValidationError{Name: "post", err: errors.New("ent: missing required edge \"post\"")}
	}
	return nil
}

func (ptc *PostThumbnailCreate) sqlSave(ctx context.Context) (*PostThumbnail, error) {
	_node, _spec := ptc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ptc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ptc *PostThumbnailCreate) createSpec() (*PostThumbnail, *sqlgraph.CreateSpec) {
	var (
		_node = &PostThumbnail{config: ptc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: postthumbnail.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: postthumbnail.FieldID,
			},
		}
	)
	if value, ok := ptc.mutation.Width(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: postthumbnail.FieldWidth,
		})
		_node.Width = value
	}
	if value, ok := ptc.mutation.Height(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: postthumbnail.FieldHeight,
		})
		_node.Height = value
	}
	if value, ok := ptc.mutation.Hash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postthumbnail.FieldHash,
		})
		_node.Hash = value
	}
	if value, ok := ptc.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postthumbnail.FieldURL,
		})
		_node.URL = value
	}
	if value, ok := ptc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: postthumbnail.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := ptc.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   postthumbnail.PostTable,
			Columns: []string{postthumbnail.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PostThumbnailCreateBulk is the builder for creating many PostThumbnail entities in bulk.
type PostThumbnailCreateBulk struct {
	config
	builders []*PostThumbnailCreate
}

// Save creates the PostThumbnail entities in the database.
func (ptcb *PostThumbnailCreateBulk) Save(ctx context.Context) ([]*PostThumbnail, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ptcb.builders))
	nodes := make([]*PostThumbnail, len(ptcb.builders))
	mutators := make([]Mutator, len(ptcb.builders))
	for i := range ptcb.builders {
		func(i int, root context.Context) {
			builder := ptcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PostThumbnailMutation)
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
					_, err = mutators[i+1].Mutate(root, ptcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ptcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ptcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ptcb *PostThumbnailCreateBulk) SaveX(ctx context.Context) []*PostThumbnail {
	v, err := ptcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
