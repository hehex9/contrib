// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/contrib/entgql/internal/todo/ent/schema/schematype"
	"entgo.io/contrib/entgql/internal/todoplugin/ent/category"
	"entgo.io/contrib/entgql/internal/todoplugin/ent/predicate"
	"entgo.io/contrib/entgql/internal/todoplugin/ent/todo"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CategoryUpdate is the builder for updating Category entities.
type CategoryUpdate struct {
	config
	hooks    []Hook
	mutation *CategoryMutation
}

// Where appends a list predicates to the CategoryUpdate builder.
func (cu *CategoryUpdate) Where(ps ...predicate.Category) *CategoryUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetText sets the "text" field.
func (cu *CategoryUpdate) SetText(s string) *CategoryUpdate {
	cu.mutation.SetText(s)
	return cu
}

// SetUUIDA sets the "uuid_a" field.
func (cu *CategoryUpdate) SetUUIDA(u uuid.UUID) *CategoryUpdate {
	cu.mutation.SetUUIDA(u)
	return cu
}

// SetNillableUUIDA sets the "uuid_a" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableUUIDA(u *uuid.UUID) *CategoryUpdate {
	if u != nil {
		cu.SetUUIDA(*u)
	}
	return cu
}

// ClearUUIDA clears the value of the "uuid_a" field.
func (cu *CategoryUpdate) ClearUUIDA() *CategoryUpdate {
	cu.mutation.ClearUUIDA()
	return cu
}

// SetStatus sets the "status" field.
func (cu *CategoryUpdate) SetStatus(c category.Status) *CategoryUpdate {
	cu.mutation.SetStatus(c)
	return cu
}

// SetConfig sets the "config" field.
func (cu *CategoryUpdate) SetConfig(sc *schematype.CategoryConfig) *CategoryUpdate {
	cu.mutation.SetConfig(sc)
	return cu
}

// ClearConfig clears the value of the "config" field.
func (cu *CategoryUpdate) ClearConfig() *CategoryUpdate {
	cu.mutation.ClearConfig()
	return cu
}

// SetDuration sets the "duration" field.
func (cu *CategoryUpdate) SetDuration(t time.Duration) *CategoryUpdate {
	cu.mutation.ResetDuration()
	cu.mutation.SetDuration(t)
	return cu
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableDuration(t *time.Duration) *CategoryUpdate {
	if t != nil {
		cu.SetDuration(*t)
	}
	return cu
}

// AddDuration adds t to the "duration" field.
func (cu *CategoryUpdate) AddDuration(t time.Duration) *CategoryUpdate {
	cu.mutation.AddDuration(t)
	return cu
}

// ClearDuration clears the value of the "duration" field.
func (cu *CategoryUpdate) ClearDuration() *CategoryUpdate {
	cu.mutation.ClearDuration()
	return cu
}

// SetCount sets the "count" field.
func (cu *CategoryUpdate) SetCount(u uint64) *CategoryUpdate {
	cu.mutation.ResetCount()
	cu.mutation.SetCount(u)
	return cu
}

// SetNillableCount sets the "count" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableCount(u *uint64) *CategoryUpdate {
	if u != nil {
		cu.SetCount(*u)
	}
	return cu
}

// AddCount adds u to the "count" field.
func (cu *CategoryUpdate) AddCount(u int64) *CategoryUpdate {
	cu.mutation.AddCount(u)
	return cu
}

// ClearCount clears the value of the "count" field.
func (cu *CategoryUpdate) ClearCount() *CategoryUpdate {
	cu.mutation.ClearCount()
	return cu
}

// SetStrings sets the "strings" field.
func (cu *CategoryUpdate) SetStrings(s []string) *CategoryUpdate {
	cu.mutation.SetStrings(s)
	return cu
}

// ClearStrings clears the value of the "strings" field.
func (cu *CategoryUpdate) ClearStrings() *CategoryUpdate {
	cu.mutation.ClearStrings()
	return cu
}

// AddTodoIDs adds the "todos" edge to the Todo entity by IDs.
func (cu *CategoryUpdate) AddTodoIDs(ids ...int) *CategoryUpdate {
	cu.mutation.AddTodoIDs(ids...)
	return cu
}

// AddTodos adds the "todos" edges to the Todo entity.
func (cu *CategoryUpdate) AddTodos(t ...*Todo) *CategoryUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.AddTodoIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cu *CategoryUpdate) Mutation() *CategoryMutation {
	return cu.mutation
}

// ClearTodos clears all "todos" edges to the Todo entity.
func (cu *CategoryUpdate) ClearTodos() *CategoryUpdate {
	cu.mutation.ClearTodos()
	return cu
}

