// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devlog/ent/predicate"
	"devlog/ent/unsavedpost"
	"devlog/ent/unsavedpostattachment"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UnsavedPostAttachmentUpdate is the builder for updating UnsavedPostAttachment entities.
type UnsavedPostAttachmentUpdate struct {
	config
	hooks    []Hook
	mutation *UnsavedPostAttachmentMutation
}

// Where adds a new predicate for the UnsavedPostAttachmentUpdate builder.
func (upau *UnsavedPostAttachmentUpdate) Where(ps ...predicate.UnsavedPostAttachment) *UnsavedPostAttachmentUpdate {
	upau.mutation.predicates = append(upau.mutation.predicates, ps...)
	return upau
}

// SetUUID sets the "uuid" field.
func (upau *UnsavedPostAttachmentUpdate) SetUUID(s string) *UnsavedPostAttachmentUpdate {
	upau.mutation.SetUUID(s)
	return upau
}

// SetValidity sets the "validity" field.
func (upau *UnsavedPostAttachmentUpdate) SetValidity(u unsavedpostattachment.Validity) *UnsavedPostAttachmentUpdate {
	upau.mutation.SetValidity(u)
	return upau
}

// SetNillableValidity sets the "validity" field if the given value is not nil.
func (upau *UnsavedPostAttachmentUpdate) SetNillableValidity(u *unsavedpostattachment.Validity) *UnsavedPostAttachmentUpdate {
	if u != nil {
		upau.SetValidity(*u)
	}
	return upau
}

