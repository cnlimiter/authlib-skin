// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TexturesColumns holds the columns for the "textures" table.
	TexturesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "texture_hash", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(100)"}},
		{Name: "texture_created_user", Type: field.TypeInt},
	}
	// TexturesTable holds the schema information for the "textures" table.
	TexturesTable = &schema.Table{
		Name:       "textures",
		Columns:    TexturesColumns,
		PrimaryKey: []*schema.Column{TexturesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "textures_users_created_user",
				Columns:    []*schema.Column{TexturesColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "texture_texture_hash",
				Unique:  true,
				Columns: []*schema.Column{TexturesColumns[1]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"mysql": "VARCHAR(30)"}},
		{Name: "password", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(80)"}},
		{Name: "salt", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(50)"}},
		{Name: "reg_ip", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(32)"}},
		{Name: "state", Type: field.TypeInt},
		{Name: "reg_time", Type: field.TypeInt64},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_email",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[1]},
			},
			{
				Name:    "user_reg_ip",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[4]},
			},
		},
	}
	// UserProfilesColumns holds the columns for the "user_profiles" table.
	UserProfilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"mysql": "VARCHAR(20)"}},
		{Name: "uuid", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(32)"}},
		{Name: "user_profile", Type: field.TypeInt, Unique: true},
	}
	// UserProfilesTable holds the schema information for the "user_profiles" table.
	UserProfilesTable = &schema.Table{
		Name:       "user_profiles",
		Columns:    UserProfilesColumns,
		PrimaryKey: []*schema.Column{UserProfilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_profiles_users_profile",
				Columns:    []*schema.Column{UserProfilesColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "userprofile_user_profile",
				Unique:  false,
				Columns: []*schema.Column{UserProfilesColumns[3]},
			},
			{
				Name:    "userprofile_name",
				Unique:  false,
				Columns: []*schema.Column{UserProfilesColumns[1]},
			},
			{
				Name:    "userprofile_uuid",
				Unique:  false,
				Columns: []*schema.Column{UserProfilesColumns[2]},
			},
		},
	}
	// UserTexturesColumns holds the columns for the "user_textures" table.
	UserTexturesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(10)"}},
		{Name: "variant", Type: field.TypeString, SchemaType: map[string]string{"mysql": "VARCHAR(10)"}},
		{Name: "user_profile_id", Type: field.TypeInt},
		{Name: "texture_id", Type: field.TypeInt},
	}
	// UserTexturesTable holds the schema information for the "user_textures" table.
	UserTexturesTable = &schema.Table{
		Name:       "user_textures",
		Columns:    UserTexturesColumns,
		PrimaryKey: []*schema.Column{UserTexturesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_textures_user_profiles_user_profile",
				Columns:    []*schema.Column{UserTexturesColumns[3]},
				RefColumns: []*schema.Column{UserProfilesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "user_textures_textures_texture",
				Columns:    []*schema.Column{UserTexturesColumns[4]},
				RefColumns: []*schema.Column{TexturesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "usertexture_user_profile_id",
				Unique:  false,
				Columns: []*schema.Column{UserTexturesColumns[3]},
			},
			{
				Name:    "usertexture_texture_id",
				Unique:  false,
				Columns: []*schema.Column{UserTexturesColumns[4]},
			},
			{
				Name:    "usertexture_texture_id_user_profile_id",
				Unique:  true,
				Columns: []*schema.Column{UserTexturesColumns[4], UserTexturesColumns[3]},
			},
		},
	}
	// UserTokensColumns holds the columns for the "user_tokens" table.
	UserTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "token_id", Type: field.TypeUint64},
		{Name: "user_token", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// UserTokensTable holds the schema information for the "user_tokens" table.
	UserTokensTable = &schema.Table{
		Name:       "user_tokens",
		Columns:    UserTokensColumns,
		PrimaryKey: []*schema.Column{UserTokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_tokens_users_token",
				Columns:    []*schema.Column{UserTokensColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "usertoken_user_token",
				Unique:  false,
				Columns: []*schema.Column{UserTokensColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TexturesTable,
		UsersTable,
		UserProfilesTable,
		UserTexturesTable,
		UserTokensTable,
	}
)

func init() {
	TexturesTable.ForeignKeys[0].RefTable = UsersTable
	UserProfilesTable.ForeignKeys[0].RefTable = UsersTable
	UserTexturesTable.ForeignKeys[0].RefTable = UserProfilesTable
	UserTexturesTable.ForeignKeys[1].RefTable = TexturesTable
	UserTokensTable.ForeignKeys[0].RefTable = UsersTable
}
