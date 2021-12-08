// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/bug/internal/ent/discordguild"
	"entgo.io/ent/dialect/sql"
)

// DiscordGuild is the model entity for the DiscordGuild schema.
type DiscordGuild struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DiscordID holds the value of the "discord_id" field.
	DiscordID string `json:"discord_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Icon holds the value of the "icon" field.
	Icon string `json:"icon,omitempty"`
	// Banner holds the value of the "banner" field.
	Banner string `json:"banner,omitempty"`
	// Permissions holds the value of the "permissions" field.
	Permissions string `json:"permissions,omitempty"`
	// JoinedAt holds the value of the "joined_at" field.
	JoinedAt time.Time `json:"joined_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DiscordGuildQuery when eager-loading is set.
	Edges DiscordGuildEdges `json:"edges"`
}

// DiscordGuildEdges holds the relations/edges for other nodes in the graph.
type DiscordGuildEdges struct {
	// Roles holds the value of the roles edge.
	Roles []*DiscordRole `json:"roles,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e DiscordGuildEdges) RolesOrErr() ([]*DiscordRole, error) {
	if e.loadedTypes[0] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DiscordGuild) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case discordguild.FieldID:
			values[i] = new(sql.NullInt64)
		case discordguild.FieldDiscordID, discordguild.FieldName, discordguild.FieldDescription, discordguild.FieldIcon, discordguild.FieldBanner, discordguild.FieldPermissions:
			values[i] = new(sql.NullString)
		case discordguild.FieldJoinedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type DiscordGuild", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DiscordGuild fields.
func (dg *DiscordGuild) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case discordguild.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dg.ID = int(value.Int64)
		case discordguild.FieldDiscordID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field discord_id", values[i])
			} else if value.Valid {
				dg.DiscordID = value.String
			}
		case discordguild.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				dg.Name = value.String
			}
		case discordguild.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				dg.Description = value.String
			}
		case discordguild.FieldIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icon", values[i])
			} else if value.Valid {
				dg.Icon = value.String
			}
		case discordguild.FieldBanner:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field banner", values[i])
			} else if value.Valid {
				dg.Banner = value.String
			}
		case discordguild.FieldPermissions:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field permissions", values[i])
			} else if value.Valid {
				dg.Permissions = value.String
			}
		case discordguild.FieldJoinedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field joined_at", values[i])
			} else if value.Valid {
				dg.JoinedAt = value.Time
			}
		}
	}
	return nil
}

// QueryRoles queries the "roles" edge of the DiscordGuild entity.
func (dg *DiscordGuild) QueryRoles() *DiscordRoleQuery {
	return (&DiscordGuildClient{config: dg.config}).QueryRoles(dg)
}

// Update returns a builder for updating this DiscordGuild.
// Note that you need to call DiscordGuild.Unwrap() before calling this method if this DiscordGuild
// was returned from a transaction, and the transaction was committed or rolled back.
func (dg *DiscordGuild) Update() *DiscordGuildUpdateOne {
	return (&DiscordGuildClient{config: dg.config}).UpdateOne(dg)
}

// Unwrap unwraps the DiscordGuild entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dg *DiscordGuild) Unwrap() *DiscordGuild {
	tx, ok := dg.config.driver.(*txDriver)
	if !ok {
		panic("ent: DiscordGuild is not a transactional entity")
	}
	dg.config.driver = tx.drv
	return dg
}

// String implements the fmt.Stringer.
func (dg *DiscordGuild) String() string {
	var builder strings.Builder
	builder.WriteString("DiscordGuild(")
	builder.WriteString(fmt.Sprintf("id=%v", dg.ID))
	builder.WriteString(", discord_id=")
	builder.WriteString(dg.DiscordID)
	builder.WriteString(", name=")
	builder.WriteString(dg.Name)
	builder.WriteString(", description=")
	builder.WriteString(dg.Description)
	builder.WriteString(", icon=")
	builder.WriteString(dg.Icon)
	builder.WriteString(", banner=")
	builder.WriteString(dg.Banner)
	builder.WriteString(", permissions=")
	builder.WriteString(dg.Permissions)
	builder.WriteString(", joined_at=")
	builder.WriteString(dg.JoinedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// DiscordGuilds is a parsable slice of DiscordGuild.
type DiscordGuilds []*DiscordGuild

func (dg DiscordGuilds) config(cfg config) {
	for _i := range dg {
		dg[_i].config = cfg
	}
}