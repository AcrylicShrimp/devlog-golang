// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devlog/ent/post"
	"devlog/ent/postvideo"
	"devlog/ent/predicate"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PostVideoUpdate is the builder for updating PostVideo entities.
type PostVideoUpdate struct {
	config
	hooks    []Hook
	mutation *PostVideoMutation
}

// Where appends a list predicates to the PostVideoUpdate builder.
func (pvu *PostVideoUpdate) Where(ps ...predicate.PostVideo) *PostVideoUpdate {
	pvu.mutation.Where(ps...)
	return pvu
}

// SetUUID sets the "uuid" field.
func (pvu *PostVideoUpdate) SetUUID(s string) *PostVideoUpdate {
	pvu.mutation.SetUUID(s)
	return pvu
}

// SetTitle sets the "title" field.
func (pvu *PostVideoUpdate) SetTitle(s string) *PostVideoUpdate {
	pvu.mutation.SetTitle(s)
	return pvu
}

// SetURL sets the "url" field.
func (pvu *PostVideoUpdate) SetURL(s string) *PostVideoUpdate {
	pvu.mutation.SetURL(s)
	return pvu
}

// SetCreatedAt sets the "created_at" field.
func (pvu *PostVideoUpdate) SetCreatedAt(t time.Time) *PostVideoUpdate {
	pvu.mutation.SetCreatedAt(t)
	return pvu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pvu *PostVideoUpdate) SetNillableCreatedAt(t *time.Time) *PostVideoUpdate {
	if t != nil {
		pvu.SetCreatedAt(*t)
	}
	return pvu
}

// SetPostID sets the "post" edge to the Post entity by ID.
func (pvu *PostVideoUpdate) SetPostID(id int) *PostVideoUpdate {
	pvu.mutation.SetPostID(id)
	return pvu
}

// SetPost sets the "post" edge to the Post entity.
func (pvu *PostVideoUpdate) SetPost(p *Post) *PostVideoUpdate {
	return pvu.SetPostID(p.ID)
}

// Mutation returns the PostVideoMutation object of the builder.
func (pvu *PostVideoUpdate) Mutation() *PostVideoMutation {
	return pvu.mutation
}

// ClearPost clears the "post" edge to the Post entity.
func (pvu *PostVideoUpdate) ClearPost() *PostVideoUpdate {
	pvu.mutation.ClearPost()
	return pvu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pvu *PostVideoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pvu.hooks) == 0 {
		if err = pvu.check(); err != nil {
			return 0, err
		}
		affected, err = pvu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PostVideoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pvu.check(); err != nil {
				return 0, err
			}
			pvu.mutation = mutation
			affected, err = pvu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pvu.hooks) - 1; i >= 0; i-- {
			if pvu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pvu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pvu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pvu *PostVideoUpdate) SaveX(ctx context.Context) int {
	affected, err := pvu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pvu *PostVideoUpdate) Exec(ctx context.Context) error {
	_, err := pvu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pvu *PostVideoUpdate) ExecX(ctx context.Context) {
	if err := pvu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pvu *PostVideoUpdate) check() error {
	if v, ok := pvu.mutation.UUID(); ok {
		if err := postvideo.UUIDValidator(v); err != nil {
			return &ValidationError{Name: "uuid", err: fmt.Errorf("ent: validator failed for field \"uuid\": %w", err)}
		}
	}
	if v, ok := pvu.mutation.Title(); ok {
		if err := postvideo.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if v, ok := pvu.mutation.URL(); ok {
		if err := postvideo.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf("ent: validator failed for field \"url\": %w", err)}
		}
	}
	if _, ok := pvu.mutation.PostID(); pvu.mutation.PostCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"post\"")
	}
	return nil
}

