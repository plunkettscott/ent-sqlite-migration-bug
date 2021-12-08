package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DiscordGuild holds the schema definition for the DiscordGuild entity.
type DiscordGuild struct {
	ent.Schema
}

// Fields of the DiscordGuild.
func (DiscordGuild) Fields() []ent.Field {
	return []ent.Field{
		field.String("discord_id"),
		field.String("name"),
		field.String("description"),
		field.String("icon").
			Optional(),
		field.String("banner").
			Optional(),
		field.String("permissions"),
		field.Time("joined_at"),
	}
}

func (DiscordGuild) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

// Edges of the DiscordGuild.
func (DiscordGuild) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", DiscordRole.Type),
	}
}
