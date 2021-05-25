// Code generated by entc, DO NOT EDIT.

package unsavedpostthumbnail

import (
	"devlog/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Width applies equality check predicate on the "width" field. It's identical to WidthEQ.
func Width(v uint32) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWidth), v))
	})
}

// Height applies equality check predicate on the "height" field. It's identical to HeightEQ.
func Height(v uint32) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHeight), v))
	})
}

// Hash applies equality check predicate on the "hash" field. It's identical to HashEQ.
func Hash(v string) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHash), v))
	})
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURL), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// WidthEQ applies the EQ predicate on the "width" field.
func WidthEQ(v uint32) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWidth), v))
	})
}

// WidthNEQ applies the NEQ predicate on the "width" field.
func WidthNEQ(v uint32) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWidth), v))
	})
}

// WidthIn applies the In predicate on the "width" field.
func WidthIn(vs ...uint32) predicate.UnsavedPostThumbnail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWidth), v...))
	})
}

// WidthNotIn applies the NotIn predicate on the "width" field.
func WidthNotIn(vs ...uint32) predicate.UnsavedPostThumbnail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWidth), v...))
	})
}

// WidthGT applies the GT predicate on the "width" field.
func WidthGT(v uint32) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWidth), v))
	})
}

// WidthGTE applies the GTE predicate on the "width" field.
func WidthGTE(v uint32) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWidth), v))
	})
}

// WidthLT applies the LT predicate on the "width" field.
func WidthLT(v uint32) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWidth), v))
	})
}

// WidthLTE applies the LTE predicate on the "width" field.
func WidthLTE(v uint32) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWidth), v))
	})
}

// HeightEQ applies the EQ predicate on the "height" field.
func HeightEQ(v uint32) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHeight), v))
	})
}

// HeightNEQ applies the NEQ predicate on the "height" field.
func HeightNEQ(v uint32) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHeight), v))
	})
}

// HeightIn applies the In predicate on the "height" field.
func HeightIn(vs ...uint32) predicate.UnsavedPostThumbnail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHeight), v...))
	})
}

// HeightNotIn applies the NotIn predicate on the "height" field.
func HeightNotIn(vs ...uint32) predicate.UnsavedPostThumbnail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHeight), v...))
	})
}

// HeightGT applies the GT predicate on the "height" field.
func HeightGT(v uint32) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHeight), v))
	})
}

// HeightGTE applies the GTE predicate on the "height" field.
func HeightGTE(v uint32) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHeight), v))
	})
}

// HeightLT applies the LT predicate on the "height" field.
func HeightLT(v uint32) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHeight), v))
	})
}

// HeightLTE applies the LTE predicate on the "height" field.
func HeightLTE(v uint32) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHeight), v))
	})
}

// HashEQ applies the EQ predicate on the "hash" field.
func HashEQ(v string) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHash), v))
	})
}

// HashNEQ applies the NEQ predicate on the "hash" field.
func HashNEQ(v string) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHash), v))
	})
}

// HashIn applies the In predicate on the "hash" field.
func HashIn(vs ...string) predicate.UnsavedPostThumbnail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHash), v...))
	})
}

// HashNotIn applies the NotIn predicate on the "hash" field.
func HashNotIn(vs ...string) predicate.UnsavedPostThumbnail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHash), v...))
	})
}

// HashGT applies the GT predicate on the "hash" field.
func HashGT(v string) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHash), v))
	})
}

// HashGTE applies the GTE predicate on the "hash" field.
func HashGTE(v string) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHash), v))
	})
}

// HashLT applies the LT predicate on the "hash" field.
func HashLT(v string) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHash), v))
	})
}

// HashLTE applies the LTE predicate on the "hash" field.
func HashLTE(v string) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHash), v))
	})
}

// HashContains applies the Contains predicate on the "hash" field.
func HashContains(v string) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHash), v))
	})
}

// HashHasPrefix applies the HasPrefix predicate on the "hash" field.
func HashHasPrefix(v string) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHash), v))
	})
}

// HashHasSuffix applies the HasSuffix predicate on the "hash" field.
func HashHasSuffix(v string) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHash), v))
	})
}

// HashEqualFold applies the EqualFold predicate on the "hash" field.
func HashEqualFold(v string) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHash), v))
	})
}

// HashContainsFold applies the ContainsFold predicate on the "hash" field.
func HashContainsFold(v string) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHash), v))
	})
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURL), v))
	})
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldURL), v))
	})
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.UnsavedPostThumbnail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldURL), v...))
	})
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.UnsavedPostThumbnail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldURL), v...))
	})
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldURL), v))
	})
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldURL), v))
	})
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldURL), v))
	})
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldURL), v))
	})
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldURL), v))
	})
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldURL), v))
	})
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldURL), v))
	})
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldURL), v))
	})
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldURL), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.UnsavedPostThumbnail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.UnsavedPostThumbnail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// HasUnsavedPost applies the HasEdge predicate on the "unsaved_post" edge.
func HasUnsavedPost() predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UnsavedPostTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UnsavedPostTable, UnsavedPostColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUnsavedPostWith applies the HasEdge predicate on the "unsaved_post" edge with a given conditions (other predicates).
func HasUnsavedPostWith(preds ...predicate.UnsavedPost) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UnsavedPostInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UnsavedPostTable, UnsavedPostColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UnsavedPostThumbnail) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UnsavedPostThumbnail) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
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
func Not(p predicate.UnsavedPostThumbnail) predicate.UnsavedPostThumbnail {
	return predicate.UnsavedPostThumbnail(func(s *sql.Selector) {
		p(s.Not())
	})
}