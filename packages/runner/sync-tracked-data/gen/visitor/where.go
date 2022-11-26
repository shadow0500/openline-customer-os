// Code generated by ent, DO NOT EDIT.

package visitor

import (
	"entgo.io/ent/dialect/sql"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-tracked-data/gen/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// NameTracker applies equality check predicate on the "name_tracker" field. It's identical to NameTrackerEQ.
func NameTracker(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNameTracker), v))
	})
}

// Tenant applies equality check predicate on the "tenant" field. It's identical to TenantEQ.
func Tenant(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTenant), v))
	})
}

// VisitorID applies equality check predicate on the "visitor_id" field. It's identical to VisitorIDEQ.
func VisitorID(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVisitorID), v))
	})
}

// CustomerOsContactID applies equality check predicate on the "customer_os_contact_id" field. It's identical to CustomerOsContactIDEQ.
func CustomerOsContactID(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCustomerOsContactID), v))
	})
}

// DomainUserid applies equality check predicate on the "domain_userid" field. It's identical to DomainUseridEQ.
func DomainUserid(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDomainUserid), v))
	})
}

// NetworkUserid applies equality check predicate on the "network_userid" field. It's identical to NetworkUseridEQ.
func NetworkUserid(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNetworkUserid), v))
	})
}

// PageViews applies equality check predicate on the "page_views" field. It's identical to PageViewsEQ.
func PageViews(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPageViews), v))
	})
}

// Sessions applies equality check predicate on the "sessions" field. It's identical to SessionsEQ.
func Sessions(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSessions), v))
	})
}

// EngagedTimeInS applies equality check predicate on the "engaged_time_in_s" field. It's identical to EngagedTimeInSEQ.
func EngagedTimeInS(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEngagedTimeInS), v))
	})
}

// SyncedToCustomerOs applies equality check predicate on the "synced_to_customer_os" field. It's identical to SyncedToCustomerOsEQ.
func SyncedToCustomerOs(v bool) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSyncedToCustomerOs), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...string) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...string) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// AppIDContains applies the Contains predicate on the "app_id" field.
func AppIDContains(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAppID), v))
	})
}

// AppIDHasPrefix applies the HasPrefix predicate on the "app_id" field.
func AppIDHasPrefix(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAppID), v))
	})
}

// AppIDHasSuffix applies the HasSuffix predicate on the "app_id" field.
func AppIDHasSuffix(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAppID), v))
	})
}

// AppIDEqualFold applies the EqualFold predicate on the "app_id" field.
func AppIDEqualFold(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAppID), v))
	})
}

// AppIDContainsFold applies the ContainsFold predicate on the "app_id" field.
func AppIDContainsFold(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAppID), v))
	})
}

// NameTrackerEQ applies the EQ predicate on the "name_tracker" field.
func NameTrackerEQ(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNameTracker), v))
	})
}

// NameTrackerNEQ applies the NEQ predicate on the "name_tracker" field.
func NameTrackerNEQ(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNameTracker), v))
	})
}

// NameTrackerIn applies the In predicate on the "name_tracker" field.
func NameTrackerIn(vs ...string) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldNameTracker), v...))
	})
}

// NameTrackerNotIn applies the NotIn predicate on the "name_tracker" field.
func NameTrackerNotIn(vs ...string) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldNameTracker), v...))
	})
}

// NameTrackerGT applies the GT predicate on the "name_tracker" field.
func NameTrackerGT(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNameTracker), v))
	})
}

// NameTrackerGTE applies the GTE predicate on the "name_tracker" field.
func NameTrackerGTE(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNameTracker), v))
	})
}

// NameTrackerLT applies the LT predicate on the "name_tracker" field.
func NameTrackerLT(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNameTracker), v))
	})
}

// NameTrackerLTE applies the LTE predicate on the "name_tracker" field.
func NameTrackerLTE(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNameTracker), v))
	})
}

