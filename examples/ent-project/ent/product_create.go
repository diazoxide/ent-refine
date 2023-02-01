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
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/schema/enums"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/vendor"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/warehouse"
	"github.com/google/uuid"
)

// ProductCreate is the builder for creating a Product entity.
type ProductCreate struct {
	config
	mutation *ProductMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *ProductCreate) SetName(s string) *ProductCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *ProductCreate) SetDescription(s string) *ProductCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetImage sets the "image" field.
func (pc *ProductCreate) SetImage(s string) *ProductCreate {
	pc.mutation.SetImage(s)
	return pc
}

// SetURL sets the "url" field.
func (pc *ProductCreate) SetURL(s string) *ProductCreate {
	pc.mutation.SetURL(s)
	return pc
}

// SetLastSell sets the "last_sell" field.
func (pc *ProductCreate) SetLastSell(t time.Time) *ProductCreate {
	pc.mutation.SetLastSell(t)
	return pc
}

// SetNillableLastSell sets the "last_sell" field if the given value is not nil.
func (pc *ProductCreate) SetNillableLastSell(t *time.Time) *ProductCreate {
	if t != nil {
		pc.SetLastSell(*t)
	}
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *ProductCreate) SetCreatedAt(t time.Time) *ProductCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *ProductCreate) SetNillableCreatedAt(t *time.Time) *ProductCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetStatus sets the "status" field.
func (pc *ProductCreate) SetStatus(es enums.ProcessStatus) *ProductCreate {
	pc.mutation.SetStatus(es)
	return pc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pc *ProductCreate) SetNillableStatus(es *enums.ProcessStatus) *ProductCreate {
	if es != nil {
		pc.SetStatus(*es)
	}
	return pc
}

// SetBuildStatus sets the "build_status" field.
func (pc *ProductCreate) SetBuildStatus(es enums.ProcessStatus) *ProductCreate {
	pc.mutation.SetBuildStatus(es)
	return pc
}

// SetNillableBuildStatus sets the "build_status" field if the given value is not nil.
func (pc *ProductCreate) SetNillableBuildStatus(es *enums.ProcessStatus) *ProductCreate {
	if es != nil {
		pc.SetBuildStatus(*es)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *ProductCreate) SetID(u uuid.UUID) *ProductCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *ProductCreate) SetNillableID(u *uuid.UUID) *ProductCreate {
	if u != nil {
		pc.SetID(*u)
	}
	return pc
}

// SetWarehouseID sets the "warehouse" edge to the Warehouse entity by ID.
func (pc *ProductCreate) SetWarehouseID(id uuid.UUID) *ProductCreate {
	pc.mutation.SetWarehouseID(id)
	return pc
}

// SetNillableWarehouseID sets the "warehouse" edge to the Warehouse entity by ID if the given value is not nil.
func (pc *ProductCreate) SetNillableWarehouseID(id *uuid.UUID) *ProductCreate {
	if id != nil {
		pc = pc.SetWarehouseID(*id)
	}
	return pc
}

// SetWarehouse sets the "warehouse" edge to the Warehouse entity.
func (pc *ProductCreate) SetWarehouse(w *Warehouse) *ProductCreate {
	return pc.SetWarehouseID(w.ID)
}

// SetVendorID sets the "vendor" edge to the Vendor entity by ID.
func (pc *ProductCreate) SetVendorID(id uuid.UUID) *ProductCreate {
	pc.mutation.SetVendorID(id)
	return pc
}

// SetNillableVendorID sets the "vendor" edge to the Vendor entity by ID if the given value is not nil.
func (pc *ProductCreate) SetNillableVendorID(id *uuid.UUID) *ProductCreate {
	if id != nil {
		pc = pc.SetVendorID(*id)
	}
	return pc
}

// SetVendor sets the "vendor" edge to the Vendor entity.
func (pc *ProductCreate) SetVendor(v *Vendor) *ProductCreate {
	return pc.SetVendorID(v.ID)
}

// Mutation returns the ProductMutation object of the builder.
func (pc *ProductCreate) Mutation() *ProductMutation {
	return pc.mutation
}

// Save creates the Product in the database.
func (pc *ProductCreate) Save(ctx context.Context) (*Product, error) {
	pc.defaults()
	return withHooks[*Product, ProductMutation](ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProductCreate) SaveX(ctx context.Context) *Product {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProductCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProductCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProductCreate) defaults() {
	if _, ok := pc.mutation.Status(); !ok {
		v := product.DefaultStatus
		pc.mutation.SetStatus(v)
	}
	if _, ok := pc.mutation.BuildStatus(); !ok {
		v := product.DefaultBuildStatus
		pc.mutation.SetBuildStatus(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := product.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProductCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Product.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := product.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Product.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Product.description"`)}
	}
	if v, ok := pc.mutation.Description(); ok {
		if err := product.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Product.description": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Image(); !ok {
		return &ValidationError{Name: "image", err: errors.New(`ent: missing required field "Product.image"`)}
	}
	if v, ok := pc.mutation.Image(); ok {
		if err := product.ImageValidator(v); err != nil {
			return &ValidationError{Name: "image", err: fmt.Errorf(`ent: validator failed for field "Product.image": %w`, err)}
		}
	}
	if _, ok := pc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Product.url"`)}
	}
	if v, ok := pc.mutation.URL(); ok {
		if err := product.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Product.url": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Product.status"`)}
	}
	if v, ok := pc.mutation.Status(); ok {
		if err := product.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Product.status": %w`, err)}
		}
	}
	if _, ok := pc.mutation.BuildStatus(); !ok {
		return &ValidationError{Name: "build_status", err: errors.New(`ent: missing required field "Product.build_status"`)}
	}
	if v, ok := pc.mutation.BuildStatus(); ok {
		if err := product.BuildStatusValidator(v); err != nil {
			return &ValidationError{Name: "build_status", err: fmt.Errorf(`ent: validator failed for field "Product.build_status": %w`, err)}
		}
	}
	return nil
}

func (pc *ProductCreate) sqlSave(ctx context.Context) (*Product, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
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
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProductCreate) createSpec() (*Product, *sqlgraph.CreateSpec) {
	var (
		_node = &Product{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: product.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: product.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(product.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(product.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.Image(); ok {
		_spec.SetField(product.FieldImage, field.TypeString, value)
		_node.Image = value
	}
	if value, ok := pc.mutation.URL(); ok {
		_spec.SetField(product.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := pc.mutation.LastSell(); ok {
		_spec.SetField(product.FieldLastSell, field.TypeTime, value)
		_node.LastSell = &value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(product.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = &value
	}
	if value, ok := pc.mutation.Status(); ok {
		_spec.SetField(product.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := pc.mutation.BuildStatus(); ok {
		_spec.SetField(product.FieldBuildStatus, field.TypeEnum, value)
		_node.BuildStatus = value
	}
	if nodes := pc.mutation.WarehouseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.WarehouseTable,
			Columns: []string{product.WarehouseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: warehouse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.warehouse_products = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.VendorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.VendorTable,
			Columns: []string{product.VendorColumn},
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
		_node.vendor_products = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProductCreateBulk is the builder for creating many Product entities in bulk.
type ProductCreateBulk struct {
	config
	builders []*ProductCreate
}

// Save creates the Product entities in the database.
func (pcb *ProductCreateBulk) Save(ctx context.Context) ([]*Product, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Product, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProductCreateBulk) SaveX(ctx context.Context) []*Product {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProductCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProductCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
