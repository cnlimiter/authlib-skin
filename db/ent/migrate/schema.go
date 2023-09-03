// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SkinsColumns holds the columns for the "skins" table.
	SkinsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "skin_hash", Type: field.TypeString},
		{Name: "type", Type: field.TypeUint8},
		{Name: "variant", Type: field.TypeString},
		{Name: "skin_created_user", Type: field.TypeInt},
	}
	// SkinsTable holds the schema information for the "skins" table.
	SkinsTable = &schema.Table{
		Name:       "skins",
		Columns:    SkinsColumns,
		PrimaryKey: []*schema.Column{SkinsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "skins_users_created_user",
				Columns:    []*schema.Column{SkinsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "skin_skin_hash",
				Unique:  false,
				Columns: []*schema.Column{SkinsColumns[1]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "salt", Type: field.TypeString},
		{Name: "state", Type: field.TypeInt},
		{Name: "reg_time", Type: field.TypeInt64},
		{Name: "user_token", Type: field.TypeInt, Nullable: true},
		{Name: "user_skin", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_user_tokens_token",
				Columns:    []*schema.Column{UsersColumns[6]},
				RefColumns: []*schema.Column{UserTokensColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "users_skins_skin",
				Columns:    []*schema.Column{UsersColumns[7]},
				RefColumns: []*schema.Column{SkinsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserProfilesColumns holds the columns for the "user_profiles" table.
	UserProfilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "uuid", Type: field.TypeString},
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
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "userprofile_name",
				Unique:  false,
				Columns: []*schema.Column{UserProfilesColumns[1]},
			},
		},
	}
	// UserTokensColumns holds the columns for the "user_tokens" table.
	UserTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "token_id", Type: field.TypeUint64},
	}
	// UserTokensTable holds the schema information for the "user_tokens" table.
	UserTokensTable = &schema.Table{
		Name:       "user_tokens",
		Columns:    UserTokensColumns,
		PrimaryKey: []*schema.Column{UserTokensColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SkinsTable,
		UsersTable,
		UserProfilesTable,
		UserTokensTable,
	}
)

func init() {
	SkinsTable.ForeignKeys[0].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = UserTokensTable
	UsersTable.ForeignKeys[1].RefTable = SkinsTable
	UserProfilesTable.ForeignKeys[0].RefTable = UsersTable
}
