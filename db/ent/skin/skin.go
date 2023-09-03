// Code generated by ent, DO NOT EDIT.

package skin

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the skin type in the database.
	Label = "skin"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSkinHash holds the string denoting the skin_hash field in the database.
	FieldSkinHash = "skin_hash"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldVariant holds the string denoting the variant field in the database.
	FieldVariant = "variant"
	// EdgeCreatedUser holds the string denoting the created_user edge name in mutations.
	EdgeCreatedUser = "created_user"
	// Table holds the table name of the skin in the database.
	Table = "skins"
	// CreatedUserTable is the table that holds the created_user relation/edge.
	CreatedUserTable = "skins"
	// CreatedUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	CreatedUserInverseTable = "users"
	// CreatedUserColumn is the table column denoting the created_user relation/edge.
	CreatedUserColumn = "skin_created_user"
)

// Columns holds all SQL columns for skin fields.
var Columns = []string{
	FieldID,
	FieldSkinHash,
	FieldType,
	FieldVariant,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "skins"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"skin_created_user",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Skin queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySkinHash orders the results by the skin_hash field.
func BySkinHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSkinHash, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByVariant orders the results by the variant field.
func ByVariant(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVariant, opts...).ToFunc()
}

// ByCreatedUserField orders the results by created_user field.
func ByCreatedUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreatedUserStep(), sql.OrderByField(field, opts...))
	}
}
func newCreatedUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreatedUserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, CreatedUserTable, CreatedUserColumn),
	)
}
