// Code generated by ent, DO NOT EDIT.

package gen

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-tracked-data/gen/session"
)

// Session is the model entity for the Session schema.
type Session struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID string `json:"app_id,omitempty"`
	// NameTracker holds the value of the "name_tracker" field.
	NameTracker string `json:"name_tracker,omitempty"`
	// Tenant holds the value of the "tenant" field.
	Tenant string `json:"tenant,omitempty"`
	// DomainSessionid holds the value of the "domain_sessionid" field.
	DomainSessionid string `json:"domain_sessionid,omitempty"`
	// DomainSessionidx holds the value of the "domain_sessionidx" field.
	DomainSessionidx int `json:"domain_sessionidx,omitempty"`
	// SyncedToCustomerOs holds the value of the "synced_to_customer_os" field.
	SyncedToCustomerOs bool `json:"synced_to_customer_os,omitempty"`
	// StartTstamp holds the value of the "start_tstamp" field.
	StartTstamp time.Time `json:"start_tstamp,omitempty"`
	// EndTstamp holds the value of the "end_tstamp" field.
	EndTstamp time.Time `json:"end_tstamp,omitempty"`
	// DomainUserid holds the value of the "domain_userid" field.
	DomainUserid string `json:"domain_userid,omitempty"`
	// NetworkUserid holds the value of the "network_userid" field.
	NetworkUserid string `json:"network_userid,omitempty"`
	// VisitorID holds the value of the "visitor_id" field.
	VisitorID string `json:"visitor_id,omitempty"`
	// CustomerOsContactID holds the value of the "customer_os_contact_id" field.
	CustomerOsContactID string `json:"customer_os_contact_id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Session) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case session.FieldSyncedToCustomerOs:
			values[i] = new(sql.NullBool)
		case session.FieldID, session.FieldDomainSessionidx:
			values[i] = new(sql.NullInt64)
		case session.FieldAppID, session.FieldNameTracker, session.FieldTenant, session.FieldDomainSessionid, session.FieldDomainUserid, session.FieldNetworkUserid, session.FieldVisitorID, session.FieldCustomerOsContactID:
			values[i] = new(sql.NullString)
		case session.FieldStartTstamp, session.FieldEndTstamp:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Session", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Session fields.
func (s *Session) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case session.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case session.FieldAppID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value.Valid {
				s.AppID = value.String
			}
		case session.FieldNameTracker:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_tracker", values[i])
			} else if value.Valid {
				s.NameTracker = value.String
			}
		case session.FieldTenant:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenant", values[i])
			} else if value.Valid {
				s.Tenant = value.String
			}
		case session.FieldDomainSessionid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field domain_sessionid", values[i])
			} else if value.Valid {
				s.DomainSessionid = value.String
			}
		case session.FieldDomainSessionidx:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field domain_sessionidx", values[i])
			} else if value.Valid {
				s.DomainSessionidx = int(value.Int64)
			}
		case session.FieldSyncedToCustomerOs:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field synced_to_customer_os", values[i])
			} else if value.Valid {
				s.SyncedToCustomerOs = value.Bool
			}
		case session.FieldStartTstamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_tstamp", values[i])
			} else if value.Valid {
				s.StartTstamp = value.Time
			}
		case session.FieldEndTstamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_tstamp", values[i])
			} else if value.Valid {
				s.EndTstamp = value.Time
			}
		case session.FieldDomainUserid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field domain_userid", values[i])
			} else if value.Valid {
				s.DomainUserid = value.String
			}
		case session.FieldNetworkUserid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field network_userid", values[i])
			} else if value.Valid {
				s.NetworkUserid = value.String
			}
		case session.FieldVisitorID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field visitor_id", values[i])
			} else if value.Valid {
				s.VisitorID = value.String
			}
		case session.FieldCustomerOsContactID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field customer_os_contact_id", values[i])
			} else if value.Valid {
				s.CustomerOsContactID = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Session.
// Note that you need to call Session.Unwrap() before calling this method if this Session
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Session) Update() *SessionUpdateOne {
	return (&SessionClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Session entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Session) Unwrap() *Session {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("gen: Session is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Session) String() string {
	var builder strings.Builder
	builder.WriteString("Session(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("app_id=")
	builder.WriteString(s.AppID)
	builder.WriteString(", ")
	builder.WriteString("name_tracker=")
	builder.WriteString(s.NameTracker)
	builder.WriteString(", ")
	builder.WriteString("tenant=")
	builder.WriteString(s.Tenant)
	builder.WriteString(", ")
	builder.WriteString("domain_sessionid=")
	builder.WriteString(s.DomainSessionid)
	builder.WriteString(", ")
	builder.WriteString("domain_sessionidx=")
	builder.WriteString(fmt.Sprintf("%v", s.DomainSessionidx))
	builder.WriteString(", ")
	builder.WriteString("synced_to_customer_os=")
	builder.WriteString(fmt.Sprintf("%v", s.SyncedToCustomerOs))
	builder.WriteString(", ")
	builder.WriteString("start_tstamp=")
	builder.WriteString(s.StartTstamp.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_tstamp=")
	builder.WriteString(s.EndTstamp.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("domain_userid=")
	builder.WriteString(s.DomainUserid)
	builder.WriteString(", ")
	builder.WriteString("network_userid=")
	builder.WriteString(s.NetworkUserid)
	builder.WriteString(", ")
	builder.WriteString("visitor_id=")
	builder.WriteString(s.VisitorID)
	builder.WriteString(", ")
	builder.WriteString("customer_os_contact_id=")
	builder.WriteString(s.CustomerOsContactID)
	builder.WriteByte(')')
	return builder.String()
}

// Sessions is a parsable slice of Session.
type Sessions []*Session

func (s Sessions) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
