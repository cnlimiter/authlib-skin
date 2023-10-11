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
	"github.com/xmdhs/authlib-skin/db/ent/texture"
	"github.com/xmdhs/authlib-skin/db/ent/userprofile"
	"github.com/xmdhs/authlib-skin/db/ent/usertexture"
)

// UserTextureQuery is the builder for querying UserTexture entities.
type UserTextureQuery struct {
	config
	ctx             *QueryContext
	order           []usertexture.OrderOption
	inters          []Interceptor
	predicates      []predicate.UserTexture
	withUserProfile *UserProfileQuery
	withTexture     *TextureQuery
	modifiers       []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserTextureQuery builder.
func (utq *UserTextureQuery) Where(ps ...predicate.UserTexture) *UserTextureQuery {
	utq.predicates = append(utq.predicates, ps...)
	return utq
}

// Limit the number of records to be returned by this query.
func (utq *UserTextureQuery) Limit(limit int) *UserTextureQuery {
	utq.ctx.Limit = &limit
	return utq
}

// Offset to start from.
func (utq *UserTextureQuery) Offset(offset int) *UserTextureQuery {
	utq.ctx.Offset = &offset
	return utq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (utq *UserTextureQuery) Unique(unique bool) *UserTextureQuery {
	utq.ctx.Unique = &unique
	return utq
}

// Order specifies how the records should be ordered.
func (utq *UserTextureQuery) Order(o ...usertexture.OrderOption) *UserTextureQuery {
	utq.order = append(utq.order, o...)
	return utq
}

// QueryUserProfile chains the current query on the "user_profile" edge.
func (utq *UserTextureQuery) QueryUserProfile() *UserProfileQuery {
	query := (&UserProfileClient{config: utq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := utq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := utq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usertexture.Table, usertexture.FieldID, selector),
			sqlgraph.To(userprofile.Table, userprofile.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, usertexture.UserProfileTable, usertexture.UserProfileColumn),
		)
		fromU = sqlgraph.SetNeighbors(utq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTexture chains the current query on the "texture" edge.
func (utq *UserTextureQuery) QueryTexture() *TextureQuery {
	query := (&TextureClient{config: utq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := utq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := utq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usertexture.Table, usertexture.FieldID, selector),
			sqlgraph.To(texture.Table, texture.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, usertexture.TextureTable, usertexture.TextureColumn),
		)
		fromU = sqlgraph.SetNeighbors(utq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserTexture entity from the query.
// Returns a *NotFoundError when no UserTexture was found.
func (utq *UserTextureQuery) First(ctx context.Context) (*UserTexture, error) {
	nodes, err := utq.Limit(1).All(setContextOp(ctx, utq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usertexture.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (utq *UserTextureQuery) FirstX(ctx context.Context) *UserTexture {
	node, err := utq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserTexture ID from the query.
// Returns a *NotFoundError when no UserTexture ID was found.
func (utq *UserTextureQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = utq.Limit(1).IDs(setContextOp(ctx, utq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usertexture.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (utq *UserTextureQuery) FirstIDX(ctx context.Context) int {
	id, err := utq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserTexture entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserTexture entity is found.
// Returns a *NotFoundError when no UserTexture entities are found.
func (utq *UserTextureQuery) Only(ctx context.Context) (*UserTexture, error) {
	nodes, err := utq.Limit(2).All(setContextOp(ctx, utq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usertexture.Label}
	default:
		return nil, &NotSingularError{usertexture.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (utq *UserTextureQuery) OnlyX(ctx context.Context) *UserTexture {
	node, err := utq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserTexture ID in the query.
// Returns a *NotSingularError when more than one UserTexture ID is found.
// Returns a *NotFoundError when no entities are found.
func (utq *UserTextureQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = utq.Limit(2).IDs(setContextOp(ctx, utq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usertexture.Label}
	default:
		err = &NotSingularError{usertexture.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (utq *UserTextureQuery) OnlyIDX(ctx context.Context) int {
	id, err := utq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserTextures.
func (utq *UserTextureQuery) All(ctx context.Context) ([]*UserTexture, error) {
	ctx = setContextOp(ctx, utq.ctx, "All")
	if err := utq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserTexture, *UserTextureQuery]()
	return withInterceptors[[]*UserTexture](ctx, utq, qr, utq.inters)
}

// AllX is like All, but panics if an error occurs.
func (utq *UserTextureQuery) AllX(ctx context.Context) []*UserTexture {
	nodes, err := utq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserTexture IDs.
func (utq *UserTextureQuery) IDs(ctx context.Context) (ids []int, err error) {
	if utq.ctx.Unique == nil && utq.path != nil {
		utq.Unique(true)
	}
	ctx = setContextOp(ctx, utq.ctx, "IDs")
	if err = utq.Select(usertexture.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (utq *UserTextureQuery) IDsX(ctx context.Context) []int {
	ids, err := utq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (utq *UserTextureQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, utq.ctx, "Count")
	if err := utq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, utq, querierCount[*UserTextureQuery](), utq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (utq *UserTextureQuery) CountX(ctx context.Context) int {
	count, err := utq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (utq *UserTextureQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, utq.ctx, "Exist")
	switch _, err := utq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (utq *UserTextureQuery) ExistX(ctx context.Context) bool {
	exist, err := utq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserTextureQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (utq *UserTextureQuery) Clone() *UserTextureQuery {
	if utq == nil {
		return nil
	}
	return &UserTextureQuery{
		config:          utq.config,
		ctx:             utq.ctx.Clone(),
		order:           append([]usertexture.OrderOption{}, utq.order...),
		inters:          append([]Interceptor{}, utq.inters...),
		predicates:      append([]predicate.UserTexture{}, utq.predicates...),
		withUserProfile: utq.withUserProfile.Clone(),
		withTexture:     utq.withTexture.Clone(),
		// clone intermediate query.
		sql:  utq.sql.Clone(),
		path: utq.path,
	}
}

// WithUserProfile tells the query-builder to eager-load the nodes that are connected to
// the "user_profile" edge. The optional arguments are used to configure the query builder of the edge.
func (utq *UserTextureQuery) WithUserProfile(opts ...func(*UserProfileQuery)) *UserTextureQuery {
	query := (&UserProfileClient{config: utq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	utq.withUserProfile = query
	return utq
}

// WithTexture tells the query-builder to eager-load the nodes that are connected to
// the "texture" edge. The optional arguments are used to configure the query builder of the edge.
func (utq *UserTextureQuery) WithTexture(opts ...func(*TextureQuery)) *UserTextureQuery {
	query := (&TextureClient{config: utq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	utq.withTexture = query
	return utq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserProfileID int `json:"user_profile_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserTexture.Query().
//		GroupBy(usertexture.FieldUserProfileID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (utq *UserTextureQuery) GroupBy(field string, fields ...string) *UserTextureGroupBy {
	utq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserTextureGroupBy{build: utq}
	grbuild.flds = &utq.ctx.Fields
	grbuild.label = usertexture.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserProfileID int `json:"user_profile_id,omitempty"`
//	}
//
//	client.UserTexture.Query().
//		Select(usertexture.FieldUserProfileID).
//		Scan(ctx, &v)
func (utq *UserTextureQuery) Select(fields ...string) *UserTextureSelect {
	utq.ctx.Fields = append(utq.ctx.Fields, fields...)
	sbuild := &UserTextureSelect{UserTextureQuery: utq}
	sbuild.label = usertexture.Label
	sbuild.flds, sbuild.scan = &utq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserTextureSelect configured with the given aggregations.
func (utq *UserTextureQuery) Aggregate(fns ...AggregateFunc) *UserTextureSelect {
	return utq.Select().Aggregate(fns...)
}

func (utq *UserTextureQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range utq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, utq); err != nil {
				return err
			}
		}
	}
	for _, f := range utq.ctx.Fields {
		if !usertexture.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if utq.path != nil {
		prev, err := utq.path(ctx)
		if err != nil {
			return err
		}
		utq.sql = prev
	}
	return nil
}

func (utq *UserTextureQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserTexture, error) {
	var (
		nodes       = []*UserTexture{}
		_spec       = utq.querySpec()
		loadedTypes = [2]bool{
			utq.withUserProfile != nil,
			utq.withTexture != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserTexture).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserTexture{config: utq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(utq.modifiers) > 0 {
		_spec.Modifiers = utq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, utq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := utq.withUserProfile; query != nil {
		if err := utq.loadUserProfile(ctx, query, nodes, nil,
			func(n *UserTexture, e *UserProfile) { n.Edges.UserProfile = e }); err != nil {
			return nil, err
		}
	}
	if query := utq.withTexture; query != nil {
		if err := utq.loadTexture(ctx, query, nodes, nil,
			func(n *UserTexture, e *Texture) { n.Edges.Texture = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (utq *UserTextureQuery) loadUserProfile(ctx context.Context, query *UserProfileQuery, nodes []*UserTexture, init func(*UserTexture), assign func(*UserTexture, *UserProfile)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserTexture)
	for i := range nodes {
		fk := nodes[i].UserProfileID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(userprofile.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_profile_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (utq *UserTextureQuery) loadTexture(ctx context.Context, query *TextureQuery, nodes []*UserTexture, init func(*UserTexture), assign func(*UserTexture, *Texture)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserTexture)
	for i := range nodes {
		fk := nodes[i].TextureID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(texture.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "texture_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (utq *UserTextureQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := utq.querySpec()
	if len(utq.modifiers) > 0 {
		_spec.Modifiers = utq.modifiers
	}
	_spec.Node.Columns = utq.ctx.Fields
	if len(utq.ctx.Fields) > 0 {
		_spec.Unique = utq.ctx.Unique != nil && *utq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, utq.driver, _spec)
}

func (utq *UserTextureQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(usertexture.Table, usertexture.Columns, sqlgraph.NewFieldSpec(usertexture.FieldID, field.TypeInt))
	_spec.From = utq.sql
	if unique := utq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if utq.path != nil {
		_spec.Unique = true
	}
	if fields := utq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usertexture.FieldID)
		for i := range fields {
			if fields[i] != usertexture.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if utq.withUserProfile != nil {
			_spec.Node.AddColumnOnce(usertexture.FieldUserProfileID)
		}
		if utq.withTexture != nil {
			_spec.Node.AddColumnOnce(usertexture.FieldTextureID)
		}
	}
	if ps := utq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := utq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := utq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := utq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (utq *UserTextureQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(utq.driver.Dialect())
	t1 := builder.Table(usertexture.Table)
	columns := utq.ctx.Fields
	if len(columns) == 0 {
		columns = usertexture.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if utq.sql != nil {
		selector = utq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if utq.ctx.Unique != nil && *utq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range utq.modifiers {
		m(selector)
	}
	for _, p := range utq.predicates {
		p(selector)
	}
	for _, p := range utq.order {
		p(selector)
	}
	if offset := utq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := utq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (utq *UserTextureQuery) ForUpdate(opts ...sql.LockOption) *UserTextureQuery {
	if utq.driver.Dialect() == dialect.Postgres {
		utq.Unique(false)
	}
	utq.modifiers = append(utq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return utq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (utq *UserTextureQuery) ForShare(opts ...sql.LockOption) *UserTextureQuery {
	if utq.driver.Dialect() == dialect.Postgres {
		utq.Unique(false)
	}
	utq.modifiers = append(utq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return utq
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (utq *UserTextureQuery) ForUpdateA(opts ...sql.LockOption) *UserTextureQuery {
	if utq.driver.Dialect() == dialect.SQLite {
		return utq
	}
	return utq.ForUpdate(opts...)
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (utq *UserTextureQuery) ForShareA(opts ...sql.LockOption) *UserTextureQuery {
	if utq.driver.Dialect() == dialect.SQLite {
		return utq
	}
	return utq.ForShare(opts...)
}

// UserTextureGroupBy is the group-by builder for UserTexture entities.
type UserTextureGroupBy struct {
	selector
	build *UserTextureQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (utgb *UserTextureGroupBy) Aggregate(fns ...AggregateFunc) *UserTextureGroupBy {
	utgb.fns = append(utgb.fns, fns...)
	return utgb
}

// Scan applies the selector query and scans the result into the given value.
func (utgb *UserTextureGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, utgb.build.ctx, "GroupBy")
	if err := utgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserTextureQuery, *UserTextureGroupBy](ctx, utgb.build, utgb, utgb.build.inters, v)
}

func (utgb *UserTextureGroupBy) sqlScan(ctx context.Context, root *UserTextureQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(utgb.fns))
	for _, fn := range utgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*utgb.flds)+len(utgb.fns))
		for _, f := range *utgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*utgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := utgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserTextureSelect is the builder for selecting fields of UserTexture entities.
type UserTextureSelect struct {
	*UserTextureQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (uts *UserTextureSelect) Aggregate(fns ...AggregateFunc) *UserTextureSelect {
	uts.fns = append(uts.fns, fns...)
	return uts
}

// Scan applies the selector query and scans the result into the given value.
func (uts *UserTextureSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uts.ctx, "Select")
	if err := uts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserTextureQuery, *UserTextureSelect](ctx, uts.UserTextureQuery, uts, uts.inters, v)
}

func (uts *UserTextureSelect) sqlScan(ctx context.Context, root *UserTextureQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(uts.fns))
	for _, fn := range uts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*uts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
