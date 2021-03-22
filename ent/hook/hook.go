// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"devlog/ent"
	"fmt"
)

// The AdminFunc type is an adapter to allow the use of ordinary
// function as Admin mutator.
type AdminFunc func(context.Context, *ent.AdminMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AdminFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AdminMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AdminMutation", m)
	}
	return f(ctx, mv)
}

// The AdminSessionFunc type is an adapter to allow the use of ordinary
// function as AdminSession mutator.
type AdminSessionFunc func(context.Context, *ent.AdminSessionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AdminSessionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AdminSessionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AdminSessionMutation", m)
	}
	return f(ctx, mv)
}

// The CategoryFunc type is an adapter to allow the use of ordinary
// function as Category mutator.
type CategoryFunc func(context.Context, *ent.CategoryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CategoryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CategoryMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CategoryMutation", m)
	}
	return f(ctx, mv)
}

// The PostFunc type is an adapter to allow the use of ordinary
// function as Post mutator.
type PostFunc func(context.Context, *ent.PostMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PostFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PostMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PostMutation", m)
	}
	return f(ctx, mv)
}

// The PostAttachmentFunc type is an adapter to allow the use of ordinary
// function as PostAttachment mutator.
type PostAttachmentFunc func(context.Context, *ent.PostAttachmentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PostAttachmentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PostAttachmentMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PostAttachmentMutation", m)
	}
	return f(ctx, mv)
}

// The PostImageFunc type is an adapter to allow the use of ordinary
// function as PostImage mutator.
type PostImageFunc func(context.Context, *ent.PostImageMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PostImageFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PostImageMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PostImageMutation", m)
	}
	return f(ctx, mv)
}

// The PostThumbnailFunc type is an adapter to allow the use of ordinary
// function as PostThumbnail mutator.
type PostThumbnailFunc func(context.Context, *ent.PostThumbnailMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PostThumbnailFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PostThumbnailMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PostThumbnailMutation", m)
	}
	return f(ctx, mv)
}

// The PostVideoFunc type is an adapter to allow the use of ordinary
// function as PostVideo mutator.
type PostVideoFunc func(context.Context, *ent.PostVideoMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PostVideoFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PostVideoMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PostVideoMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