// SetSize sets the "size" field.
func (upau *UnsavedPostAttachmentUpdate) SetSize(u uint64) *UnsavedPostAttachmentUpdate {
	upau.mutation.ResetSize()
	upau.mutation.SetSize(u)
	return upau
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (upau *UnsavedPostAttachmentUpdate) SetNillableSize(u *uint64) *UnsavedPostAttachmentUpdate {
	if u != nil {
		upau.SetSize(*u)
	}
	return upau
}

// AddSize adds u to the "size" field.
func (upau *UnsavedPostAttachmentUpdate) AddSize(u uint64) *UnsavedPostAttachmentUpdate {
	upau.mutation.AddSize(u)
	return upau
}

// ClearSize clears the value of the "size" field.
func (upau *UnsavedPostAttachmentUpdate) ClearSize() *UnsavedPostAttachmentUpdate {
	upau.mutation.ClearSize()
	return upau
}

// SetName sets the "name" field.
func (upau *UnsavedPostAttachmentUpdate) SetName(s string) *UnsavedPostAttachmentUpdate {
	upau.mutation.SetName(s)
	return upau
}

// SetNillableName sets the "name" field if the given value is not nil.
func (upau *UnsavedPostAttachmentUpdate) SetNillableName(s *string) *UnsavedPostAttachmentUpdate {
	if s != nil {
		upau.SetName(*s)
	}
	return upau
}

// ClearName clears the value of the "name" field.
func (upau *UnsavedPostAttachmentUpdate) ClearName() *UnsavedPostAttachmentUpdate {
	upau.mutation.ClearName()
	return upau
}

// SetMime sets the "mime" field.
func (upau *UnsavedPostAttachmentUpdate) SetMime(s string) *UnsavedPostAttachmentUpdate {
	upau.mutation.SetMime(s)
	return upau
}

// SetNillableMime sets the "mime" field if the given value is not nil.
func (upau *UnsavedPostAttachmentUpdate) SetNillableMime(s *string) *UnsavedPostAttachmentUpdate {
	if s != nil {
		upau.SetMime(*s)
	}
	return upau
}

// ClearMime clears the value of the "mime" field.
func (upau *UnsavedPostAttachmentUpdate) ClearMime() *UnsavedPostAttachmentUpdate {
	upau.mutation.ClearMime()
	return upau
}

// SetURL sets the "url" field.
func (upau *UnsavedPostAttachmentUpdate) SetURL(s string) *UnsavedPostAttachmentUpdate {
	upau.mutation.SetURL(s)
	return upau
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (upau *UnsavedPostAttachmentUpdate) SetNillableURL(s *string) *UnsavedPostAttachmentUpdate {
	if s != nil {
		upau.SetURL(*s)
	}
	return upau
}

// ClearURL clears the value of the "url" field.
func (upau *UnsavedPostAttachmentUpdate) ClearURL() *UnsavedPostAttachmentUpdate {
	upau.mutation.ClearURL()
	return upau
}

// SetCreatedAt sets the "created_at" field.
func (upau *UnsavedPostAttachmentUpdate) SetCreatedAt(t time.Time) *UnsavedPostAttachmentUpdate {
	upau.mutation.SetCreatedAt(t)
	return upau
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (upau *UnsavedPostAttachmentUpdate) SetNillableCreatedAt(t *time.Time) *UnsavedPostAttachmentUpdate {
	if t != nil {
		upau.SetCreatedAt(*t)
	}
	return upau
}

// SetUnsavedPostID sets the "unsaved_post" edge to the UnsavedPost entity by ID.
func (upau *UnsavedPostAttachmentUpdate) SetUnsavedPostID(id int) *UnsavedPostAttachmentUpdate {
	upau.mutation.SetUnsavedPostID(id)
	return upau
}

// SetUnsavedPost sets the "unsaved_post" edge to the UnsavedPost entity.
func (upau *UnsavedPostAttachmentUpdate) SetUnsavedPost(u *UnsavedPost) *UnsavedPostAttachmentUpdate {
	return upau.SetUnsavedPostID(u.ID)
}

// Mutation returns the UnsavedPostAttachmentMutation object of the builder.
func (upau *UnsavedPostAttachmentUpdate) Mutation() *UnsavedPostAttachmentMutation {
	return upau.mutation
}

// ClearUnsavedPost clears the "unsaved_post" edge to the UnsavedPost entity.
func (upau *UnsavedPostAttachmentUpdate) ClearUnsavedPost() *UnsavedPostAttachmentUpdate {
	upau.mutation.ClearUnsavedPost()
	return upau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (upau *UnsavedPostAttachmentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(upau.hooks) == 0 {
		if err = upau.check(); err != nil {
			return 0, err
		}
		affected, err = upau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UnsavedPostAttachmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = upau.check(); err != nil {
				return 0, err
			}
			upau.mutation = mutation
			affected, err = upau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(upau.hooks) - 1; i >= 0; i-- {
			mut = upau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, upau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (upau *UnsavedPostAttachmentUpdate) SaveX(ctx context.Context) int {
	affected, err := upau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (upau *UnsavedPostAttachmentUpdate) Exec(ctx context.Context) error {
	_, err := upau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upau *UnsavedPostAttachmentUpdate) ExecX(ctx context.Context) {
	if err := upau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upau *UnsavedPostAttachmentUpdate) check() error {
	if v, ok := upau.mutation.UUID(); ok {
		if err := unsavedpostattachment.UUIDValidator(v); err != nil {
			return &ValidationError{Name: "uuid", err: fmt.Errorf("ent: validator failed for field \"uuid\": %w", err)}
		}
	}
	if v, ok := upau.mutation.Validity(); ok {
		if err := unsavedpostattachment.ValidityValidator(v); err != nil {
			return &ValidationError{Name: "validity", err: fmt.Errorf("ent: validator failed for field \"validity\": %w", err)}
		}
	}
	if v, ok := upau.mutation.Name(); ok {
		if err := unsavedpostattachment.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := upau.mutation.Mime(); ok {
		if err := unsavedpostattachment.MimeValidator(v); err != nil {
			return &ValidationError{Name: "mime", err: fmt.Errorf("ent: validator failed for field \"mime\": %w", err)}
		}
	}
	if v, ok := upau.mutation.URL(); ok {
		if err := unsavedpostattachment.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf("ent: validator failed for field \"url\": %w", err)}
		}
	}
	if _, ok := upau.mutation.UnsavedPostID(); upau.mutation.UnsavedPostCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"unsaved_post\"")
	}
	return nil
}

func (upau *UnsavedPostAttachmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   unsavedpostattachment.Table,
			Columns: unsavedpostattachment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: unsavedpostattachment.FieldID,
			},
		},
	}
	if ps := upau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := upau.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostattachment.FieldUUID,
		})
	}
	if value, ok := upau.mutation.Validity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: unsavedpostattachment.FieldValidity,
		})
	}
	if value, ok := upau.mutation.Size(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: unsavedpostattachment.FieldSize,
		})
	}
	if value, ok := upau.mutation.AddedSize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: unsavedpostattachment.FieldSize,
		})
	}
	if upau.mutation.SizeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Column: unsavedpostattachment.FieldSize,
		})
	}
	if value, ok := upau.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostattachment.FieldName,
		})
	}
	if upau.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: unsavedpostattachment.FieldName,
		})
	}
	if value, ok := upau.mutation.Mime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostattachment.FieldMime,
		})
	}
	if upau.mutation.MimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: unsavedpostattachment.FieldMime,
		})
	}
	if value, ok := upau.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostattachment.FieldURL,
		})
	}
	if upau.mutation.URLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: unsavedpostattachment.FieldURL,
		})
	}
	if value, ok := upau.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: unsavedpostattachment.FieldCreatedAt,
		})
	}
	if upau.mutation.UnsavedPostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   unsavedpostattachment.UnsavedPostTable,
			Columns: []string{unsavedpostattachment.UnsavedPostColumn},
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
	if nodes := upau.mutation.UnsavedPostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   unsavedpostattachment.UnsavedPostTable,
			Columns: []string{unsavedpostattachment.UnsavedPostColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, upau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{unsavedpostattachment.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UnsavedPostAttachmentUpdateOne is the builder for updating a single UnsavedPostAttachment entity.
type UnsavedPostAttachmentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UnsavedPostAttachmentMutation
}

// SetUUID sets the "uuid" field.
func (upauo *UnsavedPostAttachmentUpdateOne) SetUUID(s string) *UnsavedPostAttachmentUpdateOne {
	upauo.mutation.SetUUID(s)
	return upauo
}

// SetValidity sets the "validity" field.
func (upauo *UnsavedPostAttachmentUpdateOne) SetValidity(u unsavedpostattachment.Validity) *UnsavedPostAttachmentUpdateOne {
	upauo.mutation.SetValidity(u)
	return upauo
}

// SetNillableValidity sets the "validity" field if the given value is not nil.
func (upauo *UnsavedPostAttachmentUpdateOne) SetNillableValidity(u *unsavedpostattachment.Validity) *UnsavedPostAttachmentUpdateOne {
	if u != nil {
		upauo.SetValidity(*u)
	}
	return upauo
}

// SetSize sets the "size" field.
func (upauo *UnsavedPostAttachmentUpdateOne) SetSize(u uint64) *UnsavedPostAttachmentUpdateOne {
	upauo.mutation.ResetSize()
	upauo.mutation.SetSize(u)
	return upauo
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (upauo *UnsavedPostAttachmentUpdateOne) SetNillableSize(u *uint64) *UnsavedPostAttachmentUpdateOne {
	if u != nil {
		upauo.SetSize(*u)
	}
	return upauo
}

// AddSize adds u to the "size" field.
func (upauo *UnsavedPostAttachmentUpdateOne) AddSize(u uint64) *UnsavedPostAttachmentUpdateOne {
	upauo.mutation.AddSize(u)
	return upauo
}

// ClearSize clears the value of the "size" field.
func (upauo *UnsavedPostAttachmentUpdateOne) ClearSize() *UnsavedPostAttachmentUpdateOne {
	upauo.mutation.ClearSize()
	return upauo
}

// SetName sets the "name" field.
func (upauo *UnsavedPostAttachmentUpdateOne) SetName(s string) *UnsavedPostAttachmentUpdateOne {
	upauo.mutation.SetName(s)
	return upauo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (upauo *UnsavedPostAttachmentUpdateOne) SetNillableName(s *string) *UnsavedPostAttachmentUpdateOne {
	if s != nil {
		upauo.SetName(*s)
	}
	return upauo
}

// ClearName clears the value of the "name" field.
func (upauo *UnsavedPostAttachmentUpdateOne) ClearName() *UnsavedPostAttachmentUpdateOne {
	upauo.mutation.ClearName()
	return upauo
}

// SetMime sets the "mime" field.
func (upauo *UnsavedPostAttachmentUpdateOne) SetMime(s string) *UnsavedPostAttachmentUpdateOne {
	upauo.mutation.SetMime(s)
	return upauo
}

// SetNillableMime sets the "mime" field if the given value is not nil.
func (upauo *UnsavedPostAttachmentUpdateOne) SetNillableMime(s *string) *UnsavedPostAttachmentUpdateOne {
	if s != nil {
		upauo.SetMime(*s)
	}
	return upauo
}

// ClearMime clears the value of the "mime" field.
func (upauo *UnsavedPostAttachmentUpdateOne) ClearMime() *UnsavedPostAttachmentUpdateOne {
	upauo.mutation.ClearMime()
	return upauo
}

// SetURL sets the "url" field.
func (upauo *UnsavedPostAttachmentUpdateOne) SetURL(s string) *UnsavedPostAttachmentUpdateOne {
	upauo.mutation.SetURL(s)
	return upauo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (upauo *UnsavedPostAttachmentUpdateOne) SetNillableURL(s *string) *UnsavedPostAttachmentUpdateOne {
	if s != nil {
		upauo.SetURL(*s)
	}
	return upauo
}

// ClearURL clears the value of the "url" field.
func (upauo *UnsavedPostAttachmentUpdateOne) ClearURL() *UnsavedPostAttachmentUpdateOne {
	upauo.mutation.ClearURL()
	return upauo
}

// SetCreatedAt sets the "created_at" field.
func (upauo *UnsavedPostAttachmentUpdateOne) SetCreatedAt(t time.Time) *UnsavedPostAttachmentUpdateOne {
	upauo.mutation.SetCreatedAt(t)
	return upauo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (upauo *UnsavedPostAttachmentUpdateOne) SetNillableCreatedAt(t *time.Time) *UnsavedPostAttachmentUpdateOne {
	if t != nil {
		upauo.SetCreatedAt(*t)
	}
	return upauo
}

// SetUnsavedPostID sets the "unsaved_post" edge to the UnsavedPost entity by ID.
func (upauo *UnsavedPostAttachmentUpdateOne) SetUnsavedPostID(id int) *UnsavedPostAttachmentUpdateOne {
	upauo.mutation.SetUnsavedPostID(id)
	return upauo
}

// SetUnsavedPost sets the "unsaved_post" edge to the UnsavedPost entity.
func (upauo *UnsavedPostAttachmentUpdateOne) SetUnsavedPost(u *UnsavedPost) *UnsavedPostAttachmentUpdateOne {
	return upauo.SetUnsavedPostID(u.ID)
}

// Mutation returns the UnsavedPostAttachmentMutation object of the builder.
func (upauo *UnsavedPostAttachmentUpdateOne) Mutation() *UnsavedPostAttachmentMutation {
	return upauo.mutation
}

// ClearUnsavedPost clears the "unsaved_post" edge to the UnsavedPost entity.
func (upauo *UnsavedPostAttachmentUpdateOne) ClearUnsavedPost() *UnsavedPostAttachmentUpdateOne {
	upauo.mutation.ClearUnsavedPost()
	return upauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (upauo *UnsavedPostAttachmentUpdateOne) Select(field string, fields ...string) *UnsavedPostAttachmentUpdateOne {
	upauo.fields = append([]string{field}, fields...)
	return upauo
}

// Save executes the query and returns the updated UnsavedPostAttachment entity.
func (upauo *UnsavedPostAttachmentUpdateOne) Save(ctx context.Context) (*UnsavedPostAttachment, error) {
	var (
		err  error
		node *UnsavedPostAttachment
	)
	if len(upauo.hooks) == 0 {
		if err = upauo.check(); err != nil {
			return nil, err
		}
		node, err = upauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UnsavedPostAttachmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = upauo.check(); err != nil {
				return nil, err
			}
			upauo.mutation = mutation
			node, err = upauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(upauo.hooks) - 1; i >= 0; i-- {
			mut = upauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, upauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (upauo *UnsavedPostAttachmentUpdateOne) SaveX(ctx context.Context) *UnsavedPostAttachment {
	node, err := upauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (upauo *UnsavedPostAttachmentUpdateOne) Exec(ctx context.Context) error {
	_, err := upauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upauo *UnsavedPostAttachmentUpdateOne) ExecX(ctx context.Context) {
	if err := upauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upauo *UnsavedPostAttachmentUpdateOne) check() error {
	if v, ok := upauo.mutation.UUID(); ok {
		if err := unsavedpostattachment.UUIDValidator(v); err != nil {
			return &ValidationError{Name: "uuid", err: fmt.Errorf("ent: validator failed for field \"uuid\": %w", err)}
		}
	}
	if v, ok := upauo.mutation.Validity(); ok {
		if err := unsavedpostattachment.ValidityValidator(v); err != nil {
			return &ValidationError{Name: "validity", err: fmt.Errorf("ent: validator failed for field \"validity\": %w", err)}
		}
	}
	if v, ok := upauo.mutation.Name(); ok {
		if err := unsavedpostattachment.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := upauo.mutation.Mime(); ok {
		if err := unsavedpostattachment.MimeValidator(v); err != nil {
			return &ValidationError{Name: "mime", err: fmt.Errorf("ent: validator failed for field \"mime\": %w", err)}
		}
	}
	if v, ok := upauo.mutation.URL(); ok {
		if err := unsavedpostattachment.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf("ent: validator failed for field \"url\": %w", err)}
		}
	}
	if _, ok := upauo.mutation.UnsavedPostID(); upauo.mutation.UnsavedPostCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"unsaved_post\"")
	}
	return nil
}

func (upauo *UnsavedPostAttachmentUpdateOne) sqlSave(ctx context.Context) (_node *UnsavedPostAttachment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   unsavedpostattachment.Table,
			Columns: unsavedpostattachment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: unsavedpostattachment.FieldID,
			},
		},
	}
	id, ok := upauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing UnsavedPostAttachment.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := upauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, unsavedpostattachment.FieldID)
		for _, f := range fields {
			if !unsavedpostattachment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != unsavedpostattachment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := upauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := upauo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostattachment.FieldUUID,
		})
	}
	if value, ok := upauo.mutation.Validity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: unsavedpostattachment.FieldValidity,
		})
	}
	if value, ok := upauo.mutation.Size(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: unsavedpostattachment.FieldSize,
		})
	}
	if value, ok := upauo.mutation.AddedSize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: unsavedpostattachment.FieldSize,
		})
	}
	if upauo.mutation.SizeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Column: unsavedpostattachment.FieldSize,
		})
	}
	if value, ok := upauo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostattachment.FieldName,
		})
	}
	if upauo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: unsavedpostattachment.FieldName,
		})
	}
	if value, ok := upauo.mutation.Mime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostattachment.FieldMime,
		})
	}
	if upauo.mutation.MimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: unsavedpostattachment.FieldMime,
		})
	}
	if value, ok := upauo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unsavedpostattachment.FieldURL,
		})
	}
	if upauo.mutation.URLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: unsavedpostattachment.FieldURL,
		})
	}
	if value, ok := upauo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: unsavedpostattachment.FieldCreatedAt,
		})
	}
	if upauo.mutation.UnsavedPostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   unsavedpostattachment.UnsavedPostTable,
			Columns: []string{unsavedpostattachment.UnsavedPostColumn},
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
	if nodes := upauo.mutation.UnsavedPostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   unsavedpostattachment.UnsavedPostTable,
			Columns: []string{unsavedpostattachment.UnsavedPostColumn},
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
	_node = &UnsavedPostAttachment{config: upauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, upauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{unsavedpostattachment.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
