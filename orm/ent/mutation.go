// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/yaroslavvasilenko/go_test_orm/orm/ent/order"
	"github.com/yaroslavvasilenko/go_test_orm/orm/ent/predicate"
	"github.com/yaroslavvasilenko/go_test_orm/orm/ent/user"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeOrder = "Order"
	TypeUser  = "User"
)

// OrderMutation represents an operation that mutates the Order nodes in the graph.
type OrderMutation struct {
	config
	op                Op
	typ               string
	id                *int
	name              *string
	amount            *int
	addamount         *int
	price             *int
	addprice          *int
	created_at        *time.Time
	user_orders_id    *int
	adduser_orders_id *int
	clearedFields     map[string]struct{}
	done              bool
	oldValue          func(context.Context) (*Order, error)
	predicates        []predicate.Order
}

var _ ent.Mutation = (*OrderMutation)(nil)

// orderOption allows management of the mutation configuration using functional options.
type orderOption func(*OrderMutation)

// newOrderMutation creates new mutation for the Order entity.
func newOrderMutation(c config, op Op, opts ...orderOption) *OrderMutation {
	m := &OrderMutation{
		config:        c,
		op:            op,
		typ:           TypeOrder,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withOrderID sets the ID field of the mutation.
func withOrderID(id int) orderOption {
	return func(m *OrderMutation) {
		var (
			err   error
			once  sync.Once
			value *Order
		)
		m.oldValue = func(ctx context.Context) (*Order, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Order.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withOrder sets the old Order of the mutation.
func withOrder(node *Order) orderOption {
	return func(m *OrderMutation) {
		m.oldValue = func(context.Context) (*Order, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m OrderMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m OrderMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *OrderMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *OrderMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Order.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *OrderMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *OrderMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Order entity.
// If the Order object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *OrderMutation) ResetName() {
	m.name = nil
}

// SetAmount sets the "amount" field.
func (m *OrderMutation) SetAmount(i int) {
	m.amount = &i
	m.addamount = nil
}

// Amount returns the value of the "amount" field in the mutation.
func (m *OrderMutation) Amount() (r int, exists bool) {
	v := m.amount
	if v == nil {
		return
	}
	return *v, true
}

// OldAmount returns the old "amount" field's value of the Order entity.
// If the Order object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderMutation) OldAmount(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAmount is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAmount requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAmount: %w", err)
	}
	return oldValue.Amount, nil
}

// AddAmount adds i to the "amount" field.
func (m *OrderMutation) AddAmount(i int) {
	if m.addamount != nil {
		*m.addamount += i
	} else {
		m.addamount = &i
	}
}

// AddedAmount returns the value that was added to the "amount" field in this mutation.
func (m *OrderMutation) AddedAmount() (r int, exists bool) {
	v := m.addamount
	if v == nil {
		return
	}
	return *v, true
}

// ResetAmount resets all changes to the "amount" field.
func (m *OrderMutation) ResetAmount() {
	m.amount = nil
	m.addamount = nil
}

// SetPrice sets the "price" field.
func (m *OrderMutation) SetPrice(i int) {
	m.price = &i
	m.addprice = nil
}

// Price returns the value of the "price" field in the mutation.
func (m *OrderMutation) Price() (r int, exists bool) {
	v := m.price
	if v == nil {
		return
	}
	return *v, true
}

// OldPrice returns the old "price" field's value of the Order entity.
// If the Order object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderMutation) OldPrice(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPrice is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPrice requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPrice: %w", err)
	}
	return oldValue.Price, nil
}

// AddPrice adds i to the "price" field.
func (m *OrderMutation) AddPrice(i int) {
	if m.addprice != nil {
		*m.addprice += i
	} else {
		m.addprice = &i
	}
}

// AddedPrice returns the value that was added to the "price" field in this mutation.
func (m *OrderMutation) AddedPrice() (r int, exists bool) {
	v := m.addprice
	if v == nil {
		return
	}
	return *v, true
}

// ResetPrice resets all changes to the "price" field.
func (m *OrderMutation) ResetPrice() {
	m.price = nil
	m.addprice = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *OrderMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *OrderMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Order entity.
// If the Order object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *OrderMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUserOrdersID sets the "user_orders_id" field.
func (m *OrderMutation) SetUserOrdersID(i int) {
	m.user_orders_id = &i
	m.adduser_orders_id = nil
}

// UserOrdersID returns the value of the "user_orders_id" field in the mutation.
func (m *OrderMutation) UserOrdersID() (r int, exists bool) {
	v := m.user_orders_id
	if v == nil {
		return
	}
	return *v, true
}

// OldUserOrdersID returns the old "user_orders_id" field's value of the Order entity.
// If the Order object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OrderMutation) OldUserOrdersID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUserOrdersID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUserOrdersID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserOrdersID: %w", err)
	}
	return oldValue.UserOrdersID, nil
}

// AddUserOrdersID adds i to the "user_orders_id" field.
func (m *OrderMutation) AddUserOrdersID(i int) {
	if m.adduser_orders_id != nil {
		*m.adduser_orders_id += i
	} else {
		m.adduser_orders_id = &i
	}
}

// AddedUserOrdersID returns the value that was added to the "user_orders_id" field in this mutation.
func (m *OrderMutation) AddedUserOrdersID() (r int, exists bool) {
	v := m.adduser_orders_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetUserOrdersID resets all changes to the "user_orders_id" field.
func (m *OrderMutation) ResetUserOrdersID() {
	m.user_orders_id = nil
	m.adduser_orders_id = nil
}

// Where appends a list predicates to the OrderMutation builder.
func (m *OrderMutation) Where(ps ...predicate.Order) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the OrderMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *OrderMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Order, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *OrderMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *OrderMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Order).
func (m *OrderMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *OrderMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.name != nil {
		fields = append(fields, order.FieldName)
	}
	if m.amount != nil {
		fields = append(fields, order.FieldAmount)
	}
	if m.price != nil {
		fields = append(fields, order.FieldPrice)
	}
	if m.created_at != nil {
		fields = append(fields, order.FieldCreatedAt)
	}
	if m.user_orders_id != nil {
		fields = append(fields, order.FieldUserOrdersID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *OrderMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case order.FieldName:
		return m.Name()
	case order.FieldAmount:
		return m.Amount()
	case order.FieldPrice:
		return m.Price()
	case order.FieldCreatedAt:
		return m.CreatedAt()
	case order.FieldUserOrdersID:
		return m.UserOrdersID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *OrderMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case order.FieldName:
		return m.OldName(ctx)
	case order.FieldAmount:
		return m.OldAmount(ctx)
	case order.FieldPrice:
		return m.OldPrice(ctx)
	case order.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case order.FieldUserOrdersID:
		return m.OldUserOrdersID(ctx)
	}
	return nil, fmt.Errorf("unknown Order field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *OrderMutation) SetField(name string, value ent.Value) error {
	switch name {
	case order.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case order.FieldAmount:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAmount(v)
		return nil
	case order.FieldPrice:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPrice(v)
		return nil
	case order.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case order.FieldUserOrdersID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserOrdersID(v)
		return nil
	}
	return fmt.Errorf("unknown Order field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *OrderMutation) AddedFields() []string {
	var fields []string
	if m.addamount != nil {
		fields = append(fields, order.FieldAmount)
	}
	if m.addprice != nil {
		fields = append(fields, order.FieldPrice)
	}
	if m.adduser_orders_id != nil {
		fields = append(fields, order.FieldUserOrdersID)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *OrderMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case order.FieldAmount:
		return m.AddedAmount()
	case order.FieldPrice:
		return m.AddedPrice()
	case order.FieldUserOrdersID:
		return m.AddedUserOrdersID()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *OrderMutation) AddField(name string, value ent.Value) error {
	switch name {
	case order.FieldAmount:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAmount(v)
		return nil
	case order.FieldPrice:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddPrice(v)
		return nil
	case order.FieldUserOrdersID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddUserOrdersID(v)
		return nil
	}
	return fmt.Errorf("unknown Order numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *OrderMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *OrderMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *OrderMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Order nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *OrderMutation) ResetField(name string) error {
	switch name {
	case order.FieldName:
		m.ResetName()
		return nil
	case order.FieldAmount:
		m.ResetAmount()
		return nil
	case order.FieldPrice:
		m.ResetPrice()
		return nil
	case order.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case order.FieldUserOrdersID:
		m.ResetUserOrdersID()
		return nil
	}
	return fmt.Errorf("unknown Order field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *OrderMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *OrderMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *OrderMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *OrderMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *OrderMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *OrderMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *OrderMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Order unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *OrderMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Order edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	age           *int
	addage        *int
	created_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// SetAge sets the "age" field.
func (m *UserMutation) SetAge(i int) {
	m.age = &i
	m.addage = nil
}

// Age returns the value of the "age" field in the mutation.
func (m *UserMutation) Age() (r int, exists bool) {
	v := m.age
	if v == nil {
		return
	}
	return *v, true
}

// OldAge returns the old "age" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldAge(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAge is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAge requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAge: %w", err)
	}
	return oldValue.Age, nil
}

// AddAge adds i to the "age" field.
func (m *UserMutation) AddAge(i int) {
	if m.addage != nil {
		*m.addage += i
	} else {
		m.addage = &i
	}
}

// AddedAge returns the value that was added to the "age" field in this mutation.
func (m *UserMutation) AddedAge() (r int, exists bool) {
	v := m.addage
	if v == nil {
		return
	}
	return *v, true
}

// ResetAge resets all changes to the "age" field.
func (m *UserMutation) ResetAge() {
	m.age = nil
	m.addage = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *UserMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *UserMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *UserMutation) ResetCreatedAt() {
	m.created_at = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	if m.age != nil {
		fields = append(fields, user.FieldAge)
	}
	if m.created_at != nil {
		fields = append(fields, user.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldName:
		return m.Name()
	case user.FieldAge:
		return m.Age()
	case user.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldName:
		return m.OldName(ctx)
	case user.FieldAge:
		return m.OldAge(ctx)
	case user.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case user.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAge(v)
		return nil
	case user.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	var fields []string
	if m.addage != nil {
		fields = append(fields, user.FieldAge)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case user.FieldAge:
		return m.AddedAge()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case user.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAge(v)
		return nil
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldName:
		m.ResetName()
		return nil
	case user.FieldAge:
		m.ResetAge()
		return nil
	case user.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}
