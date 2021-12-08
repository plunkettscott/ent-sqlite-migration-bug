package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DiscordRole holds the schema definition for the DiscordRole entity.
type DiscordRole struct {
	ent.Schema
}

// Fields of the DiscordRole.
func (DiscordRole) Fields() []ent.Field {
	return []ent.Field{
		field.String("discord_id"),
		field.String("name"),
		field.String("color").
			Optional().
			Nillable(),
		field.Bool("managed").
			Default(false),
	}
}

// Edges of the DiscordRole.
func (DiscordRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("guild", DiscordGuild.Type).
			Ref("roles").
			Unique().
			Annotations(),
	}
}
