// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devlog/ent/predicate"
	"devlog/ent/unsavedpost"
	"devlog/ent/unsavedpostimage"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UnsavedPostImageUpdate is the builder for updating UnsavedPostImage entities.
type UnsavedPostImageUpdate struct {
	config
	hooks    []Hook
	mutation *UnsavedPostImageMutation
}

// Where appends a list predicates to the UnsavedPostImageUpdate builder.
func (upiu *UnsavedPostImageUpdate) Where(ps ...predicate.UnsavedPostImage) *UnsavedPostImageUpdate {
	upiu.mutation.Where(ps...)
	return upiu
}

// SetUUID sets the "uuid" field.
func (upiu *UnsavedPostImageUpdate) SetUUID(s string) *UnsavedPostImageUpdate {
	upiu.mutation.SetUUID(s)
	return upiu
}

// SetValidity sets the "validity" field.
func (upiu *UnsavedPostImageUpdate) SetValidity(u unsavedpostimage.Validity) *UnsavedPostImageUpdate {
	upiu.mutation.SetValidity(u)
	return upiu
}

// SetNillableValidity sets the "validity" field if the given value is not nil.
func (upiu *UnsavedPostImageUpdate) SetNillableValidity(u *unsavedpostimage.Validity) *UnsavedPostImageUpdate {
	if u != nil {
		upiu.SetValidity(*u)
	}
	return upiu
}

// SetWidth sets the "width" field.
func (upiu *UnsavedPostImageUpdate) SetWidth(u uint32) *UnsavedPostImageUpdate {
	upiu.mutation.ResetWidth()
	upiu.mutation.SetWidth(u)
	return upiu
}

// SetNillableWidth sets the "width" field if the given value is not nil.
func (upiu *UnsavedPostImageUpdate) SetNillableWidth(u *uint32) *UnsavedPostImageUpdate {
	if u != nil {
		upiu.SetWidth(*u)
	}
	return upiu
}

// AddWidth adds u to the "width" field.
func (upiu *UnsavedPostImageUpdate) AddWidth(u uint32) *UnsavedPostImageUpdate {
	upiu.mutation.AddWidth(u)
	return upiu
}

// ClearWidth clears the value of the "width" field.
func (upiu *UnsavedPostImageUpdate) ClearWidth() *UnsavedPostImageUpdate {
	upiu.mutation.ClearWidth()
	return upiu
}

