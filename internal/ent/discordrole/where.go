// Code generated by entc, DO NOT EDIT.

package discordrole

import (
	"entgo.io/bug/internal/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// DiscordID applies equality check predicate on the "discord_id" field. It's identical to DiscordIDEQ.
func DiscordID(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiscordID), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Color applies equality check predicate on the "color" field. It's identical to ColorEQ.
func Color(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldColor), v))
	})
}

// Managed applies equality check predicate on the "managed" field. It's identical to ManagedEQ.
func Managed(v bool) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldManaged), v))
	})
}

// DiscordIDEQ applies the EQ predicate on the "discord_id" field.
func DiscordIDEQ(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiscordID), v))
	})
}

// DiscordIDNEQ applies the NEQ predicate on the "discord_id" field.
func DiscordIDNEQ(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDiscordID), v))
	})
}

// DiscordIDIn applies the In predicate on the "discord_id" field.
func DiscordIDIn(vs ...string) predicate.DiscordRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscordRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDiscordID), v...))
	})
}

// DiscordIDNotIn applies the NotIn predicate on the "discord_id" field.
func DiscordIDNotIn(vs ...string) predicate.DiscordRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscordRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDiscordID), v...))
	})
}

// DiscordIDGT applies the GT predicate on the "discord_id" field.
func DiscordIDGT(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDiscordID), v))
	})
}

// DiscordIDGTE applies the GTE predicate on the "discord_id" field.
func DiscordIDGTE(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDiscordID), v))
	})
}

// DiscordIDLT applies the LT predicate on the "discord_id" field.
func DiscordIDLT(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDiscordID), v))
	})
}

// DiscordIDLTE applies the LTE predicate on the "discord_id" field.
func DiscordIDLTE(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDiscordID), v))
	})
}

// DiscordIDContains applies the Contains predicate on the "discord_id" field.
func DiscordIDContains(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDiscordID), v))
	})
}

// DiscordIDHasPrefix applies the HasPrefix predicate on the "discord_id" field.
func DiscordIDHasPrefix(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDiscordID), v))
	})
}

// DiscordIDHasSuffix applies the HasSuffix predicate on the "discord_id" field.
func DiscordIDHasSuffix(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDiscordID), v))
	})
}

// DiscordIDEqualFold applies the EqualFold predicate on the "discord_id" field.
func DiscordIDEqualFold(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDiscordID), v))
	})
}

// DiscordIDContainsFold applies the ContainsFold predicate on the "discord_id" field.
func DiscordIDContainsFold(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDiscordID), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.DiscordRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscordRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.DiscordRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscordRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// ColorEQ applies the EQ predicate on the "color" field.
func ColorEQ(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldColor), v))
	})
}

// ColorNEQ applies the NEQ predicate on the "color" field.
func ColorNEQ(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldColor), v))
	})
}

// ColorIn applies the In predicate on the "color" field.
func ColorIn(vs ...string) predicate.DiscordRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscordRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldColor), v...))
	})
}

// ColorNotIn applies the NotIn predicate on the "color" field.
func ColorNotIn(vs ...string) predicate.DiscordRole {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscordRole(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldColor), v...))
	})
}

// ColorGT applies the GT predicate on the "color" field.
func ColorGT(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldColor), v))
	})
}

// ColorGTE applies the GTE predicate on the "color" field.
func ColorGTE(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldColor), v))
	})
}

// ColorLT applies the LT predicate on the "color" field.
func ColorLT(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldColor), v))
	})
}

// ColorLTE applies the LTE predicate on the "color" field.
func ColorLTE(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldColor), v))
	})
}

// ColorContains applies the Contains predicate on the "color" field.
func ColorContains(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldColor), v))
	})
}

// ColorHasPrefix applies the HasPrefix predicate on the "color" field.
func ColorHasPrefix(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldColor), v))
	})
}

// ColorHasSuffix applies the HasSuffix predicate on the "color" field.
func ColorHasSuffix(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldColor), v))
	})
}

// ColorIsNil applies the IsNil predicate on the "color" field.
func ColorIsNil() predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldColor)))
	})
}

// ColorNotNil applies the NotNil predicate on the "color" field.
func ColorNotNil() predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldColor)))
	})
}

// ColorEqualFold applies the EqualFold predicate on the "color" field.
func ColorEqualFold(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldColor), v))
	})
}

// ColorContainsFold applies the ContainsFold predicate on the "color" field.
func ColorContainsFold(v string) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldColor), v))
	})
}

// ManagedEQ applies the EQ predicate on the "managed" field.
func ManagedEQ(v bool) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldManaged), v))
	})
}

// ManagedNEQ applies the NEQ predicate on the "managed" field.
func ManagedNEQ(v bool) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldManaged), v))
	})
}

// HasGuild applies the HasEdge predicate on the "guild" edge.
func HasGuild() predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GuildTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GuildTable, GuildColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGuildWith applies the HasEdge predicate on the "guild" edge with a given conditions (other predicates).
func HasGuildWith(preds ...predicate.DiscordGuild) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GuildInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GuildTable, GuildColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DiscordRole) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DiscordRole) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DiscordRole) predicate.DiscordRole {
	return predicate.DiscordRole(func(s *sql.Selector) {
		p(s.Not())
	})
}
