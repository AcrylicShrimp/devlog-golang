// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devlog/ent/admin"
	"devlog/ent/adminsession"
	"devlog/ent/predicate"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// AdminUpdate is the builder for updating Admin entities.
type AdminUpdate struct {
	config
	hooks    []Hook
	mutation *AdminMutation
}

// Where adds a new predicate for the AdminUpdate builder.
func (au *AdminUpdate) Where(ps ...predicate.Admin) *AdminUpdate {
	au.mutation.predicates = append(au.mutation.predicates, ps...)
	return au
}

// SetEmail sets the "email" field.
func (au *AdminUpdate) SetEmail(s string) *AdminUpdate {
	au.mutation.SetEmail(s)
	return au
}

// SetUsername sets the "username" field.
func (au *AdminUpdate) SetUsername(s string) *AdminUpdate {
	au.mutation.SetUsername(s)
	return au
}

// SetPassword sets the "password" field.
func (au *AdminUpdate) SetPassword(s string) *AdminUpdate {
	au.mutation.SetPassword(s)
	return au
}

// SetJoinedAt sets the "joined_at" field.
func (au *AdminUpdate) SetJoinedAt(t time.Time) *AdminUpdate {
	au.mutation.SetJoinedAt(t)
	return au
}

// SetNillableJoinedAt sets the "joined_at" field if the given value is not nil.
func (au *AdminUpdate) SetNillableJoinedAt(t *time.Time) *AdminUpdate {
	if t != nil {
		au.SetJoinedAt(*t)
	}
	return au
}

// AddSessionIDs adds the "sessions" edge to the AdminSession entity by IDs.
func (au *AdminUpdate) AddSessionIDs(ids ...int) *AdminUpdate {
	au.mutation.AddSessionIDs(ids...)
	return au
}

// AddSessions adds the "sessions" edges to the AdminSession entity.
func (au *AdminUpdate) AddSessions(a ...*AdminSession) *AdminUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.AddSessionIDs(ids...)
}

// Mutation returns the AdminMutation object of the builder.
func (au *AdminUpdate) Mutation() *AdminMutation {
	return au.mutation
}

// ClearSessions clears all "sessions" edges to the AdminSession entity.
func (au *AdminUpdate) ClearSessions() *AdminUpdate {
	au.mutation.ClearSessions()
	return au
}

// RemoveSessionIDs removes the "sessions" edge to AdminSession entities by IDs.
func (au *AdminUpdate) RemoveSessionIDs(ids ...int) *AdminUpdate {
	au.mutation.RemoveSessionIDs(ids...)
	return au
}

// RemoveSessions removes "sessions" edges to AdminSession entities.
func (au *AdminUpdate) RemoveSessions(a ...*AdminSession) *AdminUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.RemoveSessionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AdminUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		if err = au.check(); err != nil {
			return 0, err
		}
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdminMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = au.check(); err != nil {
				return 0, err
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AdminUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AdminUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AdminUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AdminUpdate) check() error {
	if v, ok := au.mutation.Email(); ok {
		if err := admin.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	if v, ok := au.mutation.Username(); ok {
		if err := admin.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf("ent: validator failed for field \"username\": %w", err)}
		}
	}
	if v, ok := au.mutation.Password(); ok {
		if err := admin.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}
	return nil
}

