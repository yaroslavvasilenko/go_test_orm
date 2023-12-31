// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yaroslavvasilenko/go_test_orm/orm/ent/order"
	"github.com/yaroslavvasilenko/go_test_orm/orm/ent/predicate"
)

// OrderUpdate is the builder for updating Order entities.
type OrderUpdate struct {
	config
	hooks    []Hook
	mutation *OrderMutation
}

// Where appends a list predicates to the OrderUpdate builder.
func (ou *OrderUpdate) Where(ps ...predicate.Order) *OrderUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetName sets the "name" field.
func (ou *OrderUpdate) SetName(s string) *OrderUpdate {
	ou.mutation.SetName(s)
	return ou
}

// SetAmount sets the "amount" field.
func (ou *OrderUpdate) SetAmount(i int) *OrderUpdate {
	ou.mutation.ResetAmount()
	ou.mutation.SetAmount(i)
	return ou
}

// AddAmount adds i to the "amount" field.
func (ou *OrderUpdate) AddAmount(i int) *OrderUpdate {
	ou.mutation.AddAmount(i)
	return ou
}

// SetPrice sets the "price" field.
func (ou *OrderUpdate) SetPrice(i int) *OrderUpdate {
	ou.mutation.ResetPrice()
	ou.mutation.SetPrice(i)
	return ou
}

// AddPrice adds i to the "price" field.
func (ou *OrderUpdate) AddPrice(i int) *OrderUpdate {
	ou.mutation.AddPrice(i)
	return ou
}

// SetCreatedAt sets the "created_at" field.
func (ou *OrderUpdate) SetCreatedAt(t time.Time) *OrderUpdate {
	ou.mutation.SetCreatedAt(t)
	return ou
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableCreatedAt(t *time.Time) *OrderUpdate {
	if t != nil {
		ou.SetCreatedAt(*t)
	}
	return ou
}

// SetUserOrdersID sets the "user_orders_id" field.
func (ou *OrderUpdate) SetUserOrdersID(i int) *OrderUpdate {
	ou.mutation.ResetUserOrdersID()
	ou.mutation.SetUserOrdersID(i)
	return ou
}

// AddUserOrdersID adds i to the "user_orders_id" field.
func (ou *OrderUpdate) AddUserOrdersID(i int) *OrderUpdate {
	ou.mutation.AddUserOrdersID(i)
	return ou
}

// Mutation returns the OrderMutation object of the builder.
func (ou *OrderUpdate) Mutation() *OrderMutation {
	return ou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrderUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrderUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrderUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrderUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ou *OrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(order.Table, order.Columns, sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.Name(); ok {
		_spec.SetField(order.FieldName, field.TypeString, value)
	}
	if value, ok := ou.mutation.Amount(); ok {
		_spec.SetField(order.FieldAmount, field.TypeInt, value)
	}
	if value, ok := ou.mutation.AddedAmount(); ok {
		_spec.AddField(order.FieldAmount, field.TypeInt, value)
	}
	if value, ok := ou.mutation.Price(); ok {
		_spec.SetField(order.FieldPrice, field.TypeInt, value)
	}
	if value, ok := ou.mutation.AddedPrice(); ok {
		_spec.AddField(order.FieldPrice, field.TypeInt, value)
	}
	if value, ok := ou.mutation.CreatedAt(); ok {
		_spec.SetField(order.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ou.mutation.UserOrdersID(); ok {
		_spec.SetField(order.FieldUserOrdersID, field.TypeInt, value)
	}
	if value, ok := ou.mutation.AddedUserOrdersID(); ok {
		_spec.AddField(order.FieldUserOrdersID, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OrderUpdateOne is the builder for updating a single Order entity.
type OrderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderMutation
}

// SetName sets the "name" field.
func (ouo *OrderUpdateOne) SetName(s string) *OrderUpdateOne {
	ouo.mutation.SetName(s)
	return ouo
}

// SetAmount sets the "amount" field.
func (ouo *OrderUpdateOne) SetAmount(i int) *OrderUpdateOne {
	ouo.mutation.ResetAmount()
	ouo.mutation.SetAmount(i)
	return ouo
}

// AddAmount adds i to the "amount" field.
func (ouo *OrderUpdateOne) AddAmount(i int) *OrderUpdateOne {
	ouo.mutation.AddAmount(i)
	return ouo
}

// SetPrice sets the "price" field.
func (ouo *OrderUpdateOne) SetPrice(i int) *OrderUpdateOne {
	ouo.mutation.ResetPrice()
	ouo.mutation.SetPrice(i)
	return ouo
}

// AddPrice adds i to the "price" field.
func (ouo *OrderUpdateOne) AddPrice(i int) *OrderUpdateOne {
	ouo.mutation.AddPrice(i)
	return ouo
}

// SetCreatedAt sets the "created_at" field.
func (ouo *OrderUpdateOne) SetCreatedAt(t time.Time) *OrderUpdateOne {
	ouo.mutation.SetCreatedAt(t)
	return ouo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableCreatedAt(t *time.Time) *OrderUpdateOne {
	if t != nil {
		ouo.SetCreatedAt(*t)
	}
	return ouo
}

// SetUserOrdersID sets the "user_orders_id" field.
func (ouo *OrderUpdateOne) SetUserOrdersID(i int) *OrderUpdateOne {
	ouo.mutation.ResetUserOrdersID()
	ouo.mutation.SetUserOrdersID(i)
	return ouo
}

// AddUserOrdersID adds i to the "user_orders_id" field.
func (ouo *OrderUpdateOne) AddUserOrdersID(i int) *OrderUpdateOne {
	ouo.mutation.AddUserOrdersID(i)
	return ouo
}

// Mutation returns the OrderMutation object of the builder.
func (ouo *OrderUpdateOne) Mutation() *OrderMutation {
	return ouo.mutation
}

// Where appends a list predicates to the OrderUpdate builder.
func (ouo *OrderUpdateOne) Where(ps ...predicate.Order) *OrderUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrderUpdateOne) Select(field string, fields ...string) *OrderUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Order entity.
func (ouo *OrderUpdateOne) Save(ctx context.Context) (*Order, error) {
	return withHooks(ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrderUpdateOne) SaveX(ctx context.Context) *Order {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrderUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrderUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ouo *OrderUpdateOne) sqlSave(ctx context.Context) (_node *Order, err error) {
	_spec := sqlgraph.NewUpdateSpec(order.Table, order.Columns, sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Order.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, order.FieldID)
		for _, f := range fields {
			if !order.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != order.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.Name(); ok {
		_spec.SetField(order.FieldName, field.TypeString, value)
	}
	if value, ok := ouo.mutation.Amount(); ok {
		_spec.SetField(order.FieldAmount, field.TypeInt, value)
	}
	if value, ok := ouo.mutation.AddedAmount(); ok {
		_spec.AddField(order.FieldAmount, field.TypeInt, value)
	}
	if value, ok := ouo.mutation.Price(); ok {
		_spec.SetField(order.FieldPrice, field.TypeInt, value)
	}
	if value, ok := ouo.mutation.AddedPrice(); ok {
		_spec.AddField(order.FieldPrice, field.TypeInt, value)
	}
	if value, ok := ouo.mutation.CreatedAt(); ok {
		_spec.SetField(order.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ouo.mutation.UserOrdersID(); ok {
		_spec.SetField(order.FieldUserOrdersID, field.TypeInt, value)
	}
	if value, ok := ouo.mutation.AddedUserOrdersID(); ok {
		_spec.AddField(order.FieldUserOrdersID, field.TypeInt, value)
	}
	_node = &Order{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}
