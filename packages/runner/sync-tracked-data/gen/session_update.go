// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-tracked-data/gen/predicate"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-tracked-data/gen/session"
)

// SessionUpdate is the builder for updating Session entities.
type SessionUpdate struct {
	config
	hooks    []Hook
	mutation *SessionMutation
}

// Where appends a list predicates to the SessionUpdate builder.
func (su *SessionUpdate) Where(ps ...predicate.Session) *SessionUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetAppID sets the "app_id" field.
func (su *SessionUpdate) SetAppID(s string) *SessionUpdate {
	su.mutation.SetAppID(s)
	return su
}

// SetNameTracker sets the "name_tracker" field.
func (su *SessionUpdate) SetNameTracker(s string) *SessionUpdate {
	su.mutation.SetNameTracker(s)
	return su
}

// SetTenant sets the "tenant" field.
func (su *SessionUpdate) SetTenant(s string) *SessionUpdate {
	su.mutation.SetTenant(s)
	return su
}

// SetDomainSessionid sets the "domain_sessionid" field.
func (su *SessionUpdate) SetDomainSessionid(s string) *SessionUpdate {
	su.mutation.SetDomainSessionid(s)
	return su
}

// SetDomainSessionidx sets the "domain_sessionidx" field.
func (su *SessionUpdate) SetDomainSessionidx(i int) *SessionUpdate {
	su.mutation.ResetDomainSessionidx()
	su.mutation.SetDomainSessionidx(i)
	return su
}

// AddDomainSessionidx adds i to the "domain_sessionidx" field.
func (su *SessionUpdate) AddDomainSessionidx(i int) *SessionUpdate {
	su.mutation.AddDomainSessionidx(i)
	return su
}

// SetSyncedToCustomerOs sets the "synced_to_customer_os" field.
func (su *SessionUpdate) SetSyncedToCustomerOs(b bool) *SessionUpdate {
	su.mutation.SetSyncedToCustomerOs(b)
	return su
}

// SetStartTstamp sets the "start_tstamp" field.
func (su *SessionUpdate) SetStartTstamp(t time.Time) *SessionUpdate {
	su.mutation.SetStartTstamp(t)
	return su
}

// SetEndTstamp sets the "end_tstamp" field.
func (su *SessionUpdate) SetEndTstamp(t time.Time) *SessionUpdate {
	su.mutation.SetEndTstamp(t)
	return su
}

// SetDomainUserid sets the "domain_userid" field.
func (su *SessionUpdate) SetDomainUserid(s string) *SessionUpdate {
	su.mutation.SetDomainUserid(s)
	return su
}

// SetNetworkUserid sets the "network_userid" field.
func (su *SessionUpdate) SetNetworkUserid(s string) *SessionUpdate {
	su.mutation.SetNetworkUserid(s)
	return su
}

// SetVisitorID sets the "visitor_id" field.
func (su *SessionUpdate) SetVisitorID(s string) *SessionUpdate {
	su.mutation.SetVisitorID(s)
	return su
}

// SetNillableVisitorID sets the "visitor_id" field if the given value is not nil.
func (su *SessionUpdate) SetNillableVisitorID(s *string) *SessionUpdate {
	if s != nil {
		su.SetVisitorID(*s)
	}
	return su
}

// ClearVisitorID clears the value of the "visitor_id" field.
func (su *SessionUpdate) ClearVisitorID() *SessionUpdate {
	su.mutation.ClearVisitorID()
	return su
}

// SetCustomerOsContactID sets the "customer_os_contact_id" field.
func (su *SessionUpdate) SetCustomerOsContactID(s string) *SessionUpdate {
	su.mutation.SetCustomerOsContactID(s)
	return su
}

// SetNillableCustomerOsContactID sets the "customer_os_contact_id" field if the given value is not nil.
func (su *SessionUpdate) SetNillableCustomerOsContactID(s *string) *SessionUpdate {
	if s != nil {
		su.SetCustomerOsContactID(*s)
	}
	return su
}

// ClearCustomerOsContactID clears the value of the "customer_os_contact_id" field.
func (su *SessionUpdate) ClearCustomerOsContactID() *SessionUpdate {
	su.mutation.ClearCustomerOsContactID()
	return su
}

