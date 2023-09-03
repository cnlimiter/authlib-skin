// Code generated by ent, DO NOT EDIT.

package skin

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/xmdhs/authlib-skin/db/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Skin {
	return predicate.Skin(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Skin {
	return predicate.Skin(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Skin {
	return predicate.Skin(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Skin {
	return predicate.Skin(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Skin {
	return predicate.Skin(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Skin {
	return predicate.Skin(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Skin {
	return predicate.Skin(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Skin {
	return predicate.Skin(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Skin {
	return predicate.Skin(sql.FieldLTE(FieldID, id))
}

// SkinHash applies equality check predicate on the "skin_hash" field. It's identical to SkinHashEQ.
func SkinHash(v string) predicate.Skin {
	return predicate.Skin(sql.FieldEQ(FieldSkinHash, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v uint8) predicate.Skin {
	return predicate.Skin(sql.FieldEQ(FieldType, v))
}

// Variant applies equality check predicate on the "variant" field. It's identical to VariantEQ.
func Variant(v string) predicate.Skin {
	return predicate.Skin(sql.FieldEQ(FieldVariant, v))
}

// SkinHashEQ applies the EQ predicate on the "skin_hash" field.
func SkinHashEQ(v string) predicate.Skin {
	return predicate.Skin(sql.FieldEQ(FieldSkinHash, v))
}

// SkinHashNEQ applies the NEQ predicate on the "skin_hash" field.
func SkinHashNEQ(v string) predicate.Skin {
	return predicate.Skin(sql.FieldNEQ(FieldSkinHash, v))
}

// SkinHashIn applies the In predicate on the "skin_hash" field.
func SkinHashIn(vs ...string) predicate.Skin {
	return predicate.Skin(sql.FieldIn(FieldSkinHash, vs...))
}

// SkinHashNotIn applies the NotIn predicate on the "skin_hash" field.
func SkinHashNotIn(vs ...string) predicate.Skin {
	return predicate.Skin(sql.FieldNotIn(FieldSkinHash, vs...))
}

// SkinHashGT applies the GT predicate on the "skin_hash" field.
func SkinHashGT(v string) predicate.Skin {
	return predicate.Skin(sql.FieldGT(FieldSkinHash, v))
}

// SkinHashGTE applies the GTE predicate on the "skin_hash" field.
func SkinHashGTE(v string) predicate.Skin {
	return predicate.Skin(sql.FieldGTE(FieldSkinHash, v))
}

// SkinHashLT applies the LT predicate on the "skin_hash" field.
func SkinHashLT(v string) predicate.Skin {
	return predicate.Skin(sql.FieldLT(FieldSkinHash, v))
}

// SkinHashLTE applies the LTE predicate on the "skin_hash" field.
func SkinHashLTE(v string) predicate.Skin {
	return predicate.Skin(sql.FieldLTE(FieldSkinHash, v))
}

// SkinHashContains applies the Contains predicate on the "skin_hash" field.
func SkinHashContains(v string) predicate.Skin {
	return predicate.Skin(sql.FieldContains(FieldSkinHash, v))
}

// SkinHashHasPrefix applies the HasPrefix predicate on the "skin_hash" field.
func SkinHashHasPrefix(v string) predicate.Skin {
	return predicate.Skin(sql.FieldHasPrefix(FieldSkinHash, v))
}

// SkinHashHasSuffix applies the HasSuffix predicate on the "skin_hash" field.
func SkinHashHasSuffix(v string) predicate.Skin {
	return predicate.Skin(sql.FieldHasSuffix(FieldSkinHash, v))
}

// SkinHashEqualFold applies the EqualFold predicate on the "skin_hash" field.
func SkinHashEqualFold(v string) predicate.Skin {
	return predicate.Skin(sql.FieldEqualFold(FieldSkinHash, v))
}

// SkinHashContainsFold applies the ContainsFold predicate on the "skin_hash" field.
func SkinHashContainsFold(v string) predicate.Skin {
	return predicate.Skin(sql.FieldContainsFold(FieldSkinHash, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v uint8) predicate.Skin {
	return predicate.Skin(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v uint8) predicate.Skin {
	return predicate.Skin(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...uint8) predicate.Skin {
	return predicate.Skin(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...uint8) predicate.Skin {
	return predicate.Skin(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v uint8) predicate.Skin {
	return predicate.Skin(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v uint8) predicate.Skin {
	return predicate.Skin(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v uint8) predicate.Skin {
	return predicate.Skin(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v uint8) predicate.Skin {
	return predicate.Skin(sql.FieldLTE(FieldType, v))
}

// VariantEQ applies the EQ predicate on the "variant" field.
func VariantEQ(v string) predicate.Skin {
	return predicate.Skin(sql.FieldEQ(FieldVariant, v))
}

// VariantNEQ applies the NEQ predicate on the "variant" field.
func VariantNEQ(v string) predicate.Skin {
	return predicate.Skin(sql.FieldNEQ(FieldVariant, v))
}

// VariantIn applies the In predicate on the "variant" field.
func VariantIn(vs ...string) predicate.Skin {
	return predicate.Skin(sql.FieldIn(FieldVariant, vs...))
}

// VariantNotIn applies the NotIn predicate on the "variant" field.
func VariantNotIn(vs ...string) predicate.Skin {
	return predicate.Skin(sql.FieldNotIn(FieldVariant, vs...))
}

// VariantGT applies the GT predicate on the "variant" field.
func VariantGT(v string) predicate.Skin {
	return predicate.Skin(sql.FieldGT(FieldVariant, v))
}

// VariantGTE applies the GTE predicate on the "variant" field.
func VariantGTE(v string) predicate.Skin {
	return predicate.Skin(sql.FieldGTE(FieldVariant, v))
}

// VariantLT applies the LT predicate on the "variant" field.
func VariantLT(v string) predicate.Skin {
	return predicate.Skin(sql.FieldLT(FieldVariant, v))
}

// VariantLTE applies the LTE predicate on the "variant" field.
func VariantLTE(v string) predicate.Skin {
	return predicate.Skin(sql.FieldLTE(FieldVariant, v))
}

// VariantContains applies the Contains predicate on the "variant" field.
func VariantContains(v string) predicate.Skin {
	return predicate.Skin(sql.FieldContains(FieldVariant, v))
}

// VariantHasPrefix applies the HasPrefix predicate on the "variant" field.
func VariantHasPrefix(v string) predicate.Skin {
	return predicate.Skin(sql.FieldHasPrefix(FieldVariant, v))
}

// VariantHasSuffix applies the HasSuffix predicate on the "variant" field.
func VariantHasSuffix(v string) predicate.Skin {
	return predicate.Skin(sql.FieldHasSuffix(FieldVariant, v))
}

// VariantEqualFold applies the EqualFold predicate on the "variant" field.
func VariantEqualFold(v string) predicate.Skin {
	return predicate.Skin(sql.FieldEqualFold(FieldVariant, v))
}

// VariantContainsFold applies the ContainsFold predicate on the "variant" field.
func VariantContainsFold(v string) predicate.Skin {
	return predicate.Skin(sql.FieldContainsFold(FieldVariant, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Skin {
	return predicate.Skin(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Skin {
	return predicate.Skin(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Skin) predicate.Skin {
	return predicate.Skin(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Skin) predicate.Skin {
	return predicate.Skin(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Skin) predicate.Skin {
	return predicate.Skin(func(s *sql.Selector) {
		p(s.Not())
	})
}