// SetHeight sets the "height" field.
func (upiu *UnsavedPostImageUpdate) SetHeight(u uint32) *UnsavedPostImageUpdate {
	upiu.mutation.ResetHeight()
	upiu.mutation.SetHeight(u)
	return upiu
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (upiu *UnsavedPostImageUpdate) SetNillableHeight(u *uint32) *UnsavedPostImageUpdate {
	if u != nil {
		upiu.SetHeight(*u)
	}
	return upiu
}

// AddHeight adds u to the "height" field.
func (upiu *UnsavedPostImageUpdate) AddHeight(u uint32) *UnsavedPostImageUpdate {
	upiu.mutation.AddHeight(u)
	return upiu
}

// ClearHeight clears the value of the "height" field.
func (upiu *UnsavedPostImageUpdate) ClearHeight() *UnsavedPostImageUpdate {
	upiu.mutation.ClearHeight()
	return upiu
}

// SetHash sets the "hash" field.
func (upiu *UnsavedPostImageUpdate) SetHash(s string) *UnsavedPostImageUpdate {
	upiu.mutation.SetHash(s)
	return upiu
}

// SetNillableHash sets the "hash" field if the given value is not nil.
func (upiu *UnsavedPostImageUpdate) SetNillableHash(s *string) *UnsavedPostImageUpdate {
	if s != nil {
		upiu.SetHash(*s)
	}
	return upiu
}

// ClearHash clears the value of the "hash" field.
func (upiu *UnsavedPostImageUpdate) ClearHash() *UnsavedPostImageUpdate {
	upiu.mutation.ClearHash()
	return upiu
}

// SetTitle sets the "title" field.
func (upiu *UnsavedPostImageUpdate) SetTitle(s string) *UnsavedPostImageUpdate {
	upiu.mutation.SetTitle(s)
	return upiu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (upiu *UnsavedPostImageUpdate) SetNillableTitle(s *string) *UnsavedPostImageUpdate {
	if s != nil {
		upiu.SetTitle(*s)
	}
	return upiu
}

// ClearTitle clears the value of the "title" field.
func (upiu *UnsavedPostImageUpdate) ClearTitle() *UnsavedPostImageUpdate {
	upiu.mutation.ClearTitle()
	return upiu
}

// SetURL sets the "url" field.
func (upiu *UnsavedPostImageUpdate) SetURL(s string) *UnsavedPostImageUpdate {
	upiu.mutation.SetURL(s)
	return upiu
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (upiu *UnsavedPostImageUpdate) SetNillableURL(s *string) *UnsavedPostImageUpdate {
	if s != nil {
		upiu.SetURL(*s)
	}
	return upiu
}

// ClearURL clears the value of the "url" field.
func (upiu *UnsavedPostImageUpdate) ClearURL() *UnsavedPostImageUpdate {
	upiu.mutation.ClearURL()
	return upiu
}

// SetCreatedAt sets the "created_at" field.
func (upiu *UnsavedPostImageUpdate) SetCreatedAt(t time.Time) *UnsavedPostImageUpdate {
	upiu.mutation.SetCreatedAt(t)
	return upiu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (upiu *UnsavedPostImageUpdate) SetNillableCreatedAt(t *time.Time) *UnsavedPostImageUpdate {
	if t != nil {
		upiu.SetCreatedAt(*t)
	}
	return upiu
}

// SetUnsavedPostID sets the "unsaved_post" edge to the UnsavedPost entity by ID.
func (upiu *UnsavedPostImageUpdate) SetUnsavedPostID(id int) *UnsavedPostImageUpdate {
	upiu.mutation.SetUnsavedPostID(id)
	return upiu
}

// SetUnsavedPost sets the "unsaved_post" edge to the UnsavedPost entity.
func (upiu *UnsavedPostImageUpdate) SetUnsavedPost(u *UnsavedPost) *UnsavedPostImageUpdate {
	return upiu.SetUnsavedPostID(u.ID)
}

// Mutation returns the UnsavedPostImageMutation object of the builder.
func (upiu *UnsavedPostImageUpdate) Mutation() *UnsavedPostImageMutation {
	return upiu.mutation
}

// ClearUnsavedPost clears the "unsaved_post" edge to the UnsavedPost entity.
func (upiu *UnsavedPostImageUpdate) ClearUnsavedPost() *UnsavedPostImageUpdate {
	upiu.mutation.ClearUnsavedPost()
	return upiu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (upiu *UnsavedPostImageUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(upiu.hooks) == 0 {
		if err = upiu.check(); err != nil {
			return 0, err
		}
		affected, err = upiu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UnsavedPostImageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = upiu.check(); err != nil {
				return 0, err
			}
			upiu.mutation = mutation
			affected, err = upiu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(upiu.hooks) - 1; i >= 0; i-- {
			if upiu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = upiu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, upiu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (upiu *UnsavedPostImageUpdate) SaveX(ctx context.Context) int {
	affected, err := upiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (upiu *UnsavedPostImageUpdate) Exec(ctx context.Context) error {
	_, err := upiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upiu *UnsavedPostImageUpdate) ExecX(ctx context.Context) {
	if err := upiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upiu *UnsavedPostImageUpdate) check() error {
	if v, ok := upiu.mutation.UUID(); ok {
		if err := unsavedpostimage.UUIDValidator(v); err != nil {
			return &ValidationError{Name: "uuid", err: fmt.Errorf("ent: validator failed for field \"uuid\": %w", err)}
		}
	}
	if v, ok := upiu.mutation.Validity(); ok {
		if err := unsavedpostimage.ValidityValidator(v); err != nil {
			return &ValidationError{Name: "validity", err: fmt.Errorf("ent: validator failed for field \"validity\": %w", err)}
		}
	}
	if v, ok := upiu.mutation.Hash(); ok {
		if err := unsavedpostimage.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf("ent: validator failed for field \"hash\": %w", err)}
		}
	}
	if v, ok := upiu.mutation.Title(); ok {
		if err := unsavedpostimage.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if v, ok := upiu.mutation.URL(); ok {
		if err := unsavedpostimage.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf("ent: validator failed for field \"url\": %w", err)}
		}
	}
	if _, ok := upiu.mutation.UnsavedPostID(); upiu.mutation.UnsavedPostCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"unsaved_post\"")
	}
	return nil
}

func (upiu *UnsavedPostImageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   unsavedpostimage.Table,
			Columns: unsavedpostimage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: unsavedpostimage.FieldID,
			},
		},
	}
	if ps := upiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := upiu.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostimage.FieldUUID,
		})
	}
	if value, ok := upiu.mutation.Validity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: unsavedpostimage.FieldValidity,
		})
	}
	if value, ok := upiu.mutation.Width(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: unsavedpostimage.FieldWidth,
		})
	}
	if value, ok := upiu.mutation.AddedWidth(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: unsavedpostimage.FieldWidth,
		})
	}
	if upiu.mutation.WidthCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: unsavedpostimage.FieldWidth,
		})
	}
	if value, ok := upiu.mutation.Height(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: unsavedpostimage.FieldHeight,
		})
	}
	if value, ok := upiu.mutation.AddedHeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: unsavedpostimage.FieldHeight,
		})
	}
	if upiu.mutation.HeightCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: unsavedpostimage.FieldHeight,
		})
	}
	if value, ok := upiu.mutation.Hash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostimage.FieldHash,
		})
	}
	if upiu.mutation.HashCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: unsavedpostimage.FieldHash,
		})
	}
	if value, ok := upiu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostimage.FieldTitle,
		})
	}
	if upiu.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: unsavedpostimage.FieldTitle,
		})
	}
	if value, ok := upiu.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostimage.FieldURL,
		})
	}
	if upiu.mutation.URLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: unsavedpostimage.FieldURL,
		})
	}
	if value, ok := upiu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: unsavedpostimage.FieldCreatedAt,
		})
	}
	if upiu.mutation.UnsavedPostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   unsavedpostimage.UnsavedPostTable,
			Columns: []string{unsavedpostimage.UnsavedPostColumn},
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
	if nodes := upiu.mutation.UnsavedPostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   unsavedpostimage.UnsavedPostTable,
			Columns: []string{unsavedpostimage.UnsavedPostColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, upiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{unsavedpostimage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UnsavedPostImageUpdateOne is the builder for updating a single UnsavedPostImage entity.
type UnsavedPostImageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UnsavedPostImageMutation
}

// SetUUID sets the "uuid" field.
func (upiuo *UnsavedPostImageUpdateOne) SetUUID(s string) *UnsavedPostImageUpdateOne {
	upiuo.mutation.SetUUID(s)
	return upiuo
}

// SetValidity sets the "validity" field.
func (upiuo *UnsavedPostImageUpdateOne) SetValidity(u unsavedpostimage.Validity) *UnsavedPostImageUpdateOne {
	upiuo.mutation.SetValidity(u)
	return upiuo
}

// SetNillableValidity sets the "validity" field if the given value is not nil.
func (upiuo *UnsavedPostImageUpdateOne) SetNillableValidity(u *unsavedpostimage.Validity) *UnsavedPostImageUpdateOne {
	if u != nil {
		upiuo.SetValidity(*u)
	}
	return upiuo
}

// SetWidth sets the "width" field.
func (upiuo *UnsavedPostImageUpdateOne) SetWidth(u uint32) *UnsavedPostImageUpdateOne {
	upiuo.mutation.ResetWidth()
	upiuo.mutation.SetWidth(u)
	return upiuo
}

// SetNillableWidth sets the "width" field if the given value is not nil.
func (upiuo *UnsavedPostImageUpdateOne) SetNillableWidth(u *uint32) *UnsavedPostImageUpdateOne {
	if u != nil {
		upiuo.SetWidth(*u)
	}
	return upiuo
}

// AddWidth adds u to the "width" field.
func (upiuo *UnsavedPostImageUpdateOne) AddWidth(u uint32) *UnsavedPostImageUpdateOne {
	upiuo.mutation.AddWidth(u)
	return upiuo
}

// ClearWidth clears the value of the "width" field.
func (upiuo *UnsavedPostImageUpdateOne) ClearWidth() *UnsavedPostImageUpdateOne {
	upiuo.mutation.ClearWidth()
	return upiuo
}

// SetHeight sets the "height" field.
func (upiuo *UnsavedPostImageUpdateOne) SetHeight(u uint32) *UnsavedPostImageUpdateOne {
	upiuo.mutation.ResetHeight()
	upiuo.mutation.SetHeight(u)
	return upiuo
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (upiuo *UnsavedPostImageUpdateOne) SetNillableHeight(u *uint32) *UnsavedPostImageUpdateOne {
	if u != nil {
		upiuo.SetHeight(*u)
	}
	return upiuo
}

// AddHeight adds u to the "height" field.
func (upiuo *UnsavedPostImageUpdateOne) AddHeight(u uint32) *UnsavedPostImageUpdateOne {
	upiuo.mutation.AddHeight(u)
	return upiuo
}

// ClearHeight clears the value of the "height" field.
func (upiuo *UnsavedPostImageUpdateOne) ClearHeight() *UnsavedPostImageUpdateOne {
	upiuo.mutation.ClearHeight()
	return upiuo
}

// SetHash sets the "hash" field.
func (upiuo *UnsavedPostImageUpdateOne) SetHash(s string) *UnsavedPostImageUpdateOne {
	upiuo.mutation.SetHash(s)
	return upiuo
}

// SetNillableHash sets the "hash" field if the given value is not nil.
func (upiuo *UnsavedPostImageUpdateOne) SetNillableHash(s *string) *UnsavedPostImageUpdateOne {
	if s != nil {
		upiuo.SetHash(*s)
	}
	return upiuo
}

// ClearHash clears the value of the "hash" field.
func (upiuo *UnsavedPostImageUpdateOne) ClearHash() *UnsavedPostImageUpdateOne {
	upiuo.mutation.ClearHash()
	return upiuo
}

// SetTitle sets the "title" field.
func (upiuo *UnsavedPostImageUpdateOne) SetTitle(s string) *UnsavedPostImageUpdateOne {
	upiuo.mutation.SetTitle(s)
	return upiuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (upiuo *UnsavedPostImageUpdateOne) SetNillableTitle(s *string) *UnsavedPostImageUpdateOne {
	if s != nil {
		upiuo.SetTitle(*s)
	}
	return upiuo
}

// ClearTitle clears the value of the "title" field.
func (upiuo *UnsavedPostImageUpdateOne) ClearTitle() *UnsavedPostImageUpdateOne {
	upiuo.mutation.ClearTitle()
	return upiuo
}

// SetURL sets the "url" field.
func (upiuo *UnsavedPostImageUpdateOne) SetURL(s string) *UnsavedPostImageUpdateOne {
	upiuo.mutation.SetURL(s)
	return upiuo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (upiuo *UnsavedPostImageUpdateOne) SetNillableURL(s *string) *UnsavedPostImageUpdateOne {
	if s != nil {
		upiuo.SetURL(*s)
	}
	return upiuo
}

// ClearURL clears the value of the "url" field.
func (upiuo *UnsavedPostImageUpdateOne) ClearURL() *UnsavedPostImageUpdateOne {
	upiuo.mutation.ClearURL()
	return upiuo
}

// SetCreatedAt sets the "created_at" field.
func (upiuo *UnsavedPostImageUpdateOne) SetCreatedAt(t time.Time) *UnsavedPostImageUpdateOne {
	upiuo.mutation.SetCreatedAt(t)
	return upiuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (upiuo *UnsavedPostImageUpdateOne) SetNillableCreatedAt(t *time.Time) *UnsavedPostImageUpdateOne {
	if t != nil {
		upiuo.SetCreatedAt(*t)
	}
	return upiuo
}

// SetUnsavedPostID sets the "unsaved_post" edge to the UnsavedPost entity by ID.
func (upiuo *UnsavedPostImageUpdateOne) SetUnsavedPostID(id int) *UnsavedPostImageUpdateOne {
	upiuo.mutation.SetUnsavedPostID(id)
	return upiuo
}

// SetUnsavedPost sets the "unsaved_post" edge to the UnsavedPost entity.
func (upiuo *UnsavedPostImageUpdateOne) SetUnsavedPost(u *UnsavedPost) *UnsavedPostImageUpdateOne {
	return upiuo.SetUnsavedPostID(u.ID)
}

// Mutation returns the UnsavedPostImageMutation object of the builder.
func (upiuo *UnsavedPostImageUpdateOne) Mutation() *UnsavedPostImageMutation {
	return upiuo.mutation
}

// ClearUnsavedPost clears the "unsaved_post" edge to the UnsavedPost entity.
func (upiuo *UnsavedPostImageUpdateOne) ClearUnsavedPost() *UnsavedPostImageUpdateOne {
	upiuo.mutation.ClearUnsavedPost()
	return upiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (upiuo *UnsavedPostImageUpdateOne) Select(field string, fields ...string) *UnsavedPostImageUpdateOne {
	upiuo.fields = append([]string{field}, fields...)
	return upiuo
}

// Save executes the query and returns the updated UnsavedPostImage entity.
func (upiuo *UnsavedPostImageUpdateOne) Save(ctx context.Context) (*UnsavedPostImage, error) {
	var (
		err  error
		node *UnsavedPostImage
	)
	if len(upiuo.hooks) == 0 {
		if err = upiuo.check(); err != nil {
			return nil, err
		}
		node, err = upiuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UnsavedPostImageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = upiuo.check(); err != nil {
				return nil, err
			}
			upiuo.mutation = mutation
			node, err = upiuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(upiuo.hooks) - 1; i >= 0; i-- {
			if upiuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = upiuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, upiuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (upiuo *UnsavedPostImageUpdateOne) SaveX(ctx context.Context) *UnsavedPostImage {
	node, err := upiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (upiuo *UnsavedPostImageUpdateOne) Exec(ctx context.Context) error {
	_, err := upiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upiuo *UnsavedPostImageUpdateOne) ExecX(ctx context.Context) {
	if err := upiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upiuo *UnsavedPostImageUpdateOne) check() error {
	if v, ok := upiuo.mutation.UUID(); ok {
		if err := unsavedpostimage.UUIDValidator(v); err != nil {
			return &ValidationError{Name: "uuid", err: fmt.Errorf("ent: validator failed for field \"uuid\": %w", err)}
		}
	}
	if v, ok := upiuo.mutation.Validity(); ok {
		if err := unsavedpostimage.ValidityValidator(v); err != nil {
			return &ValidationError{Name: "validity", err: fmt.Errorf("ent: validator failed for field \"validity\": %w", err)}
		}
	}
	if v, ok := upiuo.mutation.Hash(); ok {
		if err := unsavedpostimage.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf("ent: validator failed for field \"hash\": %w", err)}
		}
	}
	if v, ok := upiuo.mutation.Title(); ok {
		if err := unsavedpostimage.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if v, ok := upiuo.mutation.URL(); ok {
		if err := unsavedpostimage.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf("ent: validator failed for field \"url\": %w", err)}
		}
	}
	if _, ok := upiuo.mutation.UnsavedPostID(); upiuo.mutation.UnsavedPostCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"unsaved_post\"")
	}
	return nil
}

func (upiuo *UnsavedPostImageUpdateOne) sqlSave(ctx context.Context) (_node *UnsavedPostImage, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   unsavedpostimage.Table,
			Columns: unsavedpostimage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: unsavedpostimage.FieldID,
			},
		},
	}
	id, ok := upiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing UnsavedPostImage.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := upiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, unsavedpostimage.FieldID)
		for _, f := range fields {
			if !unsavedpostimage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != unsavedpostimage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := upiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := upiuo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostimage.FieldUUID,
		})
	}
	if value, ok := upiuo.mutation.Validity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: unsavedpostimage.FieldValidity,
		})
	}
	if value, ok := upiuo.mutation.Width(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: unsavedpostimage.FieldWidth,
		})
	}
	if value, ok := upiuo.mutation.AddedWidth(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: unsavedpostimage.FieldWidth,
		})
	}
	if upiuo.mutation.WidthCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: unsavedpostimage.FieldWidth,
		})
	}
	if value, ok := upiuo.mutation.Height(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: unsavedpostimage.FieldHeight,
		})
	}
	if value, ok := upiuo.mutation.AddedHeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: unsavedpostimage.FieldHeight,
		})
	}
	if upiuo.mutation.HeightCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: unsavedpostimage.FieldHeight,
		})
	}
	if value, ok := upiuo.mutation.Hash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostimage.FieldHash,
		})
	}
	if upiuo.mutation.HashCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: unsavedpostimage.FieldHash,
		})
	}
	if value, ok := upiuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostimage.FieldTitle,
		})
	}
	if upiuo.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: unsavedpostimage.FieldTitle,
		})
	}
	if value, ok := upiuo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostimage.FieldURL,
		})
	}
	if upiuo.mutation.URLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: unsavedpostimage.FieldURL,
		})
	}
	if value, ok := upiuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: unsavedpostimage.FieldCreatedAt,
		})
	}
	if upiuo.mutation.UnsavedPostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   unsavedpostimage.UnsavedPostTable,
			Columns: []string{unsavedpostimage.UnsavedPostColumn},
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
	if nodes := upiuo.mutation.UnsavedPostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   unsavedpostimage.UnsavedPostTable,
			Columns: []string{unsavedpostimage.UnsavedPostColumn},
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
	_node = &UnsavedPostImage{config: upiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, upiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{unsavedpostimage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
