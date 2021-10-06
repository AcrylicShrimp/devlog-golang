// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devlog/ent/predicate"
	"devlog/ent/unsavedpost"
	"devlog/ent/unsavedpostattachment"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UnsavedPostAttachmentQuery is the builder for querying UnsavedPostAttachment entities.
type UnsavedPostAttachmentQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UnsavedPostAttachment
	// eager-loading edges.
	withUnsavedPost *UnsavedPostQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UnsavedPostAttachmentQuery builder.
func (upaq *UnsavedPostAttachmentQuery) Where(ps ...predicate.UnsavedPostAttachment) *UnsavedPostAttachmentQuery {
	upaq.predicates = append(upaq.predicates, ps...)
	return upaq
}

// Limit adds a limit step to the query.
func (upaq *UnsavedPostAttachmentQuery) Limit(limit int) *UnsavedPostAttachmentQuery {
	upaq.limit = &limit
	return upaq
}

// Offset adds an offset step to the query.
func (upaq *UnsavedPostAttachmentQuery) Offset(offset int) *UnsavedPostAttachmentQuery {
	upaq.offset = &offset
	return upaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (upaq *UnsavedPostAttachmentQuery) Unique(unique bool) *UnsavedPostAttachmentQuery {
	upaq.unique = &unique
	return upaq
}

// Order adds an order step to the query.
func (upaq *UnsavedPostAttachmentQuery) Order(o ...OrderFunc) *UnsavedPostAttachmentQuery {
	upaq.order = append(upaq.order, o...)
	return upaq
}

// QueryUnsavedPost chains the current query on the "unsaved_post" edge.
func (upaq *UnsavedPostAttachmentQuery) QueryUnsavedPost() *UnsavedPostQuery {
	query := &UnsavedPostQuery{config: upaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := upaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := upaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(unsavedpostattachment.Table, unsavedpostattachment.FieldID, selector),
			sqlgraph.To(unsavedpost.Table, unsavedpost.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, unsavedpostattachment.UnsavedPostTable, unsavedpostattachment.UnsavedPostColumn),
		)
		fromU = sqlgraph.SetNeighbors(upaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UnsavedPostAttachment entity from the query.
// Returns a *NotFoundError when no UnsavedPostAttachment was found.
func (upaq *UnsavedPostAttachmentQuery) First(ctx context.Context) (*UnsavedPostAttachment, error) {
	nodes, err := upaq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{unsavedpostattachment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (upaq *UnsavedPostAttachmentQuery) FirstX(ctx context.Context) *UnsavedPostAttachment {
	node, err := upaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UnsavedPostAttachment ID from the query.
// Returns a *NotFoundError when no UnsavedPostAttachment ID was found.
func (upaq *UnsavedPostAttachmentQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = upaq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{unsavedpostattachment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (upaq *UnsavedPostAttachmentQuery) FirstIDX(ctx context.Context) int {
	id, err := upaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UnsavedPostAttachment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one UnsavedPostAttachment entity is not found.
// Returns a *NotFoundError when no UnsavedPostAttachment entities are found.
func (upaq *UnsavedPostAttachmentQuery) Only(ctx context.Context) (*UnsavedPostAttachment, error) {
	nodes, err := upaq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{unsavedpostattachment.Label}
	default:
		return nil, &NotSingularError{unsavedpostattachment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (upaq *UnsavedPostAttachmentQuery) OnlyX(ctx context.Context) *UnsavedPostAttachment {
	node, err := upaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UnsavedPostAttachment ID in the query.
// Returns a *NotSingularError when exactly one UnsavedPostAttachment ID is not found.
// Returns a *NotFoundError when no entities are found.
func (upaq *UnsavedPostAttachmentQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = upaq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{unsavedpostattachment.Label}
	default:
		err = &NotSingularError{unsavedpostattachment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (upaq *UnsavedPostAttachmentQuery) OnlyIDX(ctx context.Context) int {
	id, err := upaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UnsavedPostAttachments.
func (upaq *UnsavedPostAttachmentQuery) All(ctx context.Context) ([]*UnsavedPostAttachment, error) {
	if err := upaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return upaq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (upaq *UnsavedPostAttachmentQuery) AllX(ctx context.Context) []*UnsavedPostAttachment {
	nodes, err := upaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UnsavedPostAttachment IDs.
func (upaq *UnsavedPostAttachmentQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := upaq.Select(unsavedpostattachment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (upaq *UnsavedPostAttachmentQuery) IDsX(ctx context.Context) []int {
	ids, err := upaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (upaq *UnsavedPostAttachmentQuery) Count(ctx context.Context) (int, error) {
	if err := upaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return upaq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (upaq *UnsavedPostAttachmentQuery) CountX(ctx context.Context) int {
	count, err := upaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (upaq *UnsavedPostAttachmentQuery) Exist(ctx context.Context) (bool, error) {
	if err := upaq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return upaq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (upaq *UnsavedPostAttachmentQuery) ExistX(ctx context.Context) bool {
	exist, err := upaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UnsavedPostAttachmentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (upaq *UnsavedPostAttachmentQuery) Clone() *UnsavedPostAttachmentQuery {
	if upaq == nil {
		return nil
	}
	return &UnsavedPostAttachmentQuery{
		config:          upaq.config,
		limit:           upaq.limit,
		offset:          upaq.offset,
		order:           append([]OrderFunc{}, upaq.order...),
		predicates:      append([]predicate.UnsavedPostAttachment{}, upaq.predicates...),
		withUnsavedPost: upaq.withUnsavedPost.Clone(),
		// clone intermediate query.
		sql:  upaq.sql.Clone(),
		path: upaq.path,
	}
}

// WithUnsavedPost tells the query-builder to eager-load the nodes that are connected to
// the "unsaved_post" edge. The optional arguments are used to configure the query builder of the edge.
func (upaq *UnsavedPostAttachmentQuery) WithUnsavedPost(opts ...func(*UnsavedPostQuery)) *UnsavedPostAttachmentQuery {
	query := &UnsavedPostQuery{config: upaq.config}
	for _, opt := range opts {
		opt(query)
	}
	upaq.withUnsavedPost = query
	return upaq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UUID string `json:"uuid,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UnsavedPostAttachment.Query().
//		GroupBy(unsavedpostattachment.FieldUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (upaq *UnsavedPostAttachmentQuery) GroupBy(field string, fields ...string) *UnsavedPostAttachmentGroupBy {
	group := &UnsavedPostAttachmentGroupBy{config: upaq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := upaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return upaq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UUID string `json:"uuid,omitempty"`
//	}
//
//	client.UnsavedPostAttachment.Query().
//		Select(unsavedpostattachment.FieldUUID).
//		Scan(ctx, &v)
//
func (upaq *UnsavedPostAttachmentQuery) Select(fields ...string) *UnsavedPostAttachmentSelect {
	upaq.fields = append(upaq.fields, fields...)
	return &UnsavedPostAttachmentSelect{UnsavedPostAttachmentQuery: upaq}
}

func (upaq *UnsavedPostAttachmentQuery) prepareQuery(ctx context.Context) error {
	for _, f := range upaq.fields {
		if !unsavedpostattachment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if upaq.path != nil {
		prev, err := upaq.path(ctx)
		if err != nil {
			return err
		}
		upaq.sql = prev
	}
	return nil
}

func (upaq *UnsavedPostAttachmentQuery) sqlAll(ctx context.Context) ([]*UnsavedPostAttachment, error) {
	var (
		nodes       = []*UnsavedPostAttachment{}
		withFKs     = upaq.withFKs
		_spec       = upaq.querySpec()
		loadedTypes = [1]bool{
			upaq.withUnsavedPost != nil,
		}
	)
	if upaq.withUnsavedPost != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, unsavedpostattachment.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &UnsavedPostAttachment{config: upaq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, upaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := upaq.withUnsavedPost; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*UnsavedPostAttachment)
		for i := range nodes {
			if nodes[i].unsaved_post_attachments == nil {
				continue
			}
			fk := *nodes[i].unsaved_post_attachments
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(unsavedpost.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "unsaved_post_attachments" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.UnsavedPost = n
			}
		}
	}

	return nodes, nil
}

func (upaq *UnsavedPostAttachmentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := upaq.querySpec()
	return sqlgraph.CountNodes(ctx, upaq.driver, _spec)
}

func (upaq *UnsavedPostAttachmentQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := upaq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (upaq *UnsavedPostAttachmentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   unsavedpostattachment.Table,
			Columns: unsavedpostattachment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: unsavedpostattachment.FieldID,
			},
		},
		From:   upaq.sql,
		Unique: true,
	}
	if unique := upaq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := upaq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, unsavedpostattachment.FieldID)
		for i := range fields {
			if fields[i] != unsavedpostattachment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := upaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := upaq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := upaq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := upaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (upaq *UnsavedPostAttachmentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(upaq.driver.Dialect())
	t1 := builder.Table(unsavedpostattachment.Table)
	columns := upaq.fields
	if len(columns) == 0 {
		columns = unsavedpostattachment.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if upaq.sql != nil {
		selector = upaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range upaq.predicates {
		p(selector)
	}
	for _, p := range upaq.order {
		p(selector)
	}
	if offset := upaq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := upaq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UnsavedPostAttachmentGroupBy is the group-by builder for UnsavedPostAttachment entities.
type UnsavedPostAttachmentGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (upagb *UnsavedPostAttachmentGroupBy) Aggregate(fns ...AggregateFunc) *UnsavedPostAttachmentGroupBy {
	upagb.fns = append(upagb.fns, fns...)
	return upagb
}

// Scan applies the group-by query and scans the result into the given value.
func (upagb *UnsavedPostAttachmentGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := upagb.path(ctx)
	if err != nil {
		return err
	}
	upagb.sql = query
	return upagb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (upagb *UnsavedPostAttachmentGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := upagb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (upagb *UnsavedPostAttachmentGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(upagb.fields) > 1 {
		return nil, errors.New("ent: UnsavedPostAttachmentGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := upagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (upagb *UnsavedPostAttachmentGroupBy) StringsX(ctx context.Context) []string {
	v, err := upagb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (upagb *UnsavedPostAttachmentGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = upagb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{unsavedpostattachment.Label}
	default:
		err = fmt.Errorf("ent: UnsavedPostAttachmentGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (upagb *UnsavedPostAttachmentGroupBy) StringX(ctx context.Context) string {
	v, err := upagb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (upagb *UnsavedPostAttachmentGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(upagb.fields) > 1 {
		return nil, errors.New("ent: UnsavedPostAttachmentGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := upagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (upagb *UnsavedPostAttachmentGroupBy) IntsX(ctx context.Context) []int {
	v, err := upagb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (upagb *UnsavedPostAttachmentGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = upagb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{unsavedpostattachment.Label}
	default:
		err = fmt.Errorf("ent: UnsavedPostAttachmentGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (upagb *UnsavedPostAttachmentGroupBy) IntX(ctx context.Context) int {
	v, err := upagb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (upagb *UnsavedPostAttachmentGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(upagb.fields) > 1 {
		return nil, errors.New("ent: UnsavedPostAttachmentGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := upagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (upagb *UnsavedPostAttachmentGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := upagb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (upagb *UnsavedPostAttachmentGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = upagb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{unsavedpostattachment.Label}
	default:
		err = fmt.Errorf("ent: UnsavedPostAttachmentGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (upagb *UnsavedPostAttachmentGroupBy) Float64X(ctx context.Context) float64 {
	v, err := upagb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (upagb *UnsavedPostAttachmentGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(upagb.fields) > 1 {
		return nil, errors.New("ent: UnsavedPostAttachmentGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := upagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (upagb *UnsavedPostAttachmentGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := upagb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (upagb *UnsavedPostAttachmentGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = upagb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{unsavedpostattachment.Label}
	default:
		err = fmt.Errorf("ent: UnsavedPostAttachmentGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (upagb *UnsavedPostAttachmentGroupBy) BoolX(ctx context.Context) bool {
	v, err := upagb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (upagb *UnsavedPostAttachmentGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range upagb.fields {
		if !unsavedpostattachment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := upagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := upagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (upagb *UnsavedPostAttachmentGroupBy) sqlQuery() *sql.Selector {
	selector := upagb.sql.Select()
	aggregation := make([]string, 0, len(upagb.fns))
	for _, fn := range upagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(upagb.fields)+len(upagb.fns))
		for _, f := range upagb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(upagb.fields...)...)
}

// UnsavedPostAttachmentSelect is the builder for selecting fields of UnsavedPostAttachment entities.
type UnsavedPostAttachmentSelect struct {
	*UnsavedPostAttachmentQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (upas *UnsavedPostAttachmentSelect) Scan(ctx context.Context, v interface{}) error {
	if err := upas.prepareQuery(ctx); err != nil {
		return err
	}
	upas.sql = upas.UnsavedPostAttachmentQuery.sqlQuery(ctx)
	return upas.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (upas *UnsavedPostAttachmentSelect) ScanX(ctx context.Context, v interface{}) {
	if err := upas.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (upas *UnsavedPostAttachmentSelect) Strings(ctx context.Context) ([]string, error) {
	if len(upas.fields) > 1 {
		return nil, errors.New("ent: UnsavedPostAttachmentSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := upas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (upas *UnsavedPostAttachmentSelect) StringsX(ctx context.Context) []string {
	v, err := upas.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (upas *UnsavedPostAttachmentSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = upas.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{unsavedpostattachment.Label}
	default:
		err = fmt.Errorf("ent: UnsavedPostAttachmentSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (upas *UnsavedPostAttachmentSelect) StringX(ctx context.Context) string {
	v, err := upas.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (upas *UnsavedPostAttachmentSelect) Ints(ctx context.Context) ([]int, error) {
	if len(upas.fields) > 1 {
		return nil, errors.New("ent: UnsavedPostAttachmentSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := upas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (upas *UnsavedPostAttachmentSelect) IntsX(ctx context.Context) []int {
	v, err := upas.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (upas *UnsavedPostAttachmentSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = upas.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{unsavedpostattachment.Label}
	default:
		err = fmt.Errorf("ent: UnsavedPostAttachmentSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (upas *UnsavedPostAttachmentSelect) IntX(ctx context.Context) int {
	v, err := upas.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (upas *UnsavedPostAttachmentSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(upas.fields) > 1 {
		return nil, errors.New("ent: UnsavedPostAttachmentSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := upas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (upas *UnsavedPostAttachmentSelect) Float64sX(ctx context.Context) []float64 {
	v, err := upas.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (upas *UnsavedPostAttachmentSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = upas.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{unsavedpostattachment.Label}
	default:
		err = fmt.Errorf("ent: UnsavedPostAttachmentSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (upas *UnsavedPostAttachmentSelect) Float64X(ctx context.Context) float64 {
	v, err := upas.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (upas *UnsavedPostAttachmentSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(upas.fields) > 1 {
		return nil, errors.New("ent: UnsavedPostAttachmentSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := upas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (upas *UnsavedPostAttachmentSelect) BoolsX(ctx context.Context) []bool {
	v, err := upas.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (upas *UnsavedPostAttachmentSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = upas.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{unsavedpostattachment.Label}
	default:
		err = fmt.Errorf("ent: UnsavedPostAttachmentSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (upas *UnsavedPostAttachmentSelect) BoolX(ctx context.Context) bool {
	v, err := upas.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (upas *UnsavedPostAttachmentSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := upas.sql.Query()
	if err := upas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
