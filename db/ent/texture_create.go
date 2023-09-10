// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/xmdhs/authlib-skin/db/ent/texture"
	"github.com/xmdhs/authlib-skin/db/ent/user"
	"github.com/xmdhs/authlib-skin/db/ent/userprofile"
	"github.com/xmdhs/authlib-skin/db/ent/usertexture"
)

// TextureCreate is the builder for creating a Texture entity.
type TextureCreate struct {
	config
	mutation *TextureMutation
	hooks    []Hook
}

// SetTextureHash sets the "texture_hash" field.
func (tc *TextureCreate) SetTextureHash(s string) *TextureCreate {
	tc.mutation.SetTextureHash(s)
	return tc
}

// SetCreatedUserID sets the "created_user" edge to the User entity by ID.
func (tc *TextureCreate) SetCreatedUserID(id int) *TextureCreate {
	tc.mutation.SetCreatedUserID(id)
	return tc
}

// SetCreatedUser sets the "created_user" edge to the User entity.
func (tc *TextureCreate) SetCreatedUser(u *User) *TextureCreate {
	return tc.SetCreatedUserID(u.ID)
}

// AddUserProfileIDs adds the "user_profile" edge to the UserProfile entity by IDs.
func (tc *TextureCreate) AddUserProfileIDs(ids ...int) *TextureCreate {
	tc.mutation.AddUserProfileIDs(ids...)
	return tc
}

// AddUserProfile adds the "user_profile" edges to the UserProfile entity.
func (tc *TextureCreate) AddUserProfile(u ...*UserProfile) *TextureCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tc.AddUserProfileIDs(ids...)
}

// AddUsertextureIDs adds the "usertexture" edge to the UserTexture entity by IDs.
func (tc *TextureCreate) AddUsertextureIDs(ids ...int) *TextureCreate {
	tc.mutation.AddUsertextureIDs(ids...)
	return tc
}

// AddUsertexture adds the "usertexture" edges to the UserTexture entity.
func (tc *TextureCreate) AddUsertexture(u ...*UserTexture) *TextureCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tc.AddUsertextureIDs(ids...)
}

// Mutation returns the TextureMutation object of the builder.
func (tc *TextureCreate) Mutation() *TextureMutation {
	return tc.mutation
}

// Save creates the Texture in the database.
func (tc *TextureCreate) Save(ctx context.Context) (*Texture, error) {
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TextureCreate) SaveX(ctx context.Context) *Texture {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TextureCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TextureCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TextureCreate) check() error {
	if _, ok := tc.mutation.TextureHash(); !ok {
		return &ValidationError{Name: "texture_hash", err: errors.New(`ent: missing required field "Texture.texture_hash"`)}
	}
	if _, ok := tc.mutation.CreatedUserID(); !ok {
		return &ValidationError{Name: "created_user", err: errors.New(`ent: missing required edge "Texture.created_user"`)}
	}
	return nil
}

func (tc *TextureCreate) sqlSave(ctx context.Context) (*Texture, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TextureCreate) createSpec() (*Texture, *sqlgraph.CreateSpec) {
	var (
		_node = &Texture{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(texture.Table, sqlgraph.NewFieldSpec(texture.FieldID, field.TypeInt))
	)
	if value, ok := tc.mutation.TextureHash(); ok {
		_spec.SetField(texture.FieldTextureHash, field.TypeString, value)
		_node.TextureHash = value
	}
	if nodes := tc.mutation.CreatedUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   texture.CreatedUserTable,
			Columns: []string{texture.CreatedUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.texture_created_user = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.UserProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   texture.UserProfileTable,
			Columns: texture.UserProfilePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userprofile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.UsertextureIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   texture.UsertextureTable,
			Columns: []string{texture.UsertextureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usertexture.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TextureCreateBulk is the builder for creating many Texture entities in bulk.
type TextureCreateBulk struct {
	config
	builders []*TextureCreate
}

// Save creates the Texture entities in the database.
func (tcb *TextureCreateBulk) Save(ctx context.Context) ([]*Texture, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Texture, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TextureMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TextureCreateBulk) SaveX(ctx context.Context) []*Texture {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TextureCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TextureCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}