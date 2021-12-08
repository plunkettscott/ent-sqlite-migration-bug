// Code generated by entc, DO NOT EDIT.

package discordguild

const (
	// Label holds the string label denoting the discordguild type in the database.
	Label = "discord_guild"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDiscordID holds the string denoting the discord_id field in the database.
	FieldDiscordID = "discord_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldIcon holds the string denoting the icon field in the database.
	FieldIcon = "icon"
	// FieldBanner holds the string denoting the banner field in the database.
	FieldBanner = "banner"
	// FieldPermissions holds the string denoting the permissions field in the database.
	FieldPermissions = "permissions"
	// FieldJoinedAt holds the string denoting the joined_at field in the database.
	FieldJoinedAt = "joined_at"
	// EdgeRoles holds the string denoting the roles edge name in mutations.
	EdgeRoles = "roles"
	// Table holds the table name of the discordguild in the database.
	Table = "discord_guilds"
	// RolesTable is the table that holds the roles relation/edge.
	RolesTable = "discord_roles"
	// RolesInverseTable is the table name for the DiscordRole entity.
	// It exists in this package in order to avoid circular dependency with the "discordrole" package.
	RolesInverseTable = "discord_roles"
	// RolesColumn is the table column denoting the roles relation/edge.
	RolesColumn = "discord_guild_roles"
)

// Columns holds all SQL columns for discordguild fields.
var Columns = []string{
	FieldID,
	FieldDiscordID,
	FieldName,
	FieldDescription,
	FieldIcon,
	FieldBanner,
	FieldPermissions,
	FieldJoinedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}