// RemoveTodoIDs removes the "todos" edge to Todo entities by IDs.
func (cu *CategoryUpdate) RemoveTodoIDs(ids ...int) *CategoryUpdate {
	cu.mutation.RemoveTodoIDs(ids...)
	return cu
}

// RemoveTodos removes "todos" edges to Todo entities.
func (cu *CategoryUpdate) RemoveTodos(t ...*Todo) *CategoryUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.RemoveTodoIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CategoryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CategoryUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CategoryUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CategoryUpdate) check() error {
	if v, ok := cu.mutation.Text(); ok {
		if err := category.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Category.text": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Status(); ok {
		if err := category.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Category.status": %w`, err)}
		}
	}
	return nil
}

func (cu *CategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   category.Table,
			Columns: category.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: category.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Text(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldText,
		})
	}
	if value, ok := cu.mutation.UUIDA(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: category.FieldUUIDA,
		})
	}
	if cu.mutation.UUIDACleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: category.FieldUUIDA,
		})
	}
	if value, ok := cu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: category.FieldStatus,
		})
	}
	if value, ok := cu.mutation.Config(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: category.FieldConfig,
		})
	}
	if cu.mutation.ConfigCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: category.FieldConfig,
		})
	}
	if value, ok := cu.mutation.Duration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: category.FieldDuration,
		})
	}
	if value, ok := cu.mutation.AddedDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: category.FieldDuration,
		})
	}
	if cu.mutation.DurationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: category.FieldDuration,
		})
	}
	if value, ok := cu.mutation.Count(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: category.FieldCount,
		})
	}
	if value, ok := cu.mutation.AddedCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: category.FieldCount,
		})
	}
	if cu.mutation.CountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Column: category.FieldCount,
		})
	}
	if value, ok := cu.mutation.Strings(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: category.FieldStrings,
		})
	}
	if cu.mutation.StringsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: category.FieldStrings,
		})
	}
	if cu.mutation.TodosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.TodosTable,
			Columns: []string{category.TodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: todo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedTodosIDs(); len(nodes) > 0 && !cu.mutation.TodosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.TodosTable,
			Columns: []string{category.TodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: todo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.TodosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.TodosTable,
			Columns: []string{category.TodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: todo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CategoryUpdateOne is the builder for updating a single Category entity.
type CategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CategoryMutation
}

// SetText sets the "text" field.
func (cuo *CategoryUpdateOne) SetText(s string) *CategoryUpdateOne {
	cuo.mutation.SetText(s)
	return cuo
}

// SetUUIDA sets the "uuid_a" field.
func (cuo *CategoryUpdateOne) SetUUIDA(u uuid.UUID) *CategoryUpdateOne {
	cuo.mutation.SetUUIDA(u)
	return cuo
}

// SetNillableUUIDA sets the "uuid_a" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableUUIDA(u *uuid.UUID) *CategoryUpdateOne {
	if u != nil {
		cuo.SetUUIDA(*u)
	}
	return cuo
}

// ClearUUIDA clears the value of the "uuid_a" field.
func (cuo *CategoryUpdateOne) ClearUUIDA() *CategoryUpdateOne {
	cuo.mutation.ClearUUIDA()
	return cuo
}

// SetStatus sets the "status" field.
func (cuo *CategoryUpdateOne) SetStatus(c category.Status) *CategoryUpdateOne {
	cuo.mutation.SetStatus(c)
	return cuo
}

// SetConfig sets the "config" field.
func (cuo *CategoryUpdateOne) SetConfig(sc *schematype.CategoryConfig) *CategoryUpdateOne {
	cuo.mutation.SetConfig(sc)
	return cuo
}

// ClearConfig clears the value of the "config" field.
func (cuo *CategoryUpdateOne) ClearConfig() *CategoryUpdateOne {
	cuo.mutation.ClearConfig()
	return cuo
}

// SetDuration sets the "duration" field.
func (cuo *CategoryUpdateOne) SetDuration(t time.Duration) *CategoryUpdateOne {
	cuo.mutation.ResetDuration()
	cuo.mutation.SetDuration(t)
	return cuo
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableDuration(t *time.Duration) *CategoryUpdateOne {
	if t != nil {
		cuo.SetDuration(*t)
	}
	return cuo
}

// AddDuration adds t to the "duration" field.
func (cuo *CategoryUpdateOne) AddDuration(t time.Duration) *CategoryUpdateOne {
	cuo.mutation.AddDuration(t)
	return cuo
}

// ClearDuration clears the value of the "duration" field.
func (cuo *CategoryUpdateOne) ClearDuration() *CategoryUpdateOne {
	cuo.mutation.ClearDuration()
	return cuo
}

// SetCount sets the "count" field.
func (cuo *CategoryUpdateOne) SetCount(u uint64) *CategoryUpdateOne {
	cuo.mutation.ResetCount()
	cuo.mutation.SetCount(u)
	return cuo
}

// SetNillableCount sets the "count" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableCount(u *uint64) *CategoryUpdateOne {
	if u != nil {
		cuo.SetCount(*u)
	}
	return cuo
}

// AddCount adds u to the "count" field.
func (cuo *CategoryUpdateOne) AddCount(u int64) *CategoryUpdateOne {
	cuo.mutation.AddCount(u)
	return cuo
}

// ClearCount clears the value of the "count" field.
func (cuo *CategoryUpdateOne) ClearCount() *CategoryUpdateOne {
	cuo.mutation.ClearCount()
	return cuo
}

// SetStrings sets the "strings" field.
func (cuo *CategoryUpdateOne) SetStrings(s []string) *CategoryUpdateOne {
	cuo.mutation.SetStrings(s)
	return cuo
}

// ClearStrings clears the value of the "strings" field.
func (cuo *CategoryUpdateOne) ClearStrings() *CategoryUpdateOne {
	cuo.mutation.ClearStrings()
	return cuo
}

// AddTodoIDs adds the "todos" edge to the Todo entity by IDs.
func (cuo *CategoryUpdateOne) AddTodoIDs(ids ...int) *CategoryUpdateOne {
	cuo.mutation.AddTodoIDs(ids...)
	return cuo
}

// AddTodos adds the "todos" edges to the Todo entity.
func (cuo *CategoryUpdateOne) AddTodos(t ...*Todo) *CategoryUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.AddTodoIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cuo *CategoryUpdateOne) Mutation() *CategoryMutation {
	return cuo.mutation
}

