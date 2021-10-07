// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"devlog/ent/admin"
	"devlog/ent/predicate"
	"devlog/ent/robotaccess"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RobotAccessQuery is the builder for querying RobotAccess entities.
type RobotAccessQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.RobotAccess
	// eager-loading edges.
	withUser *AdminQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RobotAccessQuery builder.
func (raq *RobotAccessQuery) Where(ps ...predicate.RobotAccess) *RobotAccessQuery {
	raq.predicates = append(raq.predicates, ps...)
	return raq
}

// Limit adds a limit step to the query.
func (raq *RobotAccessQuery) Limit(limit int) *RobotAccessQuery {
	raq.limit = &limit
	return raq
}

// Offset adds an offset step to the query.
func (raq *RobotAccessQuery) Offset(offset int) *RobotAccessQuery {
	raq.offset = &offset
	return raq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (raq *RobotAccessQuery) Unique(unique bool) *RobotAccessQuery {
	raq.unique = &unique
	return raq
}

// Order adds an order step to the query.
func (raq *RobotAccessQuery) Order(o ...OrderFunc) *RobotAccessQuery {
	raq.order = append(raq.order, o...)
	return raq
}

// QueryUser chains the current query on the "user" edge.
func (raq *RobotAccessQuery) QueryUser() *AdminQuery {
	query := &AdminQuery{config: raq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := raq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := raq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(robotaccess.Table, robotaccess.FieldID, selector),
			sqlgraph.To(admin.Table, admin.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, robotaccess.UserTable, robotaccess.UserPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(raq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RobotAccess entity from the query.
// Returns a *NotFoundError when no RobotAccess was found.
func (raq *RobotAccessQuery) First(ctx context.Context) (*RobotAccess, error) {
	nodes, err := raq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{robotaccess.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (raq *RobotAccessQuery) FirstX(ctx context.Context) *RobotAccess {
	node, err := raq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RobotAccess ID from the query.
// Returns a *NotFoundError when no RobotAccess ID was found.
func (raq *RobotAccessQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = raq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{robotaccess.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (raq *RobotAccessQuery) FirstIDX(ctx context.Context) int {
	id, err := raq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RobotAccess entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one RobotAccess entity is not found.
// Returns a *NotFoundError when no RobotAccess entities are found.
func (raq *RobotAccessQuery) Only(ctx context.Context) (*RobotAccess, error) {
	nodes, err := raq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{robotaccess.Label}
	default:
		return nil, &NotSingularError{robotaccess.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (raq *RobotAccessQuery) OnlyX(ctx context.Context) *RobotAccess {
	node, err := raq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RobotAccess ID in the query.
// Returns a *NotSingularError when exactly one RobotAccess ID is not found.
// Returns a *NotFoundError when no entities are found.
func (raq *RobotAccessQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = raq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{robotaccess.Label}
	default:
		err = &NotSingularError{robotaccess.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (raq *RobotAccessQuery) OnlyIDX(ctx context.Context) int {
	id, err := raq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RobotAccesses.
func (raq *RobotAccessQuery) All(ctx context.Context) ([]*RobotAccess, error) {
	if err := raq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return raq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (raq *RobotAccessQuery) AllX(ctx context.Context) []*RobotAccess {
	nodes, err := raq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RobotAccess IDs.
func (raq *RobotAccessQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := raq.Select(robotaccess.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (raq *RobotAccessQuery) IDsX(ctx context.Context) []int {
	ids, err := raq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (raq *RobotAccessQuery) Count(ctx context.Context) (int, error) {
	if err := raq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return raq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (raq *RobotAccessQuery) CountX(ctx context.Context) int {
	count, err := raq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (raq *RobotAccessQuery) Exist(ctx context.Context) (bool, error) {
	if err := raq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return raq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (raq *RobotAccessQuery) ExistX(ctx context.Context) bool {
	exist, err := raq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RobotAccessQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (raq *RobotAccessQuery) Clone() *RobotAccessQuery {
	if raq == nil {
		return nil
	}
	return &RobotAccessQuery{
		config:     raq.config,
		limit:      raq.limit,
		offset:     raq.offset,
		order:      append([]OrderFunc{}, raq.order...),
		predicates: append([]predicate.RobotAccess{}, raq.predicates...),
		withUser:   raq.withUser.Clone(),
		// clone intermediate query.
		sql:  raq.sql.Clone(),
		path: raq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (raq *RobotAccessQuery) WithUser(opts ...func(*AdminQuery)) *RobotAccessQuery {
	query := &AdminQuery{config: raq.config}
	for _, opt := range opts {
		opt(query)
	}
	raq.withUser = query
	return raq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Token string `json:"token,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RobotAccess.Query().
//		GroupBy(robotaccess.FieldToken).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (raq *RobotAccessQuery) GroupBy(field string, fields ...string) *RobotAccessGroupBy {
	group := &RobotAccessGroupBy{config: raq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := raq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return raq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Token string `json:"token,omitempty"`
//	}
//
//	client.RobotAccess.Query().
//		Select(robotaccess.FieldToken).
//		Scan(ctx, &v)
//
func (raq *RobotAccessQuery) Select(fields ...string) *RobotAccessSelect {
	raq.fields = append(raq.fields, fields...)
	return &RobotAccessSelect{RobotAccessQuery: raq}
}

func (raq *RobotAccessQuery) prepareQuery(ctx context.Context) error {
	for _, f := range raq.fields {
		if !robotaccess.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if raq.path != nil {
		prev, err := raq.path(ctx)
		if err != nil {
			return err
		}
		raq.sql = prev
	}
	return nil
}

func (raq *RobotAccessQuery) sqlAll(ctx context.Context) ([]*RobotAccess, error) {
	var (
		nodes       = []*RobotAccess{}
		_spec       = raq.querySpec()
		loadedTypes = [1]bool{
			raq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &RobotAccess{config: raq.config}
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
	if err := sqlgraph.QueryNodes(ctx, raq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := raq.withUser; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*RobotAccess, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.User = []*Admin{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*RobotAccess)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: true,
				Table:   robotaccess.UserTable,
				Columns: robotaccess.UserPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(robotaccess.UserPrimaryKey[1], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, raq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "user": %w`, err)
		}
		query.Where(admin.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "user" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = append(nodes[i].Edges.User, n)
			}
		}
	}

	return nodes, nil
}

func (raq *RobotAccessQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := raq.querySpec()
	return sqlgraph.CountNodes(ctx, raq.driver, _spec)
}

func (raq *RobotAccessQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := raq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (raq *RobotAccessQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   robotaccess.Table,
			Columns: robotaccess.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: robotaccess.FieldID,
			},
		},
		From:   raq.sql,
		Unique: true,
	}
	if unique := raq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := raq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, robotaccess.FieldID)
		for i := range fields {
			if fields[i] != robotaccess.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := raq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := raq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := raq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := raq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (raq *RobotAccessQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(raq.driver.Dialect())
	t1 := builder.Table(robotaccess.Table)
	columns := raq.fields
	if len(columns) == 0 {
		columns = robotaccess.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if raq.sql != nil {
		selector = raq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range raq.predicates {
		p(selector)
	}
	for _, p := range raq.order {
		p(selector)
	}
	if offset := raq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := raq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RobotAccessGroupBy is the group-by builder for RobotAccess entities.
type RobotAccessGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ragb *RobotAccessGroupBy) Aggregate(fns ...AggregateFunc) *RobotAccessGroupBy {
	ragb.fns = append(ragb.fns, fns...)
	return ragb
}

// Scan applies the group-by query and scans the result into the given value.
func (ragb *RobotAccessGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ragb.path(ctx)
	if err != nil {
		return err
	}
	ragb.sql = query
	return ragb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ragb *RobotAccessGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ragb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ragb *RobotAccessGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ragb.fields) > 1 {
		return nil, errors.New("ent: RobotAccessGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ragb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ragb *RobotAccessGroupBy) StringsX(ctx context.Context) []string {
	v, err := ragb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ragb *RobotAccessGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ragb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{robotaccess.Label}
	default:
		err = fmt.Errorf("ent: RobotAccessGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ragb *RobotAccessGroupBy) StringX(ctx context.Context) string {
	v, err := ragb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ragb *RobotAccessGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ragb.fields) > 1 {
		return nil, errors.New("ent: RobotAccessGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ragb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ragb *RobotAccessGroupBy) IntsX(ctx context.Context) []int {
	v, err := ragb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ragb *RobotAccessGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ragb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{robotaccess.Label}
	default:
		err = fmt.Errorf("ent: RobotAccessGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ragb *RobotAccessGroupBy) IntX(ctx context.Context) int {
	v, err := ragb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ragb *RobotAccessGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ragb.fields) > 1 {
		return nil, errors.New("ent: RobotAccessGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ragb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ragb *RobotAccessGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ragb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ragb *RobotAccessGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ragb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{robotaccess.Label}
	default:
		err = fmt.Errorf("ent: RobotAccessGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ragb *RobotAccessGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ragb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ragb *RobotAccessGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ragb.fields) > 1 {
		return nil, errors.New("ent: RobotAccessGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ragb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ragb *RobotAccessGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ragb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ragb *RobotAccessGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ragb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{robotaccess.Label}
	default:
		err = fmt.Errorf("ent: RobotAccessGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ragb *RobotAccessGroupBy) BoolX(ctx context.Context) bool {
	v, err := ragb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ragb *RobotAccessGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ragb.fields {
		if !robotaccess.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ragb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ragb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ragb *RobotAccessGroupBy) sqlQuery() *sql.Selector {
	selector := ragb.sql.Select()
	aggregation := make([]string, 0, len(ragb.fns))
	for _, fn := range ragb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ragb.fields)+len(ragb.fns))
		for _, f := range ragb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ragb.fields...)...)
}

// RobotAccessSelect is the builder for selecting fields of RobotAccess entities.
type RobotAccessSelect struct {
	*RobotAccessQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ras *RobotAccessSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ras.prepareQuery(ctx); err != nil {
		return err
	}
	ras.sql = ras.RobotAccessQuery.sqlQuery(ctx)
	return ras.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ras *RobotAccessSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ras.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ras *RobotAccessSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ras.fields) > 1 {
		return nil, errors.New("ent: RobotAccessSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ras.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ras *RobotAccessSelect) StringsX(ctx context.Context) []string {
	v, err := ras.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ras *RobotAccessSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ras.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{robotaccess.Label}
	default:
		err = fmt.Errorf("ent: RobotAccessSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ras *RobotAccessSelect) StringX(ctx context.Context) string {
	v, err := ras.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ras *RobotAccessSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ras.fields) > 1 {
		return nil, errors.New("ent: RobotAccessSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ras.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ras *RobotAccessSelect) IntsX(ctx context.Context) []int {
	v, err := ras.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ras *RobotAccessSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ras.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{robotaccess.Label}
	default:
		err = fmt.Errorf("ent: RobotAccessSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ras *RobotAccessSelect) IntX(ctx context.Context) int {
	v, err := ras.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ras *RobotAccessSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ras.fields) > 1 {
		return nil, errors.New("ent: RobotAccessSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ras.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ras *RobotAccessSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ras.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ras *RobotAccessSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ras.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{robotaccess.Label}
	default:
		err = fmt.Errorf("ent: RobotAccessSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ras *RobotAccessSelect) Float64X(ctx context.Context) float64 {
	v, err := ras.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ras *RobotAccessSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ras.fields) > 1 {
		return nil, errors.New("ent: RobotAccessSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ras.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ras *RobotAccessSelect) BoolsX(ctx context.Context) []bool {
	v, err := ras.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ras *RobotAccessSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ras.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{robotaccess.Label}
	default:
		err = fmt.Errorf("ent: RobotAccessSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ras *RobotAccessSelect) BoolX(ctx context.Context) bool {
	v, err := ras.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ras *RobotAccessSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ras.sql.Query()
	if err := ras.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}