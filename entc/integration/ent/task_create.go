// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/schema/task"
	enttask "entgo.io/ent/entc/integration/ent/task"
	"entgo.io/ent/schema/field"
)

// TaskCreate is the builder for creating a Task entity.
type TaskCreate struct {
	config
	mutation *TaskMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetPriority sets the "priority" field.
func (tc *TaskCreate) SetPriority(t task.Priority) *TaskCreate {
	tc.mutation.SetPriority(t)
	return tc
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (tc *TaskCreate) SetNillablePriority(t *task.Priority) *TaskCreate {
	if t != nil {
		tc.SetPriority(*t)
	}
	return tc
}

// SetPriorities sets the "priorities" field.
func (tc *TaskCreate) SetPriorities(m map[string]task.Priority) *TaskCreate {
	tc.mutation.SetPriorities(m)
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TaskCreate) SetCreatedAt(t time.Time) *TaskCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TaskCreate) SetNillableCreatedAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetName sets the "name" field.
func (tc *TaskCreate) SetName(s string) *TaskCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tc *TaskCreate) SetNillableName(s *string) *TaskCreate {
	if s != nil {
		tc.SetName(*s)
	}
	return tc
}

// SetOwner sets the "owner" field.
func (tc *TaskCreate) SetOwner(s string) *TaskCreate {
	tc.mutation.SetOwner(s)
	return tc
}

// SetNillableOwner sets the "owner" field if the given value is not nil.
func (tc *TaskCreate) SetNillableOwner(s *string) *TaskCreate {
	if s != nil {
		tc.SetOwner(*s)
	}
	return tc
}

