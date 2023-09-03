// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xmdhs/authlib-skin/db/ent/predicate"
	"github.com/xmdhs/authlib-skin/db/ent/skin"
	"github.com/xmdhs/authlib-skin/db/ent/user"
)

// SkinQuery is the builder for querying Skin entities.
type SkinQuery struct {
	config
	ctx             *QueryContext
	order           []skin.OrderOption
	inters          []Interceptor
	predicates      []predicate.Skin
	withCreatedUser *UserQuery
	withFKs         bool
	modifiers       []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SkinQuery builder.
func (sq *SkinQuery) Where(ps ...predicate.Skin) *SkinQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *SkinQuery) Limit(limit int) *SkinQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *SkinQuery) Offset(offset int) *SkinQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SkinQuery) Unique(unique bool) *SkinQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *SkinQuery) Order(o ...skin.OrderOption) *SkinQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryCreatedUser chains the current query on the "created_user" edge.
func (sq *SkinQuery) QueryCreatedUser() *UserQuery {
	query := (&UserClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(skin.Table, skin.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, skin.CreatedUserTable, skin.CreatedUserColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Skin entity from the query.
// Returns a *NotFoundError when no Skin was found.
func (sq *SkinQuery) First(ctx context.Context) (*Skin, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{skin.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SkinQuery) FirstX(ctx context.Context) *Skin {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Skin ID from the query.
// Returns a *NotFoundError when no Skin ID was found.
func (sq *SkinQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{skin.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SkinQuery) FirstIDX(ctx context.Context) int {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Skin entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Skin entity is found.
// Returns a *NotFoundError when no Skin entities are found.
func (sq *SkinQuery) Only(ctx context.Context) (*Skin, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{skin.Label}
	default:
		return nil, &NotSingularError{skin.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SkinQuery) OnlyX(ctx context.Context) *Skin {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Skin ID in the query.
// Returns a *NotSingularError when more than one Skin ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SkinQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{skin.Label}
	default:
		err = &NotSingularError{skin.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SkinQuery) OnlyIDX(ctx context.Context) int {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Skins.
func (sq *SkinQuery) All(ctx context.Context) ([]*Skin, error) {
	ctx = setContextOp(ctx, sq.ctx, "All")
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Skin, *SkinQuery]()
	return withInterceptors[[]*Skin](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *SkinQuery) AllX(ctx context.Context) []*Skin {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Skin IDs.
func (sq *SkinQuery) IDs(ctx context.Context) (ids []int, err error) {
	if sq.ctx.Unique == nil && sq.path != nil {
		sq.Unique(true)
	}
	ctx = setContextOp(ctx, sq.ctx, "IDs")
	if err = sq.Select(skin.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SkinQuery) IDsX(ctx context.Context) []int {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SkinQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, "Count")
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*SkinQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SkinQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SkinQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, "Exist")
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SkinQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SkinQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SkinQuery) Clone() *SkinQuery {
	if sq == nil {
		return nil
	}
	return &SkinQuery{
		config:          sq.config,
		ctx:             sq.ctx.Clone(),
		order:           append([]skin.OrderOption{}, sq.order...),
		inters:          append([]Interceptor{}, sq.inters...),
		predicates:      append([]predicate.Skin{}, sq.predicates...),
		withCreatedUser: sq.withCreatedUser.Clone(),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// WithCreatedUser tells the query-builder to eager-load the nodes that are connected to
// the "created_user" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SkinQuery) WithCreatedUser(opts ...func(*UserQuery)) *SkinQuery {
	query := (&UserClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withCreatedUser = query
	return sq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		SkinHash string `json:"skin_hash,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Skin.Query().
//		GroupBy(skin.FieldSkinHash).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *SkinQuery) GroupBy(field string, fields ...string) *SkinGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SkinGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = skin.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		SkinHash string `json:"skin_hash,omitempty"`
//	}
//
//	client.Skin.Query().
//		Select(skin.FieldSkinHash).
//		Scan(ctx, &v)
func (sq *SkinQuery) Select(fields ...string) *SkinSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &SkinSelect{SkinQuery: sq}
	sbuild.label = skin.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SkinSelect configured with the given aggregations.
func (sq *SkinQuery) Aggregate(fns ...AggregateFunc) *SkinSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *SkinQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !skin.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SkinQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Skin, error) {
	var (
		nodes       = []*Skin{}
		withFKs     = sq.withFKs
		_spec       = sq.querySpec()
		loadedTypes = [1]bool{
			sq.withCreatedUser != nil,
		}
	)
	if sq.withCreatedUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, skin.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Skin).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Skin{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withCreatedUser; query != nil {
		if err := sq.loadCreatedUser(ctx, query, nodes, nil,
			func(n *Skin, e *User) { n.Edges.CreatedUser = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *SkinQuery) loadCreatedUser(ctx context.Context, query *UserQuery, nodes []*Skin, init func(*Skin), assign func(*Skin, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Skin)
	for i := range nodes {
		if nodes[i].skin_created_user == nil {
			continue
		}
		fk := *nodes[i].skin_created_user
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "skin_created_user" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (sq *SkinQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SkinQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(skin.Table, skin.Columns, sqlgraph.NewFieldSpec(skin.FieldID, field.TypeInt))
	_spec.From = sq.sql
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sq.path != nil {
		_spec.Unique = true
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, skin.FieldID)
		for i := range fields {
			if fields[i] != skin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SkinQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(skin.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = skin.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range sq.modifiers {
		m(selector)
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (sq *SkinQuery) ForUpdate(opts ...sql.LockOption) *SkinQuery {
	if sq.driver.Dialect() == dialect.Postgres {
		sq.Unique(false)
	}
	sq.modifiers = append(sq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return sq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (sq *SkinQuery) ForShare(opts ...sql.LockOption) *SkinQuery {
	if sq.driver.Dialect() == dialect.Postgres {
		sq.Unique(false)
	}
	sq.modifiers = append(sq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return sq
}

// SkinGroupBy is the group-by builder for Skin entities.
type SkinGroupBy struct {
	selector
	build *SkinQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SkinGroupBy) Aggregate(fns ...AggregateFunc) *SkinGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *SkinGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, "GroupBy")
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SkinQuery, *SkinGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *SkinGroupBy) sqlScan(ctx context.Context, root *SkinQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SkinSelect is the builder for selecting fields of Skin entities.
type SkinSelect struct {
	*SkinQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *SkinSelect) Aggregate(fns ...AggregateFunc) *SkinSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SkinSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, "Select")
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SkinQuery, *SkinSelect](ctx, ss.SkinQuery, ss, ss.inters, v)
}

func (ss *SkinSelect) sqlScan(ctx context.Context, root *SkinQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
