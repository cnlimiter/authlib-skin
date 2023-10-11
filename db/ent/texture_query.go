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

// TextureQuery is the builder for querying Texture entities.
type TextureQuery struct {
	config
	ctx             *QueryContext
	order           []texture.OrderOption
	inters          []Interceptor
	predicates      []predicate.Texture
	withCreatedUser *UserQuery
	withUserProfile *UserProfileQuery
	withUsertexture *UserTextureQuery
	withFKs         bool
	modifiers       []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TextureQuery builder.
func (tq *TextureQuery) Where(ps ...predicate.Texture) *TextureQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TextureQuery) Limit(limit int) *TextureQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TextureQuery) Offset(offset int) *TextureQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TextureQuery) Unique(unique bool) *TextureQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TextureQuery) Order(o ...texture.OrderOption) *TextureQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryCreatedUser chains the current query on the "created_user" edge.
func (tq *TextureQuery) QueryCreatedUser() *UserQuery {
	query := (&UserClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(texture.Table, texture.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, texture.CreatedUserTable, texture.CreatedUserColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUserProfile chains the current query on the "user_profile" edge.
func (tq *TextureQuery) QueryUserProfile() *UserProfileQuery {
	query := (&UserProfileClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(texture.Table, texture.FieldID, selector),
			sqlgraph.To(userprofile.Table, userprofile.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, texture.UserProfileTable, texture.UserProfilePrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUsertexture chains the current query on the "usertexture" edge.
func (tq *TextureQuery) QueryUsertexture() *UserTextureQuery {
	query := (&UserTextureClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(texture.Table, texture.FieldID, selector),
			sqlgraph.To(usertexture.Table, usertexture.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, texture.UsertextureTable, texture.UsertextureColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Texture entity from the query.
// Returns a *NotFoundError when no Texture was found.
func (tq *TextureQuery) First(ctx context.Context) (*Texture, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{texture.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TextureQuery) FirstX(ctx context.Context) *Texture {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Texture ID from the query.
// Returns a *NotFoundError when no Texture ID was found.
func (tq *TextureQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{texture.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TextureQuery) FirstIDX(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Texture entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Texture entity is found.
// Returns a *NotFoundError when no Texture entities are found.
func (tq *TextureQuery) Only(ctx context.Context) (*Texture, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{texture.Label}
	default:
		return nil, &NotSingularError{texture.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TextureQuery) OnlyX(ctx context.Context) *Texture {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Texture ID in the query.
// Returns a *NotSingularError when more than one Texture ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TextureQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{texture.Label}
	default:
		err = &NotSingularError{texture.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TextureQuery) OnlyIDX(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Textures.
func (tq *TextureQuery) All(ctx context.Context) ([]*Texture, error) {
	ctx = setContextOp(ctx, tq.ctx, "All")
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Texture, *TextureQuery]()
	return withInterceptors[[]*Texture](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TextureQuery) AllX(ctx context.Context) []*Texture {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Texture IDs.
func (tq *TextureQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, "IDs")
	if err = tq.Select(texture.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TextureQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TextureQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, "Count")
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TextureQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TextureQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TextureQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, "Exist")
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TextureQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TextureQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TextureQuery) Clone() *TextureQuery {
	if tq == nil {
		return nil
	}
	return &TextureQuery{
		config:          tq.config,
		ctx:             tq.ctx.Clone(),
		order:           append([]texture.OrderOption{}, tq.order...),
		inters:          append([]Interceptor{}, tq.inters...),
		predicates:      append([]predicate.Texture{}, tq.predicates...),
		withCreatedUser: tq.withCreatedUser.Clone(),
		withUserProfile: tq.withUserProfile.Clone(),
		withUsertexture: tq.withUsertexture.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithCreatedUser tells the query-builder to eager-load the nodes that are connected to
// the "created_user" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TextureQuery) WithCreatedUser(opts ...func(*UserQuery)) *TextureQuery {
	query := (&UserClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withCreatedUser = query
	return tq
}

// WithUserProfile tells the query-builder to eager-load the nodes that are connected to
// the "user_profile" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TextureQuery) WithUserProfile(opts ...func(*UserProfileQuery)) *TextureQuery {
	query := (&UserProfileClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withUserProfile = query
	return tq
}

// WithUsertexture tells the query-builder to eager-load the nodes that are connected to
// the "usertexture" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TextureQuery) WithUsertexture(opts ...func(*UserTextureQuery)) *TextureQuery {
	query := (&UserTextureClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withUsertexture = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TextureHash string `json:"texture_hash,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Texture.Query().
//		GroupBy(texture.FieldTextureHash).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TextureQuery) GroupBy(field string, fields ...string) *TextureGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TextureGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = texture.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TextureHash string `json:"texture_hash,omitempty"`
//	}
//
//	client.Texture.Query().
//		Select(texture.FieldTextureHash).
//		Scan(ctx, &v)
func (tq *TextureQuery) Select(fields ...string) *TextureSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TextureSelect{TextureQuery: tq}
	sbuild.label = texture.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TextureSelect configured with the given aggregations.
func (tq *TextureQuery) Aggregate(fns ...AggregateFunc) *TextureSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TextureQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !texture.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TextureQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Texture, error) {
	var (
		nodes       = []*Texture{}
		withFKs     = tq.withFKs
		_spec       = tq.querySpec()
		loadedTypes = [3]bool{
			tq.withCreatedUser != nil,
			tq.withUserProfile != nil,
			tq.withUsertexture != nil,
		}
	)
	if tq.withCreatedUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, texture.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Texture).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Texture{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withCreatedUser; query != nil {
		if err := tq.loadCreatedUser(ctx, query, nodes, nil,
			func(n *Texture, e *User) { n.Edges.CreatedUser = e }); err != nil {
			return nil, err
		}
	}
	if query := tq.withUserProfile; query != nil {
		if err := tq.loadUserProfile(ctx, query, nodes,
			func(n *Texture) { n.Edges.UserProfile = []*UserProfile{} },
			func(n *Texture, e *UserProfile) { n.Edges.UserProfile = append(n.Edges.UserProfile, e) }); err != nil {
			return nil, err
		}
	}
	if query := tq.withUsertexture; query != nil {
		if err := tq.loadUsertexture(ctx, query, nodes,
			func(n *Texture) { n.Edges.Usertexture = []*UserTexture{} },
			func(n *Texture, e *UserTexture) { n.Edges.Usertexture = append(n.Edges.Usertexture, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TextureQuery) loadCreatedUser(ctx context.Context, query *UserQuery, nodes []*Texture, init func(*Texture), assign func(*Texture, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Texture)
	for i := range nodes {
		if nodes[i].texture_created_user == nil {
			continue
		}
		fk := *nodes[i].texture_created_user
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
			return fmt.Errorf(`unexpected foreign-key "texture_created_user" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tq *TextureQuery) loadUserProfile(ctx context.Context, query *UserProfileQuery, nodes []*Texture, init func(*Texture), assign func(*Texture, *UserProfile)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Texture)
	nids := make(map[int]map[*Texture]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(texture.UserProfileTable)
		s.Join(joinT).On(s.C(userprofile.FieldID), joinT.C(texture.UserProfilePrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(texture.UserProfilePrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(texture.UserProfilePrimaryKey[0]))
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
					nids[inValue] = map[*Texture]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*UserProfile](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "user_profile" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (tq *TextureQuery) loadUsertexture(ctx context.Context, query *UserTextureQuery, nodes []*Texture, init func(*Texture), assign func(*Texture, *UserTexture)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Texture)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(usertexture.FieldTextureID)
	}
	query.Where(predicate.UserTexture(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(texture.UsertextureColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.TextureID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "texture_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (tq *TextureQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TextureQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(texture.Table, texture.Columns, sqlgraph.NewFieldSpec(texture.FieldID, field.TypeInt))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, texture.FieldID)
		for i := range fields {
			if fields[i] != texture.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TextureQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(texture.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = texture.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range tq.modifiers {
		m(selector)
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (tq *TextureQuery) ForUpdate(opts ...sql.LockOption) *TextureQuery {
	if tq.driver.Dialect() == dialect.Postgres {
		tq.Unique(false)
	}
	tq.modifiers = append(tq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return tq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (tq *TextureQuery) ForShare(opts ...sql.LockOption) *TextureQuery {
	if tq.driver.Dialect() == dialect.Postgres {
		tq.Unique(false)
	}
	tq.modifiers = append(tq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return tq
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (tq *TextureQuery) ForUpdateA(opts ...sql.LockOption) *TextureQuery {
	if tq.driver.Dialect() == dialect.SQLite {
		return tq
	}
	return tq.ForUpdate(opts...)
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (tq *TextureQuery) ForShareA(opts ...sql.LockOption) *TextureQuery {
	if tq.driver.Dialect() == dialect.SQLite {
		return tq
	}
	return tq.ForShare(opts...)
}

// TextureGroupBy is the group-by builder for Texture entities.
type TextureGroupBy struct {
	selector
	build *TextureQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TextureGroupBy) Aggregate(fns ...AggregateFunc) *TextureGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TextureGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, "GroupBy")
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TextureQuery, *TextureGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TextureGroupBy) sqlScan(ctx context.Context, root *TextureQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TextureSelect is the builder for selecting fields of Texture entities.
type TextureSelect struct {
	*TextureQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TextureSelect) Aggregate(fns ...AggregateFunc) *TextureSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TextureSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, "Select")
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TextureQuery, *TextureSelect](ctx, ts.TextureQuery, ts, ts.inters, v)
}

func (ts *TextureSelect) sqlScan(ctx context.Context, root *TextureQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
