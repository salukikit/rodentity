package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

func (Project) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// Embed the BaseMixin in the Project schema.
		BaseMixin{},
	}
}

// Fields of the Project.
func (Project) Fields() []ent.Field {

	return []ent.Field{
		field.String("name"),
		field.String("objective").Optional(),
		field.Time("end_date").Optional(),
		field.Time("start_date").Optional(),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("operators", Operator.Type),
		edge.To("rodents", Rodent.Type),
		edge.To("routers", Router.Type),
	}
}
