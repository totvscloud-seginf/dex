// Code generated by ent, DO NOT EDIT.

package connector

import (
	"entgo.io/ent/dialect/sql"
	"github.com/dexidp/dex/storage/ent/db/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Connector {
	return predicate.Connector(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Connector {
	return predicate.Connector(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Connector {
	return predicate.Connector(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Connector {
	return predicate.Connector(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Connector {
	return predicate.Connector(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Connector {
	return predicate.Connector(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Connector {
	return predicate.Connector(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Connector {
	return predicate.Connector(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Connector {
	return predicate.Connector(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Connector {
	return predicate.Connector(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Connector {
	return predicate.Connector(sql.FieldContainsFold(FieldID, id))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Connector {
	return predicate.Connector(sql.FieldEQ(FieldType, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Connector {
	return predicate.Connector(sql.FieldEQ(FieldName, v))
}

// ResourceVersion applies equality check predicate on the "resource_version" field. It's identical to ResourceVersionEQ.
func ResourceVersion(v string) predicate.Connector {
	return predicate.Connector(sql.FieldEQ(FieldResourceVersion, v))
}

// Config applies equality check predicate on the "config" field. It's identical to ConfigEQ.
func Config(v []byte) predicate.Connector {
	return predicate.Connector(sql.FieldEQ(FieldConfig, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Connector {
	return predicate.Connector(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Connector {
	return predicate.Connector(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Connector {
	return predicate.Connector(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Connector {
	return predicate.Connector(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Connector {
	return predicate.Connector(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Connector {
	return predicate.Connector(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Connector {
	return predicate.Connector(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Connector {
	return predicate.Connector(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Connector {
	return predicate.Connector(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Connector {
	return predicate.Connector(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Connector {
	return predicate.Connector(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Connector {
	return predicate.Connector(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Connector {
	return predicate.Connector(sql.FieldContainsFold(FieldType, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Connector {
	return predicate.Connector(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Connector {
	return predicate.Connector(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Connector {
	return predicate.Connector(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Connector {
	return predicate.Connector(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Connector {
	return predicate.Connector(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Connector {
	return predicate.Connector(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Connector {
	return predicate.Connector(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Connector {
	return predicate.Connector(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Connector {
	return predicate.Connector(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Connector {
	return predicate.Connector(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Connector {
	return predicate.Connector(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Connector {
	return predicate.Connector(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Connector {
	return predicate.Connector(sql.FieldContainsFold(FieldName, v))
}

// ResourceVersionEQ applies the EQ predicate on the "resource_version" field.
func ResourceVersionEQ(v string) predicate.Connector {
	return predicate.Connector(sql.FieldEQ(FieldResourceVersion, v))
}

// ResourceVersionNEQ applies the NEQ predicate on the "resource_version" field.
func ResourceVersionNEQ(v string) predicate.Connector {
	return predicate.Connector(sql.FieldNEQ(FieldResourceVersion, v))
}

// ResourceVersionIn applies the In predicate on the "resource_version" field.
func ResourceVersionIn(vs ...string) predicate.Connector {
	return predicate.Connector(sql.FieldIn(FieldResourceVersion, vs...))
}

// ResourceVersionNotIn applies the NotIn predicate on the "resource_version" field.
func ResourceVersionNotIn(vs ...string) predicate.Connector {
	return predicate.Connector(sql.FieldNotIn(FieldResourceVersion, vs...))
}

// ResourceVersionGT applies the GT predicate on the "resource_version" field.
func ResourceVersionGT(v string) predicate.Connector {
	return predicate.Connector(sql.FieldGT(FieldResourceVersion, v))
}

// ResourceVersionGTE applies the GTE predicate on the "resource_version" field.
func ResourceVersionGTE(v string) predicate.Connector {
	return predicate.Connector(sql.FieldGTE(FieldResourceVersion, v))
}

// ResourceVersionLT applies the LT predicate on the "resource_version" field.
func ResourceVersionLT(v string) predicate.Connector {
	return predicate.Connector(sql.FieldLT(FieldResourceVersion, v))
}

// ResourceVersionLTE applies the LTE predicate on the "resource_version" field.
func ResourceVersionLTE(v string) predicate.Connector {
	return predicate.Connector(sql.FieldLTE(FieldResourceVersion, v))
}

// ResourceVersionContains applies the Contains predicate on the "resource_version" field.
func ResourceVersionContains(v string) predicate.Connector {
	return predicate.Connector(sql.FieldContains(FieldResourceVersion, v))
}

// ResourceVersionHasPrefix applies the HasPrefix predicate on the "resource_version" field.
func ResourceVersionHasPrefix(v string) predicate.Connector {
	return predicate.Connector(sql.FieldHasPrefix(FieldResourceVersion, v))
}

// ResourceVersionHasSuffix applies the HasSuffix predicate on the "resource_version" field.
func ResourceVersionHasSuffix(v string) predicate.Connector {
	return predicate.Connector(sql.FieldHasSuffix(FieldResourceVersion, v))
}

// ResourceVersionEqualFold applies the EqualFold predicate on the "resource_version" field.
func ResourceVersionEqualFold(v string) predicate.Connector {
	return predicate.Connector(sql.FieldEqualFold(FieldResourceVersion, v))
}

// ResourceVersionContainsFold applies the ContainsFold predicate on the "resource_version" field.
func ResourceVersionContainsFold(v string) predicate.Connector {
	return predicate.Connector(sql.FieldContainsFold(FieldResourceVersion, v))
}

// ConfigEQ applies the EQ predicate on the "config" field.
func ConfigEQ(v []byte) predicate.Connector {
	return predicate.Connector(sql.FieldEQ(FieldConfig, v))
}

// ConfigNEQ applies the NEQ predicate on the "config" field.
func ConfigNEQ(v []byte) predicate.Connector {
	return predicate.Connector(sql.FieldNEQ(FieldConfig, v))
}

// ConfigIn applies the In predicate on the "config" field.
func ConfigIn(vs ...[]byte) predicate.Connector {
	return predicate.Connector(sql.FieldIn(FieldConfig, vs...))
}

// ConfigNotIn applies the NotIn predicate on the "config" field.
func ConfigNotIn(vs ...[]byte) predicate.Connector {
	return predicate.Connector(sql.FieldNotIn(FieldConfig, vs...))
}

// ConfigGT applies the GT predicate on the "config" field.
func ConfigGT(v []byte) predicate.Connector {
	return predicate.Connector(sql.FieldGT(FieldConfig, v))
}

// ConfigGTE applies the GTE predicate on the "config" field.
func ConfigGTE(v []byte) predicate.Connector {
	return predicate.Connector(sql.FieldGTE(FieldConfig, v))
}

// ConfigLT applies the LT predicate on the "config" field.
func ConfigLT(v []byte) predicate.Connector {
	return predicate.Connector(sql.FieldLT(FieldConfig, v))
}

// ConfigLTE applies the LTE predicate on the "config" field.
func ConfigLTE(v []byte) predicate.Connector {
	return predicate.Connector(sql.FieldLTE(FieldConfig, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Connector) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Connector) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
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
func Not(p predicate.Connector) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		p(s.Not())
	})
}
