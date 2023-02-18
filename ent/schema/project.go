package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return nil
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("operators", Operator.Type),
		edge.To("rodents", Rodent.Type),
	}
}
