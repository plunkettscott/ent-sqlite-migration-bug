// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/bug/internal/ent/discordguild"
	"entgo.io/bug/internal/ent/discordrole"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DiscordRoleCreate is the builder for creating a DiscordRole entity.
type DiscordRoleCreate struct {
	config
	mutation *DiscordRoleMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetDiscordID sets the "discord_id" field.
func (drc *DiscordRoleCreate) SetDiscordID(s string) *DiscordRoleCreate {
	drc.mutation.SetDiscordID(s)
	return drc
}

// SetName sets the "name" field.
func (drc *DiscordRoleCreate) SetName(s string) *DiscordRoleCreate {
	drc.mutation.SetName(s)
	return drc
}

// SetColor sets the "color" field.
func (drc *DiscordRoleCreate) SetColor(s string) *DiscordRoleCreate {
	drc.mutation.SetColor(s)
	return drc
}

// SetNillableColor sets the "color" field if the given value is not nil.
func (drc *DiscordRoleCreate) SetNillableColor(s *string) *DiscordRoleCreate {
	if s != nil {
		drc.SetColor(*s)
	}
	return drc
}

// SetManaged sets the "managed" field.
func (drc *DiscordRoleCreate) SetManaged(b bool) *DiscordRoleCreate {
	drc.mutation.SetManaged(b)
	return drc
}

// SetNillableManaged sets the "managed" field if the given value is not nil.
func (drc *DiscordRoleCreate) SetNillableManaged(b *bool) *DiscordRoleCreate {
	if b != nil {
		drc.SetManaged(*b)
	}
	return drc
}

// SetGuildID sets the "guild" edge to the DiscordGuild entity by ID.
func (drc *DiscordRoleCreate) SetGuildID(id int) *DiscordRoleCreate {
	drc.mutation.SetGuildID(id)
	return drc
}

// SetNillableGuildID sets the "guild" edge to the DiscordGuild entity by ID if the given value is not nil.
func (drc *DiscordRoleCreate) SetNillableGuildID(id *int) *DiscordRoleCreate {
	if id != nil {
		drc = drc.SetGuildID(*id)
	}
	return drc
}

// SetGuild sets the "guild" edge to the DiscordGuild entity.
func (drc *DiscordRoleCreate) SetGuild(d *DiscordGuild) *DiscordRoleCreate {
	return drc.SetGuildID(d.ID)
}

// Mutation returns the DiscordRoleMutation object of the builder.
func (drc *DiscordRoleCreate) Mutation() *DiscordRoleMutation {
	return drc.mutation
}

// Save creates the DiscordRole in the database.
func (drc *DiscordRoleCreate) Save(ctx context.Context) (*DiscordRole, error) {
	var (
		err  error
		node *DiscordRole
	)
	drc.defaults()
	if len(drc.hooks) == 0 {
		if err = drc.check(); err != nil {
			return nil, err
		}
		node, err = drc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiscordRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = drc.check(); err != nil {
				return nil, err
			}
			drc.mutation = mutation
			if node, err = drc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(drc.hooks) - 1; i >= 0; i-- {
			if drc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = drc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, drc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (drc *DiscordRoleCreate) SaveX(ctx context.Context) *DiscordRole {
	v, err := drc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (drc *DiscordRoleCreate) Exec(ctx context.Context) error {
	_, err := drc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (drc *DiscordRoleCreate) ExecX(ctx context.Context) {
	if err := drc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (drc *DiscordRoleCreate) defaults() {
	if _, ok := drc.mutation.Managed(); !ok {
		v := discordrole.DefaultManaged
		drc.mutation.SetManaged(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (drc *DiscordRoleCreate) check() error {
	if _, ok := drc.mutation.DiscordID(); !ok {
		return &ValidationError{Name: "discord_id", err: errors.New(`ent: missing required field "DiscordRole.discord_id"`)}
	}
	if _, ok := drc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "DiscordRole.name"`)}
	}
	if _, ok := drc.mutation.Managed(); !ok {
		return &ValidationError{Name: "managed", err: errors.New(`ent: missing required field "DiscordRole.managed"`)}
	}
	return nil
}

func (drc *DiscordRoleCreate) sqlSave(ctx context.Context) (*DiscordRole, error) {
	_node, _spec := drc.createSpec()
	if err := sqlgraph.CreateNode(ctx, drc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (drc *DiscordRoleCreate) createSpec() (*DiscordRole, *sqlgraph.CreateSpec) {
	var (
		_node = &DiscordRole{config: drc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: discordrole.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: discordrole.FieldID,
			},
		}
	)
	_spec.OnConflict = drc.conflict
	if value, ok := drc.mutation.DiscordID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: discordrole.FieldDiscordID,
		})
		_node.DiscordID = value
	}
	if value, ok := drc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: discordrole.FieldName,
		})
		_node.Name = value
	}
	if value, ok := drc.mutation.Color(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: discordrole.FieldColor,
		})
		_node.Color = &value
	}
	if value, ok := drc.mutation.Managed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: discordrole.FieldManaged,
		})
		_node.Managed = value
	}
	if nodes := drc.mutation.GuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   discordrole.GuildTable,
			Columns: []string{discordrole.GuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: discordguild.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.discord_guild_roles = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DiscordRole.Create().
//		SetDiscordID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DiscordRoleUpsert) {
//			SetDiscordID(v+v).
//		}).
//		Exec(ctx)
//
func (drc *DiscordRoleCreate) OnConflict(opts ...sql.ConflictOption) *DiscordRoleUpsertOne {
	drc.conflict = opts
	return &DiscordRoleUpsertOne{
		create: drc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DiscordRole.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (drc *DiscordRoleCreate) OnConflictColumns(columns ...string) *DiscordRoleUpsertOne {
	drc.conflict = append(drc.conflict, sql.ConflictColumns(columns...))
	return &DiscordRoleUpsertOne{
		create: drc,
	}
}

type (
	// DiscordRoleUpsertOne is the builder for "upsert"-ing
	//  one DiscordRole node.
	DiscordRoleUpsertOne struct {
		create *DiscordRoleCreate
	}

	// DiscordRoleUpsert is the "OnConflict" setter.
	DiscordRoleUpsert struct {
		*sql.UpdateSet
	}
)

// SetDiscordID sets the "discord_id" field.
func (u *DiscordRoleUpsert) SetDiscordID(v string) *DiscordRoleUpsert {
	u.Set(discordrole.FieldDiscordID, v)
	return u
}

// UpdateDiscordID sets the "discord_id" field to the value that was provided on create.
func (u *DiscordRoleUpsert) UpdateDiscordID() *DiscordRoleUpsert {
	u.SetExcluded(discordrole.FieldDiscordID)
	return u
}

// SetName sets the "name" field.
func (u *DiscordRoleUpsert) SetName(v string) *DiscordRoleUpsert {
	u.Set(discordrole.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DiscordRoleUpsert) UpdateName() *DiscordRoleUpsert {
	u.SetExcluded(discordrole.FieldName)
	return u
}

// SetColor sets the "color" field.
func (u *DiscordRoleUpsert) SetColor(v string) *DiscordRoleUpsert {
	u.Set(discordrole.FieldColor, v)
	return u
}

// UpdateColor sets the "color" field to the value that was provided on create.
func (u *DiscordRoleUpsert) UpdateColor() *DiscordRoleUpsert {
	u.SetExcluded(discordrole.FieldColor)
	return u
}

// ClearColor clears the value of the "color" field.
func (u *DiscordRoleUpsert) ClearColor() *DiscordRoleUpsert {
	u.SetNull(discordrole.FieldColor)
	return u
}

// SetManaged sets the "managed" field.
func (u *DiscordRoleUpsert) SetManaged(v bool) *DiscordRoleUpsert {
	u.Set(discordrole.FieldManaged, v)
	return u
}

// UpdateManaged sets the "managed" field to the value that was provided on create.
func (u *DiscordRoleUpsert) UpdateManaged() *DiscordRoleUpsert {
	u.SetExcluded(discordrole.FieldManaged)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.DiscordRole.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *DiscordRoleUpsertOne) UpdateNewValues() *DiscordRoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.DiscordRole.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *DiscordRoleUpsertOne) Ignore() *DiscordRoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DiscordRoleUpsertOne) DoNothing() *DiscordRoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DiscordRoleCreate.OnConflict
// documentation for more info.
func (u *DiscordRoleUpsertOne) Update(set func(*DiscordRoleUpsert)) *DiscordRoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DiscordRoleUpsert{UpdateSet: update})
	}))
	return u
}

// SetDiscordID sets the "discord_id" field.
func (u *DiscordRoleUpsertOne) SetDiscordID(v string) *DiscordRoleUpsertOne {
	return u.Update(func(s *DiscordRoleUpsert) {
		s.SetDiscordID(v)
	})
}

// UpdateDiscordID sets the "discord_id" field to the value that was provided on create.
func (u *DiscordRoleUpsertOne) UpdateDiscordID() *DiscordRoleUpsertOne {
	return u.Update(func(s *DiscordRoleUpsert) {
		s.UpdateDiscordID()
	})
}

// SetName sets the "name" field.
func (u *DiscordRoleUpsertOne) SetName(v string) *DiscordRoleUpsertOne {
	return u.Update(func(s *DiscordRoleUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DiscordRoleUpsertOne) UpdateName() *DiscordRoleUpsertOne {
	return u.Update(func(s *DiscordRoleUpsert) {
		s.UpdateName()
	})
}

// SetColor sets the "color" field.
func (u *DiscordRoleUpsertOne) SetColor(v string) *DiscordRoleUpsertOne {
	return u.Update(func(s *DiscordRoleUpsert) {
		s.SetColor(v)
	})
}

// UpdateColor sets the "color" field to the value that was provided on create.
func (u *DiscordRoleUpsertOne) UpdateColor() *DiscordRoleUpsertOne {
	return u.Update(func(s *DiscordRoleUpsert) {
		s.UpdateColor()
	})
}

// ClearColor clears the value of the "color" field.
func (u *DiscordRoleUpsertOne) ClearColor() *DiscordRoleUpsertOne {
	return u.Update(func(s *DiscordRoleUpsert) {
		s.ClearColor()
	})
}

// SetManaged sets the "managed" field.
func (u *DiscordRoleUpsertOne) SetManaged(v bool) *DiscordRoleUpsertOne {
	return u.Update(func(s *DiscordRoleUpsert) {
		s.SetManaged(v)
	})
}

// UpdateManaged sets the "managed" field to the value that was provided on create.
func (u *DiscordRoleUpsertOne) UpdateManaged() *DiscordRoleUpsertOne {
	return u.Update(func(s *DiscordRoleUpsert) {
		s.UpdateManaged()
	})
}

// Exec executes the query.
func (u *DiscordRoleUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DiscordRoleCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DiscordRoleUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DiscordRoleUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DiscordRoleUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DiscordRoleCreateBulk is the builder for creating many DiscordRole entities in bulk.
type DiscordRoleCreateBulk struct {
	config
	builders []*DiscordRoleCreate
	conflict []sql.ConflictOption
}

// Save creates the DiscordRole entities in the database.
func (drcb *DiscordRoleCreateBulk) Save(ctx context.Context) ([]*DiscordRole, error) {
	specs := make([]*sqlgraph.CreateSpec, len(drcb.builders))
	nodes := make([]*DiscordRole, len(drcb.builders))
	mutators := make([]Mutator, len(drcb.builders))
	for i := range drcb.builders {
		func(i int, root context.Context) {
			builder := drcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DiscordRoleMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, drcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = drcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, drcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, drcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (drcb *DiscordRoleCreateBulk) SaveX(ctx context.Context) []*DiscordRole {
	v, err := drcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (drcb *DiscordRoleCreateBulk) Exec(ctx context.Context) error {
	_, err := drcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (drcb *DiscordRoleCreateBulk) ExecX(ctx context.Context) {
	if err := drcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DiscordRole.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DiscordRoleUpsert) {
//			SetDiscordID(v+v).
//		}).
//		Exec(ctx)
//
func (drcb *DiscordRoleCreateBulk) OnConflict(opts ...sql.ConflictOption) *DiscordRoleUpsertBulk {
	drcb.conflict = opts
	return &DiscordRoleUpsertBulk{
		create: drcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DiscordRole.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (drcb *DiscordRoleCreateBulk) OnConflictColumns(columns ...string) *DiscordRoleUpsertBulk {
	drcb.conflict = append(drcb.conflict, sql.ConflictColumns(columns...))
	return &DiscordRoleUpsertBulk{
		create: drcb,
	}
}

// DiscordRoleUpsertBulk is the builder for "upsert"-ing
// a bulk of DiscordRole nodes.
type DiscordRoleUpsertBulk struct {
	create *DiscordRoleCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DiscordRole.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *DiscordRoleUpsertBulk) UpdateNewValues() *DiscordRoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DiscordRole.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *DiscordRoleUpsertBulk) Ignore() *DiscordRoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DiscordRoleUpsertBulk) DoNothing() *DiscordRoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DiscordRoleCreateBulk.OnConflict
// documentation for more info.
func (u *DiscordRoleUpsertBulk) Update(set func(*DiscordRoleUpsert)) *DiscordRoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DiscordRoleUpsert{UpdateSet: update})
	}))
	return u
}

// SetDiscordID sets the "discord_id" field.
func (u *DiscordRoleUpsertBulk) SetDiscordID(v string) *DiscordRoleUpsertBulk {
	return u.Update(func(s *DiscordRoleUpsert) {
		s.SetDiscordID(v)
	})
}

// UpdateDiscordID sets the "discord_id" field to the value that was provided on create.
func (u *DiscordRoleUpsertBulk) UpdateDiscordID() *DiscordRoleUpsertBulk {
	return u.Update(func(s *DiscordRoleUpsert) {
		s.UpdateDiscordID()
	})
}

// SetName sets the "name" field.
func (u *DiscordRoleUpsertBulk) SetName(v string) *DiscordRoleUpsertBulk {
	return u.Update(func(s *DiscordRoleUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DiscordRoleUpsertBulk) UpdateName() *DiscordRoleUpsertBulk {
	return u.Update(func(s *DiscordRoleUpsert) {
		s.UpdateName()
	})
}

// SetColor sets the "color" field.
func (u *DiscordRoleUpsertBulk) SetColor(v string) *DiscordRoleUpsertBulk {
	return u.Update(func(s *DiscordRoleUpsert) {
		s.SetColor(v)
	})
}

// UpdateColor sets the "color" field to the value that was provided on create.
func (u *DiscordRoleUpsertBulk) UpdateColor() *DiscordRoleUpsertBulk {
	return u.Update(func(s *DiscordRoleUpsert) {
		s.UpdateColor()
	})
}

// ClearColor clears the value of the "color" field.
func (u *DiscordRoleUpsertBulk) ClearColor() *DiscordRoleUpsertBulk {
	return u.Update(func(s *DiscordRoleUpsert) {
		s.ClearColor()
	})
}

// SetManaged sets the "managed" field.
func (u *DiscordRoleUpsertBulk) SetManaged(v bool) *DiscordRoleUpsertBulk {
	return u.Update(func(s *DiscordRoleUpsert) {
		s.SetManaged(v)
	})
}

// UpdateManaged sets the "managed" field to the value that was provided on create.
func (u *DiscordRoleUpsertBulk) UpdateManaged() *DiscordRoleUpsertBulk {
	return u.Update(func(s *DiscordRoleUpsert) {
		s.UpdateManaged()
	})
}

// Exec executes the query.
func (u *DiscordRoleUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DiscordRoleCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DiscordRoleCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DiscordRoleUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
