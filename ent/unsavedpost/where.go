// Code generated by entc, DO NOT EDIT.

package unsavedpost

import (
	"devlog/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
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
func IDGT(id int) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUUID), v))
	})
}

// Slug applies equality check predicate on the "slug" field. It's identical to SlugEQ.
func Slug(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSlug), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// HTMLContent applies equality check predicate on the "html_content" field. It's identical to HTMLContentEQ.
func HTMLContent(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHTMLContent), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// ModifiedAt applies equality check predicate on the "modified_at" field. It's identical to ModifiedAtEQ.
func ModifiedAt(v time.Time) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModifiedAt), v))
	})
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUUID), v))
	})
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUUID), v))
	})
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...string) predicate.UnsavedPost {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UnsavedPost(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUUID), v...))
	})
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...string) predicate.UnsavedPost {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UnsavedPost(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUUID), v...))
	})
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUUID), v))
	})
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUUID), v))
	})
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUUID), v))
	})
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUUID), v))
	})
}

// UUIDContains applies the Contains predicate on the "uuid" field.
func UUIDContains(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUUID), v))
	})
}

// UUIDHasPrefix applies the HasPrefix predicate on the "uuid" field.
func UUIDHasPrefix(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUUID), v))
	})
}

// UUIDHasSuffix applies the HasSuffix predicate on the "uuid" field.
func UUIDHasSuffix(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUUID), v))
	})
}

// UUIDEqualFold applies the EqualFold predicate on the "uuid" field.
func UUIDEqualFold(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUUID), v))
	})
}

// UUIDContainsFold applies the ContainsFold predicate on the "uuid" field.
func UUIDContainsFold(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUUID), v))
	})
}

// SlugEQ applies the EQ predicate on the "slug" field.
func SlugEQ(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSlug), v))
	})
}

// SlugNEQ applies the NEQ predicate on the "slug" field.
func SlugNEQ(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSlug), v))
	})
}

// SlugIn applies the In predicate on the "slug" field.
func SlugIn(vs ...string) predicate.UnsavedPost {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UnsavedPost(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSlug), v...))
	})
}

// SlugNotIn applies the NotIn predicate on the "slug" field.
func SlugNotIn(vs ...string) predicate.UnsavedPost {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UnsavedPost(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSlug), v...))
	})
}

// SlugGT applies the GT predicate on the "slug" field.
func SlugGT(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSlug), v))
	})
}

// SlugGTE applies the GTE predicate on the "slug" field.
func SlugGTE(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSlug), v))
	})
}

// SlugLT applies the LT predicate on the "slug" field.
func SlugLT(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSlug), v))
	})
}

// SlugLTE applies the LTE predicate on the "slug" field.
func SlugLTE(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSlug), v))
	})
}

// SlugContains applies the Contains predicate on the "slug" field.
func SlugContains(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSlug), v))
	})
}

// SlugHasPrefix applies the HasPrefix predicate on the "slug" field.
func SlugHasPrefix(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSlug), v))
	})
}

// SlugHasSuffix applies the HasSuffix predicate on the "slug" field.
func SlugHasSuffix(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSlug), v))
	})
}

// SlugIsNil applies the IsNil predicate on the "slug" field.
func SlugIsNil() predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSlug)))
	})
}

// SlugNotNil applies the NotNil predicate on the "slug" field.
func SlugNotNil() predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSlug)))
	})
}

// SlugEqualFold applies the EqualFold predicate on the "slug" field.
func SlugEqualFold(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSlug), v))
	})
}

// SlugContainsFold applies the ContainsFold predicate on the "slug" field.
func SlugContainsFold(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSlug), v))
	})
}

// AccessLevelEQ applies the EQ predicate on the "access_level" field.
func AccessLevelEQ(v AccessLevel) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccessLevel), v))
	})
}

// AccessLevelNEQ applies the NEQ predicate on the "access_level" field.
func AccessLevelNEQ(v AccessLevel) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAccessLevel), v))
	})
}

// AccessLevelIn applies the In predicate on the "access_level" field.
func AccessLevelIn(vs ...AccessLevel) predicate.UnsavedPost {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UnsavedPost(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAccessLevel), v...))
	})
}

// AccessLevelNotIn applies the NotIn predicate on the "access_level" field.
func AccessLevelNotIn(vs ...AccessLevel) predicate.UnsavedPost {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UnsavedPost(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAccessLevel), v...))
	})
}

// AccessLevelIsNil applies the IsNil predicate on the "access_level" field.
func AccessLevelIsNil() predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAccessLevel)))
	})
}

// AccessLevelNotNil applies the NotNil predicate on the "access_level" field.
func AccessLevelNotNil() predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAccessLevel)))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.UnsavedPost {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UnsavedPost(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.UnsavedPost {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UnsavedPost(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleIsNil applies the IsNil predicate on the "title" field.
func TitleIsNil() predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTitle)))
	})
}

// TitleNotNil applies the NotNil predicate on the "title" field.
func TitleNotNil() predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTitle)))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContent), v))
	})
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.UnsavedPost {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UnsavedPost(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldContent), v...))
	})
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.UnsavedPost {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UnsavedPost(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldContent), v...))
	})
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContent), v))
	})
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContent), v))
	})
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContent), v))
	})
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContent), v))
	})
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContent), v))
	})
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContent), v))
	})
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContent), v))
	})
}

// ContentIsNil applies the IsNil predicate on the "content" field.
func ContentIsNil() predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldContent)))
	})
}

// ContentNotNil applies the NotNil predicate on the "content" field.
func ContentNotNil() predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldContent)))
	})
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContent), v))
	})
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContent), v))
	})
}

// HTMLContentEQ applies the EQ predicate on the "html_content" field.
func HTMLContentEQ(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHTMLContent), v))
	})
}

