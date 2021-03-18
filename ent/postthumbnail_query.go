// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"devlog/ent/post"
	"devlog/ent/postthumbnail"
	"devlog/ent/predicate"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// PostThumbnailQuery is the builder for querying PostThumbnail entities.
type PostThumbnailQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.PostThumbnail
	// eager-loading edges.
	withPost *PostQuery
	withFKs  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PostThumbnailQuery builder.
func (ptq *PostThumbnailQuery) Where(ps ...predicate.PostThumbnail) *PostThumbnailQuery {
	ptq.predicates = append(ptq.predicates, ps...)
	return ptq
}

// Limit adds a limit step to the query.
func (ptq *PostThumbnailQuery) Limit(limit int) *PostThumbnailQuery {
	ptq.limit = &limit
	return ptq
}

// Offset adds an offset step to the query.
func (ptq *PostThumbnailQuery) Offset(offset int) *PostThumbnailQuery {
	ptq.offset = &offset
	return ptq
}

// Order adds an order step to the query.
func (ptq *PostThumbnailQuery) Order(o ...OrderFunc) *PostThumbnailQuery {
	ptq.order = append(ptq.order, o...)
	return ptq
}

// QueryPost chains the current query on the "post" edge.
func (ptq *PostThumbnailQuery) QueryPost() *PostQuery {
	query := &PostQuery{config: ptq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ptq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ptq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(postthumbnail.Table, postthumbnail.FieldID, selector),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, postthumbnail.PostTable, postthumbnail.PostColumn),
		)
		fromU = sqlgraph.SetNeighbors(ptq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PostThumbnail entity from the query.
// Returns a *NotFoundError when no PostThumbnail was found.
func (ptq *PostThumbnailQuery) First(ctx context.Context) (*PostThumbnail, error) {
	nodes, err := ptq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{postthumbnail.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ptq *PostThumbnailQuery) FirstX(ctx context.Context) *PostThumbnail {
	node, err := ptq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PostThumbnail ID from the query.
// Returns a *NotFoundError when no PostThumbnail ID was found.
func (ptq *PostThumbnailQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ptq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{postthumbnail.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ptq *PostThumbnailQuery) FirstIDX(ctx context.Context) int {
	id, err := ptq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PostThumbnail entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one PostThumbnail entity is not found.
// Returns a *NotFoundError when no PostThumbnail entities are found.
func (ptq *PostThumbnailQuery) Only(ctx context.Context) (*PostThumbnail, error) {
	nodes, err := ptq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{postthumbnail.Label}
	default:
		return nil, &NotSingularError{postthumbnail.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ptq *PostThumbnailQuery) OnlyX(ctx context.Context) *PostThumbnail {
	node, err := ptq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PostThumbnail ID in the query.
// Returns a *NotSingularError when exactly one PostThumbnail ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ptq *PostThumbnailQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ptq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{postthumbnail.Label}
	default:
		err = &NotSingularError{postthumbnail.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ptq *PostThumbnailQuery) OnlyIDX(ctx context.Context) int {
	id, err := ptq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PostThumbnails.
func (ptq *PostThumbnailQuery) All(ctx context.Context) ([]*PostThumbnail, error) {
	if err := ptq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ptq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ptq *PostThumbnailQuery) AllX(ctx context.Context) []*PostThumbnail {
	nodes, err := ptq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PostThumbnail IDs.
func (ptq *PostThumbnailQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ptq.Select(postthumbnail.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ptq *PostThumbnailQuery) IDsX(ctx context.Context) []int {
	ids, err := ptq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ptq *PostThumbnailQuery) Count(ctx context.Context) (int, error) {
	if err := ptq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ptq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ptq *PostThumbnailQuery) CountX(ctx context.Context) int {
	count, err := ptq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ptq *PostThumbnailQuery) Exist(ctx context.Context) (bool, error) {
	if err := ptq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ptq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ptq *PostThumbnailQuery) ExistX(ctx context.Context) bool {
	exist, err := ptq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PostThumbnailQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ptq *PostThumbnailQuery) Clone() *PostThumbnailQuery {
	if ptq == nil {
		return nil
	}
	return &PostThumbnailQuery{
		config:     ptq.config,
		limit:      ptq.limit,
		offset:     ptq.offset,
		order:      append([]OrderFunc{}, ptq.order...),
		predicates: append([]predicate.PostThumbnail{}, ptq.predicates...),
		withPost:   ptq.withPost.Clone(),
		// clone intermediate query.
		sql:  ptq.sql.Clone(),
		path: ptq.path,
	}
}

// WithPost tells the query-builder to eager-load the nodes that are connected to
// the "post" edge. The optional arguments are used to configure the query builder of the edge.
func (ptq *PostThumbnailQuery) WithPost(opts ...func(*PostQuery)) *PostThumbnailQuery {
	query := &PostQuery{config: ptq.config}
	for _, opt := range opts {
		opt(query)
	}
	ptq.withPost = query
	return ptq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Width uint32 `json:"width,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PostThumbnail.Query().
//		GroupBy(postthumbnail.FieldWidth).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ptq *PostThumbnailQuery) GroupBy(field string, fields ...string) *PostThumbnailGroupBy {
	group := &PostThumbnailGroupBy{config: ptq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ptq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ptq.sqlQuery(), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Width uint32 `json:"width,omitempty"`
//	}
//
//	client.PostThumbnail.Query().
//		Select(postthumbnail.FieldWidth).
//		Scan(ctx, &v)
//
func (ptq *PostThumbnailQuery) Select(field string, fields ...string) *PostThumbnailSelect {
	ptq.fields = append([]string{field}, fields...)
	return &PostThumbnailSelect{PostThumbnailQuery: ptq}
}

func (ptq *PostThumbnailQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ptq.fields {
		if !postthumbnail.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ptq.path != nil {
		prev, err := ptq.path(ctx)
		if err != nil {
			return err
		}
		ptq.sql = prev
	}
	return nil
}

func (ptq *PostThumbnailQuery) sqlAll(ctx context.Context) ([]*PostThumbnail, error) {
	var (
		nodes       = []*PostThumbnail{}
		withFKs     = ptq.withFKs
		_spec       = ptq.querySpec()
		loadedTypes = [1]bool{
			ptq.withPost != nil,
		}
	)
	if ptq.withPost != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, postthumbnail.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &PostThumbnail{config: ptq.config}
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
	if err := sqlgraph.QueryNodes(ctx, ptq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ptq.withPost; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*PostThumbnail)
		for i := range nodes {
			if fk := nodes[i].post_thumbnail; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(post.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "post_thumbnail" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Post = n
			}
		}
	}

	return nodes, nil
}

func (ptq *PostThumbnailQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ptq.querySpec()
	return sqlgraph.CountNodes(ctx, ptq.driver, _spec)
}

func (ptq *PostThumbnailQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ptq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (ptq *PostThumbnailQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   postthumbnail.Table,
			Columns: postthumbnail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: postthumbnail.FieldID,
			},
		},
		From:   ptq.sql,
		Unique: true,
	}
	if fields := ptq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, postthumbnail.FieldID)
		for i := range fields {
			if fields[i] != postthumbnail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ptq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ptq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ptq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ptq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, postthumbnail.ValidColumn)
			}
		}
	}
	return _spec
}

func (ptq *PostThumbnailQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(ptq.driver.Dialect())
	t1 := builder.Table(postthumbnail.Table)
	selector := builder.Select(t1.Columns(postthumbnail.Columns...)...).From(t1)
	if ptq.sql != nil {
		selector = ptq.sql
		selector.Select(selector.Columns(postthumbnail.Columns...)...)
	}
	for _, p := range ptq.predicates {
		p(selector)
	}
	for _, p := range ptq.order {
		p(selector, postthumbnail.ValidColumn)
	}
	if offset := ptq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ptq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PostThumbnailGroupBy is the group-by builder for PostThumbnail entities.
type PostThumbnailGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ptgb *PostThumbnailGroupBy) Aggregate(fns ...AggregateFunc) *PostThumbnailGroupBy {
	ptgb.fns = append(ptgb.fns, fns...)
	return ptgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ptgb *PostThumbnailGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ptgb.path(ctx)
	if err != nil {
		return err
	}
	ptgb.sql = query
	return ptgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ptgb *PostThumbnailGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ptgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ptgb *PostThumbnailGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ptgb.fields) > 1 {
		return nil, errors.New("ent: PostThumbnailGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ptgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ptgb *PostThumbnailGroupBy) StringsX(ctx context.Context) []string {
	v, err := ptgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ptgb *PostThumbnailGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ptgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{postthumbnail.Label}
	default:
		err = fmt.Errorf("ent: PostThumbnailGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ptgb *PostThumbnailGroupBy) StringX(ctx context.Context) string {
	v, err := ptgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ptgb *PostThumbnailGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ptgb.fields) > 1 {
		return nil, errors.New("ent: PostThumbnailGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ptgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ptgb *PostThumbnailGroupBy) IntsX(ctx context.Context) []int {
	v, err := ptgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ptgb *PostThumbnailGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ptgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{postthumbnail.Label}
	default:
		err = fmt.Errorf("ent: PostThumbnailGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ptgb *PostThumbnailGroupBy) IntX(ctx context.Context) int {
	v, err := ptgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ptgb *PostThumbnailGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ptgb.fields) > 1 {
		return nil, errors.New("ent: PostThumbnailGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ptgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ptgb *PostThumbnailGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ptgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ptgb *PostThumbnailGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ptgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{postthumbnail.Label}
	default:
		err = fmt.Errorf("ent: PostThumbnailGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ptgb *PostThumbnailGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ptgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ptgb *PostThumbnailGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ptgb.fields) > 1 {
		return nil, errors.New("ent: PostThumbnailGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ptgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ptgb *PostThumbnailGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ptgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ptgb *PostThumbnailGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ptgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{postthumbnail.Label}
	default:
		err = fmt.Errorf("ent: PostThumbnailGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ptgb *PostThumbnailGroupBy) BoolX(ctx context.Context) bool {
	v, err := ptgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ptgb *PostThumbnailGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ptgb.fields {
		if !postthumbnail.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ptgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ptgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ptgb *PostThumbnailGroupBy) sqlQuery() *sql.Selector {
	selector := ptgb.sql
	columns := make([]string, 0, len(ptgb.fields)+len(ptgb.fns))
	columns = append(columns, ptgb.fields...)
	for _, fn := range ptgb.fns {
		columns = append(columns, fn(selector, postthumbnail.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(ptgb.fields...)
}

// PostThumbnailSelect is the builder for selecting fields of PostThumbnail entities.
type PostThumbnailSelect struct {
	*PostThumbnailQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pts *PostThumbnailSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pts.prepareQuery(ctx); err != nil {
		return err
	}
	pts.sql = pts.PostThumbnailQuery.sqlQuery()
	return pts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pts *PostThumbnailSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (pts *PostThumbnailSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pts.fields) > 1 {
		return nil, errors.New("ent: PostThumbnailSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pts *PostThumbnailSelect) StringsX(ctx context.Context) []string {
	v, err := pts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (pts *PostThumbnailSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{postthumbnail.Label}
	default:
		err = fmt.Errorf("ent: PostThumbnailSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pts *PostThumbnailSelect) StringX(ctx context.Context) string {
	v, err := pts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (pts *PostThumbnailSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pts.fields) > 1 {
		return nil, errors.New("ent: PostThumbnailSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pts *PostThumbnailSelect) IntsX(ctx context.Context) []int {
	v, err := pts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (pts *PostThumbnailSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{postthumbnail.Label}
	default:
		err = fmt.Errorf("ent: PostThumbnailSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pts *PostThumbnailSelect) IntX(ctx context.Context) int {
	v, err := pts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (pts *PostThumbnailSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pts.fields) > 1 {
		return nil, errors.New("ent: PostThumbnailSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pts *PostThumbnailSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (pts *PostThumbnailSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{postthumbnail.Label}
	default:
		err = fmt.Errorf("ent: PostThumbnailSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pts *PostThumbnailSelect) Float64X(ctx context.Context) float64 {
	v, err := pts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (pts *PostThumbnailSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pts.fields) > 1 {
		return nil, errors.New("ent: PostThumbnailSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pts *PostThumbnailSelect) BoolsX(ctx context.Context) []bool {
	v, err := pts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (pts *PostThumbnailSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{postthumbnail.Label}
	default:
		err = fmt.Errorf("ent: PostThumbnailSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pts *PostThumbnailSelect) BoolX(ctx context.Context) bool {
	v, err := pts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pts *PostThumbnailSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pts.sqlQuery().Query()
	if err := pts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pts *PostThumbnailSelect) sqlQuery() sql.Querier {
	selector := pts.sql
	selector.Select(selector.Columns(pts.fields...)...)
	return selector
}
