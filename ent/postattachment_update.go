// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devlog/ent/post"
	"devlog/ent/postattachment"
	"devlog/ent/predicate"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PostAttachmentUpdate is the builder for updating PostAttachment entities.
type PostAttachmentUpdate struct {
	config
	hooks    []Hook
	mutation *PostAttachmentMutation
}

// Where adds a new predicate for the PostAttachmentUpdate builder.
func (pau *PostAttachmentUpdate) Where(ps ...predicate.PostAttachment) *PostAttachmentUpdate {
	pau.mutation.predicates = append(pau.mutation.predicates, ps...)
	return pau
}

// SetUUID sets the "uuid" field.
func (pau *PostAttachmentUpdate) SetUUID(s string) *PostAttachmentUpdate {
	pau.mutation.SetUUID(s)
	return pau
}

// SetSize sets the "size" field.
func (pau *PostAttachmentUpdate) SetSize(u uint64) *PostAttachmentUpdate {
	pau.mutation.ResetSize()
	pau.mutation.SetSize(u)
	return pau
}

// AddSize adds u to the "size" field.
func (pau *PostAttachmentUpdate) AddSize(u uint64) *PostAttachmentUpdate {
	pau.mutation.AddSize(u)
	return pau
}

// SetName sets the "name" field.
func (pau *PostAttachmentUpdate) SetName(s string) *PostAttachmentUpdate {
	pau.mutation.SetName(s)
	return pau
}

// SetMime sets the "mime" field.
func (pau *PostAttachmentUpdate) SetMime(s string) *PostAttachmentUpdate {
	pau.mutation.SetMime(s)
	return pau
}

// SetURL sets the "url" field.
func (pau *PostAttachmentUpdate) SetURL(s string) *PostAttachmentUpdate {
	pau.mutation.SetURL(s)
	return pau
}

// SetCreatedAt sets the "created_at" field.
func (pau *PostAttachmentUpdate) SetCreatedAt(t time.Time) *PostAttachmentUpdate {
	pau.mutation.SetCreatedAt(t)
	return pau
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pau *PostAttachmentUpdate) SetNillableCreatedAt(t *time.Time) *PostAttachmentUpdate {
	if t != nil {
		pau.SetCreatedAt(*t)
	}
	return pau
}

// SetPostID sets the "post" edge to the Post entity by ID.
func (pau *PostAttachmentUpdate) SetPostID(id int) *PostAttachmentUpdate {
	pau.mutation.SetPostID(id)
	return pau
}

// SetPost sets the "post" edge to the Post entity.
func (pau *PostAttachmentUpdate) SetPost(p *Post) *PostAttachmentUpdate {
	return pau.SetPostID(p.ID)
}

// Mutation returns the PostAttachmentMutation object of the builder.
func (pau *PostAttachmentUpdate) Mutation() *PostAttachmentMutation {
	return pau.mutation
}

// ClearPost clears the "post" edge to the Post entity.
func (pau *PostAttachmentUpdate) ClearPost() *PostAttachmentUpdate {
	pau.mutation.ClearPost()
	return pau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pau *PostAttachmentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pau.hooks) == 0 {
		if err = pau.check(); err != nil {
			return 0, err
		}
		affected, err = pau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PostAttachmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pau.check(); err != nil {
				return 0, err
			}
			pau.mutation = mutation
			affected, err = pau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pau.hooks) - 1; i >= 0; i-- {
			mut = pau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pau *PostAttachmentUpdate) SaveX(ctx context.Context) int {
	affected, err := pau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pau *PostAttachmentUpdate) Exec(ctx context.Context) error {
	_, err := pau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pau *PostAttachmentUpdate) ExecX(ctx context.Context) {
	if err := pau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pau *PostAttachmentUpdate) check() error {
	if v, ok := pau.mutation.UUID(); ok {
		if err := postattachment.UUIDValidator(v); err != nil {
			return &ValidationError{Name: "uuid", err: fmt.Errorf("ent: validator failed for field \"uuid\": %w", err)}
		}
	}
	if v, ok := pau.mutation.Name(); ok {
		if err := postattachment.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := pau.mutation.Mime(); ok {
		if err := postattachment.MimeValidator(v); err != nil {
			return &ValidationError{Name: "mime", err: fmt.Errorf("ent: validator failed for field \"mime\": %w", err)}
		}
	}
	if v, ok := pau.mutation.URL(); ok {
		if err := postattachment.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf("ent: validator failed for field \"url\": %w", err)}
		}
	}
	if _, ok := pau.mutation.PostID(); pau.mutation.PostCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"post\"")
	}
	return nil
}

