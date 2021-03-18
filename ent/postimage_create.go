// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devlog/ent/post"
	"devlog/ent/postimage"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// PostImageCreate is the builder for creating a PostImage entity.
type PostImageCreate struct {
	config
	mutation *PostImageMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (pic *PostImageCreate) SetUUID(s string) *PostImageCreate {
	pic.mutation.SetUUID(s)
	return pic
}

// SetIndex sets the "index" field.
func (pic *PostImageCreate) SetIndex(u uint64) *PostImageCreate {
	pic.mutation.SetIndex(u)
	return pic
}

// SetWidth sets the "width" field.
func (pic *PostImageCreate) SetWidth(u uint32) *PostImageCreate {
	pic.mutation.SetWidth(u)
	return pic
}

// SetHeight sets the "height" field.
func (pic *PostImageCreate) SetHeight(u uint32) *PostImageCreate {
	pic.mutation.SetHeight(u)
	return pic
}

// SetHash sets the "hash" field.
func (pic *PostImageCreate) SetHash(s string) *PostImageCreate {
	pic.mutation.SetHash(s)
	return pic
}

// SetURL sets the "url" field.
func (pic *PostImageCreate) SetURL(s string) *PostImageCreate {
	pic.mutation.SetURL(s)
	return pic
}

// SetCreatedAt sets the "created_at" field.
func (pic *PostImageCreate) SetCreatedAt(t time.Time) *PostImageCreate {
	pic.mutation.SetCreatedAt(t)
	return pic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pic *PostImageCreate) SetNillableCreatedAt(t *time.Time) *PostImageCreate {
	if t != nil {
		pic.SetCreatedAt(*t)
	}
	return pic
}

// SetPostID sets the "post" edge to the Post entity by ID.
func (pic *PostImageCreate) SetPostID(id int) *PostImageCreate {
	pic.mutation.SetPostID(id)
	return pic
}

// SetPost sets the "post" edge to the Post entity.
func (pic *PostImageCreate) SetPost(p *Post) *PostImageCreate {
	return pic.SetPostID(p.ID)
}

// Mutation returns the PostImageMutation object of the builder.
func (pic *PostImageCreate) Mutation() *PostImageMutation {
	return pic.mutation
}

// Save creates the PostImage in the database.
func (pic *PostImageCreate) Save(ctx context.Context) (*PostImage, error) {
	var (
		err  error
		node *PostImage
	)
	pic.defaults()
	if len(pic.hooks) == 0 {
		if err = pic.check(); err != nil {
			return nil, err
		}
		node, err = pic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PostImageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pic.check(); err != nil {
				return nil, err
			}
			pic.mutation = mutation
			node, err = pic.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pic.hooks) - 1; i >= 0; i-- {
			mut = pic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pic *PostImageCreate) SaveX(ctx context.Context) *PostImage {
	v, err := pic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (pic *PostImageCreate) defaults() {
	if _, ok := pic.mutation.CreatedAt(); !ok {
		v := postimage.DefaultCreatedAt()
		pic.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pic *PostImageCreate) check() error {
	if _, ok := pic.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New("ent: missing required field \"uuid\"")}
	}
	if _, ok := pic.mutation.Index(); !ok {
		return &ValidationError{Name: "index", err: errors.New("ent: missing required field \"index\"")}
	}
	if _, ok := pic.mutation.Width(); !ok {
		return &ValidationError{Name: "width", err: errors.New("ent: missing required field \"width\"")}
	}
	if _, ok := pic.mutation.Height(); !ok {
		return &ValidationError{Name: "height", err: errors.New("ent: missing required field \"height\"")}
	}
	if _, ok := pic.mutation.Hash(); !ok {
		return &ValidationError{Name: "hash", err: errors.New("ent: missing required field \"hash\"")}
	}
	if v, ok := pic.mutation.Hash(); ok {
		if err := postimage.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf("ent: validator failed for field \"hash\": %w", err)}
		}
	}
	if _, ok := pic.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New("ent: missing required field \"url\"")}
	}
	if v, ok := pic.mutation.URL(); ok {
		if err := postimage.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf("ent: validator failed for field \"url\": %w", err)}
		}
	}
	if _, ok := pic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := pic.mutation.PostID(); !ok {
		return &ValidationError{Name: "post", err: errors.New("ent: missing required edge \"post\"")}
	}
	return nil
}

func (pic *PostImageCreate) sqlSave(ctx context.Context) (*PostImage, error) {
	_node, _spec := pic.createSpec()
	if err := sqlgraph.CreateNode(ctx, pic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pic *PostImageCreate) createSpec() (*PostImage, *sqlgraph.CreateSpec) {
	var (
		_node = &PostImage{config: pic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: postimage.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: postimage.FieldID,
			},
		}
	)
	if value, ok := pic.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postimage.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := pic.mutation.Index(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: postimage.FieldIndex,
		})
		_node.Index = value
	}
	if value, ok := pic.mutation.Width(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: postimage.FieldWidth,
		})
		_node.Width = value
	}
	if value, ok := pic.mutation.Height(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: postimage.FieldHeight,
		})
		_node.Height = value
	}
	if value, ok := pic.mutation.Hash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postimage.FieldHash,
		})
		_node.Hash = value
	}
	if value, ok := pic.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postimage.FieldURL,
		})
		_node.URL = value
	}
	if value, ok := pic.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: postimage.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := pic.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   postimage.PostTable,
			Columns: []string{postimage.PostColumn},
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

// PostImageCreateBulk is the builder for creating many PostImage entities in bulk.
type PostImageCreateBulk struct {
	config
	builders []*PostImageCreate
}

// Save creates the PostImage entities in the database.
func (picb *PostImageCreateBulk) Save(ctx context.Context) ([]*PostImage, error) {
	specs := make([]*sqlgraph.CreateSpec, len(picb.builders))
	nodes := make([]*PostImage, len(picb.builders))
	mutators := make([]Mutator, len(picb.builders))
	for i := range picb.builders {
		func(i int, root context.Context) {
			builder := picb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PostImageMutation)
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
					_, err = mutators[i+1].Mutate(root, picb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, picb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, picb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (picb *PostImageCreateBulk) SaveX(ctx context.Context) []*PostImage {
	v, err := picb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}