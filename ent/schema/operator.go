package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Operator holds the schema definition for the Operator entity.
type Operator struct {
	ent.Schema
}

// Fields of the Operator.
func (Operator) Fields() []ent.Field {
	return []ent.Field{
		field.String("username"),
		field.Bytes("privkey").Optional(),
		field.Bytes("cert").Optional(),
	}
}

// Edges of the Operator.
func (Operator) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("projects", Project.Type).
			Ref("operators"),
	}
}