func (pau *PostAttachmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   postattachment.Table,
			Columns: postattachment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: postattachment.FieldID,
			},
		},
	}
	if ps := pau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pau.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postattachment.FieldUUID,
		})
	}
	if value, ok := pau.mutation.Size(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: postattachment.FieldSize,
		})
	}
	if value, ok := pau.mutation.AddedSize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: postattachment.FieldSize,
		})
	}
	if value, ok := pau.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postattachment.FieldName,
		})
	}
	if value, ok := pau.mutation.Mime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postattachment.FieldMime,
		})
	}
	if value, ok := pau.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postattachment.FieldURL,
		})
	}
	if value, ok := pau.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: postattachment.FieldCreatedAt,
		})
	}
	if pau.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   postattachment.PostTable,
			Columns: []string{postattachment.PostColumn},
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
	if nodes := pau.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   postattachment.PostTable,
			Columns: []string{postattachment.PostColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, pau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{postattachment.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PostAttachmentUpdateOne is the builder for updating a single PostAttachment entity.
type PostAttachmentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PostAttachmentMutation
}

// SetUUID sets the "uuid" field.
func (pauo *PostAttachmentUpdateOne) SetUUID(s string) *PostAttachmentUpdateOne {
	pauo.mutation.SetUUID(s)
	return pauo
}

// SetSize sets the "size" field.
func (pauo *PostAttachmentUpdateOne) SetSize(u uint64) *PostAttachmentUpdateOne {
	pauo.mutation.ResetSize()
	pauo.mutation.SetSize(u)
	return pauo
}

// AddSize adds u to the "size" field.
func (pauo *PostAttachmentUpdateOne) AddSize(u uint64) *PostAttachmentUpdateOne {
	pauo.mutation.AddSize(u)
	return pauo
}

// SetName sets the "name" field.
func (pauo *PostAttachmentUpdateOne) SetName(s string) *PostAttachmentUpdateOne {
	pauo.mutation.SetName(s)
	return pauo
}

// SetMime sets the "mime" field.
func (pauo *PostAttachmentUpdateOne) SetMime(s string) *PostAttachmentUpdateOne {
	pauo.mutation.SetMime(s)
	return pauo
}

// SetURL sets the "url" field.
func (pauo *PostAttachmentUpdateOne) SetURL(s string) *PostAttachmentUpdateOne {
	pauo.mutation.SetURL(s)
	return pauo
}

// SetCreatedAt sets the "created_at" field.
func (pauo *PostAttachmentUpdateOne) SetCreatedAt(t time.Time) *PostAttachmentUpdateOne {
	pauo.mutation.SetCreatedAt(t)
	return pauo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pauo *PostAttachmentUpdateOne) SetNillableCreatedAt(t *time.Time) *PostAttachmentUpdateOne {
	if t != nil {
		pauo.SetCreatedAt(*t)
	}
	return pauo
}

// SetPostID sets the "post" edge to the Post entity by ID.
func (pauo *PostAttachmentUpdateOne) SetPostID(id int) *PostAttachmentUpdateOne {
	pauo.mutation.SetPostID(id)
	return pauo
}

// SetPost sets the "post" edge to the Post entity.
func (pauo *PostAttachmentUpdateOne) SetPost(p *Post) *PostAttachmentUpdateOne {
	return pauo.SetPostID(p.ID)
}

// Mutation returns the PostAttachmentMutation object of the builder.
func (pauo *PostAttachmentUpdateOne) Mutation() *PostAttachmentMutation {
	return pauo.mutation
}

// ClearPost clears the "post" edge to the Post entity.
func (pauo *PostAttachmentUpdateOne) ClearPost() *PostAttachmentUpdateOne {
	pauo.mutation.ClearPost()
	return pauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pauo *PostAttachmentUpdateOne) Select(field string, fields ...string) *PostAttachmentUpdateOne {
	pauo.fields = append([]string{field}, fields...)
	return pauo
}

// Save executes the query and returns the updated PostAttachment entity.
func (pauo *PostAttachmentUpdateOne) Save(ctx context.Context) (*PostAttachment, error) {
	var (
		err  error
		node *PostAttachment
	)
	if len(pauo.hooks) == 0 {
		if err = pauo.check(); err != nil {
			return nil, err
		}
		node, err = pauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PostAttachmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pauo.check(); err != nil {
				return nil, err
			}
			pauo.mutation = mutation
			node, err = pauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pauo.hooks) - 1; i >= 0; i-- {
			mut = pauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pauo *PostAttachmentUpdateOne) SaveX(ctx context.Context) *PostAttachment {
	node, err := pauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pauo *PostAttachmentUpdateOne) Exec(ctx context.Context) error {
	_, err := pauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pauo *PostAttachmentUpdateOne) ExecX(ctx context.Context) {
	if err := pauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pauo *PostAttachmentUpdateOne) check() error {
	if v, ok := pauo.mutation.UUID(); ok {
		if err := postattachment.UUIDValidator(v); err != nil {
			return &ValidationError{Name: "uuid", err: fmt.Errorf("ent: validator failed for field \"uuid\": %w", err)}
		}
	}
	if v, ok := pauo.mutation.Name(); ok {
		if err := postattachment.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := pauo.mutation.Mime(); ok {
		if err := postattachment.MimeValidator(v); err != nil {
			return &ValidationError{Name: "mime", err: fmt.Errorf("ent: validator failed for field \"mime\": %w", err)}
		}
	}
	if v, ok := pauo.mutation.URL(); ok {
		if err := postattachment.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf("ent: validator failed for field \"url\": %w", err)}
		}
	}
	if _, ok := pauo.mutation.PostID(); pauo.mutation.PostCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"post\"")
	}
	return nil
}

func (pauo *PostAttachmentUpdateOne) sqlSave(ctx context.Context) (_node *PostAttachment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   postattachment.Table,
			Columns: postattachment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: postattachment.FieldID,
			},
		},
	}
	id, ok := pauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing PostAttachment.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := pauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, postattachment.FieldID)
		for _, f := range fields {
			if !postattachment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != postattachment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pauo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postattachment.FieldUUID,
		})
	}
	if value, ok := pauo.mutation.Size(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: postattachment.FieldSize,
		})
	}
	if value, ok := pauo.mutation.AddedSize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: postattachment.FieldSize,
		})
	}
	if value, ok := pauo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postattachment.FieldName,
		})
	}
	if value, ok := pauo.mutation.Mime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postattachment.FieldMime,
		})
	}
	if value, ok := pauo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: postattachment.FieldURL,
		})
	}
	if value, ok := pauo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: postattachment.FieldCreatedAt,
		})
	}
	if pauo.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   postattachment.PostTable,
			Columns: []string{postattachment.PostColumn},
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
	if nodes := pauo.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   postattachment.PostTable,
			Columns: []string{postattachment.PostColumn},
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
	_node = &PostAttachment{config: pauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{postattachment.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}