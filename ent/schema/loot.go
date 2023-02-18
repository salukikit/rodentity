package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Loot holds the schema definition for the Loot entity.
type Loot struct {
	ent.Schema
}

// Fields of the Loot.
func (Loot) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").Values("cred", "key", "cert", "enum", "objective", "pii", "other"),
		field.String("location").Default("unknown"),
		field.Bytes("data"),
		field.Time("collectedon"),
	}
}

// Edges of the Loot.
func (Loot) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("rodent", Rodent.Type).Ref("loot").Unique(),
		edge.From("task", Task.Type).Ref("loot").Unique(),
	}
}
