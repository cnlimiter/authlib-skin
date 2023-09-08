// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xmdhs/authlib-skin/db/ent/predicate"
	"github.com/xmdhs/authlib-skin/db/ent/texture"
	"github.com/xmdhs/authlib-skin/db/ent/user"
	"github.com/xmdhs/authlib-skin/db/ent/userprofile"
	"github.com/xmdhs/authlib-skin/db/ent/usertexture"
)

// UserProfileQuery is the builder for querying UserProfile entities.
type UserProfileQuery struct {
	config
	ctx             *QueryContext
	order           []userprofile.OrderOption
	inters          []Interceptor
	predicates      []predicate.UserProfile
	withUser        *UserQuery
	withTexture     *TextureQuery
	withUsertexture *UserTextureQuery
	withFKs         bool
	modifiers       []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserProfileQuery builder.
func (upq *UserProfileQuery) Where(ps ...predicate.UserProfile) *UserProfileQuery {
	upq.predicates = append(upq.predicates, ps...)
	return upq
}

// Limit the number of records to be returned by this query.
func (upq *UserProfileQuery) Limit(limit int) *UserProfileQuery {
	upq.ctx.Limit = &limit
	return upq
}

// Offset to start from.
func (upq *UserProfileQuery) Offset(offset int) *UserProfileQuery {
	upq.ctx.Offset = &offset
	return upq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (upq *UserProfileQuery) Unique(unique bool) *UserProfileQuery {
	upq.ctx.Unique = &unique
	return upq
}

// Order specifies how the records should be ordered.
func (upq *UserProfileQuery) Order(o ...userprofile.OrderOption) *UserProfileQuery {
	upq.order = append(upq.order, o...)
	return upq
}

// QueryUser chains the current query on the "user" edge.
func (upq *UserProfileQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: upq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := upq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := upq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userprofile.Table, userprofile.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, userprofile.UserTable, userprofile.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(upq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTexture chains the current query on the "texture" edge.
func (upq *UserProfileQuery) QueryTexture() *TextureQuery {
	query := (&TextureClient{config: upq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := upq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := upq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userprofile.Table, userprofile.FieldID, selector),
			sqlgraph.To(texture.Table, texture.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, userprofile.TextureTable, userprofile.TexturePrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(upq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUsertexture chains the current query on the "usertexture" edge.
func (upq *UserProfileQuery) QueryUsertexture() *UserTextureQuery {
	query := (&UserTextureClient{config: upq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := upq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := upq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userprofile.Table, userprofile.FieldID, selector),
			sqlgraph.To(usertexture.Table, usertexture.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, userprofile.UsertextureTable, userprofile.UsertextureColumn),
		)
		fromU = sqlgraph.SetNeighbors(upq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserProfile entity from the query.
// Returns a *NotFoundError when no UserProfile was found.
func (upq *UserProfileQuery) First(ctx context.Context) (*UserProfile, error) {
	nodes, err := upq.Limit(1).All(setContextOp(ctx, upq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userprofile.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (upq *UserProfileQuery) FirstX(ctx context.Context) *UserProfile {
	node, err := upq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserProfile ID from the query.
// Returns a *NotFoundError when no UserProfile ID was found.
func (upq *UserProfileQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = upq.Limit(1).IDs(setContextOp(ctx, upq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userprofile.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (upq *UserProfileQuery) FirstIDX(ctx context.Context) int {
	id, err := upq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserProfile entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserProfile entity is found.
// Returns a *NotFoundError when no UserProfile entities are found.
func (upq *UserProfileQuery) Only(ctx context.Context) (*UserProfile, error) {
	nodes, err := upq.Limit(2).All(setContextOp(ctx, upq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userprofile.Label}
	default:
		return nil, &NotSingularError{userprofile.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (upq *UserProfileQuery) OnlyX(ctx context.Context) *UserProfile {
	node, err := upq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserProfile ID in the query.
// Returns a *NotSingularError when more than one UserProfile ID is found.
// Returns a *NotFoundError when no entities are found.
func (upq *UserProfileQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = upq.Limit(2).IDs(setContextOp(ctx, upq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userprofile.Label}
	default:
		err = &NotSingularError{userprofile.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (upq *UserProfileQuery) OnlyIDX(ctx context.Context) int {
	id, err := upq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserProfiles.
func (upq *UserProfileQuery) All(ctx context.Context) ([]*UserProfile, error) {
	ctx = setContextOp(ctx, upq.ctx, "All")
	if err := upq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserProfile, *UserProfileQuery]()
	return withInterceptors[[]*UserProfile](ctx, upq, qr, upq.inters)
}

// AllX is like All, but panics if an error occurs.
func (upq *UserProfileQuery) AllX(ctx context.Context) []*UserProfile {
	nodes, err := upq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserProfile IDs.
func (upq *UserProfileQuery) IDs(ctx context.Context) (ids []int, err error) {
	if upq.ctx.Unique == nil && upq.path != nil {
		upq.Unique(true)
	}
	ctx = setContextOp(ctx, upq.ctx, "IDs")
	if err = upq.Select(userprofile.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (upq *UserProfileQuery) IDsX(ctx context.Context) []int {
	ids, err := upq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (upq *UserProfileQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, upq.ctx, "Count")
	if err := upq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, upq, querierCount[*UserProfileQuery](), upq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (upq *UserProfileQuery) CountX(ctx context.Context) int {
	count, err := upq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (upq *UserProfileQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, upq.ctx, "Exist")
	switch _, err := upq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (upq *UserProfileQuery) ExistX(ctx context.Context) bool {
	exist, err := upq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserProfileQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (upq *UserProfileQuery) Clone() *UserProfileQuery {
	if upq == nil {
		return nil
	}
	return &UserProfileQuery{
		config:          upq.config,
		ctx:             upq.ctx.Clone(),
		order:           append([]userprofile.OrderOption{}, upq.order...),
		inters:          append([]Interceptor{}, upq.inters...),
		predicates:      append([]predicate.UserProfile{}, upq.predicates...),
		withUser:        upq.withUser.Clone(),
		withTexture:     upq.withTexture.Clone(),
		withUsertexture: upq.withUsertexture.Clone(),
		// clone intermediate query.
		sql:  upq.sql.Clone(),
		path: upq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (upq *UserProfileQuery) WithUser(opts ...func(*UserQuery)) *UserProfileQuery {
	query := (&UserClient{config: upq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	upq.withUser = query
	return upq
}

// WithTexture tells the query-builder to eager-load the nodes that are connected to
// the "texture" edge. The optional arguments are used to configure the query builder of the edge.
func (upq *UserProfileQuery) WithTexture(opts ...func(*TextureQuery)) *UserProfileQuery {
	query := (&TextureClient{config: upq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	upq.withTexture = query
	return upq
}

// WithUsertexture tells the query-builder to eager-load the nodes that are connected to
// the "usertexture" edge. The optional arguments are used to configure the query builder of the edge.
func (upq *UserProfileQuery) WithUsertexture(opts ...func(*UserTextureQuery)) *UserProfileQuery {
	query := (&UserTextureClient{config: upq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	upq.withUsertexture = query
	return upq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserProfile.Query().
//		GroupBy(userprofile.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (upq *UserProfileQuery) GroupBy(field string, fields ...string) *UserProfileGroupBy {
	upq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserProfileGroupBy{build: upq}
	grbuild.flds = &upq.ctx.Fields
	grbuild.label = userprofile.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.UserProfile.Query().
//		Select(userprofile.FieldName).
//		Scan(ctx, &v)
func (upq *UserProfileQuery) Select(fields ...string) *UserProfileSelect {
	upq.ctx.Fields = append(upq.ctx.Fields, fields...)
	sbuild := &UserProfileSelect{UserProfileQuery: upq}
	sbuild.label = userprofile.Label
	sbuild.flds, sbuild.scan = &upq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserProfileSelect configured with the given aggregations.
func (upq *UserProfileQuery) Aggregate(fns ...AggregateFunc) *UserProfileSelect {
	return upq.Select().Aggregate(fns...)
}

func (upq *UserProfileQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range upq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, upq); err != nil {
				return err
			}
		}
	}
	for _, f := range upq.ctx.Fields {
		if !userprofile.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if upq.path != nil {
		prev, err := upq.path(ctx)
		if err != nil {
			return err
		}
		upq.sql = prev
	}
	return nil
}

func (upq *UserProfileQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserProfile, error) {
	var (
		nodes       = []*UserProfile{}
		withFKs     = upq.withFKs
		_spec       = upq.querySpec()
		loadedTypes = [3]bool{
			upq.withUser != nil,
			upq.withTexture != nil,
			upq.withUsertexture != nil,
		}
	)
	if upq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, userprofile.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserProfile).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserProfile{config: upq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(upq.modifiers) > 0 {
		_spec.Modifiers = upq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, upq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := upq.withUser; query != nil {
		if err := upq.loadUser(ctx, query, nodes, nil,
			func(n *UserProfile, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := upq.withTexture; query != nil {
		if err := upq.loadTexture(ctx, query, nodes,
			func(n *UserProfile) { n.Edges.Texture = []*Texture{} },
			func(n *UserProfile, e *Texture) { n.Edges.Texture = append(n.Edges.Texture, e) }); err != nil {
			return nil, err
		}
	}
	if query := upq.withUsertexture; query != nil {
		if err := upq.loadUsertexture(ctx, query, nodes,
			func(n *UserProfile) { n.Edges.Usertexture = []*UserTexture{} },
			func(n *UserProfile, e *UserTexture) { n.Edges.Usertexture = append(n.Edges.Usertexture, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (upq *UserProfileQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserProfile, init func(*UserProfile), assign func(*UserProfile, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserProfile)
	for i := range nodes {
		if nodes[i].user_profile == nil {
			continue
		}
		fk := *nodes[i].user_profile
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
			return fmt.Errorf(`unexpected foreign-key "user_profile" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (upq *UserProfileQuery) loadTexture(ctx context.Context, query *TextureQuery, nodes []*UserProfile, init func(*UserProfile), assign func(*UserProfile, *Texture)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*UserProfile)
	nids := make(map[int]map[*UserProfile]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(userprofile.TextureTable)
		s.Join(joinT).On(s.C(texture.FieldID), joinT.C(userprofile.TexturePrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(userprofile.TexturePrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(userprofile.TexturePrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*UserProfile]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Texture](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "texture" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (upq *UserProfileQuery) loadUsertexture(ctx context.Context, query *UserTextureQuery, nodes []*UserProfile, init func(*UserProfile), assign func(*UserProfile, *UserTexture)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*UserProfile)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(usertexture.FieldUserProfileID)
	}
	query.Where(predicate.UserTexture(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(userprofile.UsertextureColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.UserProfileID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_profile_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (upq *UserProfileQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := upq.querySpec()
	if len(upq.modifiers) > 0 {
		_spec.Modifiers = upq.modifiers
	}
	_spec.Node.Columns = upq.ctx.Fields
	if len(upq.ctx.Fields) > 0 {
		_spec.Unique = upq.ctx.Unique != nil && *upq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, upq.driver, _spec)
}

func (upq *UserProfileQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(userprofile.Table, userprofile.Columns, sqlgraph.NewFieldSpec(userprofile.FieldID, field.TypeInt))
	_spec.From = upq.sql
	if unique := upq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if upq.path != nil {
		_spec.Unique = true
	}
	if fields := upq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userprofile.FieldID)
		for i := range fields {
			if fields[i] != userprofile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := upq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := upq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := upq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := upq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (upq *UserProfileQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(upq.driver.Dialect())
	t1 := builder.Table(userprofile.Table)
	columns := upq.ctx.Fields
	if len(columns) == 0 {
		columns = userprofile.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if upq.sql != nil {
		selector = upq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if upq.ctx.Unique != nil && *upq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range upq.modifiers {
		m(selector)
	}
	for _, p := range upq.predicates {
		p(selector)
	}
	for _, p := range upq.order {
		p(selector)
	}
	if offset := upq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := upq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (upq *UserProfileQuery) ForUpdate(opts ...sql.LockOption) *UserProfileQuery {
	if upq.driver.Dialect() == dialect.Postgres {
		upq.Unique(false)
	}
	upq.modifiers = append(upq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return upq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (upq *UserProfileQuery) ForShare(opts ...sql.LockOption) *UserProfileQuery {
	if upq.driver.Dialect() == dialect.Postgres {
		upq.Unique(false)
	}
	upq.modifiers = append(upq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return upq
}

// UserProfileGroupBy is the group-by builder for UserProfile entities.
type UserProfileGroupBy struct {
	selector
	build *UserProfileQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (upgb *UserProfileGroupBy) Aggregate(fns ...AggregateFunc) *UserProfileGroupBy {
	upgb.fns = append(upgb.fns, fns...)
	return upgb
}

// Scan applies the selector query and scans the result into the given value.
func (upgb *UserProfileGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, upgb.build.ctx, "GroupBy")
	if err := upgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserProfileQuery, *UserProfileGroupBy](ctx, upgb.build, upgb, upgb.build.inters, v)
}

func (upgb *UserProfileGroupBy) sqlScan(ctx context.Context, root *UserProfileQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(upgb.fns))
	for _, fn := range upgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*upgb.flds)+len(upgb.fns))
		for _, f := range *upgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*upgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := upgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserProfileSelect is the builder for selecting fields of UserProfile entities.
type UserProfileSelect struct {
	*UserProfileQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ups *UserProfileSelect) Aggregate(fns ...AggregateFunc) *UserProfileSelect {
	ups.fns = append(ups.fns, fns...)
	return ups
}

// Scan applies the selector query and scans the result into the given value.
func (ups *UserProfileSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ups.ctx, "Select")
	if err := ups.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserProfileQuery, *UserProfileSelect](ctx, ups.UserProfileQuery, ups, ups.inters, v)
}

func (ups *UserProfileSelect) sqlScan(ctx context.Context, root *UserProfileQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ups.fns))
	for _, fn := range ups.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ups.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ups.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
