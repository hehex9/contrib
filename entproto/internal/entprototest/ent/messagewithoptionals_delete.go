// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/messagewithoptionals"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageWithOptionalsDelete is the builder for deleting a MessageWithOptionals entity.
type MessageWithOptionalsDelete struct {
	config
	hooks    []Hook
	mutation *MessageWithOptionalsMutation
}

// Where appends a list predicates to the MessageWithOptionalsDelete builder.
func (mwod *MessageWithOptionalsDelete) Where(ps ...predicate.MessageWithOptionals) *MessageWithOptionalsDelete {
	mwod.mutation.Where(ps...)
	return mwod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mwod *MessageWithOptionalsDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mwod.hooks) == 0 {
		affected, err = mwod.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageWithOptionalsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mwod.mutation = mutation
			affected, err = mwod.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mwod.hooks) - 1; i >= 0; i-- {
			if mwod.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mwod.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mwod.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwod *MessageWithOptionalsDelete) ExecX(ctx context.Context) int {
	n, err := mwod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mwod *MessageWithOptionalsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: messagewithoptionals.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: messagewithoptionals.FieldID,
			},
		},
	}
	if ps := mwod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, mwod.driver, _spec)
}

// MessageWithOptionalsDeleteOne is the builder for deleting a single MessageWithOptionals entity.
type MessageWithOptionalsDeleteOne struct {
	mwod *MessageWithOptionalsDelete
}

// Exec executes the deletion query.
func (mwodo *MessageWithOptionalsDeleteOne) Exec(ctx context.Context) error {
	n, err := mwodo.mwod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{messagewithoptionals.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mwodo *MessageWithOptionalsDeleteOne) ExecX(ctx context.Context) {
	mwodo.mwod.ExecX(ctx)
}