// HTMLContentNEQ applies the NEQ predicate on the "html_content" field.
func HTMLContentNEQ(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHTMLContent), v))
	})
}

// HTMLContentIn applies the In predicate on the "html_content" field.
func HTMLContentIn(vs ...string) predicate.UnsavedPost {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UnsavedPost(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHTMLContent), v...))
	})
}

// HTMLContentNotIn applies the NotIn predicate on the "html_content" field.
func HTMLContentNotIn(vs ...string) predicate.UnsavedPost {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UnsavedPost(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHTMLContent), v...))
	})
}

// HTMLContentGT applies the GT predicate on the "html_content" field.
func HTMLContentGT(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHTMLContent), v))
	})
}

// HTMLContentGTE applies the GTE predicate on the "html_content" field.
func HTMLContentGTE(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHTMLContent), v))
	})
}

// HTMLContentLT applies the LT predicate on the "html_content" field.
func HTMLContentLT(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHTMLContent), v))
	})
}

// HTMLContentLTE applies the LTE predicate on the "html_content" field.
func HTMLContentLTE(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHTMLContent), v))
	})
}

// HTMLContentContains applies the Contains predicate on the "html_content" field.
func HTMLContentContains(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHTMLContent), v))
	})
}

// HTMLContentHasPrefix applies the HasPrefix predicate on the "html_content" field.
func HTMLContentHasPrefix(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHTMLContent), v))
	})
}

// HTMLContentHasSuffix applies the HasSuffix predicate on the "html_content" field.
func HTMLContentHasSuffix(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHTMLContent), v))
	})
}

// HTMLContentIsNil applies the IsNil predicate on the "html_content" field.
func HTMLContentIsNil() predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldHTMLContent)))
	})
}

// HTMLContentNotNil applies the NotNil predicate on the "html_content" field.
func HTMLContentNotNil() predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldHTMLContent)))
	})
}

// HTMLContentEqualFold applies the EqualFold predicate on the "html_content" field.
func HTMLContentEqualFold(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHTMLContent), v))
	})
}

// HTMLContentContainsFold applies the ContainsFold predicate on the "html_content" field.
func HTMLContentContainsFold(v string) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHTMLContent), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.UnsavedPost {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UnsavedPost(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.UnsavedPost {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UnsavedPost(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// ModifiedAtEQ applies the EQ predicate on the "modified_at" field.
func ModifiedAtEQ(v time.Time) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModifiedAt), v))
	})
}

// ModifiedAtNEQ applies the NEQ predicate on the "modified_at" field.
func ModifiedAtNEQ(v time.Time) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldModifiedAt), v))
	})
}

// ModifiedAtIn applies the In predicate on the "modified_at" field.
func ModifiedAtIn(vs ...time.Time) predicate.UnsavedPost {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UnsavedPost(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldModifiedAt), v...))
	})
}

// ModifiedAtNotIn applies the NotIn predicate on the "modified_at" field.
func ModifiedAtNotIn(vs ...time.Time) predicate.UnsavedPost {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UnsavedPost(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldModifiedAt), v...))
	})
}

// ModifiedAtGT applies the GT predicate on the "modified_at" field.
func ModifiedAtGT(v time.Time) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldModifiedAt), v))
	})
}

// ModifiedAtGTE applies the GTE predicate on the "modified_at" field.
func ModifiedAtGTE(v time.Time) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldModifiedAt), v))
	})
}

// ModifiedAtLT applies the LT predicate on the "modified_at" field.
func ModifiedAtLT(v time.Time) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldModifiedAt), v))
	})
}

// ModifiedAtLTE applies the LTE predicate on the "modified_at" field.
func ModifiedAtLTE(v time.Time) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldModifiedAt), v))
	})
}

// HasAuthor applies the HasEdge predicate on the "author" edge.
func HasAuthor() predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthorTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorWith applies the HasEdge predicate on the "author" edge with a given conditions (other predicates).
func HasAuthorWith(preds ...predicate.Admin) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasThumbnail applies the HasEdge predicate on the "thumbnail" edge.
func HasThumbnail() predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ThumbnailTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ThumbnailTable, ThumbnailColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasThumbnailWith applies the HasEdge predicate on the "thumbnail" edge with a given conditions (other predicates).
func HasThumbnailWith(preds ...predicate.UnsavedPostThumbnail) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ThumbnailInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ThumbnailTable, ThumbnailColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasImages applies the HasEdge predicate on the "images" edge.
func HasImages() predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ImagesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ImagesTable, ImagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasImagesWith applies the HasEdge predicate on the "images" edge with a given conditions (other predicates).
func HasImagesWith(preds ...predicate.UnsavedPostImage) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ImagesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ImagesTable, ImagesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasVideos applies the HasEdge predicate on the "videos" edge.
func HasVideos() predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VideosTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, VideosTable, VideosColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVideosWith applies the HasEdge predicate on the "videos" edge with a given conditions (other predicates).
func HasVideosWith(preds ...predicate.UnsavedPostVideo) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(VideosInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, VideosTable, VideosColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAttachments applies the HasEdge predicate on the "attachments" edge.
func HasAttachments() predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AttachmentsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AttachmentsTable, AttachmentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAttachmentsWith applies the HasEdge predicate on the "attachments" edge with a given conditions (other predicates).
func HasAttachmentsWith(preds ...predicate.UnsavedPostAttachment) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AttachmentsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AttachmentsTable, AttachmentsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UnsavedPost) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UnsavedPost) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
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
func Not(p predicate.UnsavedPost) predicate.UnsavedPost {
	return predicate.UnsavedPost(func(s *sql.Selector) {
		p(s.Not())
	})
}