// SetOrder sets the "order" field.
func (tc *TaskCreate) SetOrder(i int) *TaskCreate {
	tc.mutation.SetOrder(i)
	return tc
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (tc *TaskCreate) SetNillableOrder(i *int) *TaskCreate {
	if i != nil {
		tc.SetOrder(*i)
	}
	return tc
}

// SetOrderOption sets the "order_option" field.
func (tc *TaskCreate) SetOrderOption(i int) *TaskCreate {
	tc.mutation.SetOrderOption(i)
	return tc
}

// SetNillableOrderOption sets the "order_option" field if the given value is not nil.
func (tc *TaskCreate) SetNillableOrderOption(i *int) *TaskCreate {
	if i != nil {
		tc.SetOrderOption(*i)
	}
	return tc
}

// Mutation returns the TaskMutation object of the builder.
func (tc *TaskCreate) Mutation() *TaskMutation {
	return tc.mutation
}

// Save creates the Task in the database.
func (tc *TaskCreate) Save(ctx context.Context) (*Task, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TaskCreate) SaveX(ctx context.Context) *Task {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TaskCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TaskCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TaskCreate) defaults() {
	if _, ok := tc.mutation.Priority(); !ok {
		v := enttask.DefaultPriority
		tc.mutation.SetPriority(v)
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := enttask.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TaskCreate) check() error {
	if _, ok := tc.mutation.Priority(); !ok {
		return &ValidationError{Name: "priority", err: errors.New(`ent: missing required field "Task.priority"`)}
	}
	if v, ok := tc.mutation.Priority(); ok {
		if err := v.Validate(); err != nil {
			return &ValidationError{Name: "priority", err: fmt.Errorf(`ent: validator failed for field "Task.priority": %w`, err)}
		}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Task.created_at"`)}
	}
	return nil
}

func (tc *TaskCreate) sqlSave(ctx context.Context) (*Task, error) {
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

func (tc *TaskCreate) createSpec() (*Task, *sqlgraph.CreateSpec) {
	var (
		_node = &Task{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(enttask.Table, sqlgraph.NewFieldSpec(enttask.FieldID, field.TypeInt))
	)
	_spec.OnConflict = tc.conflict
	if value, ok := tc.mutation.Priority(); ok {
		_spec.SetField(enttask.FieldPriority, field.TypeInt, value)
		_node.Priority = value
	}
	if value, ok := tc.mutation.Priorities(); ok {
		_spec.SetField(enttask.FieldPriorities, field.TypeJSON, value)
		_node.Priorities = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(enttask.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = &value
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(enttask.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tc.mutation.Owner(); ok {
		_spec.SetField(enttask.FieldOwner, field.TypeString, value)
		_node.Owner = value
	}
	if value, ok := tc.mutation.Order(); ok {
		_spec.SetField(enttask.FieldOrder, field.TypeInt, value)
		_node.Order = value
	}
	if value, ok := tc.mutation.OrderOption(); ok {
		_spec.SetField(enttask.FieldOrderOption, field.TypeInt, value)
		_node.OrderOption = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Task.Create().
//		SetPriority(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TaskUpsert) {
//			SetPriority(v+v).
//		}).
//		Exec(ctx)
func (tc *TaskCreate) OnConflict(opts ...sql.ConflictOption) *TaskUpsertOne {
	tc.conflict = opts
	return &TaskUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Task.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tc *TaskCreate) OnConflictColumns(columns ...string) *TaskUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TaskUpsertOne{
		create: tc,
	}
}

type (
	// TaskUpsertOne is the builder for "upsert"-ing
	//  one Task node.
	TaskUpsertOne struct {
		create *TaskCreate
	}

	// TaskUpsert is the "OnConflict" setter.
	TaskUpsert struct {
		*sql.UpdateSet
	}
)

// SetPriority sets the "priority" field.
func (u *TaskUpsert) SetPriority(v task.Priority) *TaskUpsert {
	u.Set(enttask.FieldPriority, v)
	return u
}

// UpdatePriority sets the "priority" field to the value that was provided on create.
func (u *TaskUpsert) UpdatePriority() *TaskUpsert {
	u.SetExcluded(enttask.FieldPriority)
	return u
}

// AddPriority adds v to the "priority" field.
func (u *TaskUpsert) AddPriority(v task.Priority) *TaskUpsert {
	u.Add(enttask.FieldPriority, v)
	return u
}

// SetPriorities sets the "priorities" field.
func (u *TaskUpsert) SetPriorities(v map[string]task.Priority) *TaskUpsert {
	u.Set(enttask.FieldPriorities, v)
	return u
}

// UpdatePriorities sets the "priorities" field to the value that was provided on create.
func (u *TaskUpsert) UpdatePriorities() *TaskUpsert {
	u.SetExcluded(enttask.FieldPriorities)
	return u
}

// ClearPriorities clears the value of the "priorities" field.
func (u *TaskUpsert) ClearPriorities() *TaskUpsert {
	u.SetNull(enttask.FieldPriorities)
	return u
}

// SetName sets the "name" field.
func (u *TaskUpsert) SetName(v string) *TaskUpsert {
	u.Set(enttask.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TaskUpsert) UpdateName() *TaskUpsert {
	u.SetExcluded(enttask.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *TaskUpsert) ClearName() *TaskUpsert {
	u.SetNull(enttask.FieldName)
	return u
}

// SetOwner sets the "owner" field.
func (u *TaskUpsert) SetOwner(v string) *TaskUpsert {
	u.Set(enttask.FieldOwner, v)
	return u
}

// UpdateOwner sets the "owner" field to the value that was provided on create.
func (u *TaskUpsert) UpdateOwner() *TaskUpsert {
	u.SetExcluded(enttask.FieldOwner)
	return u
}

// ClearOwner clears the value of the "owner" field.
func (u *TaskUpsert) ClearOwner() *TaskUpsert {
	u.SetNull(enttask.FieldOwner)
	return u
}

// SetOrder sets the "order" field.
func (u *TaskUpsert) SetOrder(v int) *TaskUpsert {
	u.Set(enttask.FieldOrder, v)
	return u
}

// UpdateOrder sets the "order" field to the value that was provided on create.
func (u *TaskUpsert) UpdateOrder() *TaskUpsert {
	u.SetExcluded(enttask.FieldOrder)
	return u
}

// AddOrder adds v to the "order" field.
func (u *TaskUpsert) AddOrder(v int) *TaskUpsert {
	u.Add(enttask.FieldOrder, v)
	return u
}

// ClearOrder clears the value of the "order" field.
func (u *TaskUpsert) ClearOrder() *TaskUpsert {
	u.SetNull(enttask.FieldOrder)
	return u
}

// SetOrderOption sets the "order_option" field.
func (u *TaskUpsert) SetOrderOption(v int) *TaskUpsert {
	u.Set(enttask.FieldOrderOption, v)
	return u
}

// UpdateOrderOption sets the "order_option" field to the value that was provided on create.
func (u *TaskUpsert) UpdateOrderOption() *TaskUpsert {
	u.SetExcluded(enttask.FieldOrderOption)
	return u
}

// AddOrderOption adds v to the "order_option" field.
func (u *TaskUpsert) AddOrderOption(v int) *TaskUpsert {
	u.Add(enttask.FieldOrderOption, v)
	return u
}

// ClearOrderOption clears the value of the "order_option" field.
func (u *TaskUpsert) ClearOrderOption() *TaskUpsert {
	u.SetNull(enttask.FieldOrderOption)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Task.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *TaskUpsertOne) UpdateNewValues() *TaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(enttask.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Task.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TaskUpsertOne) Ignore() *TaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TaskUpsertOne) DoNothing() *TaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TaskCreate.OnConflict
// documentation for more info.
func (u *TaskUpsertOne) Update(set func(*TaskUpsert)) *TaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TaskUpsert{UpdateSet: update})
	}))
	return u
}

// SetPriority sets the "priority" field.
func (u *TaskUpsertOne) SetPriority(v task.Priority) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetPriority(v)
	})
}

// AddPriority adds v to the "priority" field.
func (u *TaskUpsertOne) AddPriority(v task.Priority) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.AddPriority(v)
	})
}

