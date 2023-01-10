package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Rodent holds the schema definition for the Rodent entity.
type Rodent struct {
	ent.Schema
}

// Fields of the Rodent.
func (Rodent) Fields() []ent.Field {
	return []ent.Field{
		field.String("xid"),
		field.String("type").Default("FancyRat"),
		field.String("codename"),
		field.String("key"),
		field.String("usercontext").Optional(),
		field.String("beacontime").Optional(),
		field.Bool("burned").Default(false),
		field.Bool("alive").Default(true),
		field.Time("joined"),
		field.Time("lastseen"),
	}
}

// Edges of the Rodent.
func (Rodent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("device", Device.Type).
			Ref("rodents").
			Unique(),
		edge.From("user", User.Type).
			Ref("rodents").
			Unique(),
		edge.To("tasks", Task.Type),
		edge.To("loot", Loot.Type),
	}
}