// NameTrackerContains applies the Contains predicate on the "name_tracker" field.
func NameTrackerContains(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNameTracker), v))
	})
}

// NameTrackerHasPrefix applies the HasPrefix predicate on the "name_tracker" field.
func NameTrackerHasPrefix(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNameTracker), v))
	})
}

// NameTrackerHasSuffix applies the HasSuffix predicate on the "name_tracker" field.
func NameTrackerHasSuffix(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNameTracker), v))
	})
}

// NameTrackerEqualFold applies the EqualFold predicate on the "name_tracker" field.
func NameTrackerEqualFold(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNameTracker), v))
	})
}

// NameTrackerContainsFold applies the ContainsFold predicate on the "name_tracker" field.
func NameTrackerContainsFold(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNameTracker), v))
	})
}

// TenantEQ applies the EQ predicate on the "tenant" field.
func TenantEQ(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTenant), v))
	})
}

// TenantNEQ applies the NEQ predicate on the "tenant" field.
func TenantNEQ(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTenant), v))
	})
}

// TenantIn applies the In predicate on the "tenant" field.
func TenantIn(vs ...string) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTenant), v...))
	})
}

// TenantNotIn applies the NotIn predicate on the "tenant" field.
func TenantNotIn(vs ...string) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTenant), v...))
	})
}

// TenantGT applies the GT predicate on the "tenant" field.
func TenantGT(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTenant), v))
	})
}

// TenantGTE applies the GTE predicate on the "tenant" field.
func TenantGTE(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTenant), v))
	})
}

// TenantLT applies the LT predicate on the "tenant" field.
func TenantLT(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTenant), v))
	})
}

// TenantLTE applies the LTE predicate on the "tenant" field.
func TenantLTE(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTenant), v))
	})
}

// TenantContains applies the Contains predicate on the "tenant" field.
func TenantContains(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTenant), v))
	})
}

// TenantHasPrefix applies the HasPrefix predicate on the "tenant" field.
func TenantHasPrefix(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTenant), v))
	})
}

// TenantHasSuffix applies the HasSuffix predicate on the "tenant" field.
func TenantHasSuffix(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTenant), v))
	})
}

// TenantEqualFold applies the EqualFold predicate on the "tenant" field.
func TenantEqualFold(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTenant), v))
	})
}

// TenantContainsFold applies the ContainsFold predicate on the "tenant" field.
func TenantContainsFold(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTenant), v))
	})
}

// VisitorIDEQ applies the EQ predicate on the "visitor_id" field.
func VisitorIDEQ(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVisitorID), v))
	})
}

// VisitorIDNEQ applies the NEQ predicate on the "visitor_id" field.
func VisitorIDNEQ(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVisitorID), v))
	})
}

// VisitorIDIn applies the In predicate on the "visitor_id" field.
func VisitorIDIn(vs ...string) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldVisitorID), v...))
	})
}

// VisitorIDNotIn applies the NotIn predicate on the "visitor_id" field.
func VisitorIDNotIn(vs ...string) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldVisitorID), v...))
	})
}

// VisitorIDGT applies the GT predicate on the "visitor_id" field.
func VisitorIDGT(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVisitorID), v))
	})
}

// VisitorIDGTE applies the GTE predicate on the "visitor_id" field.
func VisitorIDGTE(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVisitorID), v))
	})
}

// VisitorIDLT applies the LT predicate on the "visitor_id" field.
func VisitorIDLT(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVisitorID), v))
	})
}

// VisitorIDLTE applies the LTE predicate on the "visitor_id" field.
func VisitorIDLTE(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVisitorID), v))
	})
}

// VisitorIDContains applies the Contains predicate on the "visitor_id" field.
func VisitorIDContains(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldVisitorID), v))
	})
}

// VisitorIDHasPrefix applies the HasPrefix predicate on the "visitor_id" field.
func VisitorIDHasPrefix(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldVisitorID), v))
	})
}