// UpdatePriority sets the "priority" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdatePriority() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdatePriority()
	})
}

// SetPriorities sets the "priorities" field.
func (u *TaskUpsertOne) SetPriorities(v map[string]task.Priority) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetPriorities(v)
	})
}

// UpdatePriorities sets the "priorities" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdatePriorities() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdatePriorities()
	})
}

// ClearPriorities clears the value of the "priorities" field.
func (u *TaskUpsertOne) ClearPriorities() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.ClearPriorities()
	})
}

// SetName sets the "name" field.
func (u *TaskUpsertOne) SetName(v string) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateName() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *TaskUpsertOne) ClearName() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.ClearName()
	})
}

// SetOwner sets the "owner" field.
func (u *TaskUpsertOne) SetOwner(v string) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetOwner(v)
	})
}

// UpdateOwner sets the "owner" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateOwner() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateOwner()
	})
}

// ClearOwner clears the value of the "owner" field.
func (u *TaskUpsertOne) ClearOwner() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.ClearOwner()
	})
}

// SetOrder sets the "order" field.
func (u *TaskUpsertOne) SetOrder(v int) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetOrder(v)
	})
}

// AddOrder adds v to the "order" field.
func (u *TaskUpsertOne) AddOrder(v int) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.AddOrder(v)
	})
}

// UpdateOrder sets the "order" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateOrder() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateOrder()
	})
}

// ClearOrder clears the value of the "order" field.
func (u *TaskUpsertOne) ClearOrder() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.ClearOrder()
	})
}

// SetOrderOption sets the "order_option" field.
func (u *TaskUpsertOne) SetOrderOption(v int) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetOrderOption(v)
	})
}

// AddOrderOption adds v to the "order_option" field.
func (u *TaskUpsertOne) AddOrderOption(v int) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.AddOrderOption(v)
	})
}

// UpdateOrderOption sets the "order_option" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateOrderOption() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateOrderOption()
	})
}

// ClearOrderOption clears the value of the "order_option" field.
func (u *TaskUpsertOne) ClearOrderOption() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.ClearOrderOption()
	})
}

