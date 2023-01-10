package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Domain holds the schema definition for the Domain entity.
type Domain struct {
	ent.Schema
}

// Fields of the Domain.
func (Domain) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Bool("AD").Default(true),
		field.Bool("owned").Default(false),
		field.String("cloud").Default("unknown"),
	}
}

// Edges of the Domain.
func (Domain) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("devices", Device.Type),
		edge.To("users", User.Type),
		edge.To("groups", Group.Type),
		edge.To("childdomains", Domain.Type),
		edge.From("parentdomain", Domain.Type).Ref("childdomains").Unique(),
	}
}
