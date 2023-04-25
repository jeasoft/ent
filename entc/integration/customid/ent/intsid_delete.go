// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/intsid"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/schema/field"
)

// IntSIDDelete is the builder for deleting a IntSID entity.
type IntSIDDelete struct {
	config
	hooks    []Hook
	mutation *IntSIDMutation
}

// Where appends a list predicates to the IntSIDDelete builder.
func (isd *IntSIDDelete) Where(ps ...predicate.IntSID) *IntSIDDelete {
	isd.mutation.Where(ps...)
	return isd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (isd *IntSIDDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, isd.sqlExec, isd.mutation, isd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (isd *IntSIDDelete) ExecX(ctx context.Context) int {
	n, err := isd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (isd *IntSIDDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(intsid.Table, sqlgraph.NewFieldSpec(intsid.FieldID, field.TypeInt64))
	if ps := isd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, isd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	isd.mutation.done = true
	return affected, err
}

// IntSIDDeleteOne is the builder for deleting a single IntSID entity.
type IntSIDDeleteOne struct {
	isd *IntSIDDelete
}

// Where appends a list predicates to the IntSIDDelete builder.
func (isdo *IntSIDDeleteOne) Where(ps ...predicate.IntSID) *IntSIDDeleteOne {
	isdo.isd.mutation.Where(ps...)
	return isdo
}

// Exec executes the deletion query.
func (isdo *IntSIDDeleteOne) Exec(ctx context.Context) error {
	n, err := isdo.isd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{intsid.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (isdo *IntSIDDeleteOne) ExecX(ctx context.Context) {
	if err := isdo.Exec(ctx); err != nil {
		panic(err)
	}
}
