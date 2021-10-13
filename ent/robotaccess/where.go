// Code generated by entc, DO NOT EDIT.

package robotaccess

import (
	"devlog/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
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
func IDGT(id int) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v string) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldToken), v))
	})
}

// Memo applies equality check predicate on the "memo" field. It's identical to MemoEQ.
func Memo(v string) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMemo), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// ExpiresAt applies equality check predicate on the "expires_at" field. It's identical to ExpiresAtEQ.
func ExpiresAt(v time.Time) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiresAt), v))
	})
}

// LastAccessAt applies equality check predicate on the "last_access_at" field. It's identical to LastAccessAtEQ.
func LastAccessAt(v time.Time) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastAccessAt), v))
	})
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v string) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldToken), v))
	})
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v string) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldToken), v))
	})
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...string) predicate.RobotAccess {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RobotAccess(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldToken), v...))
	})
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...string) predicate.RobotAccess {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RobotAccess(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldToken), v...))
	})
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v string) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldToken), v))
	})
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v string) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldToken), v))
	})
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v string) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldToken), v))
	})
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v string) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldToken), v))
	})
}

// TokenContains applies the Contains predicate on the "token" field.
func TokenContains(v string) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldToken), v))
	})
}

// TokenHasPrefix applies the HasPrefix predicate on the "token" field.
func TokenHasPrefix(v string) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldToken), v))
	})
}

// TokenHasSuffix applies the HasSuffix predicate on the "token" field.
func TokenHasSuffix(v string) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldToken), v))
	})
}

// TokenEqualFold applies the EqualFold predicate on the "token" field.
func TokenEqualFold(v string) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldToken), v))
	})
}

// TokenContainsFold applies the ContainsFold predicate on the "token" field.
func TokenContainsFold(v string) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldToken), v))
	})
}

// MemoEQ applies the EQ predicate on the "memo" field.
func MemoEQ(v string) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMemo), v))
	})
}

// MemoNEQ applies the NEQ predicate on the "memo" field.
func MemoNEQ(v string) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMemo), v))
	})
}

// MemoIn applies the In predicate on the "memo" field.
func MemoIn(vs ...string) predicate.RobotAccess {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RobotAccess(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMemo), v...))
	})
}

// MemoNotIn applies the NotIn predicate on the "memo" field.
func MemoNotIn(vs ...string) predicate.RobotAccess {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RobotAccess(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMemo), v...))
	})
}

// MemoGT applies the GT predicate on the "memo" field.
func MemoGT(v string) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMemo), v))
	})
}

// MemoGTE applies the GTE predicate on the "memo" field.
func MemoGTE(v string) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMemo), v))
	})
}

// MemoLT applies the LT predicate on the "memo" field.
func MemoLT(v string) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMemo), v))
	})
}

// MemoLTE applies the LTE predicate on the "memo" field.
func MemoLTE(v string) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMemo), v))
	})
}

// MemoContains applies the Contains predicate on the "memo" field.
func MemoContains(v string) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMemo), v))
	})
}

// MemoHasPrefix applies the HasPrefix predicate on the "memo" field.
func MemoHasPrefix(v string) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMemo), v))
	})
}

// MemoHasSuffix applies the HasSuffix predicate on the "memo" field.
func MemoHasSuffix(v string) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMemo), v))
	})
}

// MemoIsNil applies the IsNil predicate on the "memo" field.
func MemoIsNil() predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMemo)))
	})
}

// MemoNotNil applies the NotNil predicate on the "memo" field.
func MemoNotNil() predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMemo)))
	})
}

// MemoEqualFold applies the EqualFold predicate on the "memo" field.
func MemoEqualFold(v string) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMemo), v))
	})
}

// MemoContainsFold applies the ContainsFold predicate on the "memo" field.
func MemoContainsFold(v string) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMemo), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.RobotAccess {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RobotAccess(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.RobotAccess {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RobotAccess(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// ExpiresAtEQ applies the EQ predicate on the "expires_at" field.
func ExpiresAtEQ(v time.Time) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiresAt), v))
	})
}

// ExpiresAtNEQ applies the NEQ predicate on the "expires_at" field.
func ExpiresAtNEQ(v time.Time) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExpiresAt), v))
	})
}

// ExpiresAtIn applies the In predicate on the "expires_at" field.
func ExpiresAtIn(vs ...time.Time) predicate.RobotAccess {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RobotAccess(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExpiresAt), v...))
	})
}

// ExpiresAtNotIn applies the NotIn predicate on the "expires_at" field.
func ExpiresAtNotIn(vs ...time.Time) predicate.RobotAccess {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RobotAccess(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExpiresAt), v...))
	})
}

// ExpiresAtGT applies the GT predicate on the "expires_at" field.
func ExpiresAtGT(v time.Time) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExpiresAt), v))
	})
}

// ExpiresAtGTE applies the GTE predicate on the "expires_at" field.
func ExpiresAtGTE(v time.Time) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExpiresAt), v))
	})
}

// ExpiresAtLT applies the LT predicate on the "expires_at" field.
func ExpiresAtLT(v time.Time) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExpiresAt), v))
	})
}

// ExpiresAtLTE applies the LTE predicate on the "expires_at" field.
func ExpiresAtLTE(v time.Time) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExpiresAt), v))
	})
}

// ExpiresAtIsNil applies the IsNil predicate on the "expires_at" field.
func ExpiresAtIsNil() predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldExpiresAt)))
	})
}

// ExpiresAtNotNil applies the NotNil predicate on the "expires_at" field.
func ExpiresAtNotNil() predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldExpiresAt)))
	})
}

// LastAccessAtEQ applies the EQ predicate on the "last_access_at" field.
func LastAccessAtEQ(v time.Time) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastAccessAt), v))
	})
}

// LastAccessAtNEQ applies the NEQ predicate on the "last_access_at" field.
func LastAccessAtNEQ(v time.Time) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastAccessAt), v))
	})
}

// LastAccessAtIn applies the In predicate on the "last_access_at" field.
func LastAccessAtIn(vs ...time.Time) predicate.RobotAccess {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RobotAccess(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastAccessAt), v...))
	})
}

// LastAccessAtNotIn applies the NotIn predicate on the "last_access_at" field.
func LastAccessAtNotIn(vs ...time.Time) predicate.RobotAccess {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RobotAccess(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastAccessAt), v...))
	})
}

// LastAccessAtGT applies the GT predicate on the "last_access_at" field.
func LastAccessAtGT(v time.Time) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastAccessAt), v))
	})
}

// LastAccessAtGTE applies the GTE predicate on the "last_access_at" field.
func LastAccessAtGTE(v time.Time) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastAccessAt), v))
	})
}

// LastAccessAtLT applies the LT predicate on the "last_access_at" field.
func LastAccessAtLT(v time.Time) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastAccessAt), v))
	})
}

// LastAccessAtLTE applies the LTE predicate on the "last_access_at" field.
func LastAccessAtLTE(v time.Time) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastAccessAt), v))
	})
}

// LastAccessAtIsNil applies the IsNil predicate on the "last_access_at" field.
func LastAccessAtIsNil() predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLastAccessAt)))
	})
}

// LastAccessAtNotNil applies the NotNil predicate on the "last_access_at" field.
func LastAccessAtNotNil() predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLastAccessAt)))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.Admin) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RobotAccess) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RobotAccess) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
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
func Not(p predicate.RobotAccess) predicate.RobotAccess {
	return predicate.RobotAccess(func(s *sql.Selector) {
		p(s.Not())
	})
}
