// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/bug/internal/ent/discordguild"
	"entgo.io/bug/internal/ent/discordrole"
	"entgo.io/bug/internal/ent/predicate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DiscordRoleQuery is the builder for querying DiscordRole entities.
type DiscordRoleQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.DiscordRole
	// eager-loading edges.
	withGuild *DiscordGuildQuery
	withFKs   bool
	modifiers []func(s *sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DiscordRoleQuery builder.
func (drq *DiscordRoleQuery) Where(ps ...predicate.DiscordRole) *DiscordRoleQuery {
	drq.predicates = append(drq.predicates, ps...)
	return drq
}

// Limit adds a limit step to the query.
func (drq *DiscordRoleQuery) Limit(limit int) *DiscordRoleQuery {
	drq.limit = &limit
	return drq
}

// Offset adds an offset step to the query.
func (drq *DiscordRoleQuery) Offset(offset int) *DiscordRoleQuery {
	drq.offset = &offset
	return drq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (drq *DiscordRoleQuery) Unique(unique bool) *DiscordRoleQuery {
	drq.unique = &unique
	return drq
}

// Order adds an order step to the query.
func (drq *DiscordRoleQuery) Order(o ...OrderFunc) *DiscordRoleQuery {
	drq.order = append(drq.order, o...)
	return drq
}

// QueryGuild chains the current query on the "guild" edge.
func (drq *DiscordRoleQuery) QueryGuild() *DiscordGuildQuery {
	query := &DiscordGuildQuery{config: drq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := drq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := drq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(discordrole.Table, discordrole.FieldID, selector),
			sqlgraph.To(discordguild.Table, discordguild.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, discordrole.GuildTable, discordrole.GuildColumn),
		)
		fromU = sqlgraph.SetNeighbors(drq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DiscordRole entity from the query.
// Returns a *NotFoundError when no DiscordRole was found.
func (drq *DiscordRoleQuery) First(ctx context.Context) (*DiscordRole, error) {
	nodes, err := drq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{discordrole.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (drq *DiscordRoleQuery) FirstX(ctx context.Context) *DiscordRole {
	node, err := drq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DiscordRole ID from the query.
// Returns a *NotFoundError when no DiscordRole ID was found.
func (drq *DiscordRoleQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = drq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{discordrole.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (drq *DiscordRoleQuery) FirstIDX(ctx context.Context) int {
	id, err := drq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DiscordRole entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one DiscordRole entity is not found.
// Returns a *NotFoundError when no DiscordRole entities are found.
func (drq *DiscordRoleQuery) Only(ctx context.Context) (*DiscordRole, error) {
	nodes, err := drq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{discordrole.Label}
	default:
		return nil, &NotSingularError{discordrole.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (drq *DiscordRoleQuery) OnlyX(ctx context.Context) *DiscordRole {
	node, err := drq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DiscordRole ID in the query.
// Returns a *NotSingularError when exactly one DiscordRole ID is not found.
// Returns a *NotFoundError when no entities are found.
func (drq *DiscordRoleQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = drq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{discordrole.Label}
	default:
		err = &NotSingularError{discordrole.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (drq *DiscordRoleQuery) OnlyIDX(ctx context.Context) int {
	id, err := drq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DiscordRoles.
func (drq *DiscordRoleQuery) All(ctx context.Context) ([]*DiscordRole, error) {
	if err := drq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return drq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (drq *DiscordRoleQuery) AllX(ctx context.Context) []*DiscordRole {
	nodes, err := drq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DiscordRole IDs.
func (drq *DiscordRoleQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := drq.Select(discordrole.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (drq *DiscordRoleQuery) IDsX(ctx context.Context) []int {
	ids, err := drq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (drq *DiscordRoleQuery) Count(ctx context.Context) (int, error) {
	if err := drq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return drq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (drq *DiscordRoleQuery) CountX(ctx context.Context) int {
	count, err := drq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (drq *DiscordRoleQuery) Exist(ctx context.Context) (bool, error) {
	if err := drq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return drq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (drq *DiscordRoleQuery) ExistX(ctx context.Context) bool {
	exist, err := drq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DiscordRoleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (drq *DiscordRoleQuery) Clone() *DiscordRoleQuery {
	if drq == nil {
		return nil
	}
	return &DiscordRoleQuery{
		config:     drq.config,
		limit:      drq.limit,
		offset:     drq.offset,
		order:      append([]OrderFunc{}, drq.order...),
		predicates: append([]predicate.DiscordRole{}, drq.predicates...),
		withGuild:  drq.withGuild.Clone(),
		// clone intermediate query.
		sql:  drq.sql.Clone(),
		path: drq.path,
	}
}

// WithGuild tells the query-builder to eager-load the nodes that are connected to
// the "guild" edge. The optional arguments are used to configure the query builder of the edge.
func (drq *DiscordRoleQuery) WithGuild(opts ...func(*DiscordGuildQuery)) *DiscordRoleQuery {
	query := &DiscordGuildQuery{config: drq.config}
	for _, opt := range opts {
		opt(query)
	}
	drq.withGuild = query
	return drq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DiscordID string `json:"discord_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DiscordRole.Query().
//		GroupBy(discordrole.FieldDiscordID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (drq *DiscordRoleQuery) GroupBy(field string, fields ...string) *DiscordRoleGroupBy {
	group := &DiscordRoleGroupBy{config: drq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := drq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return drq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DiscordID string `json:"discord_id,omitempty"`
//	}
//
//	client.DiscordRole.Query().
//		Select(discordrole.FieldDiscordID).
//		Scan(ctx, &v)
//
func (drq *DiscordRoleQuery) Select(fields ...string) *DiscordRoleSelect {
	drq.fields = append(drq.fields, fields...)
	return &DiscordRoleSelect{DiscordRoleQuery: drq}
}

func (drq *DiscordRoleQuery) prepareQuery(ctx context.Context) error {
	for _, f := range drq.fields {
		if !discordrole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if drq.path != nil {
		prev, err := drq.path(ctx)
		if err != nil {
			return err
		}
		drq.sql = prev
	}
	return nil
}

func (drq *DiscordRoleQuery) sqlAll(ctx context.Context) ([]*DiscordRole, error) {
	var (
		nodes       = []*DiscordRole{}
		withFKs     = drq.withFKs
		_spec       = drq.querySpec()
		loadedTypes = [1]bool{
			drq.withGuild != nil,
		}
	)
	if drq.withGuild != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, discordrole.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &DiscordRole{config: drq.config}
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
	if len(drq.modifiers) > 0 {
		_spec.Modifiers = drq.modifiers
	}
	if err := sqlgraph.QueryNodes(ctx, drq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := drq.withGuild; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*DiscordRole)
		for i := range nodes {
			if nodes[i].discord_guild_roles == nil {
				continue
			}
			fk := *nodes[i].discord_guild_roles
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(discordguild.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "discord_guild_roles" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Guild = n
			}
		}
	}

	return nodes, nil
}

func (drq *DiscordRoleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := drq.querySpec()
	if len(drq.modifiers) > 0 {
		_spec.Modifiers = drq.modifiers
	}
	_spec.Node.Columns = drq.fields
	if len(drq.fields) > 0 {
		_spec.Unique = drq.unique != nil && *drq.unique
	}
	return sqlgraph.CountNodes(ctx, drq.driver, _spec)
}

func (drq *DiscordRoleQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := drq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (drq *DiscordRoleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   discordrole.Table,
			Columns: discordrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: discordrole.FieldID,
			},
		},
		From:   drq.sql,
		Unique: true,
	}
	if unique := drq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := drq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, discordrole.FieldID)
		for i := range fields {
			if fields[i] != discordrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := drq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := drq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := drq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := drq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (drq *DiscordRoleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(drq.driver.Dialect())
	t1 := builder.Table(discordrole.Table)
	columns := drq.fields
	if len(columns) == 0 {
		columns = discordrole.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if drq.sql != nil {
		selector = drq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if drq.unique != nil && *drq.unique {
		selector.Distinct()
	}
	for _, m := range drq.modifiers {
		m(selector)
	}
	for _, p := range drq.predicates {
		p(selector)
	}
	for _, p := range drq.order {
		p(selector)
	}
	if offset := drq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := drq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (drq *DiscordRoleQuery) ForUpdate(opts ...sql.LockOption) *DiscordRoleQuery {
	if drq.driver.Dialect() == dialect.Postgres {
		drq.Unique(false)
	}
	drq.modifiers = append(drq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return drq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (drq *DiscordRoleQuery) ForShare(opts ...sql.LockOption) *DiscordRoleQuery {
	if drq.driver.Dialect() == dialect.Postgres {
		drq.Unique(false)
	}
	drq.modifiers = append(drq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return drq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (drq *DiscordRoleQuery) Modify(modifiers ...func(s *sql.Selector)) *DiscordRoleSelect {
	drq.modifiers = append(drq.modifiers, modifiers...)
	return drq.Select()
}

// DiscordRoleGroupBy is the group-by builder for DiscordRole entities.
type DiscordRoleGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (drgb *DiscordRoleGroupBy) Aggregate(fns ...AggregateFunc) *DiscordRoleGroupBy {
	drgb.fns = append(drgb.fns, fns...)
	return drgb
}

// Scan applies the group-by query and scans the result into the given value.
func (drgb *DiscordRoleGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := drgb.path(ctx)
	if err != nil {
		return err
	}
	drgb.sql = query
	return drgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (drgb *DiscordRoleGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := drgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (drgb *DiscordRoleGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(drgb.fields) > 1 {
		return nil, errors.New("ent: DiscordRoleGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := drgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (drgb *DiscordRoleGroupBy) StringsX(ctx context.Context) []string {
	v, err := drgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (drgb *DiscordRoleGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = drgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discordrole.Label}
	default:
		err = fmt.Errorf("ent: DiscordRoleGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (drgb *DiscordRoleGroupBy) StringX(ctx context.Context) string {
	v, err := drgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (drgb *DiscordRoleGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(drgb.fields) > 1 {
		return nil, errors.New("ent: DiscordRoleGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := drgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (drgb *DiscordRoleGroupBy) IntsX(ctx context.Context) []int {
	v, err := drgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (drgb *DiscordRoleGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = drgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discordrole.Label}
	default:
		err = fmt.Errorf("ent: DiscordRoleGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (drgb *DiscordRoleGroupBy) IntX(ctx context.Context) int {
	v, err := drgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (drgb *DiscordRoleGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(drgb.fields) > 1 {
		return nil, errors.New("ent: DiscordRoleGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := drgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (drgb *DiscordRoleGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := drgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (drgb *DiscordRoleGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = drgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discordrole.Label}
	default:
		err = fmt.Errorf("ent: DiscordRoleGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (drgb *DiscordRoleGroupBy) Float64X(ctx context.Context) float64 {
	v, err := drgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (drgb *DiscordRoleGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(drgb.fields) > 1 {
		return nil, errors.New("ent: DiscordRoleGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := drgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (drgb *DiscordRoleGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := drgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (drgb *DiscordRoleGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = drgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discordrole.Label}
	default:
		err = fmt.Errorf("ent: DiscordRoleGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (drgb *DiscordRoleGroupBy) BoolX(ctx context.Context) bool {
	v, err := drgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (drgb *DiscordRoleGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range drgb.fields {
		if !discordrole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := drgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := drgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (drgb *DiscordRoleGroupBy) sqlQuery() *sql.Selector {
	selector := drgb.sql.Select()
	aggregation := make([]string, 0, len(drgb.fns))
	for _, fn := range drgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(drgb.fields)+len(drgb.fns))
		for _, f := range drgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(drgb.fields...)...)
}

// DiscordRoleSelect is the builder for selecting fields of DiscordRole entities.
type DiscordRoleSelect struct {
	*DiscordRoleQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (drs *DiscordRoleSelect) Scan(ctx context.Context, v interface{}) error {
	if err := drs.prepareQuery(ctx); err != nil {
		return err
	}
	drs.sql = drs.DiscordRoleQuery.sqlQuery(ctx)
	return drs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (drs *DiscordRoleSelect) ScanX(ctx context.Context, v interface{}) {
	if err := drs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (drs *DiscordRoleSelect) Strings(ctx context.Context) ([]string, error) {
	if len(drs.fields) > 1 {
		return nil, errors.New("ent: DiscordRoleSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := drs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (drs *DiscordRoleSelect) StringsX(ctx context.Context) []string {
	v, err := drs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (drs *DiscordRoleSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = drs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discordrole.Label}
	default:
		err = fmt.Errorf("ent: DiscordRoleSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (drs *DiscordRoleSelect) StringX(ctx context.Context) string {
	v, err := drs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (drs *DiscordRoleSelect) Ints(ctx context.Context) ([]int, error) {
	if len(drs.fields) > 1 {
		return nil, errors.New("ent: DiscordRoleSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := drs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (drs *DiscordRoleSelect) IntsX(ctx context.Context) []int {
	v, err := drs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (drs *DiscordRoleSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = drs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discordrole.Label}
	default:
		err = fmt.Errorf("ent: DiscordRoleSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (drs *DiscordRoleSelect) IntX(ctx context.Context) int {
	v, err := drs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (drs *DiscordRoleSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(drs.fields) > 1 {
		return nil, errors.New("ent: DiscordRoleSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := drs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (drs *DiscordRoleSelect) Float64sX(ctx context.Context) []float64 {
	v, err := drs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (drs *DiscordRoleSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = drs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discordrole.Label}
	default:
		err = fmt.Errorf("ent: DiscordRoleSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (drs *DiscordRoleSelect) Float64X(ctx context.Context) float64 {
	v, err := drs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (drs *DiscordRoleSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(drs.fields) > 1 {
		return nil, errors.New("ent: DiscordRoleSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := drs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (drs *DiscordRoleSelect) BoolsX(ctx context.Context) []bool {
	v, err := drs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (drs *DiscordRoleSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = drs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discordrole.Label}
	default:
		err = fmt.Errorf("ent: DiscordRoleSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (drs *DiscordRoleSelect) BoolX(ctx context.Context) bool {
	v, err := drs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (drs *DiscordRoleSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := drs.sql.Query()
	if err := drs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (drs *DiscordRoleSelect) Modify(modifiers ...func(s *sql.Selector)) *DiscordRoleSelect {
	drs.modifiers = append(drs.modifiers, modifiers...)
	return drs
}