// VisitorIDHasSuffix applies the HasSuffix predicate on the "visitor_id" field.
func VisitorIDHasSuffix(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldVisitorID), v))
	})
}

// VisitorIDIsNil applies the IsNil predicate on the "visitor_id" field.
func VisitorIDIsNil() predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldVisitorID)))
	})
}

// VisitorIDNotNil applies the NotNil predicate on the "visitor_id" field.
func VisitorIDNotNil() predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldVisitorID)))
	})
}

// VisitorIDEqualFold applies the EqualFold predicate on the "visitor_id" field.
func VisitorIDEqualFold(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldVisitorID), v))
	})
}

// VisitorIDContainsFold applies the ContainsFold predicate on the "visitor_id" field.
func VisitorIDContainsFold(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldVisitorID), v))
	})
}

// CustomerOsContactIDEQ applies the EQ predicate on the "customer_os_contact_id" field.
func CustomerOsContactIDEQ(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCustomerOsContactID), v))
	})
}

// CustomerOsContactIDNEQ applies the NEQ predicate on the "customer_os_contact_id" field.
func CustomerOsContactIDNEQ(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCustomerOsContactID), v))
	})
}

// CustomerOsContactIDIn applies the In predicate on the "customer_os_contact_id" field.
func CustomerOsContactIDIn(vs ...string) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCustomerOsContactID), v...))
	})
}

// CustomerOsContactIDNotIn applies the NotIn predicate on the "customer_os_contact_id" field.
func CustomerOsContactIDNotIn(vs ...string) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCustomerOsContactID), v...))
	})
}

// CustomerOsContactIDGT applies the GT predicate on the "customer_os_contact_id" field.
func CustomerOsContactIDGT(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCustomerOsContactID), v))
	})
}

// CustomerOsContactIDGTE applies the GTE predicate on the "customer_os_contact_id" field.
func CustomerOsContactIDGTE(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCustomerOsContactID), v))
	})
}

// CustomerOsContactIDLT applies the LT predicate on the "customer_os_contact_id" field.
func CustomerOsContactIDLT(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCustomerOsContactID), v))
	})
}

// CustomerOsContactIDLTE applies the LTE predicate on the "customer_os_contact_id" field.
func CustomerOsContactIDLTE(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCustomerOsContactID), v))
	})
}

// CustomerOsContactIDContains applies the Contains predicate on the "customer_os_contact_id" field.
func CustomerOsContactIDContains(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCustomerOsContactID), v))
	})
}

// CustomerOsContactIDHasPrefix applies the HasPrefix predicate on the "customer_os_contact_id" field.
func CustomerOsContactIDHasPrefix(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCustomerOsContactID), v))
	})
}

// CustomerOsContactIDHasSuffix applies the HasSuffix predicate on the "customer_os_contact_id" field.
func CustomerOsContactIDHasSuffix(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCustomerOsContactID), v))
	})
}

// CustomerOsContactIDIsNil applies the IsNil predicate on the "customer_os_contact_id" field.
func CustomerOsContactIDIsNil() predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCustomerOsContactID)))
	})
}

// CustomerOsContactIDNotNil applies the NotNil predicate on the "customer_os_contact_id" field.
func CustomerOsContactIDNotNil() predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCustomerOsContactID)))
	})
}

// CustomerOsContactIDEqualFold applies the EqualFold predicate on the "customer_os_contact_id" field.
func CustomerOsContactIDEqualFold(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCustomerOsContactID), v))
	})
}

// CustomerOsContactIDContainsFold applies the ContainsFold predicate on the "customer_os_contact_id" field.
func CustomerOsContactIDContainsFold(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCustomerOsContactID), v))
	})
}

// DomainUseridEQ applies the EQ predicate on the "domain_userid" field.
func DomainUseridEQ(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDomainUserid), v))
	})
}

// DomainUseridNEQ applies the NEQ predicate on the "domain_userid" field.
func DomainUseridNEQ(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDomainUserid), v))
	})
}

