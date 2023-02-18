package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Router holds the schema definition for the Router entity.
type Router struct {
	ent.Schema
}

// Fields of the Router.
func (Router) Fields() []ent.Field {
	return []ent.Field{
		field.String("username"),
		field.Bytes("privkey").Optional(),
		field.Bytes("cert").Optional(),
		field.Strings("commands").Optional(),
	}
}

// Edges of the Router.
func (Router) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("rodents", Rodent.Type),
	}
}