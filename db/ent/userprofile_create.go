// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xmdhs/authlib-skin/db/ent/user"
	"github.com/xmdhs/authlib-skin/db/ent/userprofile"
)

// UserProfileCreate is the builder for creating a UserProfile entity.
type UserProfileCreate struct {
	config
	mutation *UserProfileMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (upc *UserProfileCreate) SetName(s string) *UserProfileCreate {
	upc.mutation.SetName(s)
	return upc
}

// SetUUID sets the "uuid" field.
func (upc *UserProfileCreate) SetUUID(s string) *UserProfileCreate {
	upc.mutation.SetUUID(s)
	return upc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (upc *UserProfileCreate) SetUserID(id int) *UserProfileCreate {
	upc.mutation.SetUserID(id)
	return upc
}

// SetUser sets the "user" edge to the User entity.
func (upc *UserProfileCreate) SetUser(u *User) *UserProfileCreate {
	return upc.SetUserID(u.ID)
}

// Mutation returns the UserProfileMutation object of the builder.
func (upc *UserProfileCreate) Mutation() *UserProfileMutation {
	return upc.mutation
}

// Save creates the UserProfile in the database.
func (upc *UserProfileCreate) Save(ctx context.Context) (*UserProfile, error) {
	return withHooks(ctx, upc.sqlSave, upc.mutation, upc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (upc *UserProfileCreate) SaveX(ctx context.Context) *UserProfile {
	v, err := upc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (upc *UserProfileCreate) Exec(ctx context.Context) error {
	_, err := upc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upc *UserProfileCreate) ExecX(ctx context.Context) {
	if err := upc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upc *UserProfileCreate) check() error {
	if _, ok := upc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "UserProfile.name"`)}
	}
	if _, ok := upc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "UserProfile.uuid"`)}
	}
	if _, ok := upc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "UserProfile.user"`)}
	}
	return nil
}

func (upc *UserProfileCreate) sqlSave(ctx context.Context) (*UserProfile, error) {
	if err := upc.check(); err != nil {
		return nil, err
	}
	_node, _spec := upc.createSpec()
	if err := sqlgraph.CreateNode(ctx, upc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	upc.mutation.id = &_node.ID
	upc.mutation.done = true
	return _node, nil
}

func (upc *UserProfileCreate) createSpec() (*UserProfile, *sqlgraph.CreateSpec) {
	var (
		_node = &UserProfile{config: upc.config}
		_spec = sqlgraph.NewCreateSpec(userprofile.Table, sqlgraph.NewFieldSpec(userprofile.FieldID, field.TypeInt))
	)
	if value, ok := upc.mutation.Name(); ok {
		_spec.SetField(userprofile.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := upc.mutation.UUID(); ok {
		_spec.SetField(userprofile.FieldUUID, field.TypeString, value)
		_node.UUID = value
	}
	if nodes := upc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userprofile.UserTable,
			Columns: []string{userprofile.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_profile = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserProfileCreateBulk is the builder for creating many UserProfile entities in bulk.
type UserProfileCreateBulk struct {
	config
	builders []*UserProfileCreate
}

// Save creates the UserProfile entities in the database.
func (upcb *UserProfileCreateBulk) Save(ctx context.Context) ([]*UserProfile, error) {
	specs := make([]*sqlgraph.CreateSpec, len(upcb.builders))
	nodes := make([]*UserProfile, len(upcb.builders))
	mutators := make([]Mutator, len(upcb.builders))
	for i := range upcb.builders {
		func(i int, root context.Context) {
			builder := upcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserProfileMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, upcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, upcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, upcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (upcb *UserProfileCreateBulk) SaveX(ctx context.Context) []*UserProfile {
	v, err := upcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (upcb *UserProfileCreateBulk) Exec(ctx context.Context) error {
	_, err := upcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upcb *UserProfileCreateBulk) ExecX(ctx context.Context) {
	if err := upcb.Exec(ctx); err != nil {
		panic(err)
	}
}