// ClearTodos clears all "todos" edges to the Todo entity.
func (cuo *CategoryUpdateOne) ClearTodos() *CategoryUpdateOne {
	cuo.mutation.ClearTodos()
	return cuo
}

// RemoveTodoIDs removes the "todos" edge to Todo entities by IDs.
func (cuo *CategoryUpdateOne) RemoveTodoIDs(ids ...int) *CategoryUpdateOne {
	cuo.mutation.RemoveTodoIDs(ids...)
	return cuo
}

// RemoveTodos removes "todos" edges to Todo entities.
func (cuo *CategoryUpdateOne) RemoveTodos(t ...*Todo) *CategoryUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.RemoveTodoIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CategoryUpdateOne) Select(field string, fields ...string) *CategoryUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Category entity.
func (cuo *CategoryUpdateOne) Save(ctx context.Context) (*Category, error) {
	var (
		err  error
		node *Category
	)
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CategoryUpdateOne) SaveX(ctx context.Context) *Category {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CategoryUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CategoryUpdateOne) check() error {
	if v, ok := cuo.mutation.Text(); ok {
		if err := category.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Category.text": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Status(); ok {
		if err := category.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Category.status": %w`, err)}
		}
	}
	return nil
}

func (cuo *CategoryUpdateOne) sqlSave(ctx context.Context) (_node *Category, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   category.Table,
			Columns: category.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: category.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Category.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, category.FieldID)
		for _, f := range fields {
			if !category.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != category.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Text(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldText,
		})
	}
	if value, ok := cuo.mutation.UUIDA(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: category.FieldUUIDA,
		})
	}
	if cuo.mutation.UUIDACleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: category.FieldUUIDA,
		})
	}
	if value, ok := cuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: category.FieldStatus,
		})
	}
	if value, ok := cuo.mutation.Config(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: category.FieldConfig,
		})
	}
	if cuo.mutation.ConfigCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: category.FieldConfig,
		})
	}
	if value, ok := cuo.mutation.Duration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: category.FieldDuration,
		})
	}
	if value, ok := cuo.mutation.AddedDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: category.FieldDuration,
		})
	}
	if cuo.mutation.DurationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: category.FieldDuration,
		})
	}
	if value, ok := cuo.mutation.Count(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: category.FieldCount,
		})
	}
	if value, ok := cuo.mutation.AddedCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: category.FieldCount,
		})
	}
	if cuo.mutation.CountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Column: category.FieldCount,
		})
	}
	if value, ok := cuo.mutation.Strings(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: category.FieldStrings,
		})
	}
	if cuo.mutation.StringsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: category.FieldStrings,
		})
	}
	if cuo.mutation.TodosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.TodosTable,
			Columns: []string{category.TodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: todo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedTodosIDs(); len(nodes) > 0 && !cuo.mutation.TodosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.TodosTable,
			Columns: []string{category.TodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: todo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.TodosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.TodosTable,
			Columns: []string{category.TodosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: todo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Category{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