func (pvu *PostVideoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   postvideo.Table,
			Columns: postvideo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: postvideo.FieldID,
			},
		},
	}
	if ps := pvu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pvu.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postvideo.FieldUUID,
		})
	}
	if value, ok := pvu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postvideo.FieldTitle,
		})
	}
	if value, ok := pvu.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postvideo.FieldURL,
		})
	}
	if value, ok := pvu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: postvideo.FieldCreatedAt,
		})
	}
	if pvu.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   postvideo.PostTable,
			Columns: []string{postvideo.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvu.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   postvideo.PostTable,
			Columns: []string{postvideo.PostColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pvu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{postvideo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// PostVideoUpdateOne is the builder for updating a single PostVideo entity.
type PostVideoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PostVideoMutation
}

// SetUUID sets the "uuid" field.
func (pvuo *PostVideoUpdateOne) SetUUID(s string) *PostVideoUpdateOne {
	pvuo.mutation.SetUUID(s)
	return pvuo
}

// SetTitle sets the "title" field.
func (pvuo *PostVideoUpdateOne) SetTitle(s string) *PostVideoUpdateOne {
	pvuo.mutation.SetTitle(s)
	return pvuo
}

// SetURL sets the "url" field.
func (pvuo *PostVideoUpdateOne) SetURL(s string) *PostVideoUpdateOne {
	pvuo.mutation.SetURL(s)
	return pvuo
}

// SetCreatedAt sets the "created_at" field.
func (pvuo *PostVideoUpdateOne) SetCreatedAt(t time.Time) *PostVideoUpdateOne {
	pvuo.mutation.SetCreatedAt(t)
	return pvuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pvuo *PostVideoUpdateOne) SetNillableCreatedAt(t *time.Time) *PostVideoUpdateOne {
	if t != nil {
		pvuo.SetCreatedAt(*t)
	}
	return pvuo
}

// SetPostID sets the "post" edge to the Post entity by ID.
func (pvuo *PostVideoUpdateOne) SetPostID(id int) *PostVideoUpdateOne {
	pvuo.mutation.SetPostID(id)
	return pvuo
}

// SetPost sets the "post" edge to the Post entity.
func (pvuo *PostVideoUpdateOne) SetPost(p *Post) *PostVideoUpdateOne {
	return pvuo.SetPostID(p.ID)
}

// Mutation returns the PostVideoMutation object of the builder.
func (pvuo *PostVideoUpdateOne) Mutation() *PostVideoMutation {
	return pvuo.mutation
}

// ClearPost clears the "post" edge to the Post entity.
func (pvuo *PostVideoUpdateOne) ClearPost() *PostVideoUpdateOne {
	pvuo.mutation.ClearPost()
	return pvuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pvuo *PostVideoUpdateOne) Select(field string, fields ...string) *PostVideoUpdateOne {
	pvuo.fields = append([]string{field}, fields...)
	return pvuo
}

// Save executes the query and returns the updated PostVideo entity.
func (pvuo *PostVideoUpdateOne) Save(ctx context.Context) (*PostVideo, error) {
	var (
		err  error
		node *PostVideo
	)
	if len(pvuo.hooks) == 0 {
		if err = pvuo.check(); err != nil {
			return nil, err
		}
		node, err = pvuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PostVideoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pvuo.check(); err != nil {
				return nil, err
			}
			pvuo.mutation = mutation
			node, err = pvuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pvuo.hooks) - 1; i >= 0; i-- {
			if pvuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pvuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pvuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pvuo *PostVideoUpdateOne) SaveX(ctx context.Context) *PostVideo {
	node, err := pvuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pvuo *PostVideoUpdateOne) Exec(ctx context.Context) error {
	_, err := pvuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pvuo *PostVideoUpdateOne) ExecX(ctx context.Context) {
	if err := pvuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pvuo *PostVideoUpdateOne) check() error {
	if v, ok := pvuo.mutation.UUID(); ok {
		if err := postvideo.UUIDValidator(v); err != nil {
			return &ValidationError{Name: "uuid", err: fmt.Errorf("ent: validator failed for field \"uuid\": %w", err)}
		}
	}
	if v, ok := pvuo.mutation.Title(); ok {
		if err := postvideo.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if v, ok := pvuo.mutation.URL(); ok {
		if err := postvideo.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf("ent: validator failed for field \"url\": %w", err)}
		}
	}
	if _, ok := pvuo.mutation.PostID(); pvuo.mutation.PostCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"post\"")
	}
	return nil
}

func (pvuo *PostVideoUpdateOne) sqlSave(ctx context.Context) (_node *PostVideo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   postvideo.Table,
			Columns: postvideo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: postvideo.FieldID,
			},
		},
	}
	id, ok := pvuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing PostVideo.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := pvuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, postvideo.FieldID)
		for _, f := range fields {
			if !postvideo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != postvideo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pvuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pvuo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postvideo.FieldUUID,
		})
	}
	if value, ok := pvuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postvideo.FieldTitle,
		})
	}
	if value, ok := pvuo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postvideo.FieldURL,
		})
	}
	if value, ok := pvuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: postvideo.FieldCreatedAt,
		})
	}
	if pvuo.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   postvideo.PostTable,
			Columns: []string{postvideo.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvuo.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   postvideo.PostTable,
			Columns: []string{postvideo.PostColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PostVideo{config: pvuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pvuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{postvideo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
