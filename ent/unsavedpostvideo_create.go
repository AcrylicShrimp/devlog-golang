// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devlog/ent/unsavedpost"
	"devlog/ent/unsavedpostvideo"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UnsavedPostVideoCreate is the builder for creating a UnsavedPostVideo entity.
type UnsavedPostVideoCreate struct {
	config
	mutation *UnsavedPostVideoMutation
	hooks    []Hook
}

// SetUUID sets the "uuid" field.
func (upvc *UnsavedPostVideoCreate) SetUUID(s string) *UnsavedPostVideoCreate {
	upvc.mutation.SetUUID(s)
	return upvc
}

// SetValidity sets the "validity" field.
func (upvc *UnsavedPostVideoCreate) SetValidity(u unsavedpostvideo.Validity) *UnsavedPostVideoCreate {
	upvc.mutation.SetValidity(u)
	return upvc
}

// SetNillableValidity sets the "validity" field if the given value is not nil.
func (upvc *UnsavedPostVideoCreate) SetNillableValidity(u *unsavedpostvideo.Validity) *UnsavedPostVideoCreate {
	if u != nil {
		upvc.SetValidity(*u)
	}
	return upvc
}

// SetTitle sets the "title" field.
func (upvc *UnsavedPostVideoCreate) SetTitle(s string) *UnsavedPostVideoCreate {
	upvc.mutation.SetTitle(s)
	return upvc
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (upvc *UnsavedPostVideoCreate) SetNillableTitle(s *string) *UnsavedPostVideoCreate {
	if s != nil {
		upvc.SetTitle(*s)
	}
	return upvc
}

// SetURL sets the "url" field.
func (upvc *UnsavedPostVideoCreate) SetURL(s string) *UnsavedPostVideoCreate {
	upvc.mutation.SetURL(s)
	return upvc
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (upvc *UnsavedPostVideoCreate) SetNillableURL(s *string) *UnsavedPostVideoCreate {
	if s != nil {
		upvc.SetURL(*s)
	}
	return upvc
}

// SetCreatedAt sets the "created_at" field.
func (upvc *UnsavedPostVideoCreate) SetCreatedAt(t time.Time) *UnsavedPostVideoCreate {
	upvc.mutation.SetCreatedAt(t)
	return upvc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (upvc *UnsavedPostVideoCreate) SetNillableCreatedAt(t *time.Time) *UnsavedPostVideoCreate {
	if t != nil {
		upvc.SetCreatedAt(*t)
	}
	return upvc
}

// SetUnsavedPostID sets the "unsaved_post" edge to the UnsavedPost entity by ID.
func (upvc *UnsavedPostVideoCreate) SetUnsavedPostID(id int) *UnsavedPostVideoCreate {
	upvc.mutation.SetUnsavedPostID(id)
	return upvc
}

// SetUnsavedPost sets the "unsaved_post" edge to the UnsavedPost entity.
func (upvc *UnsavedPostVideoCreate) SetUnsavedPost(u *UnsavedPost) *UnsavedPostVideoCreate {
	return upvc.SetUnsavedPostID(u.ID)
}

// Mutation returns the UnsavedPostVideoMutation object of the builder.
func (upvc *UnsavedPostVideoCreate) Mutation() *UnsavedPostVideoMutation {
	return upvc.mutation
}

// Save creates the UnsavedPostVideo in the database.
func (upvc *UnsavedPostVideoCreate) Save(ctx context.Context) (*UnsavedPostVideo, error) {
	var (
		err  error
		node *UnsavedPostVideo
	)
	upvc.defaults()
	if len(upvc.hooks) == 0 {
		if err = upvc.check(); err != nil {
			return nil, err
		}
		node, err = upvc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UnsavedPostVideoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = upvc.check(); err != nil {
				return nil, err
			}
			upvc.mutation = mutation
			node, err = upvc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(upvc.hooks) - 1; i >= 0; i-- {
			mut = upvc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, upvc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (upvc *UnsavedPostVideoCreate) SaveX(ctx context.Context) *UnsavedPostVideo {
	v, err := upvc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (upvc *UnsavedPostVideoCreate) defaults() {
	if _, ok := upvc.mutation.Validity(); !ok {
		v := unsavedpostvideo.DefaultValidity
		upvc.mutation.SetValidity(v)
	}
	if _, ok := upvc.mutation.CreatedAt(); !ok {
		v := unsavedpostvideo.DefaultCreatedAt()
		upvc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upvc *UnsavedPostVideoCreate) check() error {
	if _, ok := upvc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New("ent: missing required field \"uuid\"")}
	}
	if v, ok := upvc.mutation.UUID(); ok {
		if err := unsavedpostvideo.UUIDValidator(v); err != nil {
			return &ValidationError{Name: "uuid", err: fmt.Errorf("ent: validator failed for field \"uuid\": %w", err)}
		}
	}
	if _, ok := upvc.mutation.Validity(); !ok {
		return &ValidationError{Name: "validity", err: errors.New("ent: missing required field \"validity\"")}
	}
	if v, ok := upvc.mutation.Validity(); ok {
		if err := unsavedpostvideo.ValidityValidator(v); err != nil {
			return &ValidationError{Name: "validity", err: fmt.Errorf("ent: validator failed for field \"validity\": %w", err)}
		}
	}
	if v, ok := upvc.mutation.Title(); ok {
		if err := unsavedpostvideo.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if v, ok := upvc.mutation.URL(); ok {
		if err := unsavedpostvideo.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf("ent: validator failed for field \"url\": %w", err)}
		}
	}
	if _, ok := upvc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := upvc.mutation.UnsavedPostID(); !ok {
		return &ValidationError{Name: "unsaved_post", err: errors.New("ent: missing required edge \"unsaved_post\"")}
	}
	return nil
}

func (upvc *UnsavedPostVideoCreate) sqlSave(ctx context.Context) (*UnsavedPostVideo, error) {
	_node, _spec := upvc.createSpec()
	if err := sqlgraph.CreateNode(ctx, upvc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (upvc *UnsavedPostVideoCreate) createSpec() (*UnsavedPostVideo, *sqlgraph.CreateSpec) {
	var (
		_node = &UnsavedPostVideo{config: upvc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: unsavedpostvideo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: unsavedpostvideo.FieldID,
			},
		}
	)
	if value, ok := upvc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostvideo.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := upvc.mutation.Validity(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: unsavedpostvideo.FieldValidity,
		})
		_node.Validity = value
	}
	if value, ok := upvc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostvideo.FieldTitle,
		})
		_node.Title = &value
	}
	if value, ok := upvc.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostvideo.FieldURL,
		})
		_node.URL = &value
	}
	if value, ok := upvc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: unsavedpostvideo.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := upvc.mutation.UnsavedPostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   unsavedpostvideo.UnsavedPostTable,
			Columns: []string{unsavedpostvideo.UnsavedPostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: unsavedpost.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.unsaved_post_videos = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UnsavedPostVideoCreateBulk is the builder for creating many UnsavedPostVideo entities in bulk.
type UnsavedPostVideoCreateBulk struct {
	config
	builders []*UnsavedPostVideoCreate
}

// Save creates the UnsavedPostVideo entities in the database.
func (upvcb *UnsavedPostVideoCreateBulk) Save(ctx context.Context) ([]*UnsavedPostVideo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(upvcb.builders))
	nodes := make([]*UnsavedPostVideo, len(upvcb.builders))
	mutators := make([]Mutator, len(upvcb.builders))
	for i := range upvcb.builders {
		func(i int, root context.Context) {
			builder := upvcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UnsavedPostVideoMutation)
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
					_, err = mutators[i+1].Mutate(root, upvcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, upvcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, upvcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (upvcb *UnsavedPostVideoCreateBulk) SaveX(ctx context.Context) []*UnsavedPostVideo {
	v, err := upvcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
