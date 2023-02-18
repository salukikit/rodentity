package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("xid"),
		field.String("type").Default("cmd"),
		field.Strings("args").Optional(),
		field.Bytes("data").Optional(),
		field.Bytes("result").Optional(),
		field.Bool("Executed").Default(false),
		field.Bool("looted").Default(false),
		field.Time("requestedat"),
		field.Time("completedat").Optional(),
		field.Strings("TTPs").Optional(),
	}

}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("rodent", Rodent.Type).Ref("tasks").Unique(),
		edge.From("operator", Operator.Type).Ref("tasks").Unique(),
		edge.To("loot", Loot.Type),
	}
}