// DomainUseridIn applies the In predicate on the "domain_userid" field.
func DomainUseridIn(vs ...string) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDomainUserid), v...))
	})
}

// DomainUseridNotIn applies the NotIn predicate on the "domain_userid" field.
func DomainUseridNotIn(vs ...string) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDomainUserid), v...))
	})
}

// DomainUseridGT applies the GT predicate on the "domain_userid" field.
func DomainUseridGT(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDomainUserid), v))
	})
}

// DomainUseridGTE applies the GTE predicate on the "domain_userid" field.
func DomainUseridGTE(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDomainUserid), v))
	})
}

// DomainUseridLT applies the LT predicate on the "domain_userid" field.
func DomainUseridLT(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDomainUserid), v))
	})
}

// DomainUseridLTE applies the LTE predicate on the "domain_userid" field.
func DomainUseridLTE(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDomainUserid), v))
	})
}

// DomainUseridContains applies the Contains predicate on the "domain_userid" field.
func DomainUseridContains(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDomainUserid), v))
	})
}

// DomainUseridHasPrefix applies the HasPrefix predicate on the "domain_userid" field.
func DomainUseridHasPrefix(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDomainUserid), v))
	})
}

// DomainUseridHasSuffix applies the HasSuffix predicate on the "domain_userid" field.
func DomainUseridHasSuffix(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDomainUserid), v))
	})
}

// DomainUseridEqualFold applies the EqualFold predicate on the "domain_userid" field.
func DomainUseridEqualFold(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDomainUserid), v))
	})
}

// DomainUseridContainsFold applies the ContainsFold predicate on the "domain_userid" field.
func DomainUseridContainsFold(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDomainUserid), v))
	})
}

// NetworkUseridEQ applies the EQ predicate on the "network_userid" field.
func NetworkUseridEQ(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNetworkUserid), v))
	})
}

// NetworkUseridNEQ applies the NEQ predicate on the "network_userid" field.
func NetworkUseridNEQ(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNetworkUserid), v))
	})
}

// NetworkUseridIn applies the In predicate on the "network_userid" field.
func NetworkUseridIn(vs ...string) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldNetworkUserid), v...))
	})
}

// NetworkUseridNotIn applies the NotIn predicate on the "network_userid" field.
func NetworkUseridNotIn(vs ...string) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldNetworkUserid), v...))
	})
}

// NetworkUseridGT applies the GT predicate on the "network_userid" field.
func NetworkUseridGT(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNetworkUserid), v))
	})
}

// NetworkUseridGTE applies the GTE predicate on the "network_userid" field.
func NetworkUseridGTE(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNetworkUserid), v))
	})
}

// NetworkUseridLT applies the LT predicate on the "network_userid" field.
func NetworkUseridLT(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNetworkUserid), v))
	})
}

// NetworkUseridLTE applies the LTE predicate on the "network_userid" field.
func NetworkUseridLTE(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNetworkUserid), v))
	})
}

// NetworkUseridContains applies the Contains predicate on the "network_userid" field.
func NetworkUseridContains(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNetworkUserid), v))
	})
}

// NetworkUseridHasPrefix applies the HasPrefix predicate on the "network_userid" field.
func NetworkUseridHasPrefix(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNetworkUserid), v))
	})
}

// NetworkUseridHasSuffix applies the HasSuffix predicate on the "network_userid" field.
func NetworkUseridHasSuffix(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNetworkUserid), v))
	})
}

// NetworkUseridEqualFold applies the EqualFold predicate on the "network_userid" field.
func NetworkUseridEqualFold(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNetworkUserid), v))
	})
}

// NetworkUseridContainsFold applies the ContainsFold predicate on the "network_userid" field.
func NetworkUseridContainsFold(v string) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNetworkUserid), v))
	})
}

// PageViewsEQ applies the EQ predicate on the "page_views" field.
func PageViewsEQ(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPageViews), v))
	})
}