// Mutation returns the SessionMutation object of the builder.
func (su *SessionUpdate) Mutation() *SessionMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SessionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SessionUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SessionUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SessionUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SessionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   session.Table,
			Columns: session.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: session.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.AppID(); ok {
		_spec.SetField(session.FieldAppID, field.TypeString, value)
	}
	if value, ok := su.mutation.NameTracker(); ok {
		_spec.SetField(session.FieldNameTracker, field.TypeString, value)
	}
	if value, ok := su.mutation.Tenant(); ok {
		_spec.SetField(session.FieldTenant, field.TypeString, value)
	}
	if value, ok := su.mutation.DomainSessionid(); ok {
		_spec.SetField(session.FieldDomainSessionid, field.TypeString, value)
	}
	if value, ok := su.mutation.DomainSessionidx(); ok {
		_spec.SetField(session.FieldDomainSessionidx, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedDomainSessionidx(); ok {
		_spec.AddField(session.FieldDomainSessionidx, field.TypeInt, value)
	}
	if value, ok := su.mutation.SyncedToCustomerOs(); ok {
		_spec.SetField(session.FieldSyncedToCustomerOs, field.TypeBool, value)
	}
	if value, ok := su.mutation.StartTstamp(); ok {
		_spec.SetField(session.FieldStartTstamp, field.TypeTime, value)
	}
	if value, ok := su.mutation.EndTstamp(); ok {
		_spec.SetField(session.FieldEndTstamp, field.TypeTime, value)
	}
	if value, ok := su.mutation.DomainUserid(); ok {
		_spec.SetField(session.FieldDomainUserid, field.TypeString, value)
	}
	if value, ok := su.mutation.NetworkUserid(); ok {
		_spec.SetField(session.FieldNetworkUserid, field.TypeString, value)
	}
	if value, ok := su.mutation.VisitorID(); ok {
		_spec.SetField(session.FieldVisitorID, field.TypeString, value)
	}
	if su.mutation.VisitorIDCleared() {
		_spec.ClearField(session.FieldVisitorID, field.TypeString)
	}
	if value, ok := su.mutation.CustomerOsContactID(); ok {
		_spec.SetField(session.FieldCustomerOsContactID, field.TypeString, value)
	}
	if su.mutation.CustomerOsContactIDCleared() {
		_spec.ClearField(session.FieldCustomerOsContactID, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SessionUpdateOne is the builder for updating a single Session entity.
type SessionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SessionMutation
}

// SetAppID sets the "app_id" field.
func (suo *SessionUpdateOne) SetAppID(s string) *SessionUpdateOne {
	suo.mutation.SetAppID(s)
	return suo
}

// SetNameTracker sets the "name_tracker" field.
func (suo *SessionUpdateOne) SetNameTracker(s string) *SessionUpdateOne {
	suo.mutation.SetNameTracker(s)
	return suo
}

// SetTenant sets the "tenant" field.
func (suo *SessionUpdateOne) SetTenant(s string) *SessionUpdateOne {
	suo.mutation.SetTenant(s)
	return suo
}

// SetDomainSessionid sets the "domain_sessionid" field.
func (suo *SessionUpdateOne) SetDomainSessionid(s string) *SessionUpdateOne {
	suo.mutation.SetDomainSessionid(s)
	return suo
}

// SetDomainSessionidx sets the "domain_sessionidx" field.
func (suo *SessionUpdateOne) SetDomainSessionidx(i int) *SessionUpdateOne {
	suo.mutation.ResetDomainSessionidx()
	suo.mutation.SetDomainSessionidx(i)
	return suo
}

// AddDomainSessionidx adds i to the "domain_sessionidx" field.
func (suo *SessionUpdateOne) AddDomainSessionidx(i int) *SessionUpdateOne {
	suo.mutation.AddDomainSessionidx(i)
	return suo
}

// SetSyncedToCustomerOs sets the "synced_to_customer_os" field.
func (suo *SessionUpdateOne) SetSyncedToCustomerOs(b bool) *SessionUpdateOne {
	suo.mutation.SetSyncedToCustomerOs(b)
	return suo
}

// SetStartTstamp sets the "start_tstamp" field.
func (suo *SessionUpdateOne) SetStartTstamp(t time.Time) *SessionUpdateOne {
	suo.mutation.SetStartTstamp(t)
	return suo
}

// SetEndTstamp sets the "end_tstamp" field.
func (suo *SessionUpdateOne) SetEndTstamp(t time.Time) *SessionUpdateOne {
	suo.mutation.SetEndTstamp(t)
	return suo
}

// SetDomainUserid sets the "domain_userid" field.
func (suo *SessionUpdateOne) SetDomainUserid(s string) *SessionUpdateOne {
	suo.mutation.SetDomainUserid(s)
	return suo
}

// SetNetworkUserid sets the "network_userid" field.
func (suo *SessionUpdateOne) SetNetworkUserid(s string) *SessionUpdateOne {
	suo.mutation.SetNetworkUserid(s)
	return suo
}

// SetVisitorID sets the "visitor_id" field.
func (suo *SessionUpdateOne) SetVisitorID(s string) *SessionUpdateOne {
	suo.mutation.SetVisitorID(s)
	return suo
}

// SetNillableVisitorID sets the "visitor_id" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableVisitorID(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetVisitorID(*s)
	}
	return suo
}

// ClearVisitorID clears the value of the "visitor_id" field.
func (suo *SessionUpdateOne) ClearVisitorID() *SessionUpdateOne {
	suo.mutation.ClearVisitorID()
	return suo
}

// SetCustomerOsContactID sets the "customer_os_contact_id" field.
func (suo *SessionUpdateOne) SetCustomerOsContactID(s string) *SessionUpdateOne {
	suo.mutation.SetCustomerOsContactID(s)
	return suo
}

// SetNillableCustomerOsContactID sets the "customer_os_contact_id" field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableCustomerOsContactID(s *string) *SessionUpdateOne {
	if s != nil {
		suo.SetCustomerOsContactID(*s)
	}
	return suo
}

