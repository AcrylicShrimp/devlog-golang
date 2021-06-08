// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devlog/ent/predicate"
	"devlog/ent/unsavedpost"
	"devlog/ent/unsavedpostthumbnail"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UnsavedPostThumbnailUpdate is the builder for updating UnsavedPostThumbnail entities.
type UnsavedPostThumbnailUpdate struct {
	config
	hooks    []Hook
	mutation *UnsavedPostThumbnailMutation
}

// Where adds a new predicate for the UnsavedPostThumbnailUpdate builder.
func (uptu *UnsavedPostThumbnailUpdate) Where(ps ...predicate.UnsavedPostThumbnail) *UnsavedPostThumbnailUpdate {
	uptu.mutation.predicates = append(uptu.mutation.predicates, ps...)
	return uptu
}

// SetWidth sets the "width" field.
func (uptu *UnsavedPostThumbnailUpdate) SetWidth(u uint32) *UnsavedPostThumbnailUpdate {
	uptu.mutation.ResetWidth()
	uptu.mutation.SetWidth(u)
	return uptu
}

// AddWidth adds u to the "width" field.
func (uptu *UnsavedPostThumbnailUpdate) AddWidth(u uint32) *UnsavedPostThumbnailUpdate {
	uptu.mutation.AddWidth(u)
	return uptu
}

// SetHeight sets the "height" field.
func (uptu *UnsavedPostThumbnailUpdate) SetHeight(u uint32) *UnsavedPostThumbnailUpdate {
	uptu.mutation.ResetHeight()
	uptu.mutation.SetHeight(u)
	return uptu
}

// AddHeight adds u to the "height" field.
func (uptu *UnsavedPostThumbnailUpdate) AddHeight(u uint32) *UnsavedPostThumbnailUpdate {
	uptu.mutation.AddHeight(u)
	return uptu
}

// SetHash sets the "hash" field.
func (uptu *UnsavedPostThumbnailUpdate) SetHash(s string) *UnsavedPostThumbnailUpdate {
	uptu.mutation.SetHash(s)
	return uptu
}

// SetURL sets the "url" field.
func (uptu *UnsavedPostThumbnailUpdate) SetURL(s string) *UnsavedPostThumbnailUpdate {
	uptu.mutation.SetURL(s)
	return uptu
}

// SetCreatedAt sets the "created_at" field.
func (uptu *UnsavedPostThumbnailUpdate) SetCreatedAt(t time.Time) *UnsavedPostThumbnailUpdate {
	uptu.mutation.SetCreatedAt(t)
	return uptu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uptu *UnsavedPostThumbnailUpdate) SetNillableCreatedAt(t *time.Time) *UnsavedPostThumbnailUpdate {
	if t != nil {
		uptu.SetCreatedAt(*t)
	}
	return uptu
}

// SetUnsavedPostID sets the "unsaved_post" edge to the UnsavedPost entity by ID.
func (uptu *UnsavedPostThumbnailUpdate) SetUnsavedPostID(id int) *UnsavedPostThumbnailUpdate {
	uptu.mutation.SetUnsavedPostID(id)
	return uptu
}

// SetUnsavedPost sets the "unsaved_post" edge to the UnsavedPost entity.
func (uptu *UnsavedPostThumbnailUpdate) SetUnsavedPost(u *UnsavedPost) *UnsavedPostThumbnailUpdate {
	return uptu.SetUnsavedPostID(u.ID)
}

// Mutation returns the UnsavedPostThumbnailMutation object of the builder.
func (uptu *UnsavedPostThumbnailUpdate) Mutation() *UnsavedPostThumbnailMutation {
	return uptu.mutation
}

// ClearUnsavedPost clears the "unsaved_post" edge to the UnsavedPost entity.
func (uptu *UnsavedPostThumbnailUpdate) ClearUnsavedPost() *UnsavedPostThumbnailUpdate {
	uptu.mutation.ClearUnsavedPost()
	return uptu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uptu *UnsavedPostThumbnailUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uptu.hooks) == 0 {
		if err = uptu.check(); err != nil {
			return 0, err
		}
		affected, err = uptu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UnsavedPostThumbnailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uptu.check(); err != nil {
				return 0, err
			}
			uptu.mutation = mutation
			affected, err = uptu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uptu.hooks) - 1; i >= 0; i-- {
			mut = uptu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uptu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uptu *UnsavedPostThumbnailUpdate) SaveX(ctx context.Context) int {
	affected, err := uptu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uptu *UnsavedPostThumbnailUpdate) Exec(ctx context.Context) error {
	_, err := uptu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uptu *UnsavedPostThumbnailUpdate) ExecX(ctx context.Context) {
	if err := uptu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uptu *UnsavedPostThumbnailUpdate) check() error {
	if v, ok := uptu.mutation.Hash(); ok {
		if err := unsavedpostthumbnail.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf("ent: validator failed for field \"hash\": %w", err)}
		}
	}
	if v, ok := uptu.mutation.URL(); ok {
		if err := unsavedpostthumbnail.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf("ent: validator failed for field \"url\": %w", err)}
		}
	}
	if _, ok := uptu.mutation.UnsavedPostID(); uptu.mutation.UnsavedPostCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"unsaved_post\"")
	}
	return nil
}

func (uptu *UnsavedPostThumbnailUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   unsavedpostthumbnail.Table,
			Columns: unsavedpostthumbnail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: unsavedpostthumbnail.FieldID,
			},
		},
	}
	if ps := uptu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uptu.mutation.Width(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: unsavedpostthumbnail.FieldWidth,
		})
	}
	if value, ok := uptu.mutation.AddedWidth(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: unsavedpostthumbnail.FieldWidth,
		})
	}
	if value, ok := uptu.mutation.Height(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: unsavedpostthumbnail.FieldHeight,
		})
	}
	if value, ok := uptu.mutation.AddedHeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: unsavedpostthumbnail.FieldHeight,
		})
	}
	if value, ok := uptu.mutation.Hash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostthumbnail.FieldHash,
		})
	}
	if value, ok := uptu.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostthumbnail.FieldURL,
		})
	}
	if value, ok := uptu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: unsavedpostthumbnail.FieldCreatedAt,
		})
	}
	if uptu.mutation.UnsavedPostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   unsavedpostthumbnail.UnsavedPostTable,
			Columns: []string{unsavedpostthumbnail.UnsavedPostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: unsavedpost.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uptu.mutation.UnsavedPostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   unsavedpostthumbnail.UnsavedPostTable,
			Columns: []string{unsavedpostthumbnail.UnsavedPostColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uptu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{unsavedpostthumbnail.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UnsavedPostThumbnailUpdateOne is the builder for updating a single UnsavedPostThumbnail entity.
type UnsavedPostThumbnailUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UnsavedPostThumbnailMutation
}

// SetWidth sets the "width" field.
func (uptuo *UnsavedPostThumbnailUpdateOne) SetWidth(u uint32) *UnsavedPostThumbnailUpdateOne {
	uptuo.mutation.ResetWidth()
	uptuo.mutation.SetWidth(u)
	return uptuo
}

// AddWidth adds u to the "width" field.
func (uptuo *UnsavedPostThumbnailUpdateOne) AddWidth(u uint32) *UnsavedPostThumbnailUpdateOne {
	uptuo.mutation.AddWidth(u)
	return uptuo
}

// SetHeight sets the "height" field.
func (uptuo *UnsavedPostThumbnailUpdateOne) SetHeight(u uint32) *UnsavedPostThumbnailUpdateOne {
	uptuo.mutation.ResetHeight()
	uptuo.mutation.SetHeight(u)
	return uptuo
}

// AddHeight adds u to the "height" field.
func (uptuo *UnsavedPostThumbnailUpdateOne) AddHeight(u uint32) *UnsavedPostThumbnailUpdateOne {
	uptuo.mutation.AddHeight(u)
	return uptuo
}

// SetHash sets the "hash" field.
func (uptuo *UnsavedPostThumbnailUpdateOne) SetHash(s string) *UnsavedPostThumbnailUpdateOne {
	uptuo.mutation.SetHash(s)
	return uptuo
}

// SetURL sets the "url" field.
func (uptuo *UnsavedPostThumbnailUpdateOne) SetURL(s string) *UnsavedPostThumbnailUpdateOne {
	uptuo.mutation.SetURL(s)
	return uptuo
}

// SetCreatedAt sets the "created_at" field.
func (uptuo *UnsavedPostThumbnailUpdateOne) SetCreatedAt(t time.Time) *UnsavedPostThumbnailUpdateOne {
	uptuo.mutation.SetCreatedAt(t)
	return uptuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uptuo *UnsavedPostThumbnailUpdateOne) SetNillableCreatedAt(t *time.Time) *UnsavedPostThumbnailUpdateOne {
	if t != nil {
		uptuo.SetCreatedAt(*t)
	}
	return uptuo
}

// SetUnsavedPostID sets the "unsaved_post" edge to the UnsavedPost entity by ID.
func (uptuo *UnsavedPostThumbnailUpdateOne) SetUnsavedPostID(id int) *UnsavedPostThumbnailUpdateOne {
	uptuo.mutation.SetUnsavedPostID(id)
	return uptuo
}

// SetUnsavedPost sets the "unsaved_post" edge to the UnsavedPost entity.
func (uptuo *UnsavedPostThumbnailUpdateOne) SetUnsavedPost(u *UnsavedPost) *UnsavedPostThumbnailUpdateOne {
	return uptuo.SetUnsavedPostID(u.ID)
}

// Mutation returns the UnsavedPostThumbnailMutation object of the builder.
func (uptuo *UnsavedPostThumbnailUpdateOne) Mutation() *UnsavedPostThumbnailMutation {
	return uptuo.mutation
}

// ClearUnsavedPost clears the "unsaved_post" edge to the UnsavedPost entity.
func (uptuo *UnsavedPostThumbnailUpdateOne) ClearUnsavedPost() *UnsavedPostThumbnailUpdateOne {
	uptuo.mutation.ClearUnsavedPost()
	return uptuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uptuo *UnsavedPostThumbnailUpdateOne) Select(field string, fields ...string) *UnsavedPostThumbnailUpdateOne {
	uptuo.fields = append([]string{field}, fields...)
	return uptuo
}

// Save executes the query and returns the updated UnsavedPostThumbnail entity.
func (uptuo *UnsavedPostThumbnailUpdateOne) Save(ctx context.Context) (*UnsavedPostThumbnail, error) {
	var (
		err  error
		node *UnsavedPostThumbnail
	)
	if len(uptuo.hooks) == 0 {
		if err = uptuo.check(); err != nil {
			return nil, err
		}
		node, err = uptuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UnsavedPostThumbnailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uptuo.check(); err != nil {
				return nil, err
			}
			uptuo.mutation = mutation
			node, err = uptuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uptuo.hooks) - 1; i >= 0; i-- {
			mut = uptuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uptuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uptuo *UnsavedPostThumbnailUpdateOne) SaveX(ctx context.Context) *UnsavedPostThumbnail {
	node, err := uptuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uptuo *UnsavedPostThumbnailUpdateOne) Exec(ctx context.Context) error {
	_, err := uptuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uptuo *UnsavedPostThumbnailUpdateOne) ExecX(ctx context.Context) {
	if err := uptuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uptuo *UnsavedPostThumbnailUpdateOne) check() error {
	if v, ok := uptuo.mutation.Hash(); ok {
		if err := unsavedpostthumbnail.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf("ent: validator failed for field \"hash\": %w", err)}
		}
	}
	if v, ok := uptuo.mutation.URL(); ok {
		if err := unsavedpostthumbnail.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf("ent: validator failed for field \"url\": %w", err)}
		}
	}
	if _, ok := uptuo.mutation.UnsavedPostID(); uptuo.mutation.UnsavedPostCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"unsaved_post\"")
	}
	return nil
}