func (au *AdminUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   admin.Table,
			Columns: admin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: admin.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admin.FieldEmail,
		})
	}
	if value, ok := au.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admin.FieldUsername,
		})
	}
	if value, ok := au.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admin.FieldPassword,
		})
	}
	if value, ok := au.mutation.JoinedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: admin.FieldJoinedAt,
		})
	}
	if au.mutation.SessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   admin.SessionsTable,
			Columns: []string{admin.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adminsession.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedSessionsIDs(); len(nodes) > 0 && !au.mutation.SessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   admin.SessionsTable,
			Columns: []string{admin.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adminsession.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.SessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   admin.SessionsTable,
			Columns: []string{admin.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adminsession.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{admin.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AdminUpdateOne is the builder for updating a single Admin entity.
type AdminUpdateOne struct {
	config
	hooks    []Hook
	mutation *AdminMutation
}

// SetEmail sets the "email" field.
func (auo *AdminUpdateOne) SetEmail(s string) *AdminUpdateOne {
	auo.mutation.SetEmail(s)
	return auo
}

// SetUsername sets the "username" field.
func (auo *AdminUpdateOne) SetUsername(s string) *AdminUpdateOne {
	auo.mutation.SetUsername(s)
	return auo
}

// SetPassword sets the "password" field.
func (auo *AdminUpdateOne) SetPassword(s string) *AdminUpdateOne {
	auo.mutation.SetPassword(s)
	return auo
}

// SetJoinedAt sets the "joined_at" field.
func (auo *AdminUpdateOne) SetJoinedAt(t time.Time) *AdminUpdateOne {
	auo.mutation.SetJoinedAt(t)
	return auo
}

// SetNillableJoinedAt sets the "joined_at" field if the given value is not nil.
func (auo *AdminUpdateOne) SetNillableJoinedAt(t *time.Time) *AdminUpdateOne {
	if t != nil {
		auo.SetJoinedAt(*t)
	}
	return auo
}

// AddSessionIDs adds the "sessions" edge to the AdminSession entity by IDs.
func (auo *AdminUpdateOne) AddSessionIDs(ids ...int) *AdminUpdateOne {
	auo.mutation.AddSessionIDs(ids...)
	return auo
}

// AddSessions adds the "sessions" edges to the AdminSession entity.
func (auo *AdminUpdateOne) AddSessions(a ...*AdminSession) *AdminUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.AddSessionIDs(ids...)
}

// Mutation returns the AdminMutation object of the builder.
func (auo *AdminUpdateOne) Mutation() *AdminMutation {
	return auo.mutation
}

// ClearSessions clears all "sessions" edges to the AdminSession entity.
func (auo *AdminUpdateOne) ClearSessions() *AdminUpdateOne {
	auo.mutation.ClearSessions()
	return auo
}

// RemoveSessionIDs removes the "sessions" edge to AdminSession entities by IDs.
func (auo *AdminUpdateOne) RemoveSessionIDs(ids ...int) *AdminUpdateOne {
	auo.mutation.RemoveSessionIDs(ids...)
	return auo
}

// RemoveSessions removes "sessions" edges to AdminSession entities.
func (auo *AdminUpdateOne) RemoveSessions(a ...*AdminSession) *AdminUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.RemoveSessionIDs(ids...)
}

// Save executes the query and returns the updated Admin entity.
func (auo *AdminUpdateOne) Save(ctx context.Context) (*Admin, error) {
	var (
		err  error
		node *Admin
	)
	if len(auo.hooks) == 0 {
		if err = auo.check(); err != nil {
			return nil, err
		}
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdminMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auo.check(); err != nil {
				return nil, err
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AdminUpdateOne) SaveX(ctx context.Context) *Admin {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AdminUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AdminUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AdminUpdateOne) check() error {
	if v, ok := auo.mutation.Email(); ok {
		if err := admin.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	if v, ok := auo.mutation.Username(); ok {
		if err := admin.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf("ent: validator failed for field \"username\": %w", err)}
		}
	}
	if v, ok := auo.mutation.Password(); ok {
		if err := admin.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}
	return nil
}

func (auo *AdminUpdateOne) sqlSave(ctx context.Context) (_node *Admin, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   admin.Table,
			Columns: admin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: admin.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Admin.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := auo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admin.FieldEmail,
		})
	}
	if value, ok := auo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admin.FieldUsername,
		})
	}
	if value, ok := auo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admin.FieldPassword,
		})
	}
	if value, ok := auo.mutation.JoinedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: admin.FieldJoinedAt,
		})
	}
	if auo.mutation.SessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   admin.SessionsTable,
			Columns: []string{admin.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adminsession.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedSessionsIDs(); len(nodes) > 0 && !auo.mutation.SessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   admin.SessionsTable,
			Columns: []string{admin.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adminsession.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.SessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   admin.SessionsTable,
			Columns: []string{admin.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: adminsession.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Admin{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{admin.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}