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
		field.String("rname"),
		field.Bytes("privkey"),
		field.Bytes("cert"),
		field.Strings("commands").Optional(),
		field.Strings("interfaces").Optional(),
	}
}

// Edges of the Router.
func (Router) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("rodents", Rodent.Type),
		edge.From("project", Project.Type).
			Ref("routers").
			Unique(),
	}
}
