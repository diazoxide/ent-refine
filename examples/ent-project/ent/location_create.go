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

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/company"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/country"
	"github.com/diazoxide/ent-refine/examples/ent-project/ent/location"
	"github.com/google/uuid"
)

// LocationCreate is the builder for creating a Location entity.
type LocationCreate struct {
	config
	mutation *LocationMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (lc *LocationCreate) SetTitle(s string) *LocationCreate {
	lc.mutation.SetTitle(s)
	return lc
}

// SetDescription sets the "description" field.
func (lc *LocationCreate) SetDescription(s string) *LocationCreate {
	lc.mutation.SetDescription(s)
	return lc
}

// SetLatitude sets the "latitude" field.
func (lc *LocationCreate) SetLatitude(f float64) *LocationCreate {
	lc.mutation.SetLatitude(f)
	return lc
}

// SetLongitude sets the "longitude" field.
func (lc *LocationCreate) SetLongitude(f float64) *LocationCreate {
	lc.mutation.SetLongitude(f)
	return lc
}

// SetAddress sets the "address" field.
func (lc *LocationCreate) SetAddress(s string) *LocationCreate {
	lc.mutation.SetAddress(s)
	return lc
}

// SetPostcode sets the "postcode" field.
func (lc *LocationCreate) SetPostcode(s string) *LocationCreate {
	lc.mutation.SetPostcode(s)
	return lc
}

// SetType sets the "type" field.
func (lc *LocationCreate) SetType(s string) *LocationCreate {
	lc.mutation.SetType(s)
	return lc
}

// SetState sets the "state" field.
func (lc *LocationCreate) SetState(s string) *LocationCreate {
	lc.mutation.SetState(s)
	return lc
}

// SetSuburb sets the "suburb" field.
func (lc *LocationCreate) SetSuburb(s string) *LocationCreate {
	lc.mutation.SetSuburb(s)
	return lc
}

// SetStreetType sets the "street_type" field.
func (lc *LocationCreate) SetStreetType(s string) *LocationCreate {
	lc.mutation.SetStreetType(s)
	return lc
}

// SetStreetName sets the "street_name" field.
func (lc *LocationCreate) SetStreetName(s string) *LocationCreate {
	lc.mutation.SetStreetName(s)
	return lc
}

// SetID sets the "id" field.
func (lc *LocationCreate) SetID(u uuid.UUID) *LocationCreate {
	lc.mutation.SetID(u)
	return lc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (lc *LocationCreate) SetNillableID(u *uuid.UUID) *LocationCreate {
	if u != nil {
		lc.SetID(*u)
	}
	return lc
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (lc *LocationCreate) SetCompanyID(id uuid.UUID) *LocationCreate {
	lc.mutation.SetCompanyID(id)
	return lc
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (lc *LocationCreate) SetNillableCompanyID(id *uuid.UUID) *LocationCreate {
	if id != nil {
		lc = lc.SetCompanyID(*id)
	}
	return lc
}

// SetCompany sets the "company" edge to the Company entity.
func (lc *LocationCreate) SetCompany(c *Company) *LocationCreate {
	return lc.SetCompanyID(c.ID)
}

// SetCountryID sets the "country" edge to the Country entity by ID.
func (lc *LocationCreate) SetCountryID(id uuid.UUID) *LocationCreate {
	lc.mutation.SetCountryID(id)
	return lc
}

// SetNillableCountryID sets the "country" edge to the Country entity by ID if the given value is not nil.
func (lc *LocationCreate) SetNillableCountryID(id *uuid.UUID) *LocationCreate {
	if id != nil {
		lc = lc.SetCountryID(*id)
	}
	return lc
}

// SetCountry sets the "country" edge to the Country entity.
func (lc *LocationCreate) SetCountry(c *Country) *LocationCreate {
	return lc.SetCountryID(c.ID)
}

// Mutation returns the LocationMutation object of the builder.
func (lc *LocationCreate) Mutation() *LocationMutation {
	return lc.mutation
}

// Save creates the Location in the database.
func (lc *LocationCreate) Save(ctx context.Context) (*Location, error) {
	lc.defaults()
	return withHooks[*Location, LocationMutation](ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LocationCreate) SaveX(ctx context.Context) *Location {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LocationCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LocationCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LocationCreate) defaults() {
	if _, ok := lc.mutation.ID(); !ok {
		v := location.DefaultID()
		lc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LocationCreate) check() error {
	if _, ok := lc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Location.title"`)}
	}
	if v, ok := lc.mutation.Title(); ok {
		if err := location.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Location.title": %w`, err)}
		}
	}
	if _, ok := lc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Location.description"`)}
	}
	if v, ok := lc.mutation.Description(); ok {
		if err := location.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Location.description": %w`, err)}
		}
	}
	if _, ok := lc.mutation.Latitude(); !ok {
		return &ValidationError{Name: "latitude", err: errors.New(`ent: missing required field "Location.latitude"`)}
	}
	if _, ok := lc.mutation.Longitude(); !ok {
		return &ValidationError{Name: "longitude", err: errors.New(`ent: missing required field "Location.longitude"`)}
	}
	if _, ok := lc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "Location.address"`)}
	}
	if v, ok := lc.mutation.Address(); ok {
		if err := location.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Location.address": %w`, err)}
		}
	}
	if _, ok := lc.mutation.Postcode(); !ok {
		return &ValidationError{Name: "postcode", err: errors.New(`ent: missing required field "Location.postcode"`)}
	}
	if v, ok := lc.mutation.Postcode(); ok {
		if err := location.PostcodeValidator(v); err != nil {
			return &ValidationError{Name: "postcode", err: fmt.Errorf(`ent: validator failed for field "Location.postcode": %w`, err)}
		}
	}
	if _, ok := lc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Location.type"`)}
	}
	if v, ok := lc.mutation.GetType(); ok {
		if err := location.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Location.type": %w`, err)}
		}
	}
	if _, ok := lc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "Location.state"`)}
	}
	if v, ok := lc.mutation.State(); ok {
		if err := location.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "Location.state": %w`, err)}
		}
	}
	if _, ok := lc.mutation.Suburb(); !ok {
		return &ValidationError{Name: "suburb", err: errors.New(`ent: missing required field "Location.suburb"`)}
	}
	if v, ok := lc.mutation.Suburb(); ok {
		if err := location.SuburbValidator(v); err != nil {
			return &ValidationError{Name: "suburb", err: fmt.Errorf(`ent: validator failed for field "Location.suburb": %w`, err)}
		}
	}
	if _, ok := lc.mutation.StreetType(); !ok {
		return &ValidationError{Name: "street_type", err: errors.New(`ent: missing required field "Location.street_type"`)}
	}
	if v, ok := lc.mutation.StreetType(); ok {
		if err := location.StreetTypeValidator(v); err != nil {
			return &ValidationError{Name: "street_type", err: fmt.Errorf(`ent: validator failed for field "Location.street_type": %w`, err)}
		}
	}
	if _, ok := lc.mutation.StreetName(); !ok {
		return &ValidationError{Name: "street_name", err: errors.New(`ent: missing required field "Location.street_name"`)}
	}
	if v, ok := lc.mutation.StreetName(); ok {
		if err := location.StreetNameValidator(v); err != nil {
			return &ValidationError{Name: "street_name", err: fmt.Errorf(`ent: validator failed for field "Location.street_name": %w`, err)}
		}
	}
	return nil
}