// PageViewsNEQ applies the NEQ predicate on the "page_views" field.
func PageViewsNEQ(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPageViews), v))
	})
}

// PageViewsIn applies the In predicate on the "page_views" field.
func PageViewsIn(vs ...int) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPageViews), v...))
	})
}

// PageViewsNotIn applies the NotIn predicate on the "page_views" field.
func PageViewsNotIn(vs ...int) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPageViews), v...))
	})
}

// PageViewsGT applies the GT predicate on the "page_views" field.
func PageViewsGT(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPageViews), v))
	})
}

// PageViewsGTE applies the GTE predicate on the "page_views" field.
func PageViewsGTE(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPageViews), v))
	})
}

// PageViewsLT applies the LT predicate on the "page_views" field.
func PageViewsLT(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPageViews), v))
	})
}

// PageViewsLTE applies the LTE predicate on the "page_views" field.
func PageViewsLTE(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPageViews), v))
	})
}

// SessionsEQ applies the EQ predicate on the "sessions" field.
func SessionsEQ(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSessions), v))
	})
}

// SessionsNEQ applies the NEQ predicate on the "sessions" field.
func SessionsNEQ(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSessions), v))
	})
}

// SessionsIn applies the In predicate on the "sessions" field.
func SessionsIn(vs ...int) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSessions), v...))
	})
}

// SessionsNotIn applies the NotIn predicate on the "sessions" field.
func SessionsNotIn(vs ...int) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSessions), v...))
	})
}

// SessionsGT applies the GT predicate on the "sessions" field.
func SessionsGT(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSessions), v))
	})
}

// SessionsGTE applies the GTE predicate on the "sessions" field.
func SessionsGTE(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSessions), v))
	})
}

// SessionsLT applies the LT predicate on the "sessions" field.
func SessionsLT(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSessions), v))
	})
}

// SessionsLTE applies the LTE predicate on the "sessions" field.
func SessionsLTE(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSessions), v))
	})
}

// EngagedTimeInSEQ applies the EQ predicate on the "engaged_time_in_s" field.
func EngagedTimeInSEQ(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEngagedTimeInS), v))
	})
}

// EngagedTimeInSNEQ applies the NEQ predicate on the "engaged_time_in_s" field.
func EngagedTimeInSNEQ(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEngagedTimeInS), v))
	})
}

// EngagedTimeInSIn applies the In predicate on the "engaged_time_in_s" field.
func EngagedTimeInSIn(vs ...int) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEngagedTimeInS), v...))
	})
}

// EngagedTimeInSNotIn applies the NotIn predicate on the "engaged_time_in_s" field.
func EngagedTimeInSNotIn(vs ...int) predicate.Visitor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEngagedTimeInS), v...))
	})
}

// EngagedTimeInSGT applies the GT predicate on the "engaged_time_in_s" field.
func EngagedTimeInSGT(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEngagedTimeInS), v))
	})
}

// EngagedTimeInSGTE applies the GTE predicate on the "engaged_time_in_s" field.
func EngagedTimeInSGTE(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEngagedTimeInS), v))
	})
}

// EngagedTimeInSLT applies the LT predicate on the "engaged_time_in_s" field.
func EngagedTimeInSLT(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEngagedTimeInS), v))
	})
}

// EngagedTimeInSLTE applies the LTE predicate on the "engaged_time_in_s" field.
func EngagedTimeInSLTE(v int) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEngagedTimeInS), v))
	})
}

// SyncedToCustomerOsEQ applies the EQ predicate on the "synced_to_customer_os" field.
func SyncedToCustomerOsEQ(v bool) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSyncedToCustomerOs), v))
	})
}

// SyncedToCustomerOsNEQ applies the NEQ predicate on the "synced_to_customer_os" field.
func SyncedToCustomerOsNEQ(v bool) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSyncedToCustomerOs), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Visitor) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Visitor) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Visitor) predicate.Visitor {
	return predicate.Visitor(func(s *sql.Selector) {
		p(s.Not())
	})
}