func (uptuo *UnsavedPostThumbnailUpdateOne) sqlSave(ctx context.Context) (_node *UnsavedPostThumbnail, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   unsavedpostthumbnail.Table,
			Columns: unsavedpostthumbnail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: unsavedpostthumbnail.FieldID,
			},
		},
	}
	id, ok := uptuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing UnsavedPostThumbnail.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := uptuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, unsavedpostthumbnail.FieldID)
		for _, f := range fields {
			if !unsavedpostthumbnail.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != unsavedpostthumbnail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uptuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uptuo.mutation.Width(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: unsavedpostthumbnail.FieldWidth,
		})
	}
	if value, ok := uptuo.mutation.AddedWidth(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: unsavedpostthumbnail.FieldWidth,
		})
	}
	if value, ok := uptuo.mutation.Height(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: unsavedpostthumbnail.FieldHeight,
		})
	}
	if value, ok := uptuo.mutation.AddedHeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: unsavedpostthumbnail.FieldHeight,
		})
	}
	if value, ok := uptuo.mutation.Hash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostthumbnail.FieldHash,
		})
	}
	if value, ok := uptuo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostthumbnail.FieldURL,
		})
	}
	if value, ok := uptuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: unsavedpostthumbnail.FieldCreatedAt,
		})
	}
	if uptuo.mutation.UnsavedPostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   unsavedpostthumbnail.UnsavedPostTable,
			Columns: []string{unsavedpostthumbnail.UnsavedPostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: unsavedpost.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uptuo.mutation.UnsavedPostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   unsavedpostthumbnail.UnsavedPostTable,
			Columns: []string{unsavedpostthumbnail.UnsavedPostColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UnsavedPostThumbnail{config: uptuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uptuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{unsavedpostthumbnail.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