// ClearCustomerOsContactID clears the value of the "customer_os_contact_id" field.
func (suo *SessionUpdateOne) ClearCustomerOsContactID() *SessionUpdateOne {
	suo.mutation.ClearCustomerOsContactID()
	return suo
}

// Mutation returns the SessionMutation object of the builder.
func (suo *SessionUpdateOne) Mutation() *SessionMutation {
	return suo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SessionUpdateOne) Select(field string, fields ...string) *SessionUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Session entity.
func (suo *SessionUpdateOne) Save(ctx context.Context) (*Session, error) {
	var (
		err  error
		node *Session
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, suo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Session)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SessionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SessionUpdateOne) SaveX(ctx context.Context) *Session {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SessionUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SessionUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SessionUpdateOne) sqlSave(ctx context.Context) (_node *Session, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   session.Table,
			Columns: session.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: session.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`gen: missing "Session.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, session.FieldID)
		for _, f := range fields {
			if !session.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("gen: invalid field %q for query", f)}
			}
			if f != session.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.AppID(); ok {
		_spec.SetField(session.FieldAppID, field.TypeString, value)
	}
	if value, ok := suo.mutation.NameTracker(); ok {
		_spec.SetField(session.FieldNameTracker, field.TypeString, value)
	}
	if value, ok := suo.mutation.Tenant(); ok {
		_spec.SetField(session.FieldTenant, field.TypeString, value)
	}
	if value, ok := suo.mutation.DomainSessionid(); ok {
		_spec.SetField(session.FieldDomainSessionid, field.TypeString, value)
	}
	if value, ok := suo.mutation.DomainSessionidx(); ok {
		_spec.SetField(session.FieldDomainSessionidx, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedDomainSessionidx(); ok {
		_spec.AddField(session.FieldDomainSessionidx, field.TypeInt, value)
	}
	if value, ok := suo.mutation.SyncedToCustomerOs(); ok {
		_spec.SetField(session.FieldSyncedToCustomerOs, field.TypeBool, value)
	}
	if value, ok := suo.mutation.StartTstamp(); ok {
		_spec.SetField(session.FieldStartTstamp, field.TypeTime, value)
	}
	if value, ok := suo.mutation.EndTstamp(); ok {
		_spec.SetField(session.FieldEndTstamp, field.TypeTime, value)
	}
	if value, ok := suo.mutation.DomainUserid(); ok {
		_spec.SetField(session.FieldDomainUserid, field.TypeString, value)
	}
	if value, ok := suo.mutation.NetworkUserid(); ok {
		_spec.SetField(session.FieldNetworkUserid, field.TypeString, value)
	}
	if value, ok := suo.mutation.VisitorID(); ok {
		_spec.SetField(session.FieldVisitorID, field.TypeString, value)
	}
	if suo.mutation.VisitorIDCleared() {
		_spec.ClearField(session.FieldVisitorID, field.TypeString)
	}
	if value, ok := suo.mutation.CustomerOsContactID(); ok {
		_spec.SetField(session.FieldCustomerOsContactID, field.TypeString, value)
	}
	if suo.mutation.CustomerOsContactIDCleared() {
		_spec.ClearField(session.FieldCustomerOsContactID, field.TypeString)
	}
	_node = &Session{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
