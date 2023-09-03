// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/xmdhs/authlib-skin/db/ent/user"
	"github.com/xmdhs/authlib-skin/db/ent/userprofile"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Salt holds the value of the "salt" field.
	Salt string `json:"salt,omitempty"`
	// State holds the value of the "state" field.
	State int `json:"state,omitempty"`
	// RegTime holds the value of the "reg_time" field.
	RegTime int64 `json:"reg_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Skin holds the value of the skin edge.
	Skin []*Skin `json:"skin,omitempty"`
	// Profile holds the value of the profile edge.
	Profile *UserProfile `json:"profile,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SkinOrErr returns the Skin value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) SkinOrErr() ([]*Skin, error) {
	if e.loadedTypes[0] {
		return e.Skin, nil
	}
	return nil, &NotLoadedError{edge: "skin"}
}

// ProfileOrErr returns the Profile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) ProfileOrErr() (*UserProfile, error) {
	if e.loadedTypes[1] {
		if e.Profile == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: userprofile.Label}
		}
		return e.Profile, nil
	}
	return nil, &NotLoadedError{edge: "profile"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID, user.FieldState, user.FieldRegTime:
			values[i] = new(sql.NullInt64)
		case user.FieldEmail, user.FieldPassword, user.FieldSalt:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldSalt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field salt", values[i])
			} else if value.Valid {
				u.Salt = value.String
			}
		case user.FieldState:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				u.State = int(value.Int64)
			}
		case user.FieldRegTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field reg_time", values[i])
			} else if value.Valid {
				u.RegTime = value.Int64
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QuerySkin queries the "skin" edge of the User entity.
func (u *User) QuerySkin() *SkinQuery {
	return NewUserClient(u.config).QuerySkin(u)
}

// QueryProfile queries the "profile" edge of the User entity.
func (u *User) QueryProfile() *UserProfileQuery {
	return NewUserClient(u.config).QueryProfile(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(u.Password)
	builder.WriteString(", ")
	builder.WriteString("salt=")
	builder.WriteString(u.Salt)
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", u.State))
	builder.WriteString(", ")
	builder.WriteString("reg_time=")
	builder.WriteString(fmt.Sprintf("%v", u.RegTime))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
