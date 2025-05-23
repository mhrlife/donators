// Code generated by ent, DO NOT EDIT.

package ent

import (
	"DonateNotifier/ent/predicate"
	"DonateNotifier/ent/processeddonate"
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeProcessedDonate = "ProcessedDonate"
)

// ProcessedDonateMutation represents an operation that mutates the ProcessedDonate nodes in the graph.
type ProcessedDonateMutation struct {
	config
	op            Op
	typ           string
	id            *string
	display_name  *string
	created_at    *int64
	addcreated_at *int64
	amount        *int64
	addamount     *int64
	currency      *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*ProcessedDonate, error)
	predicates    []predicate.ProcessedDonate
}

var _ ent.Mutation = (*ProcessedDonateMutation)(nil)

// processeddonateOption allows management of the mutation configuration using functional options.
type processeddonateOption func(*ProcessedDonateMutation)

// newProcessedDonateMutation creates new mutation for the ProcessedDonate entity.
func newProcessedDonateMutation(c config, op Op, opts ...processeddonateOption) *ProcessedDonateMutation {
	m := &ProcessedDonateMutation{
		config:        c,
		op:            op,
		typ:           TypeProcessedDonate,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withProcessedDonateID sets the ID field of the mutation.
func withProcessedDonateID(id string) processeddonateOption {
	return func(m *ProcessedDonateMutation) {
		var (
			err   error
			once  sync.Once
			value *ProcessedDonate
		)
		m.oldValue = func(ctx context.Context) (*ProcessedDonate, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().ProcessedDonate.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withProcessedDonate sets the old ProcessedDonate of the mutation.
func withProcessedDonate(node *ProcessedDonate) processeddonateOption {
	return func(m *ProcessedDonateMutation) {
		m.oldValue = func(context.Context) (*ProcessedDonate, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ProcessedDonateMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ProcessedDonateMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of ProcessedDonate entities.
func (m *ProcessedDonateMutation) SetID(id string) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ProcessedDonateMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ProcessedDonateMutation) IDs(ctx context.Context) ([]string, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []string{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().ProcessedDonate.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetDisplayName sets the "display_name" field.
func (m *ProcessedDonateMutation) SetDisplayName(s string) {
	m.display_name = &s
}

// DisplayName returns the value of the "display_name" field in the mutation.
func (m *ProcessedDonateMutation) DisplayName() (r string, exists bool) {
	v := m.display_name
	if v == nil {
		return
	}
	return *v, true
}

// OldDisplayName returns the old "display_name" field's value of the ProcessedDonate entity.
// If the ProcessedDonate object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProcessedDonateMutation) OldDisplayName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDisplayName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDisplayName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDisplayName: %w", err)
	}
	return oldValue.DisplayName, nil
}

// ResetDisplayName resets all changes to the "display_name" field.
func (m *ProcessedDonateMutation) ResetDisplayName() {
	m.display_name = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *ProcessedDonateMutation) SetCreatedAt(i int64) {
	m.created_at = &i
	m.addcreated_at = nil
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *ProcessedDonateMutation) CreatedAt() (r int64, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the ProcessedDonate entity.
// If the ProcessedDonate object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProcessedDonateMutation) OldCreatedAt(ctx context.Context) (v int64, err error) {
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

// AddCreatedAt adds i to the "created_at" field.
func (m *ProcessedDonateMutation) AddCreatedAt(i int64) {
	if m.addcreated_at != nil {
		*m.addcreated_at += i
	} else {
		m.addcreated_at = &i
	}
}

// AddedCreatedAt returns the value that was added to the "created_at" field in this mutation.
func (m *ProcessedDonateMutation) AddedCreatedAt() (r int64, exists bool) {
	v := m.addcreated_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *ProcessedDonateMutation) ResetCreatedAt() {
	m.created_at = nil
	m.addcreated_at = nil
}

// SetAmount sets the "amount" field.
func (m *ProcessedDonateMutation) SetAmount(i int64) {
	m.amount = &i
	m.addamount = nil
}

// Amount returns the value of the "amount" field in the mutation.
func (m *ProcessedDonateMutation) Amount() (r int64, exists bool) {
	v := m.amount
	if v == nil {
		return
	}
	return *v, true
}

// OldAmount returns the old "amount" field's value of the ProcessedDonate entity.
// If the ProcessedDonate object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProcessedDonateMutation) OldAmount(ctx context.Context) (v int64, err error) {
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
func (m *ProcessedDonateMutation) AddAmount(i int64) {
	if m.addamount != nil {
		*m.addamount += i
	} else {
		m.addamount = &i
	}
}

// AddedAmount returns the value that was added to the "amount" field in this mutation.
func (m *ProcessedDonateMutation) AddedAmount() (r int64, exists bool) {
	v := m.addamount
	if v == nil {
		return
	}
	return *v, true
}

// ResetAmount resets all changes to the "amount" field.
func (m *ProcessedDonateMutation) ResetAmount() {
	m.amount = nil
	m.addamount = nil
}

// SetCurrency sets the "currency" field.
func (m *ProcessedDonateMutation) SetCurrency(s string) {
	m.currency = &s
}

// Currency returns the value of the "currency" field in the mutation.
func (m *ProcessedDonateMutation) Currency() (r string, exists bool) {
	v := m.currency
	if v == nil {
		return
	}
	return *v, true
}

// OldCurrency returns the old "currency" field's value of the ProcessedDonate entity.
// If the ProcessedDonate object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProcessedDonateMutation) OldCurrency(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCurrency is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCurrency requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCurrency: %w", err)
	}
	return oldValue.Currency, nil
}

// ResetCurrency resets all changes to the "currency" field.
func (m *ProcessedDonateMutation) ResetCurrency() {
	m.currency = nil
}

// Where appends a list predicates to the ProcessedDonateMutation builder.
func (m *ProcessedDonateMutation) Where(ps ...predicate.ProcessedDonate) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ProcessedDonateMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ProcessedDonateMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.ProcessedDonate, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ProcessedDonateMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ProcessedDonateMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (ProcessedDonate).
func (m *ProcessedDonateMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ProcessedDonateMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.display_name != nil {
		fields = append(fields, processeddonate.FieldDisplayName)
	}
	if m.created_at != nil {
		fields = append(fields, processeddonate.FieldCreatedAt)
	}
	if m.amount != nil {
		fields = append(fields, processeddonate.FieldAmount)
	}
	if m.currency != nil {
		fields = append(fields, processeddonate.FieldCurrency)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ProcessedDonateMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case processeddonate.FieldDisplayName:
		return m.DisplayName()
	case processeddonate.FieldCreatedAt:
		return m.CreatedAt()
	case processeddonate.FieldAmount:
		return m.Amount()
	case processeddonate.FieldCurrency:
		return m.Currency()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ProcessedDonateMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case processeddonate.FieldDisplayName:
		return m.OldDisplayName(ctx)
	case processeddonate.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case processeddonate.FieldAmount:
		return m.OldAmount(ctx)
	case processeddonate.FieldCurrency:
		return m.OldCurrency(ctx)
	}
	return nil, fmt.Errorf("unknown ProcessedDonate field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ProcessedDonateMutation) SetField(name string, value ent.Value) error {
	switch name {
	case processeddonate.FieldDisplayName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDisplayName(v)
		return nil
	case processeddonate.FieldCreatedAt:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case processeddonate.FieldAmount:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAmount(v)
		return nil
	case processeddonate.FieldCurrency:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCurrency(v)
		return nil
	}
	return fmt.Errorf("unknown ProcessedDonate field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ProcessedDonateMutation) AddedFields() []string {
	var fields []string
	if m.addcreated_at != nil {
		fields = append(fields, processeddonate.FieldCreatedAt)
	}
	if m.addamount != nil {
		fields = append(fields, processeddonate.FieldAmount)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ProcessedDonateMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case processeddonate.FieldCreatedAt:
		return m.AddedCreatedAt()
	case processeddonate.FieldAmount:
		return m.AddedAmount()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ProcessedDonateMutation) AddField(name string, value ent.Value) error {
	switch name {
	case processeddonate.FieldCreatedAt:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddCreatedAt(v)
		return nil
	case processeddonate.FieldAmount:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAmount(v)
		return nil
	}
	return fmt.Errorf("unknown ProcessedDonate numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ProcessedDonateMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ProcessedDonateMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ProcessedDonateMutation) ClearField(name string) error {
	return fmt.Errorf("unknown ProcessedDonate nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ProcessedDonateMutation) ResetField(name string) error {
	switch name {
	case processeddonate.FieldDisplayName:
		m.ResetDisplayName()
		return nil
	case processeddonate.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case processeddonate.FieldAmount:
		m.ResetAmount()
		return nil
	case processeddonate.FieldCurrency:
		m.ResetCurrency()
		return nil
	}
	return fmt.Errorf("unknown ProcessedDonate field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ProcessedDonateMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ProcessedDonateMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ProcessedDonateMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ProcessedDonateMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ProcessedDonateMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ProcessedDonateMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ProcessedDonateMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown ProcessedDonate unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ProcessedDonateMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown ProcessedDonate edge %s", name)
}