func (lc *LocationCreate) sqlSave(ctx context.Context) (*Location, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
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
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LocationCreate) createSpec() (*Location, *sqlgraph.CreateSpec) {
	var (
		_node = &Location{config: lc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: location.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: location.FieldID,
			},
		}
	)
	if id, ok := lc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := lc.mutation.Title(); ok {
		_spec.SetField(location.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := lc.mutation.Description(); ok {
		_spec.SetField(location.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := lc.mutation.Latitude(); ok {
		_spec.SetField(location.FieldLatitude, field.TypeFloat64, value)
		_node.Latitude = value
	}
	if value, ok := lc.mutation.Longitude(); ok {
		_spec.SetField(location.FieldLongitude, field.TypeFloat64, value)
		_node.Longitude = value
	}
	if value, ok := lc.mutation.Address(); ok {
		_spec.SetField(location.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := lc.mutation.Postcode(); ok {
		_spec.SetField(location.FieldPostcode, field.TypeString, value)
		_node.Postcode = value
	}
	if value, ok := lc.mutation.GetType(); ok {
		_spec.SetField(location.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := lc.mutation.State(); ok {
		_spec.SetField(location.FieldState, field.TypeString, value)
		_node.State = value
	}
	if value, ok := lc.mutation.Suburb(); ok {
		_spec.SetField(location.FieldSuburb, field.TypeString, value)
		_node.Suburb = value
	}
	if value, ok := lc.mutation.StreetType(); ok {
		_spec.SetField(location.FieldStreetType, field.TypeString, value)
		_node.StreetType = value
	}
	if value, ok := lc.mutation.StreetName(); ok {
		_spec.SetField(location.FieldStreetName, field.TypeString, value)
		_node.StreetName = value
	}
	if nodes := lc.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   location.CompanyTable,
			Columns: []string{location.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: company.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.company_locations = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.CountryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   location.CountryTable,
			Columns: []string{location.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: country.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.country_locations = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LocationCreateBulk is the builder for creating many Location entities in bulk.
type LocationCreateBulk struct {
	config
	builders []*LocationCreate
}

// Save creates the Location entities in the database.
func (lcb *LocationCreateBulk) Save(ctx context.Context) ([]*Location, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Location, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LocationMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LocationCreateBulk) SaveX(ctx context.Context) []*Location {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LocationCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LocationCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
