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

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/product"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/vendor"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/warehouse"
	"github.com/google/uuid"
)

// WarehouseCreate is the builder for creating a Warehouse entity.
type WarehouseCreate struct {
	config
	mutation *WarehouseMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (wc *WarehouseCreate) SetName(s string) *WarehouseCreate {
	wc.mutation.SetName(s)
	return wc
}

// SetLastUpdate sets the "last_update" field.
func (wc *WarehouseCreate) SetLastUpdate(t time.Time) *WarehouseCreate {
	wc.mutation.SetLastUpdate(t)
	return wc
}

// SetNillableLastUpdate sets the "last_update" field if the given value is not nil.
func (wc *WarehouseCreate) SetNillableLastUpdate(t *time.Time) *WarehouseCreate {
	if t != nil {
		wc.SetLastUpdate(*t)
	}
	return wc
}

// SetOriginalData sets the "original_data" field.
func (wc *WarehouseCreate) SetOriginalData(s string) *WarehouseCreate {
	wc.mutation.SetOriginalData(s)
	return wc
}

// SetNillableOriginalData sets the "original_data" field if the given value is not nil.
func (wc *WarehouseCreate) SetNillableOriginalData(s *string) *WarehouseCreate {
	if s != nil {
		wc.SetOriginalData(*s)
	}
	return wc
}

// SetEnabled sets the "enabled" field.
func (wc *WarehouseCreate) SetEnabled(b bool) *WarehouseCreate {
	wc.mutation.SetEnabled(b)
	return wc
}

// SetNillableEnabled sets the "enabled" field if the given value is not nil.
func (wc *WarehouseCreate) SetNillableEnabled(b *bool) *WarehouseCreate {
	if b != nil {
		wc.SetEnabled(*b)
	}
	return wc
}

// SetFilters sets the "filters" field.
func (wc *WarehouseCreate) SetFilters(s []string) *WarehouseCreate {
	wc.mutation.SetFilters(s)
	return wc
}

// SetID sets the "id" field.
func (wc *WarehouseCreate) SetID(u uuid.UUID) *WarehouseCreate {
	wc.mutation.SetID(u)
	return wc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wc *WarehouseCreate) SetNillableID(u *uuid.UUID) *WarehouseCreate {
	if u != nil {
		wc.SetID(*u)
	}
	return wc
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (wc *WarehouseCreate) AddProductIDs(ids ...uuid.UUID) *WarehouseCreate {
	wc.mutation.AddProductIDs(ids...)
	return wc
}

// AddProducts adds the "products" edges to the Product entity.
func (wc *WarehouseCreate) AddProducts(p ...*Product) *WarehouseCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return wc.AddProductIDs(ids...)
}

// SetVendorID sets the "vendor" edge to the Vendor entity by ID.
func (wc *WarehouseCreate) SetVendorID(id uuid.UUID) *WarehouseCreate {
	wc.mutation.SetVendorID(id)
	return wc
}

// SetNillableVendorID sets the "vendor" edge to the Vendor entity by ID if the given value is not nil.
func (wc *WarehouseCreate) SetNillableVendorID(id *uuid.UUID) *WarehouseCreate {
	if id != nil {
		wc = wc.SetVendorID(*id)
	}
	return wc
}

// SetVendor sets the "vendor" edge to the Vendor entity.
func (wc *WarehouseCreate) SetVendor(v *Vendor) *WarehouseCreate {
	return wc.SetVendorID(v.ID)
}

// Mutation returns the WarehouseMutation object of the builder.
func (wc *WarehouseCreate) Mutation() *WarehouseMutation {
	return wc.mutation
}

// Save creates the Warehouse in the database.
func (wc *WarehouseCreate) Save(ctx context.Context) (*Warehouse, error) {
	wc.defaults()
	return withHooks[*Warehouse, WarehouseMutation](ctx, wc.sqlSave, wc.mutation, wc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WarehouseCreate) SaveX(ctx context.Context) *Warehouse {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WarehouseCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WarehouseCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wc *WarehouseCreate) defaults() {
	if _, ok := wc.mutation.Enabled(); !ok {
		v := warehouse.DefaultEnabled
		wc.mutation.SetEnabled(v)
	}
	if _, ok := wc.mutation.ID(); !ok {
		v := warehouse.DefaultID()
		wc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WarehouseCreate) check() error {
	if _, ok := wc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Warehouse.name"`)}
	}
	if v, ok := wc.mutation.Name(); ok {
		if err := warehouse.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Warehouse.name": %w`, err)}
		}
	}
	if _, ok := wc.mutation.Enabled(); !ok {
		return &ValidationError{Name: "enabled", err: errors.New(`ent: missing required field "Warehouse.enabled"`)}
	}
	return nil
}

func (wc *WarehouseCreate) sqlSave(ctx context.Context) (*Warehouse, error) {
	if err := wc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	wc.mutation.id = &_node.ID
	wc.mutation.done = true
	return _node, nil
}

func (wc *WarehouseCreate) createSpec() (*Warehouse, *sqlgraph.CreateSpec) {
	var (
		_node = &Warehouse{config: wc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: warehouse.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: warehouse.FieldID,
			},
		}
	)
	if id, ok := wc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := wc.mutation.Name(); ok {
		_spec.SetField(warehouse.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := wc.mutation.LastUpdate(); ok {
		_spec.SetField(warehouse.FieldLastUpdate, field.TypeTime, value)
		_node.LastUpdate = &value
	}
	if value, ok := wc.mutation.OriginalData(); ok {
		_spec.SetField(warehouse.FieldOriginalData, field.TypeString, value)
		_node.OriginalData = &value
	}
	if value, ok := wc.mutation.Enabled(); ok {
		_spec.SetField(warehouse.FieldEnabled, field.TypeBool, value)
		_node.Enabled = value
	}
	if value, ok := wc.mutation.Filters(); ok {
		_spec.SetField(warehouse.FieldFilters, field.TypeJSON, value)
		_node.Filters = value
	}
	if nodes := wc.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   warehouse.ProductsTable,
			Columns: []string{warehouse.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.VendorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   warehouse.VendorTable,
			Columns: []string{warehouse.VendorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: vendor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.vendor_warehouses = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WarehouseCreateBulk is the builder for creating many Warehouse entities in bulk.
type WarehouseCreateBulk struct {
	config
	builders []*WarehouseCreate
}

// Save creates the Warehouse entities in the database.
func (wcb *WarehouseCreateBulk) Save(ctx context.Context) ([]*Warehouse, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Warehouse, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WarehouseMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WarehouseCreateBulk) SaveX(ctx context.Context) []*Warehouse {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WarehouseCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WarehouseCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}