// Exec executes the query.
func (u *TaskUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TaskCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TaskUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TaskUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TaskUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TaskCreateBulk is the builder for creating many Task entities in bulk.
type TaskCreateBulk struct {
	config
	builders []*TaskCreate
	conflict []sql.ConflictOption
}

// Save creates the Task entities in the database.
func (tcb *TaskCreateBulk) Save(ctx context.Context) ([]*Task, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Task, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TaskMutation)
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
					spec.OnConflict = tcb.conflict
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
func (tcb *TaskCreateBulk) SaveX(ctx context.Context) []*Task {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TaskCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TaskCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Task.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TaskUpsert) {
//			SetPriority(v+v).
//		}).
//		Exec(ctx)
func (tcb *TaskCreateBulk) OnConflict(opts ...sql.ConflictOption) *TaskUpsertBulk {
	tcb.conflict = opts
	return &TaskUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Task.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tcb *TaskCreateBulk) OnConflictColumns(columns ...string) *TaskUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TaskUpsertBulk{
		create: tcb,
	}
}

// TaskUpsertBulk is the builder for "upsert"-ing
// a bulk of Task nodes.
type TaskUpsertBulk struct {
	create *TaskCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Task.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *TaskUpsertBulk) UpdateNewValues() *TaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(enttask.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Task.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TaskUpsertBulk) Ignore() *TaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TaskUpsertBulk) DoNothing() *TaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TaskCreateBulk.OnConflict
// documentation for more info.
func (u *TaskUpsertBulk) Update(set func(*TaskUpsert)) *TaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TaskUpsert{UpdateSet: update})
	}))
	return u
}

// SetPriority sets the "priority" field.
func (u *TaskUpsertBulk) SetPriority(v task.Priority) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetPriority(v)
	})
}

// AddPriority adds v to the "priority" field.
func (u *TaskUpsertBulk) AddPriority(v task.Priority) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.AddPriority(v)
	})
}

// UpdatePriority sets the "priority" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdatePriority() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdatePriority()
	})
}

// SetPriorities sets the "priorities" field.
func (u *TaskUpsertBulk) SetPriorities(v map[string]task.Priority) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetPriorities(v)
	})
}

// UpdatePriorities sets the "priorities" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdatePriorities() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdatePriorities()
	})
}

// ClearPriorities clears the value of the "priorities" field.
func (u *TaskUpsertBulk) ClearPriorities() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.ClearPriorities()
	})
}

// SetName sets the "name" field.
func (u *TaskUpsertBulk) SetName(v string) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateName() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *TaskUpsertBulk) ClearName() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.ClearName()
	})
}

// SetOwner sets the "owner" field.
func (u *TaskUpsertBulk) SetOwner(v string) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetOwner(v)
	})
}

// UpdateOwner sets the "owner" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateOwner() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateOwner()
	})
}

// ClearOwner clears the value of the "owner" field.
func (u *TaskUpsertBulk) ClearOwner() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.ClearOwner()
	})
}

// SetOrder sets the "order" field.
func (u *TaskUpsertBulk) SetOrder(v int) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetOrder(v)
	})
}

// AddOrder adds v to the "order" field.
func (u *TaskUpsertBulk) AddOrder(v int) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.AddOrder(v)
	})
}

// UpdateOrder sets the "order" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateOrder() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateOrder()
	})
}

// ClearOrder clears the value of the "order" field.
func (u *TaskUpsertBulk) ClearOrder() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.ClearOrder()
	})
}

// SetOrderOption sets the "order_option" field.
func (u *TaskUpsertBulk) SetOrderOption(v int) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetOrderOption(v)
	})
}

// AddOrderOption adds v to the "order_option" field.
func (u *TaskUpsertBulk) AddOrderOption(v int) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.AddOrderOption(v)
	})
}

// UpdateOrderOption sets the "order_option" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateOrderOption() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateOrderOption()
	})
}

// ClearOrderOption clears the value of the "order_option" field.
func (u *TaskUpsertBulk) ClearOrderOption() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.ClearOrderOption()
	})
}

// Exec executes the query.
func (u *TaskUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TaskCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TaskCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TaskUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
